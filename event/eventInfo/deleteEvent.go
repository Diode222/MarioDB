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

type DeleteEvent struct {
	BasicInfo *event.BasicEventInfo
	Key       []byte
}

func (e *DeleteEvent) Process() (*response.ResponseDBEventPackage, error) {
	dbName := e.BasicInfo.DBName
	var db *leveldb.DB
	var ok bool

	if db, ok = manager.DBManger.Get(dbName); !ok {
		return nil, errors.New(fmt.Sprintf("No such opened db, dbName: %s", dbName))
	}

	// levelDB delete will not return err when the key is not existed
	err := db.Delete(e.Key, nil)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Delete failed, key: %s, dbName: %s", e.Key, dbName))
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

func (e *DeleteEvent) GetBasicEventInfo() *event.BasicEventInfo {
	return e.BasicInfo
}

func deleteEventParse(requestDBEventPackage *request.RequestDBEventPackage) (event.Event, error) {
	return &DeleteEvent{
		BasicInfo: &event.BasicEventInfo{
			Method: event.DELETE,
			DBName: string(requestDBEventPackage.DBName),
		},
		Key: requestDBEventPackage.Keys,
	}, nil
}
