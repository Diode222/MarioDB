package eventInfo

import (
	"github.com/Diode222/MarioDB/dbEventPackage/response"
	"github.com/Diode222/MarioDB/event"
)

type PrefixRangeEvent struct {
	BasicInfo *event.BasicEventInfo
	Prefix    []byte
}

func (e *PrefixRangeEvent) Process() (*response.ResponseDBEventPackage, error) {
	return nil, nil
}

func (e *PrefixRangeEvent) GetBasicEventInfo() *event.BasicEventInfo {
	return e.BasicInfo
}
