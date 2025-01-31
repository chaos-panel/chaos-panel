package terminal

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strings"
	"sync"
	"time"

	v1 "github.com/chaos-plus/chaos-plus/api/terminals/v1"
	"github.com/chaos-plus/chaos-plus/internal/dao"
	"github.com/chaos-plus/chaos-plus/internal/model/do"
	"github.com/chaos-plus/chaos-plus/internal/model/entity"
	"github.com/chaos-plus/chaos-plus/internal/service"
	"github.com/chaos-plus/chaos-plus/utility/converter"
	"github.com/chaos-plus/chaos-plus/utility/middleware"
	"github.com/chaos-plus/chaos-plus/utility/page"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/grpool"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gorilla/websocket"
)

type sTerminal struct{}

type TerminalMessage struct {
	Rows int    `json:"rows"`
	Cols int    `json:"cols"`
	Data string `json:"data"`
}

type TerminalHandler interface {
	Open(rows int, cols int) (err error)

	Rezize(rows int, cols int) (b bool, err error)

	ReadIn(buf []byte) (n int, err error)

	ReadErr(buf []byte) (n int, err error)

	WriteOut(buf []byte) (n int, err error)

	Wait() (err error)

	Close() (err error)
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func init() {
	service.RegisterTerminal(&sTerminal{})
}

func New() service.ITerminal {
	return &sTerminal{}
}

func (s *sTerminal) Create(ctx context.Context, req *v1.CreateReq) (res *v1.CreateRes, err error) {
	do, err := converter.ToAny[do.Terminals](*req)
	if err != nil {
		return
	}
	r := g.RequestFromCtx(ctx)
	uid := middleware.GetUserId(r)
	do.Id = service.Guid().NextId()
	do.TenantId = 0
	do.CreatedBy = uid
	do.UpdatedBy = uid
	do.CreatedAt = gtime.Now()
	do.UpdatedAt = gtime.Now()
	do.DeletedBy = 0
	do.DeletedAt = gtime.NewFromTime(time.Unix(0, 0))
	err = dao.Terminals.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		_, err := tx.Model(&entity.Terminals{}).Insert(do)
		return err
	})
	if err != nil {
		return
	}
	dao.Terminals.Ctx(ctx).Where("id = ?", do.Id).Scan(&res)
	return
}

func (s *sTerminal) GetOne(ctx context.Context, req *v1.GetOneReq) (res *v1.GetOneRes, err error) {
	m := dao.Terminals.Ctx(ctx)
	err = m.Where("id = ?", req.Id).Scan(&res)
	return
}

func (s *sTerminal) GetList(ctx context.Context, req *v1.GetListReq) (res *v1.GetListRes, err error) {
	var entities []entity.Terminals
	var count int
	m := dao.Terminals.Ctx(ctx)
	m = page.ParseQueriesFromCtx(ctx, m, req)
	m = page.NonDeleted(m)
	// m = page.NonLocked(m)
	err = m.ScanAndCount(&entities, &count, true)
	res = &v1.GetListRes{}
	res.Offset = req.Offset
	res.Limit = req.Limit
	res.Count = len(entities)
	res.Total = count
	res.List = entities
	return
}

func (s *sTerminal) Delete(ctx context.Context, req *v1.DeleteReq) (res *v1.DeleteRes, err error) {
	r := g.RequestFromCtx(ctx)
	uid := middleware.GetUserId(r)
	err = dao.Terminals.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		result, err := tx.Model(&entity.Terminals{}).Where("id = ?", req.Id).Update(&do.Terminals{
			DeletedAt: gtime.Now(),
			DeletedBy: uid,
		})
		if err != nil {
			return err
		}
		rows, err := result.RowsAffected()
		if err != nil {
			return err
		}
		res = (*v1.DeleteRes)(&rows)
		return err
	})
	return
}

func (s *sTerminal) GetSession(ctx context.Context, req *v1.GetSessionReq) (res *v1.GetSessionRes, err error) {
	r := g.RequestFromCtx(ctx)

	terminal := &entity.Terminals{}
	m := dao.Terminals.Ctx(ctx)
	err = m.Where("id = ?", req.Id).Scan(&terminal)
	if err != nil {
		g.Log().Error(ctx, err)
		return
	}

	if terminal == nil {
		g.Log().Error(ctx, "terminal is nil")
		return
	}

	if !terminal.DeletedAt.IsZero() {
		g.Log().Error(ctx, "terminal is deleted")
		return
	}

	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	conn, err := upgrader.Upgrade(r.Response.Writer, r.Request, nil)

	if err != nil {
		return nil, err
	}
	defer conn.Close()

	err = HandleImpl(ctx, conn, terminal)
	return
}

func HandleImpl(ctx context.Context, ws *websocket.Conn, terminal *entity.Terminals) (err error) {
	r := g.RequestFromCtx(ctx)
	cols := r.GetParam("cols", "80").Int() // w
	rows := r.GetParam("rows", "40").Int() // h

	var handler TerminalHandler = nil
	if strings.ToLower(terminal.Type) == "agent" {
		handler = NewHandlerByAgent()
	} else if strings.ToLower(terminal.Type) == "auth" {
		handler = NewHandlerByAuth(ctx, terminal)
	} else if strings.ToLower(terminal.Type) == "host" {
		handler = NewHandlerByHost()
	} else if strings.ToLower(terminal.Type) == "plugin" {
		handler = NewHandlerByPlugin()
	} else {
		err = errors.New("unknown terminal type:" + terminal.Type)
	}

	if err != nil || handler == nil {
		g.Log().Errorf(ctx, "terminal init failed: %s", err)
		return err
	}

	err = handler.Open(rows, cols)
	if err != nil {
		g.Log().Errorf(ctx, "terminal open error: %s", err)
		return err
	}
	defer handler.Close()

	ws.SetCloseHandler(func(code int, text string) error {
		return handler.Close()
	})

	grpool.AddWithRecover(ctx, func(ctx context.Context) {
		if err = handler.Wait(); err != nil {
			g.Log().Error(ctx, err)
		}
		g.Log().Info(ctx, "terminal session closed")
		ws.Close()
	}, func(ctx context.Context, err error) {
		g.Log().Error(ctx, err)
	})

	var wg sync.WaitGroup
	var mutex sync.Mutex

	writeWsMessage := func(messageType int, data []byte) error {
		mutex.Lock()
		defer mutex.Unlock()
		if err := ws.WriteMessage(messageType, data); err != nil {
			g.Log().Error(ctx, err)
		}
		return err
	}

	grpool.AddWithRecover(ctx, func(ctx context.Context) {
		defer wg.Done()
		for {
			typ, msg, err := ws.ReadMessage()
			if err != nil {
				g.Log().Error(ctx, "websocket read error:", err)
				return
			}
			if typ == websocket.CloseMessage {
				err := writeWsMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
				if err != nil {
					g.Log().Error(ctx, "websocket write error:", err)
				}
				ws.Close()
				return
			}
			if typ == websocket.PongMessage {
				continue
			}
			if typ == websocket.PingMessage {
				err = writeWsMessage(websocket.PongMessage, nil)
				if err != nil {
					g.Log().Error(ctx, "websocket write error:", err)
					return
				}
				continue
			}
			if typ == websocket.BinaryMessage {
				g.Log().Error(ctx, "Unsupported message type:", typ)
				return
			}
			if typ != websocket.TextMessage {
				g.Log().Error(ctx, "Unsupported message type:", typ)
				return
			}
			var resmsg TerminalMessage
			var reqmsg TerminalMessage
			jsonErr := json.Unmarshal(msg, &reqmsg)
			isJson := jsonErr == nil && reqmsg.Data != ""
			if !isJson {
				resmsg.Data = string(msg)
			} else {
				resmsg.Data = reqmsg.Data
			}
			if reqmsg.Cols > 0 && reqmsg.Rows > 0 {
				b, err := handler.Rezize(reqmsg.Rows, reqmsg.Cols)
				if err != nil {
					g.Log().Error(ctx, "terminal rezise error:", err)
				}
				if err == nil && b {
					resmsg.Cols = reqmsg.Cols
					resmsg.Rows = reqmsg.Rows
				}
			}
			_, err = handler.WriteOut([]byte(resmsg.Data))
			if err != nil {
				g.Log().Error(ctx, "terminal write error:", err)
				return
			}
			if resmsg.Cols > 0 && resmsg.Rows > 0 {
				resmsg.Data = ""
				resjson, err := json.Marshal(resmsg)
				if err != nil {
					g.Log().Error(ctx, "terminal marshal error:", err)
					continue
				}
				err = writeWsMessage(websocket.TextMessage, resjson)
				if err != nil {
					return
				}
			}
		}
	}, func(ctx context.Context, err error) {
		g.Log().Error(ctx, err)
	})

	grpool.AddWithRecover(ctx, func(ctx context.Context) {
		defer wg.Done()
		buf := make([]byte, 1024)
		for {
			n, err := handler.ReadIn(buf)
			if err != nil {
				g.Log().Error(ctx, "terminal read error:", err)
				break
			}
			err = writeWsMessage(websocket.TextMessage, buf[:n])
			if err != nil {
				break
			}
		}
		handler.Close()
	}, func(ctx context.Context, err error) {
		g.Log().Error(ctx, err)
	})

	grpool.AddWithRecover(ctx, func(ctx context.Context) {
		defer wg.Done()
		buf := make([]byte, 1024)
		for {
			n, err := handler.ReadErr(buf)
			if err != nil {
				g.Log().Error(ctx, "terminal read error:", err)
				break
			}
			err = writeWsMessage(websocket.TextMessage, buf[:n])
			if err != nil {
				break
			}
		}
	}, func(ctx context.Context, err error) {
		g.Log().Error(ctx, err)
	})

	wg.Add(3)
	wg.Wait()
	g.Log().Info(ctx, "terminal session exited")
	return
}
