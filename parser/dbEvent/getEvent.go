package dbEvent

import "github.com/Diode222/MarioDB/parser/dbEventPackage/response"

type GetEvent struct {
	BasicInfo *BasicEventInfo
	Key       []byte
}

func (e *GetEvent) Process() (*response.ResponseDBEventPackage, error) {
	return nil, nil
}

func (e *GetEvent) GetBasicEventInfo() *BasicEventInfo {
	return e.BasicInfo
}
