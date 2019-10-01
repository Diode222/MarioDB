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

type PrefixRangeEvent struct {
	BasicInfo *event.BasicEventInfo
	Prefix    []byte
}

func (e *PrefixRangeEvent) Process() (*response.ResponseDBEventPackage, error) {
	var keys []byte
	var values []byte
	dbName := e.BasicInfo.DBName
	var db *leveldb.DB
	var ok bool

	if db, ok = manager.DBManger.Get(dbName); !ok {
		return nil, errors.New(fmt.Sprintf("No such opened db, dbName: %s", dbName))
	}

	iter := db.NewIterator(util.BytesPrefix(e.Prefix), nil)
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
		return nil, errors.New(fmt.Sprintf("Prefix range from start to limit failed, prefix: %s, dbName: %s", e.Prefix, testutil.DBNone))
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

func (e *PrefixRangeEvent) GetBasicEventInfo() *event.BasicEventInfo {
	return e.BasicInfo
}

func prefixRangeEventParse(requestDBEventPackage *request.RequestDBEventPackage) (event.Event, error) {
	var prefix []byte
	if len(requestDBEventPackage.Prefix) > 0 {
		prefix = requestDBEventPackage.Prefix
	}
	return &PrefixRangeEvent{
		BasicInfo: &event.BasicEventInfo{
			Method: event.PREFIXRANGE,
			DBName: string(requestDBEventPackage.DBName),
		},
		Prefix: prefix,
	}, nil
}
