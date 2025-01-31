package terminal

type HandlerByAgent struct {
}

func NewHandlerByAgent() *HandlerByAgent {
	return &HandlerByAgent{}
}

func (h *HandlerByAgent) Open(rows int, cols int) (err error) {
	return
}

func (h *HandlerByAgent) Rezize(rows int, cols int) (b bool, err error) {
	return
}

func (h *HandlerByAgent) ReadIn(buf []byte) (n int, err error) {
	return
}

func (h *HandlerByAgent) ReadErr(buf []byte) (n int, err error) {
	return
}

func (h *HandlerByAgent) WriteOut(buf []byte) (n int, err error) {
	return
}

func (h *HandlerByAgent) Wait() (err error) {
	return
}

func (h *HandlerByAgent) Close() (err error) {
	return
}
