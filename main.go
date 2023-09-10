package main

import (
	"VectronixDB/server"
	"fmt"
)

// System main function
func main() {
	// TODO: Handle command line arguments
	s := server.NewServer("", 8080)
	fmt.Println(s)
}
