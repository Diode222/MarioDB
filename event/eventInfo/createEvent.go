package eventInfo

import (
	"errors"
	"fmt"
	"github.com/Diode222/MarioDB/dbEventPackage/request"
	"github.com/Diode222/MarioDB/dbEventPackage/response"
	"github.com/Diode222/MarioDB/event"
	"github.com/Diode222/MarioDB/utils"
	"github.com/syndtr/goleveldb/leveldb"
	"log"
)

type CreateEvent struct {
	BasicInfo *event.BasicEventInfo
}

func (e *CreateEvent) Process() (*response.ResponseDBEventPackage, error) {
	dbName := e.GetBasicEventInfo().DBName
	db, err := leveldb.OpenFile(utils.GetAbsoluteOfDB(dbName), nil)
	if err != nil {
		log.Printf("DB %s create failed.", e.GetBasicEventInfo().DBName)
		return nil, errors.New(fmt.Sprintf("Open db failed, dbname: %s", dbName))
	}
	defer db.Close()

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
