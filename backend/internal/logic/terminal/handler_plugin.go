package terminal

type HandlerByPlugin struct {
}

func NewHandlerByPlugin() *HandlerByPlugin {
	return &HandlerByPlugin{}
}

func (h *HandlerByPlugin) Open(rows int, cols int) (err error) {
	return
}

func (h *HandlerByPlugin) Rezize(rows int, cols int) (b bool, err error) {
	return
}

func (h *HandlerByPlugin) ReadIn(buf []byte) (n int, err error) {
	return
}

func (h *HandlerByPlugin) ReadErr(buf []byte) (n int, err error) {
	return
}

func (h *HandlerByPlugin) WriteOut(buf []byte) (n int, err error) {
	return
}

func (h *HandlerByPlugin) Wait() (err error) {
	return
}

func (h *HandlerByPlugin) Close() (err error) {
	return
}
