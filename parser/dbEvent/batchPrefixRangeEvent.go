package dbEvent

import "github.com/Diode222/MarioDB/parser/dbEventPackage/response"

type BatchPrefixRangeEvent struct {
	BasicInfo *BasicEventInfo
	Prefixes  [][]byte
}

func (e *BatchPrefixRangeEvent) Process() (*response.ResponseDBEventPackage, error) {
	return nil, nil
}

func (e *BatchPrefixRangeEvent) GetBasicEventInfo() *BasicEventInfo {
	return e.BasicInfo
}
