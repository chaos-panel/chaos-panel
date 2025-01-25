package guid

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
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
	"golang.org/x/crypto/ssh"
)

type sTerminal struct{}

type TerminalMessage struct {
	Rows int    `json:"rows"`
	Cols int    `json:"cols"`
	Data string `json:"data"`
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

	if strings.ToLower(terminal.Type) == "agent" {
		err = handleByAgent(conn, terminal)
	} else if strings.ToLower(terminal.Type) == "auth" {
		err = handleByAuth(ctx, conn, terminal)
	} else {
		err = errors.New("unknown terminal type:" + terminal.Type)
	}
	return
}

func handleByAuth(ctx context.Context, conn *websocket.Conn, Terminal *entity.Terminals) (err error) {
	addr := fmt.Sprintf("%s:%d", Terminal.Host, Terminal.Port)
	config := &ssh.ClientConfig{
		User: Terminal.Username,
		Auth: []ssh.AuthMethod{
			ssh.Password(Terminal.Password),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	client, err := ssh.Dial("tcp", addr, config)
	if err != nil {
		g.Log().Errorf(ctx, "Failed to dial: %s", err)
		return err
	}
	defer client.Close()

	session, err := client.NewSession()
	if err != nil {
		g.Log().Errorf(ctx, "Failed to create session: %s", err)
		return err
	}
	defer session.Close()

	if err := session.RequestPty("xterm", 80, 40, ssh.TerminalModes{
		ssh.ECHO: 1, // 启用回显
	}); err != nil {
		g.Log().Error(ctx, err)
		return err
	}

	stdin, err := session.StdinPipe()
	if err != nil {
		g.Log().Errorf(ctx, "Failed to create stdin pipe: %s", err)
		return err
	}
	stdout, err := session.StdoutPipe()
	if err != nil {
		g.Log().Errorf(ctx, "Failed to create stdout pipe: %s", err)
		return err
	}
	stderr, err := session.StderrPipe()
	if err != nil {
		g.Log().Errorf(ctx, "Failed to create stderr pipe: %s", err)
		return err
	}

	conn.SetCloseHandler(func(code int, text string) error {
		session.Close()
		return nil
	})

	grpool.AddWithRecover(ctx, func(ctx context.Context) {
		if err = session.Wait(); err != nil {
			g.Log().Error(ctx, err)
		}
		g.Log().Info(ctx, "Session wait closed")
		conn.Close()
	}, func(ctx context.Context, err error) {
		g.Log().Error(ctx, err)
	})

	var wg sync.WaitGroup
	var mutex sync.Mutex

	// write message to websocket connection
	writeMessage := func(messageType int, data []byte) error {
		mutex.Lock()
		defer mutex.Unlock()
		if err := conn.WriteMessage(messageType, data); err != nil {
			g.Log().Error(ctx, err)
		}
		return err
	}

	// WebSocket -> SSH stdin
	grpool.AddWithRecover(ctx, func(ctx context.Context) {
		defer wg.Done()
		for {
			typ, msg, err := conn.ReadMessage()
			if err != nil {
				g.Log().Error(ctx, "WebSocket read error:", err)
				return
			}
			if typ == websocket.CloseMessage {
				err := writeMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
				if err != nil {
					g.Log().Error(ctx, "WebSocket write error:", err)
				}
				conn.Close()
				return
			}
			if typ == websocket.PongMessage {
				continue
			}
			if typ == websocket.PingMessage {
				err = writeMessage(websocket.PongMessage, nil)
				if err != nil {
					g.Log().Error(ctx, "WebSocket write error:", err)
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
				b, err := session.SendRequest("window-change", false, []byte{
					0, 0, byte(reqmsg.Rows), byte(reqmsg.Cols), // window change size: rows = 100, columns = 200
				})
				if err != nil {
					g.Log().Error(ctx, "Failed to send window-change request:", err)
				}
				if err == nil && b {
					resmsg.Cols = reqmsg.Cols
					resmsg.Rows = reqmsg.Rows
				}
			}
			_, err = stdin.Write([]byte(resmsg.Data))
			if err != nil {
				g.Log().Error(ctx, "SSH stdin write error:", err)
				return
			}
			if resmsg.Cols > 0 && resmsg.Rows > 0 {
				resmsg.Data = ""
				resjson, err := json.Marshal(resmsg)
				if err != nil {
					g.Log().Error(ctx, "Failed to marshal resmsg:", err)
					continue
				}
				err = writeMessage(websocket.TextMessage, resjson)
				if err != nil {
					g.Log().Error(ctx, "WebSocket write error (stderr):", err)
					return
				}
			}
		}
	}, func(ctx context.Context, err error) {
		g.Log().Error(ctx, err)
	})

	// SSH stdout -> WebSocket
	grpool.AddWithRecover(ctx, func(ctx context.Context) {
		defer wg.Done()
		buf := make([]byte, 1024)
		for {
			n, err := stdout.Read(buf)
			if err != nil {
				g.Log().Error(ctx, "SSH stdout read error:", err)
				return
			}
			err = writeMessage(websocket.TextMessage, buf[:n])
			if err != nil {
				g.Log().Error(ctx, "WebSocket write error:", err)
				return
			}
		}
	}, func(ctx context.Context, err error) {
		g.Log().Error(ctx, err)
	})

	// SSH stderr -> WebSocket (用于错误信息)
	grpool.AddWithRecover(ctx, func(ctx context.Context) {
		defer wg.Done()
		buf := make([]byte, 1024)
		for {
			n, err := stderr.Read(buf)
			if err != nil {
				g.Log().Error(ctx, "SSH stderr read error:", err)
				return
			}
			err = writeMessage(websocket.TextMessage, buf[:n])
			if err != nil {
				g.Log().Error(ctx, "WebSocket write error (stderr):", err)
				return
			}
		}
	}, func(ctx context.Context, err error) {
		g.Log().Error(ctx, err)
	})

	wg.Add(3)
	wg.Wait()
	g.Log().Info(ctx, "Terminal session exited")
	return
}

func handleByAgent(conn *websocket.Conn, Terminal *entity.Terminals) (err error) {

	return
}
