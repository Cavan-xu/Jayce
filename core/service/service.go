package service

type Server interface {
	Start() error
	Stop()
}
