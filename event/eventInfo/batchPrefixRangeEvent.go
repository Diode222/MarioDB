package eventInfo

import (
	"github.com/Diode222/MarioDB/dbEventPackage/response"
	"github.com/Diode222/MarioDB/event"
)

type BatchPrefixRangeEvent struct {
	BasicInfo *event.BasicEventInfo
	Prefixes  [][]byte
}

func (e *BatchPrefixRangeEvent) Process() (*response.ResponseDBEventPackage, error) {
	return nil, nil
}

func (e *BatchPrefixRangeEvent) GetBasicEventInfo() *event.BasicEventInfo {
	return e.BasicInfo
}
