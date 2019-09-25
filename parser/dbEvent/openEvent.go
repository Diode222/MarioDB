package dbEvent

import (
	"github.com/Diode222/MarioDB/parser/dbEventPackage/request"
	"github.com/Diode222/MarioDB/parser/dbEventPackage/response"
)

type OpenEvent struct {
	BasicInfo *BasicEventInfo
}

func (e *OpenEvent) Process() (*response.ResponseDBEventPackage, error) {
	return nil, nil
}

func (e *OpenEvent) GetBasicEventInfo() *BasicEventInfo {
	return e.BasicInfo
}

func openEventParse(requestDBEventPackage *request.RequestDBEventPackage) (DBEvent, error) {
	return &OpenEvent{BasicInfo: &BasicEventInfo{
		Method: OPEN,
		DBName: string(requestDBEventPackage.DBName),
	}}, nil
}
