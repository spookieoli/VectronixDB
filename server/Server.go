package server

type Server struct {
	ip          string
	port        int
	httpHandler *HttpHandler
	vectorServe *VectorServer
}

func NewServer(ip string, port int) *Server {
	server := &Server{
		ip:          ip,
		port:        port,
		httpHandler: NewHttpHandler(port),
		vectorServe: NewVectorServer(ip, port),
	}
	return server
}
