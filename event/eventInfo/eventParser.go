package eventInfo

import (
	"errors"
	"fmt"
	"github.com/Diode222/MarioDB/dbEventPackage/request"
	"github.com/Diode222/MarioDB/event"
	"sync"
)

type eventParser struct{}

var parser *eventParser
var once sync.Once

func EventParser() *eventParser {
	once.Do(func() {
		parser = new(eventParser)
	})
	return parser
}

func (p *eventParser) Parse(requestDBEventPackage *request.RequestDBEventPackage) (event.Event, error) {
	method := string(requestDBEventPackage.Method)
	switch method {
	case "CREATE":
		return createEventParse(requestDBEventPackage)
	case "GET":
		return getEventParse(requestDBEventPackage)
	case "BATCHGET":
		return batchGetEventParse(requestDBEventPackage)
	case "PUT":
		return putEventParse(requestDBEventPackage)
	case "BATCHPUT":

	case "DELETE":

	case "BATCHDELETE":

	case "RANGE":

	case "BATCHRANGE":

	case "SEEKRANGE":

	case "PREFIXRANGE":

	case "BATCHPREFIXRANGE":

	}

	return nil, errors.New(fmt.Sprintf("Undefined event method: %s", method))
}
