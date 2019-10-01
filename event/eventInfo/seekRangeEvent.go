package eventInfo

import (
	"errors"
	"fmt"
	"github.com/Diode222/MarioDB/dbEventPackage/request"
	"github.com/Diode222/MarioDB/dbEventPackage/response"
	"github.com/Diode222/MarioDB/event"
	"github.com/Diode222/MarioDB/global"
	"github.com/Diode222/MarioDB/manager"
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/testutil"
)

type SeekRangeEvent struct {
	BasicInfo *event.BasicEventInfo
	Key       []byte
}

func (e *SeekRangeEvent) Process() (*response.ResponseDBEventPackage, error) {
	var keys []byte
	var values []byte
	dbName := e.BasicInfo.DBName
	var db *leveldb.DB
	var ok bool

	if db, ok = manager.DBManger.Get(dbName); !ok {
		return nil, errors.New(fmt.Sprintf("No such opened db, dbName: %s", dbName))
	}

	iter := db.NewIterator(nil, nil)
	for ok := iter.Seek(e.Key); ok; ok = iter.Next() {
		keys = append(keys, iter.Key()...)
		keys = append(keys, global.DB_FIELD_SEPARATOR...)
		values = append(values, iter.Value()...)
		values = append(values, global.DB_FIELD_SEPARATOR...)
	}
	if len(keys) > 0 && len(values) > 0 {
		keys = keys[:len(keys)-2]
		values = values[:len(values)-2]
	}
	iter.Release()
	err := iter.Error()
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Key range from start to limit failed, start key: %s, dbName: %s", e.Key, testutil.DBNone))
	}
	return &response.ResponseDBEventPackage{
		Version:        [2]byte{'V', '1'},
		StatusLength:   2,
		ErrorLength:    0,
		ValuesLength:   uint16(len(values)),
		ReservedLength: uint16(len(keys)),
		Status:         []byte("OK"),
		Error:          nil,
		Values:         values,
		Reserved:       keys,
	}, nil
}

func (e *SeekRangeEvent) GetBasicEventInfo() *event.BasicEventInfo {
	return e.BasicInfo
}

func seekRangeEventParse(requestDBEventPackage *request.RequestDBEventPackage) (event.Event, error) {
	return &SeekRangeEvent{
		BasicInfo: &event.BasicEventInfo{
			Method: event.SEEKRANGE,
			DBName: string(requestDBEventPackage.DBName),
		},
		Key: requestDBEventPackage.Keys,
	}, nil
}
