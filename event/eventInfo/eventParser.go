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
		return batchPutEventParse(requestDBEventPackage)
	case "DELETE":
		return deleteEventParse(requestDBEventPackage)
	case "BATCHDELETE":
		return batchDeleteEventParse(requestDBEventPackage)
	case "RANGE":
		return rangeEventParse(requestDBEventPackage)
	case "SEEKRANGE":
		return seekRangeEventParse(requestDBEventPackage)
	case "PREFIXRANGE":
		return prefixRangeEventParse(requestDBEventPackage)
	}

	return nil, errors.New(fmt.Sprintf("Undefined event method: %s", method))
}
