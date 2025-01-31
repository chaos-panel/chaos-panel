package terminal

import (
	"context"
	"fmt"
	"io"

	"github.com/chaos-plus/chaos-plus/internal/model/entity"
	"github.com/chaos-plus/chaos-plus/utility/combine"
	"github.com/gogf/gf/v2/frame/g"
	"golang.org/x/crypto/ssh"
)

type HandlerByAuth struct {
	ctx    context.Context
	config *entity.Terminals

	client    *ssh.Client
	session   *ssh.Session
	writerOut io.WriteCloser
	readerIn  io.Reader
	readerErr io.Reader
}

func NewHandlerByAuth(ctx context.Context, config *entity.Terminals) *HandlerByAuth {
	return &HandlerByAuth{
		config: config,
	}
}

func (h *HandlerByAuth) Open(rows int, cols int) (err error) {
	addr := fmt.Sprintf("%s:%d", h.config.Host, h.config.Port)
	config := &ssh.ClientConfig{
		User: h.config.Username,
		Auth: []ssh.AuthMethod{
			ssh.Password(h.config.Password),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	client, err := ssh.Dial("tcp", addr, config)
	if err != nil {
		g.Log().Errorf(h.ctx, "Failed to dial: %s", err)
		return err
	}
	h.client = client

	term, err := client.NewSession()
	if err != nil {
		g.Log().Errorf(h.ctx, "Failed to create session: %s", err)
		return err
	}
	h.session = term

	if err := term.RequestPty("xterm", rows, cols, ssh.TerminalModes{
		ssh.ECHO: 1, // 启用回显
	}); err != nil {
		g.Log().Error(h.ctx, err)
		return err
	}
	stdin, err := term.StdinPipe()
	if err != nil {
		g.Log().Errorf(h.ctx, "Failed to create stdin pipe: %s", err)
		return err
	}
	stdout, err := term.StdoutPipe()
	if err != nil {
		g.Log().Errorf(h.ctx, "Failed to create stdout pipe: %s", err)
		return err
	}
	stderr, err := term.StderrPipe()
	if err != nil {
		g.Log().Errorf(h.ctx, "Failed to create stderr pipe: %s", err)
		return err
	}

	h.writerOut = stdin
	h.readerIn = stdout
	h.readerErr = stderr

	if err := term.Shell(); err != nil {
		g.Log().Error(h.ctx, err)
		return err
	}

	return
}

func (h *HandlerByAuth) Rezize(rows int, cols int) (b bool, err error) {
	return h.session.SendRequest("window-change", false, []byte{
		0, 0, byte(rows), byte(cols), // window change size: rows = 100, columns = 200
	})
}

func (h *HandlerByAuth) ReadIn(buf []byte) (n int, err error) {
	return h.readerIn.Read(buf)
}

func (h *HandlerByAuth) ReadErr(buf []byte) (n int, err error) {
	return h.readerErr.Read(buf)
}

func (h *HandlerByAuth) WriteOut(buf []byte) (n int, err error) {
	return h.writerOut.Write(buf)
}

func (h *HandlerByAuth) Wait() (err error) {
	return h.session.Wait()
}

func (h *HandlerByAuth) Close() (err error) {
	return combine.Errors(
		h.session.Close(),
		h.client.Close(),
	)
}
