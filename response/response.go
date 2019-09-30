package response

import (
	"github.com/Diode222/MarioDB/dbEventPackage/request"
	"github.com/Diode222/MarioDB/event/eventInfo"
	"github.com/Diode222/MarioDB/response/responseErrors"
	"log"
	"sync"
)

type responseDataSync struct {
	Lock sync.RWMutex
	Data []byte
}

func Response(packages []*request.RequestDBEventPackage) []byte {
	responseDataSyncObj := &responseDataSync{
		Lock: sync.RWMutex{},
		Data: []byte{},
	}
	var wg sync.WaitGroup
	for _, p := range packages {
		wg.Add(1)
		go func(p *request.RequestDBEventPackage) {
			defer wg.Done()
			eventObj, err := eventInfo.EventParser().Parse(p)
			if err != nil {
				log.Printf(err.Error())
				protocolErrorDataBinary := responseErrors.ResponseErrorBinary(p.Version, []byte(err.Error()))
				responseDataSyncObj.Lock.Lock()
				responseDataSyncObj.Data = append(responseDataSyncObj.Data, protocolErrorDataBinary...)
				responseDataSyncObj.Lock.Unlock()
				return
			}

			responsePackage, err := eventObj.Process()
			if err != nil {
				log.Printf(err.Error())
				protocolErrorDataBinary := responseErrors.ResponseErrorBinary(p.Version, []byte(err.Error()))
				responseDataSyncObj.Lock.Lock()
				responseDataSyncObj.Data = append(responseDataSyncObj.Data, protocolErrorDataBinary...)
				responseDataSyncObj.Lock.Unlock()
				return
			}

			responseDataBinary, err := responsePackage.PackToBinary()
			if err != nil {
				log.Print(err)
				wg.Done()
				return
			}

			responseDataSyncObj.Lock.Lock()
			responseDataSyncObj.Data = append(responseDataSyncObj.Data, responseDataBinary...)
			responseDataSyncObj.Lock.Unlock()
		}(p)
	}
	wg.Wait()

	return responseDataSyncObj.Data
}
