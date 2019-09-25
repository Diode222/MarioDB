package dbEvent

import "github.com/Diode222/MarioDB/parser/dbEventPackage/response"

type BatchPutEvent struct {
	BasicInfo *BasicEventInfo
	Keys      [][]byte
	Values    [][]byte
}

func (e *BatchPutEvent) Process() (*response.ResponseDBEventPackage, error) {
	return nil, nil
}

func (e *BatchPutEvent) GetBasicEventInfo() *BasicEventInfo {
	return e.BasicInfo
}
