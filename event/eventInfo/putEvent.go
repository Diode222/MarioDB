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

type PutEvent struct {
	BasicInfo *event.BasicEventInfo
	Key       []byte
	Value     []byte
}

func (e *PutEvent) Process() (*response.ResponseDBEventPackage, error) {
	dbName := e.BasicInfo.DBName
	var db *leveldb.DB
	var ok bool

	if db, ok = manager.DBManger.Get(dbName); !ok {
		return nil, errors.New(fmt.Sprintf("No such db, dbName: %s", dbName))
	}

	err := db.Put(e.Key, e.Value, nil)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("PUT failed to db, dbName: %s, err: %s, key: %s, values: %s", dbName, err.Error(), e.Key, e.Value))
	}

	return &response.ResponseDBEventPackage{
		Version:        [2]byte{'V', '1'},
		StatusLength:   2,
		ErrorLength:    0,
		ValuesLength:   0,
		ReservedLength: 0,
		Status:         []byte("OK"),
		Error:          nil,
		Values:         nil,
		Reserved:       nil,
	}, nil
}

func (e *PutEvent) GetBasicEventInfo() *event.BasicEventInfo {
	return e.BasicInfo
}

func putEventParse(requestDBEventPackage *request.RequestDBEventPackage) (event.Event, error) {
	return &PutEvent{
		BasicInfo: &event.BasicEventInfo{
			Method: event.GET,
			DBName: string(requestDBEventPackage.DBName),
		},
		Key:   requestDBEventPackage.Keys,
		Value: requestDBEventPackage.Values,
	}, nil
}
