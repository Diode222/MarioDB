package eventInfo

import (
	"github.com/Diode222/MarioDB/dbEventPackage/response"
	"github.com/Diode222/MarioDB/event"
)

type SeekRangeEvent struct {
	BasicInfo *event.BasicEventInfo
	Key       []byte
}

func (e *SeekRangeEvent) Process() (*response.ResponseDBEventPackage, error) {
	return nil, nil
}

func (e *SeekRangeEvent) GetBasicEventInfo() *event.BasicEventInfo {
	return e.BasicInfo
}
