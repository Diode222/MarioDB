package dbEvent

import "github.com/Diode222/MarioDB/parser/dbEventPackage/response"

type BatchRangeEvent struct {
	BasicInfo *BasicEventInfo
	// Need to promise len(Starts) == len(Limits)
	Starts [][]byte
	Limits [][]byte // Range contains this key
}

func (e *BatchRangeEvent) Process() (*response.ResponseDBEventPackage, error) {
	return nil, nil
}

func (e *BatchRangeEvent) GetBasicEventInfo() *BasicEventInfo {
	return e.BasicInfo
}
