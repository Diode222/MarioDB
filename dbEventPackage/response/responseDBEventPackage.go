package response

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"log"
)

var responseDBEventPackageHeaderLength = 10 // bytes

type ResponseDBEventPackage struct {
	Version        [2]byte
	StatusLength   uint16
	ErrorLength    uint16
	ValuesLength   uint16
	ReservedLength uint16
	Status         []byte
	Error          []byte
	Values         []byte
	Reserved       []byte
}

func (p *ResponseDBEventPackage) Pack(writer io.Writer) error {
	var err error
	err = binary.Write(writer, binary.BigEndian, &p.Version)
	err = binary.Write(writer, binary.BigEndian, &p.StatusLength)
	err = binary.Write(writer, binary.BigEndian, &p.ErrorLength)
	err = binary.Write(writer, binary.BigEndian, &p.ValuesLength)
	err = binary.Write(writer, binary.BigEndian, &p.ReservedLength)
	err = binary.Write(writer, binary.BigEndian, &p.Status)
	err = binary.Write(writer, binary.BigEndian, &p.Error)
	err = binary.Write(writer, binary.BigEndian, &p.Values)
	err = binary.Write(writer, binary.BigEndian, &p.Reserved)
	return err
}

func (p *ResponseDBEventPackage) Unpack(reader io.Reader) error {
	var err error
	err = binary.Read(reader, binary.BigEndian, &p.Version)
	err = binary.Read(reader, binary.BigEndian, &p.StatusLength)
	err = binary.Read(reader, binary.BigEndian, &p.ErrorLength)
	err = binary.Read(reader, binary.BigEndian, &p.ValuesLength)
	err = binary.Read(reader, binary.BigEndian, &p.ReservedLength)
	p.Status = make([]byte, p.StatusLength)
	err = binary.Read(reader, binary.BigEndian, &p.Status)
	p.Error = make([]byte, p.ErrorLength)
	err = binary.Read(reader, binary.BigEndian, &p.Error)
	p.Values = make([]byte, p.ValuesLength)
	err = binary.Read(reader, binary.BigEndian, &p.Values)
	p.Reserved = make([]byte, p.ReservedLength)
	err = binary.Read(reader, binary.BigEndian, &p.Reserved)
	return err
}

func (p *ResponseDBEventPackage) PackToBinary() ([]byte, error) {
	var err error
	buf := new(bytes.Buffer)

	err = binary.Write(buf, binary.BigEndian, &p.Version)
	err = binary.Write(buf, binary.BigEndian, &p.StatusLength)
	err = binary.Write(buf, binary.BigEndian, &p.ErrorLength)
	err = binary.Write(buf, binary.BigEndian, &p.ValuesLength)
	err = binary.Write(buf, binary.BigEndian, &p.ReservedLength)
	err = binary.Write(buf, binary.BigEndian, &p.Status)
	err = binary.Write(buf, binary.BigEndian, &p.Error)
	err = binary.Write(buf, binary.BigEndian, &p.Values)
	err = binary.Write(buf, binary.BigEndian, &p.Reserved)
	return buf.Bytes(), err
}

func ScannerSplit(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if !atEOF && len(data) >= 2 && data[0] == 'V' && data[1] >= '1' && data[1] <= '9' {
		if len(data) >= responseDBEventPackageHeaderLength {
			switch data[1] {
			case '1':
				return sannerSplit(data, atEOF)
			case '2':
			case '3':
			default:

			}
		}
	}

	log.Printf("Wrong server response protocol, data: %s", string(data))
	return -1, nil, errors.New(fmt.Sprintf("Wrong server protocol, data: %s", data))
}

func sannerSplit(data []byte, atEOF bool) (advance int, token []byte, err error) {
	statusLength := uint16(0)
	valuesLength := uint16(0)
	reservedLength := uint16(0)

	err = binary.Read(bytes.NewReader(data[2:4]), binary.BigEndian, &statusLength)
	err = binary.Read(bytes.NewReader(data[4:6]), binary.BigEndian, &valuesLength)
	err = binary.Read(bytes.NewReader(data[6:8]), binary.BigEndian, &reservedLength)

	totalLength := int(statusLength + valuesLength + reservedLength + uint16(responseDBEventPackageHeaderLength))
	if totalLength <= len(data) {
		return totalLength, data[:totalLength], nil
	}

	return -1, nil, errors.New(fmt.Sprintf("Wrong server protocol version 1, data: %s", data))
}
