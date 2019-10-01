package utils

import (
	"github.com/Diode222/MarioDB/global"
	"strings"
)

func ProtocolFieldParse(field []byte) [][]byte {
	fieldStrSlice := strings.Split(string(field), global.DB_FIELD_SEPARATOR)
	var fields [][]byte
	for _, fieldStr := range fieldStrSlice {
		fields = append(fields, []byte(fieldStr))
	}
	return fields
}
