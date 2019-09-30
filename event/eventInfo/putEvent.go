package eventInfo

import (
	"github.com/Diode222/MarioDB/dbEventPackage/response"
	"github.com/Diode222/MarioDB/event"
)

type PutEvent struct {
	BasicInfo *event.BasicEventInfo
	Key       []byte
	Value     []byte
}

func (e *PutEvent) Process() (*response.ResponseDBEventPackage, error) {
	return nil, nil
}

func (e *PutEvent) GetBasicEventInfo() *event.BasicEventInfo {
	return e.BasicInfo
}
