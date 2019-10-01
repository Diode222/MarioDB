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
	"github.com/syndtr/goleveldb/leveldb/util"
)

type RangeEvent struct {
	BasicInfo *event.BasicEventInfo
	Start     []byte
	Limit     []byte // Range contains this key
}

func (e *RangeEvent) Process() (*response.ResponseDBEventPackage, error) {
	dbName := e.BasicInfo.DBName
	var db *leveldb.DB
	var ok bool

	if db, ok = manager.DBManger.Get(dbName); !ok {
		return nil, errors.New(fmt.Sprintf("No such opened db, dbName: %s", dbName))
	}

	return e.rangeFromStartToLimitProcess(db)
}

func (e *RangeEvent) GetBasicEventInfo() *event.BasicEventInfo {
	return e.BasicInfo
}

func rangeEventParse(requestDBEventPackage *request.RequestDBEventPackage) (event.Event, error) {
	var start []byte
	var limit []byte
	if len(requestDBEventPackage.Start) > 0 {
		start = requestDBEventPackage.Start
	}
	if len(requestDBEventPackage.Limit) > 0 {
		limit = requestDBEventPackage.Limit
	}
	return &RangeEvent{
		BasicInfo: &event.BasicEventInfo{
			Method: event.RANGE,
			DBName: string(requestDBEventPackage.DBName),
		},
		Start: start,
		Limit: limit,
	}, nil
}

func (e *RangeEvent) rangeFromStartToLimitProcess(db *leveldb.DB) (*response.ResponseDBEventPackage, error) {
	var keys []byte
	var values []byte
	iter := db.NewIterator(&util.Range{Start: e.Start, Limit: e.Limit}, nil)
	for iter.Next() {
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
		return nil, errors.New(fmt.Sprintf("Range from start to limit failed, start: %s, limit: %s, dbName: %s", e.Start, e.Limit, testutil.DBNone))
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
