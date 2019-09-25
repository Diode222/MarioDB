package responseErrors

import (
	"github.com/Diode222/MarioDB/parser/dbEventPackage/response"
	"log"
)

func ProtocolError(version [2]byte) []byte {
	var protocolErrorPackage *response.ResponseDBEventPackage = &response.ResponseDBEventPackage{
		Version:        version,
		StatusLength:   5,
		ErrorLength:    0,
		ValuesLength:   0,
		ReservedLength: 0,
		Status:         []byte("Error"),
		Error:          nil,
		Values:         nil,
		Reserved:       nil,
	}

	errString := []byte("Protocol error")
	protocolErrorPackage.ErrorLength = uint16(len(errString))
	protocolErrorPackage.Error = errString
	errDataBinary, err := protocolErrorPackage.PackToBinary()
	if err != nil {
		log.Print(err)
	}

	return errDataBinary
}
