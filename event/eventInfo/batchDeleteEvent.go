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

type BatchDeleteEvent struct {
	BasicInfo *event.BasicEventInfo
	Keys      [][]byte
}

func (e *BatchDeleteEvent) Process() (*response.ResponseDBEventPackage, error) {
	dbName := e.BasicInfo.DBName
	var db *leveldb.DB
	var ok bool

	if db, ok = manager.DBManger.Get(dbName); !ok {
		return nil, errors.New(fmt.Sprintf("No such opened db, dbName: %s", dbName))
	}

	batch := new(leveldb.Batch)
	for i := 0; i < len(e.Keys); i++ {
		batch.Delete(e.Keys[i])
	}
	err := db.Write(batch, nil)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Batch delete failed, keys: %s, dbName: %s", e.Keys, dbName))
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

func (e *BatchDeleteEvent) GetBasicEventInfo() *event.BasicEventInfo {
	return e.BasicInfo
}

func batchDeleteEventParse(requestDBEventPackage *request.RequestDBEventPackage) (event.Event, error) {
	keys := utils.ProtocolFieldParse(requestDBEventPackage.Keys)
	return &BatchDeleteEvent{
		BasicInfo: &event.BasicEventInfo{
			Method: event.BATCHDELETE,
			DBName: string(requestDBEventPackage.DBName),
		},
		Keys: keys,
	}, nil
}
