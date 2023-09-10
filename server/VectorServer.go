package server

type Collection struct {
	name       string
	collection *[]float64
}

type VectorServer struct {
	ip          string
	port        int
	collections *[]Collection
}

func NewVectorServer(ip string, port int) *VectorServer {
	vectorServer := &VectorServer{
		ip:          ip,
		port:        port,
		collections: &[]Collection{},
	}
	return vectorServer
}
