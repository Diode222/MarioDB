package utils

import "strings"

var SEPARATOR string = "##"

func ProtocolFieldParse(field string) []string {
	return strings.Split(field, SEPARATOR)
}
