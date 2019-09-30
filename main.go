package main

import (
	"github.com/Diode222/MarioDB/global"
	"github.com/Diode222/MarioDB/server"
)

func main() {
	global.DB_ROOT_PATH = "/home/diode/levelDB_database_root/"
	global.IP = "127.0.0.1"
	global.PORT = 50000
	listener := server.NewServer(global.IP, global.PORT, false, 1)
	listener.Init()
}
