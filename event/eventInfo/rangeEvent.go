package eventInfo

import (
	"github.com/Diode222/MarioDB/dbEventPackage/response"
	"github.com/Diode222/MarioDB/event"
)

type RangeEvent struct {
	BasicInfo *event.BasicEventInfo
	Start     []byte
	Limit     []byte // Range contains this key
}

func (e *RangeEvent) Process() (*response.ResponseDBEventPackage, error) {
	return nil, nil
}

func (e *RangeEvent) GetBasicEventInfo() *event.BasicEventInfo {
	return e.BasicInfo
}
