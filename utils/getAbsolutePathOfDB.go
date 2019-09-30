package utils

import (
	"github.com/Diode222/MarioDB/global"
	"path"
)

func GetAbsoluteOfDB(dbName string) string {
	return path.Join(global.DB_ROOT_PATH, dbName)
}
