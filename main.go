package main

import (
	"github.com/Diode222/MarioDB/global"
	"github.com/Diode222/MarioDB/manager"
	"github.com/Diode222/MarioDB/server"
)

func main() {
	global.DB_ROOT_PATH = "/home/diode/levelDB_database_root/"
	global.IP = "127.0.0.1"
	global.PORT = 50000
	global.DB_LRU_CACHE_MAX_COUNT = 10
	global.MAX_CLIENT_COUNT = 10

	manager.DBManger = manager.NewDBManger(global.DB_LRU_CACHE_MAX_COUNT)
	manager.ClientManger = manager.NewClientManger(global.MAX_CLIENT_COUNT)

	listener := server.NewServer(global.IP, global.PORT, false, 1)
	listener.Init()
}
