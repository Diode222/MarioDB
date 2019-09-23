package event

import (
	"errors"
	"fmt"
	"github.com/Diode222/MarioDB/parser/dbEvent"
	"github.com/Diode222/MarioDB/parser/dbEventPackage/request"
	"sync"
)

type dbEventParser struct {}

var parser *dbEventParser
var once sync.Once

func DBEventParser() *dbEventParser {
	once.Do(func() {
		parser = new(dbEventParser)
	})
	return parser
}

func (p *dbEventParser) Parse(requestDBEventPackage *request.RequestDBEventPackage) (dbEvent.DBEvent, error) {
	method := string(requestDBEventPackage.Method)
	switch method {
	case "OPEN":
		return openEventParse(requestDBEventPackage)
	case "GET":

	case "BATCHGET":

	case "PUT":

	case "BATCHPUT":

	case "DELETE":

	case "BATCHDELETE":

	case "RANGE":

	case "BATCHRANGE":

	case "SEEKRANGE":

	case "PREFIXRANGE":

	case "BATCHPREFIXRANGE":

	default:
		return nil, errors.New(fmt.Sprintf("Undefined event method: %s", method))
	}
}
