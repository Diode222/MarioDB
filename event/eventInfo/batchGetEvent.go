package eventInfo

import (
	"github.com/Diode222/MarioDB/dbEventPackage/response"
	"github.com/Diode222/MarioDB/event"
)

type BatchGetEvent struct {
	BasicInfo *event.BasicEventInfo
	Keys      [][]byte
}

func (e *BatchGetEvent) Process() (*response.ResponseDBEventPackage, error) {
	return nil, nil
}

func (e *BatchGetEvent) GetBasicEventInfo() *event.BasicEventInfo {
	return e.BasicInfo
}
