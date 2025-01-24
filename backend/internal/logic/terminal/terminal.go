package guid

import (
	"context"
	"errors"
	"fmt"
	"log"
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

type WsMessage struct {
	Type string `json:"type"`
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
		log.Fatalf("Failed to dial: %s", err)
	}
	defer client.Close()

	session, err := client.NewSession()
	if err != nil {
		log.Fatalf("Failed to create session: %s", err)
	}
	defer session.Close()

	conn.SetCloseHandler(func(code int, text string) error {
		session.Close()
		return nil
	})
	grpool.AddWithRecover(ctx, func(ctx context.Context) {
		session.Wait()
		conn.Close()
	}, func(ctx context.Context, err error) {
		g.Log().Error(ctx, err)
	})

	var wg sync.WaitGroup

	grpool.AddWithRecover(ctx, func(ctx context.Context) {
		defer wg.Done()
		for {
			_, msg, err := conn.ReadMessage()
			if err != nil {
				return
			}
			_, err = session.Stdout.Write(msg)
			if err != nil {
				return
			}
		}
	}, func(ctx context.Context, err error) {
		g.Log().Error(ctx, err)
	})

	grpool.AddWithRecover(ctx, func(ctx context.Context) {
		defer wg.Done()
		buf := make([]byte, 1024)
		for {
			n, err := session.Stdin.Read(buf[:0])
			if err != nil {
				return
			}
			err = conn.WriteMessage(websocket.TextMessage, buf[:n])
			if err != nil {
				return
			}
		}
	}, func(ctx context.Context, err error) {
		g.Log().Error(ctx, err)
	})

	// 等待所有 goroutine 完成
	wg.Add(2) // 添加两个 goroutine 的等待
	wg.Wait() // 阻塞直到所有 goroutine 完成

	return
}

func handleByAgent(conn *websocket.Conn, Terminal *entity.Terminals) (err error) {

	return
}
