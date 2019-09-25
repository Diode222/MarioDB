package dbEvent

import "github.com/Diode222/MarioDB/parser/dbEventPackage/response"

type BatchGetEvent struct {
	BasicInfo *BasicEventInfo
	Keys      [][]byte
}

func (e *BatchGetEvent) Process() (*response.ResponseDBEventPackage, error) {
	return nil, nil
}

func (e *BatchGetEvent) GetBasicEventInfo() *BasicEventInfo {
	return e.BasicInfo
}
