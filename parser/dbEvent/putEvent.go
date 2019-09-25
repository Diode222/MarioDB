package dbEvent

import "github.com/Diode222/MarioDB/parser/dbEventPackage/response"

type PutEvent struct {
	BasicInfo *BasicEventInfo
	Key       []byte
	Value     []byte
}

func (e *PutEvent) Process() (*response.ResponseDBEventPackage, error) {
	return nil, nil
}

func (e *PutEvent) GetBasicEventInfo() *BasicEventInfo {
	return e.BasicInfo
}
