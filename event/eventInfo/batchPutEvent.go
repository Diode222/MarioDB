package eventInfo

import (
	"errors"
	"fmt"
	"github.com/Diode222/MarioDB/dbEventPackage/request"
	"github.com/Diode222/MarioDB/dbEventPackage/response"
	"github.com/Diode222/MarioDB/event"
	"github.com/Diode222/MarioDB/manager"
	"github.com/Diode222/MarioDB/utils"
	"github.com/syndtr/goleveldb/leveldb"
)

type BatchPutEvent struct {
	BasicInfo *event.BasicEventInfo
	Keys      [][]byte
	Values    [][]byte
}

func (e *BatchPutEvent) Process() (*response.ResponseDBEventPackage, error) {
	dbName := e.BasicInfo.DBName
	var db *leveldb.DB
	var ok bool

	if db, ok = manager.DBManger.Get(dbName); !ok {
		return nil, errors.New(fmt.Sprintf("No such db, dbName: %s", dbName))
	}

	if len(e.Keys) != len(e.Values) {
		return nil, errors.New(fmt.Sprintf("K/V pairs' length are not consistent, keys: %s, values: %s", e.Keys, e.Values))
	}

	batch := new(leveldb.Batch)
	for i := 0; i < len(e.Keys); i++ {
		batch.Put(e.Keys[i], e.Values[i])
	}
	err := db.Write(batch, nil)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Batch put failed, keys: %s, values: %s", e.Keys, e.Values))
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

func (e *BatchPutEvent) GetBasicEventInfo() *event.BasicEventInfo {
	return e.BasicInfo
}

func batchPutEventParse(requestDBEventPackage *request.RequestDBEventPackage) (event.Event, error) {
	keys := utils.ProtocolFieldParse(requestDBEventPackage.Keys)
	values := utils.ProtocolFieldParse(requestDBEventPackage.Values)
	return &BatchPutEvent{
		BasicInfo: &event.BasicEventInfo{
			Method: event.GET,
			DBName: string(requestDBEventPackage.DBName),
		},
		Keys:   keys,
		Values: values,
	}, nil
}
