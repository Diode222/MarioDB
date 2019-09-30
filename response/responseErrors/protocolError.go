package responseErrors

import (
	"github.com/Diode222/MarioDB/dbEventPackage/response"
	"log"
)

func ResponseErrorBinary(version [2]byte, errBytes []byte) []byte {
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

	protocolErrorPackage.ErrorLength = uint16(len(errBytes))
	protocolErrorPackage.Error = errBytes
	errDataBinary, err := protocolErrorPackage.PackToBinary()
	if err != nil {
		log.Print(err)
	}

	return errDataBinary
}
