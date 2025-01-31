package terminal

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"runtime"

	"github.com/aymanbagabas/go-pty"
	"github.com/chaos-plus/chaos-plus/utility/combine"
)

type HandlerByHost struct {
	pty pty.Pty
	cmd *pty.Cmd
}

func NewHandlerByHost() *HandlerByHost {
	return &HandlerByHost{}
}

func (h *HandlerByHost) Open(rows int, cols int) (err error) {
	// 根据操作系统选择不同的shell
	shell, args := getShellCommand()

	// cmd := exec.Command(shell, args...)
	// // 设置环境变量
	// cmd.Env = append(os.Environ(),
	// 	"TERM=xterm-256color",
	// 	fmt.Sprintf("COLUMNS=%d", cols),
	// 	fmt.Sprintf("LINES=%d", rows),
	// )
	// // 启动子进程
	// if err := cmd.Start(); err != nil {
	// 	return fmt.Errorf("failed to start command: %w", err)
	// }

	// h.cmd = *cmd

	pty, err := pty.New()
	if err != nil {
		return
	}
	h.pty = pty

	// 创建PTY
	cmd := pty.Command(shell, args...)

	// 设置初始窗口大小
	if err = pty.Resize(cols, rows); err != nil {
		return fmt.Errorf("failed to set window size: %w", err)
	}
	if err = cmd.Start(); err != nil {
		return
	}
	h.cmd = cmd
	return
}

func (h *HandlerByHost) Rezize(rows int, cols int) (b bool, err error) {
	// 设置初始窗口大小
	err = h.pty.Resize(cols, rows)
	b = err != nil
	return
}

func (h *HandlerByHost) ReadIn(buf []byte) (n int, err error) {
	return h.pty.Read(buf)
}

func (h *HandlerByHost) ReadErr(buf []byte) (n int, err error) {
	return -1, errors.New("stderr deafult rediect to stdout")
}

func (h *HandlerByHost) WriteOut(buf []byte) (n int, err error) {
	return h.pty.Write(buf)
}

func (h *HandlerByHost) Wait() (err error) {
	return h.cmd.Wait()
}

func (h *HandlerByHost) Close() (err error) {
	return combine.Errors(
		h.pty.Close(),
		h.cmd.Process.Kill(),
	)
}

// getShellCommand 根据操作系统返回对应的shell命令
func getShellCommand() (string, []string) {
	switch runtime.GOOS {
	case "windows":
		if _, err := exec.LookPath("powershell.exe"); err == nil {
			return "powershell.exe", []string{"-NoLogo"}
		} else {
			return "cmd.exe", []string{"/c"}
		}
	default:
		shell := os.Getenv("SHELL")
		if shell == "" {
			if _, err := os.Stat("/bin/bash"); err == nil {
				shell = "/bin/bash"
			} else {
				shell = "/bin/sh"
			}
		}
		return shell, []string{"-l"}
	}
}
