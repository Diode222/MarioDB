package eventInfo

import (
	"github.com/Diode222/MarioDB/dbEventPackage/request"
	"github.com/Diode222/MarioDB/dbEventPackage/response"
	"github.com/Diode222/MarioDB/event"
)

type OpenEvent struct {
	BasicInfo *event.BasicEventInfo
}

func (e *OpenEvent) Process() (*response.ResponseDBEventPackage, error) {
	return nil, nil
}

func (e *OpenEvent) GetBasicEventInfo() *event.BasicEventInfo {
	return e.BasicInfo
}

func openEventParse(requestDBEventPackage *request.RequestDBEventPackage) (event.Event, error) {
	return &OpenEvent{BasicInfo: &event.BasicEventInfo{
		Method: event.OPEN,
		DBName: string(requestDBEventPackage.DBName),
	}}, nil
}
