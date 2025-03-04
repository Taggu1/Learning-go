package iomanager

type IOManager interface {
	ReadLines() ([]string, error)
	WriteResult(data interface{}, errChan chan error)
}
