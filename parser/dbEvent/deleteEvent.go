package dbEvent

import "github.com/Diode222/MarioDB/parser/dbEventPackage/response"

type DeleteEvent struct {
	BasicInfo *BasicEventInfo
	Key       []byte
}

func (e *DeleteEvent) Process() (*response.ResponseDBEventPackage, error) {
	return nil, nil
}

func (e *DeleteEvent) GetBasicEventInfo() *BasicEventInfo {
	return e.BasicInfo
}
