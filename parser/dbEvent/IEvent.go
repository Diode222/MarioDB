package dbEvent

import "github.com/Diode222/MarioDB/parser/dbEventPackage/response"

type DBEvent interface {
	Process() (*response.ResponseDBEventPackage, error)

	GetBasicEventInfo() *BasicEventInfo
}
