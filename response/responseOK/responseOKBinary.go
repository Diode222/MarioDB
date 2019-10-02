package responseOK

import (
	"github.com/Diode222/MarioDB/dbEventPackage/response"
	"log"
)

func ResponseOKBinary(version [2]byte) []byte {
	var protocolOKPackage *response.ResponseDBEventPackage = &response.ResponseDBEventPackage{
		Version:        version,
		StatusLength:   2,
		ErrorLength:    0,
		ValuesLength:   0,
		ReservedLength: 0,
		Status:         []byte("OK"),
		Error:          nil,
		Values:         nil,
		Reserved:       nil,
	}

	errDataBinary, err := protocolOKPackage.PackToBinary()
	if err != nil {
		log.Print(err)
	}

	return errDataBinary
}
