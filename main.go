package main

import (
	"flag"
	"fmt"
	"github.com/Diode222/MarioDB/global"
	"github.com/Diode222/MarioDB/manager"
	"github.com/Diode222/MarioDB/server"
	"os"
)

func main() {
	//global.DB_ROOT_PATH = "/home/diode/levelDB_database_root/"
	defaultDBPath := os.Getenv("HOME")
	dbPath := flag.String("dbPath", defaultDBPath, "Database's data path")
	ip := flag.String("ip", "127.0.0.1", "IP")
	port := flag.Uint("port", 50000, "Port")
	dbLRUCacheMaxCount := flag.Uint("dbLRUMax", 100, "Max count of db caches in LRU")
	maxClientCount := flag.Uint("maxClient", 100, "Max count of connected clients")
	needHelp := flag.Bool("h", false, "Help")
	flag.Parse()

	if *needHelp {
		printHelpInfo(defaultDBPath)
		return
	}

	global.DB_ROOT_PATH = *dbPath
	global.IP = *ip
	global.PORT = *port
	global.DB_LRU_CACHE_MAX_COUNT = *dbLRUCacheMaxCount
	global.MAX_CLIENT_COUNT = *maxClientCount

	manager.DBManger = manager.NewDBManger(global.DB_LRU_CACHE_MAX_COUNT)
	manager.ClientManger = manager.NewClientManger(global.MAX_CLIENT_COUNT)

	listener := server.NewServer(global.IP, global.PORT, false, 1)
	listener.Init()
}

func printHelpInfo(defaultDBPath string) {
	fmt.Println("Usage:")
	fmt.Println("    -ip           string    IP                              (Default: 127.0.0.1)")
	fmt.Println("    -port         string    Port                            (Default: 50000)")
	fmt.Println("    -dbPath       string    Database's data path            (Default: " + defaultDBPath + ")")
	fmt.Println("    -maxClient    string    Max count of connected clients  (Default: 100)")
	fmt.Println("    -dbLRUMax     string    Max count of db caches in LRU   (Default: 100)")
	fmt.Println()
}
