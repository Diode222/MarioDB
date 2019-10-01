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

type CreateEvent struct {
	BasicInfo *event.BasicEventInfo
}

func (e *CreateEvent) Process() (*response.ResponseDBEventPackage, error) {
	var db *leveldb.DB
	var ok bool
	dbName := e.GetBasicEventInfo().DBName

	// The current process can only have one reference globally, so if the db is existed, response OK directly.
	if db, ok = manager.DBManger.Get(dbName); ok {
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

	db, err := leveldb.OpenFile(utils.GetAbsoluteOfDB(dbName), nil)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Create db failed, dbname: %s", dbName))
	}
	manager.DBManger.Add(dbName, db)

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

func (e *CreateEvent) GetBasicEventInfo() *event.BasicEventInfo {
	return e.BasicInfo
}

func createEventParse(requestDBEventPackage *request.RequestDBEventPackage) (event.Event, error) {
	return &CreateEvent{BasicInfo: &event.BasicEventInfo{
		Method: event.OPEN,
		DBName: string(requestDBEventPackage.DBName),
	}}, nil
}
