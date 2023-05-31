package server

type Server struct {
	Hostname string
	Port     int
}

func NewServer(host string, port int) *Server {
	server := Server{}
	server.Hostname = host
	server.Port = port
	return &server
}

func (server *Server) Run() error {
	return nil
}
