package eventInfo

import (
	"github.com/Diode222/MarioDB/dbEventPackage/response"
	"github.com/Diode222/MarioDB/event"
)

type DeleteEvent struct {
	BasicInfo *event.BasicEventInfo
	Key       []byte
}

func (e *DeleteEvent) Process() (*response.ResponseDBEventPackage, error) {
	return nil, nil
}

func (e *DeleteEvent) GetBasicEventInfo() *event.BasicEventInfo {
	return e.BasicInfo
}
