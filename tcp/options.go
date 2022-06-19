package tcp

type option func(s *server)

func WithLog() option {
	return func(s *server) {
		//s.Ip = "89"
	}
}
