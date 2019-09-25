package dbEvent

import "github.com/Diode222/MarioDB/parser/dbEventPackage/response"

type BatchDeleteEvent struct {
	BasicInfo *BasicEventInfo
	Keys      [][]byte
}

func (e *BatchDeleteEvent) Process() (*response.ResponseDBEventPackage, error) {
	return nil, nil
}

func (e *BatchDeleteEvent) GetBasicEventInfo() *BasicEventInfo {
	return e.BasicInfo
}
