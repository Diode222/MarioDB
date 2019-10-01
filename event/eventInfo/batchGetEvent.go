package eventInfo

import (
	"errors"
	"fmt"
	"github.com/Diode222/MarioDB/dbEventPackage/request"
	"github.com/Diode222/MarioDB/dbEventPackage/response"
	"github.com/Diode222/MarioDB/event"
	"github.com/Diode222/MarioDB/global"
	"github.com/Diode222/MarioDB/manager"
	"github.com/Diode222/MarioDB/utils"
	"github.com/syndtr/goleveldb/leveldb"
)

type BatchGetEvent struct {
	BasicInfo *event.BasicEventInfo
	Keys      [][]byte
}

// BatchGetEvent process() will only response keys not founded error when all keys are not exists
func (e *BatchGetEvent) Process() (*response.ResponseDBEventPackage, error) {
	values := []byte{}
	dbName := e.BasicInfo.DBName
	var db *leveldb.DB
	var ok bool

	if db, ok = manager.DBManger.Get(dbName); !ok {
		return nil, errors.New(fmt.Sprintf("No such db, dbName: %s", dbName))
	}

	ok = false
	for index, key := range e.Keys {
		value := []byte{}
		var err error
		if index != len(e.Keys)-1 {
			value, err = db.Get(key, nil)
			value = append(value, []byte(global.DB_FIELD_SEPARATOR)...)
		} else {
			value, err = db.Get(key, nil)
		}
		if err == nil {
			ok = true
		}
		values = append(values, value...)
	}

	if !ok {
		return nil, errors.New(fmt.Sprintf("No such keys in db, keys: %s, dbName: %s", e.Keys, dbName))
	}

	return &response.ResponseDBEventPackage{
		Version:        [2]byte{'V', '1'},
		StatusLength:   2,
		ErrorLength:    0,
		ValuesLength:   uint16(len(values)),
		ReservedLength: 0,
		Status:         []byte("OK"),
		Error:          nil,
		Values:         values,
		Reserved:       nil,
	}, nil
}

func (e *BatchGetEvent) GetBasicEventInfo() *event.BasicEventInfo {
	return e.BasicInfo
}

func batchGetEventParse(requestDBEventPackage *request.RequestDBEventPackage) (event.Event, error) {
	keys := utils.ProtocolFieldParse(requestDBEventPackage.Keys)
	return &BatchGetEvent{
		BasicInfo: &event.BasicEventInfo{
			Method: event.BATCHGET,
			DBName: string(requestDBEventPackage.DBName),
		},
		Keys: keys,
	}, nil
}
