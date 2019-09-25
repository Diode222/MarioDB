package dbEvent

import "github.com/Diode222/MarioDB/parser/dbEventPackage/response"

type RangeEvent struct {
	BasicInfo *BasicEventInfo
	Start     []byte
	Limit     []byte // Range contains this key
}

func (e *RangeEvent) Process() (*response.ResponseDBEventPackage, error) {
	return nil, nil
}

func (e *RangeEvent) GetBasicEventInfo() *BasicEventInfo {
	return e.BasicInfo
}
