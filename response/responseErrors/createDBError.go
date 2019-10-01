package responseErrors

import (
	"github.com/Diode222/MarioDB/dbEventPackage/response"
)

func CreateDBError(version [2]byte) *response.ResponseDBEventPackage {
	var createDBErrorPackage *response.ResponseDBEventPackage = &response.ResponseDBEventPackage{
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

	return createDBErrorPackage
}
