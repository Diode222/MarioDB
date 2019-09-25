package dbEvent

import "github.com/Diode222/MarioDB/parser/dbEventPackage/response"

type PrefixRangeEvent struct {
	BasicInfo *BasicEventInfo
	Prefix    []byte
}

func (e *PrefixRangeEvent) Process() (*response.ResponseDBEventPackage, error) {
	return nil, nil
}

func (e *PrefixRangeEvent) GetBasicEventInfo() *BasicEventInfo {
	return e.BasicInfo
}
