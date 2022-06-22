package service

type Server interface {
	Server() error
	Stop()
}
