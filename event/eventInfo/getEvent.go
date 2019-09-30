package eventInfo

import (
	"github.com/Diode222/MarioDB/dbEventPackage/response"
	"github.com/Diode222/MarioDB/event"
)

type GetEvent struct {
	BasicInfo *event.BasicEventInfo
	Key       []byte
}

func (e *GetEvent) Process() (*response.ResponseDBEventPackage, error) {
	return &response.ResponseDBEventPackage{
		Version:        [2]byte{'V', '1'},
		StatusLength:   2,
		ErrorLength:    0,
		ValuesLength:   10,
		ReservedLength: 0,
		Status:         []byte("OK"),
		Error:          nil,
		Values:         []byte("masiwei##h"),
		Reserved:       nil,
	}, nil
}

func (e *GetEvent) GetBasicEventInfo() *event.BasicEventInfo {
	return e.BasicInfo
}
