package dbEvent

import "github.com/Diode222/MarioDB/parser/dbEvent/event"

type DBEvent interface {
	Process() ([]byte, error)

	GetBasicEventInfo() *event.BasicEventInfo
}
