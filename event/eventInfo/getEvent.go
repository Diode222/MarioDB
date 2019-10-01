package eventInfo

import (
	"errors"
	"fmt"
	"github.com/Diode222/MarioDB/dbEventPackage/request"
	"github.com/Diode222/MarioDB/dbEventPackage/response"
	"github.com/Diode222/MarioDB/event"
	"github.com/Diode222/MarioDB/manager"
	"github.com/syndtr/goleveldb/leveldb"
)

type GetEvent struct {
	BasicInfo *event.BasicEventInfo
	Key       []byte
}

func (e *GetEvent) Process() (*response.ResponseDBEventPackage, error) {
	dbName := e.BasicInfo.DBName
	var db *leveldb.DB
	var ok bool

	if db, ok = manager.DBManger.Get(dbName); !ok {
		return nil, errors.New(fmt.Sprintf("No such opened db, dbName: %s", dbName))
	}

	value, err := db.Get(e.Key, nil)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("No such key, key: %s, dbName: %s", string(e.Key), dbName))
	}

	return &response.ResponseDBEventPackage{
		Version:        [2]byte{'V', '1'},
		StatusLength:   2,
		ErrorLength:    0,
		ValuesLength:   uint16(len(value)),
		ReservedLength: 0,
		Status:         []byte("OK"),
		Error:          nil,
		Values:         value,
		Reserved:       nil,
	}, nil
}

func (e *GetEvent) GetBasicEventInfo() *event.BasicEventInfo {
	return e.BasicInfo
}

func getEventParse(requestDBEventPackage *request.RequestDBEventPackage) (event.Event, error) {
	return &GetEvent{
		BasicInfo: &event.BasicEventInfo{
			Method: event.GET,
			DBName: string(requestDBEventPackage.DBName),
		},
		Key: requestDBEventPackage.Keys,
	}, nil
}
