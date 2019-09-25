package server

import (
	"github.com/Diode222/MarioDB/parser/dbEvent"
	"github.com/Diode222/MarioDB/parser/dbEventPackage/request"
	"github.com/Diode222/MarioDB/responseErrors"
	"log"
	"sync"
)

type responseDataSync struct {
	Lock sync.RWMutex
	Data []byte
}

func response(packages []*request.RequestDBEventPackage) []byte {
	responseDataSyncObj := &responseDataSync{
		Lock: sync.RWMutex{},
		Data: []byte{},
	}
	var wg sync.WaitGroup
	for _, p := range packages {
		wg.Add(1)
		go func(p *request.RequestDBEventPackage) {
			dbEventObj, err := dbEvent.DBEventParser().Parse(p)
			if err != nil {
				log.Printf("Protocol error")
				protocolErrorDataBinary := responseErrors.ProtocolError(p.Version)
				responseDataSyncObj.Lock.Lock()
				responseDataSyncObj.Data = append(responseDataSyncObj.Data, protocolErrorDataBinary...)
				responseDataSyncObj.Lock.Unlock()
				wg.Done()
				return
			}

			responseData, err := dbEventObj.Process()
			if err != nil {
				log.Printf("Protocol error")
				protocolErrorDataBinary := responseErrors.ProtocolError(p.Version)
				responseDataSyncObj.Lock.Lock()
				responseDataSyncObj.Data = append(responseDataSyncObj.Data, protocolErrorDataBinary...)
				responseDataSyncObj.Lock.Unlock()
				wg.Done()
				return
			}

			responseDataBinary, err := responseData.PackToBinary()
			if err != nil {
				log.Print(err)
				wg.Done()
				return
			}

			responseDataSyncObj.Lock.Lock()
			responseDataSyncObj.Data = append(responseDataSyncObj.Data, responseDataBinary...)
			responseDataSyncObj.Lock.Unlock()

			wg.Done()
		}(p)
	}
	wg.Wait()

	return responseDataSyncObj.Data
}
