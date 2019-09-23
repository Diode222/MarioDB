package event

import (
	"github.com/Diode222/MarioDB/parser/dbEvent"
	"github.com/Diode222/MarioDB/parser/dbEventPackage/request"
)

type OpenEvent struct {
	BasicInfo *BasicEventInfo
}

func (e *OpenEvent) Process() ([]byte, error) {
	return nil, nil
}

func (e *OpenEvent) GetBasicEventInfo() *BasicEventInfo {
	return e.BasicInfo
}

func openEventParse(requestDBEventPackage *request.RequestDBEventPackage) (dbEvent.DBEvent, error) {
	return &OpenEvent{BasicInfo: &BasicEventInfo{
		Method: OPEN,
		DBName: string(requestDBEventPackage.DBName),
	}}, nil
}
