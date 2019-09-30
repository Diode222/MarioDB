package eventInfo

import (
	"github.com/Diode222/MarioDB/dbEventPackage/response"
	"github.com/Diode222/MarioDB/event"
)

type BatchRangeEvent struct {
	BasicInfo *event.BasicEventInfo
	// Need to promise len(Starts) == len(Limits)
	Starts [][]byte
	Limits [][]byte // Range contains this key
}

func (e *BatchRangeEvent) Process() (*response.ResponseDBEventPackage, error) {
	return nil, nil
}

func (e *BatchRangeEvent) GetBasicEventInfo() *event.BasicEventInfo {
	return e.BasicInfo
}
