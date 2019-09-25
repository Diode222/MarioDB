package dbEvent

import "github.com/Diode222/MarioDB/parser/dbEventPackage/response"

type SeekRangeEvent struct {
	BasicInfo *BasicEventInfo
	Key       []byte
}

func (e *SeekRangeEvent) Process() (*response.ResponseDBEventPackage, error) {
	return nil, nil
}

func (e *SeekRangeEvent) GetBasicEventInfo() *BasicEventInfo {
	return e.BasicInfo
}
