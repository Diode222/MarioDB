package main

import "github.com/Diode222/MarioDB/server"

func main() {
	listener := server.NewServer("127.0.0.1", 50000, false, 1)
	listener.Init()
}
