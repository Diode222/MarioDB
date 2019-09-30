package eventInfo

import (
	"github.com/Diode222/MarioDB/dbEventPackage/response"
	"github.com/Diode222/MarioDB/event"
)

type BatchPutEvent struct {
	BasicInfo *event.BasicEventInfo
	Keys      [][]byte
	Values    [][]byte
}

func (e *BatchPutEvent) Process() (*response.ResponseDBEventPackage, error) {
	return nil, nil
}

func (e *BatchPutEvent) GetBasicEventInfo() *event.BasicEventInfo {
	return e.BasicInfo
}
