package eventInfo

import (
	"github.com/Diode222/MarioDB/dbEventPackage/response"
	"github.com/Diode222/MarioDB/event"
)

type BatchDeleteEvent struct {
	BasicInfo *event.BasicEventInfo
	Keys      [][]byte
}

func (e *BatchDeleteEvent) Process() (*response.ResponseDBEventPackage, error) {
	return nil, nil
}

func (e *BatchDeleteEvent) GetBasicEventInfo() *event.BasicEventInfo {
	return e.BasicInfo
}
