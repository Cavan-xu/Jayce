package tcp

import (
	"net"
	"strconv"

	"Jayce/core/log"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
)

type server struct {
	Config
	Logger *zap.Logger
}

func NewServer(c Config, opts ...option) (*server, error) {
	s := &server{
		Config: c,
	}

	if err := s.setUp(); err != nil {
		return nil, err
	}

	for _, opt := range opts {
		opt(s)
	}

	return s, nil
}

func (s *server) setUp() error {
	err := s.Config.SetUp()
	if err != nil {
		return err
	}

	logConf := &lumberjack.Logger{
		Filename:   s.Log.FileName,
		MaxSize:    s.Log.MaxSize,
		MaxBackups: s.Log.MaxBackups,
		MaxAge:     s.Log.MaxAges,
		Compress:   s.Log.Compress,
	}

	s.Logger = log.Init(logConf)
	return nil
}

// start run the tcp server and accept tcp connection
func (s *server) start() error {
	tcpAddr, err := net.ResolveTCPAddr(s.Network, s.Address())
	if err != nil {
		return err
	}

	listener, err := net.ListenTCP(s.Network, tcpAddr)
	if err != nil {
		return err
	}

	s.Logger.Info("start listen tcp server..", zap.String("ip", s.Ip), zap.Int("port", s.Port))

	go func() {
		for {
			// block until accept a tcp connection
			conn, err := listener.AcceptTCP()
			if err != nil {
				s.Logger.Error("accept tcp connection err", zap.Error(err))
				return
			}

			s.Logger.Info("accept a tcp connection from", zap.String("address", conn.RemoteAddr().String()))

			_ = conn.SetReadBuffer(s.ReadBuffer)
			_ = conn.SetWriteBuffer(s.WriteBuffer)
		}
	}()

	return nil
}

// Stop tcp server and cancel all the connections、Goroutines and channels
func (s *server) Stop() {

}

func (s *server) Server() {
	if err := s.start(); err != nil {
		s.Logger.Error("start tcp server err", zap.Error(err))
		return
	}

	// block
	select {}
}

func (s *server) Address() string {
	return s.Ip + ":" + strconv.Itoa(s.Port)
}
