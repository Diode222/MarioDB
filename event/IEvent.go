package event

import (
	"github.com/Diode222/MarioDB/dbEventPackage/response"
)

type Event interface {
	Process() (*response.ResponseDBEventPackage, error)

	GetBasicEventInfo() *BasicEventInfo
}
