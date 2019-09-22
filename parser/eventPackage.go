package parser

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"log"
)

// header length: 38 bytes
var requestDBEventPackageHeaderLength int = 20

// all keys, values and starts..., will seperated by
type RequestDBEventPackage struct {
	Version        [2]byte // package version, now is [2]byte{'V', '1'}
	MethodLength   uint16
	DBNameLength   uint16
	KeysLength     uint16
	ValuesLength   uint16
	StartsLength   uint16
	LimitsLength   uint16
	PrefixesLength uint16
	SettingsLength uint16
	ReservedLength uint16
	Method         []byte // method name
	DBName         []byte
	Keys           []byte
	Values         []byte
	Starts         []byte
	Limits         []byte
	Prefixes       []byte
	Settings       []byte
	Reserved       []byte // reserved position
}

func (p *RequestDBEventPackage) Pack(writer io.Writer) error {
	var err error
	err = binary.Write(writer, binary.BigEndian, &p.Version)
	err = binary.Write(writer, binary.BigEndian, &p.MethodLength)
	err = binary.Write(writer, binary.BigEndian, &p.DBNameLength)
	err = binary.Write(writer, binary.BigEndian, &p.KeysLength)
	err = binary.Write(writer, binary.BigEndian, &p.ValuesLength)
	err = binary.Write(writer, binary.BigEndian, &p.StartsLength)
	err = binary.Write(writer, binary.BigEndian, &p.LimitsLength)
	err = binary.Write(writer, binary.BigEndian, &p.PrefixesLength)
	err = binary.Write(writer, binary.BigEndian, &p.SettingsLength)
	err = binary.Write(writer, binary.BigEndian, &p.ReservedLength)
	err = binary.Write(writer, binary.BigEndian, &p.Method)
	err = binary.Write(writer, binary.BigEndian, &p.DBName)
	err = binary.Write(writer, binary.BigEndian, &p.Keys)
	err = binary.Write(writer, binary.BigEndian, &p.Values)
	err = binary.Write(writer, binary.BigEndian, &p.Starts)
	err = binary.Write(writer, binary.BigEndian, &p.Limits)
	err = binary.Write(writer, binary.BigEndian, &p.Prefixes)
	err = binary.Write(writer, binary.BigEndian, &p.Settings)
	err = binary.Write(writer, binary.BigEndian, &p.Reserved)
	return err
}

func (p *RequestDBEventPackage) Unpack(reader io.Reader) error {
	var err error
	err = binary.Read(reader, binary.BigEndian, &p.Version)
	err = binary.Read(reader, binary.BigEndian, &p.MethodLength)
	err = binary.Read(reader, binary.BigEndian, &p.DBNameLength)
	err = binary.Read(reader, binary.BigEndian, &p.KeysLength)
	err = binary.Read(reader, binary.BigEndian, &p.ValuesLength)
	err = binary.Read(reader, binary.BigEndian, &p.StartsLength)
	err = binary.Read(reader, binary.BigEndian, &p.LimitsLength)
	err = binary.Read(reader, binary.BigEndian, &p.PrefixesLength)
	err = binary.Read(reader, binary.BigEndian, &p.SettingsLength)
	err = binary.Read(reader, binary.BigEndian, &p.ReservedLength)
	p.Method = make([]byte, p.MethodLength)
	err = binary.Read(reader, binary.BigEndian, &p.Method)
	p.DBName = make([]byte, p.DBNameLength)
	err = binary.Read(reader, binary.BigEndian, &p.DBName)
	p.Keys = make([]byte, p.KeysLength)
	err = binary.Read(reader, binary.BigEndian, &p.Keys)
	p.Values = make([]byte, p.ValuesLength)
	err = binary.Read(reader, binary.BigEndian, &p.Values)
	p.Starts = make([]byte, p.StartsLength)
	err = binary.Read(reader, binary.BigEndian, &p.Starts)
	p.Limits = make([]byte, p.LimitsLength)
	err = binary.Read(reader, binary.BigEndian, &p.Limits)
	p.Prefixes = make([]byte, p.PrefixesLength)
	err = binary.Read(reader, binary.BigEndian, &p.Prefixes)
	p.Settings = make([]byte, p.SettingsLength)
	err = binary.Read(reader, binary.BigEndian, &p.Settings)
	p.Reserved = make([]byte, p.ReservedLength)
	err = binary.Read(reader, binary.BigEndian, &p.Reserved)
	return err
}

func ScannerSplit(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if !atEOF && len(data) >= 2 && data[0] == 'V' && data[1] >= '1' && data[1] <= '9' {
		if len(data) >= requestDBEventPackageHeaderLength {
			switch data[1] {
			case 1: // version 1
				return scannerSplitVersion1(data, atEOF)
			case 2:

			case 3:

			default: // default version 1
				return scannerSplitVersion1(data, atEOF)
			}
		}
	}
	log.Printf("Wrong client request protocol, data: %s", string(data))
	return -1, nil, errors.New(fmt.Sprintf("Wrong protocol, data: %s", data))
}

func scannerSplitVersion1(data []byte, atEOF bool) (advance int, token []byte, err error) {
	methodLength := uint16(0)
	DBNameLength := uint16(0)
	KeysLength := uint16(0)
	ValuesLength := uint16(0)
	StartsLength := uint16(0)
	LimitsLength := uint16(0)
	PrefixesLength := uint16(0)
	SettingsLength := uint16(0)
	ReservedLength := uint16(0)
	err = binary.Read(bytes.NewReader(data[2:4]), binary.BigEndian, &methodLength)
	err = binary.Read(bytes.NewReader(data[4:6]), binary.BigEndian, &DBNameLength)
	err = binary.Read(bytes.NewReader(data[6:8]), binary.BigEndian, &KeysLength)
	err = binary.Read(bytes.NewReader(data[8:10]), binary.BigEndian, &ValuesLength)
	err = binary.Read(bytes.NewReader(data[10:12]), binary.BigEndian, &StartsLength)
	err = binary.Read(bytes.NewReader(data[12:14]), binary.BigEndian, &LimitsLength)
	err = binary.Read(bytes.NewReader(data[14:16]), binary.BigEndian, &PrefixesLength)
	err = binary.Read(bytes.NewReader(data[16:18]), binary.BigEndian, &SettingsLength)
	err = binary.Read(bytes.NewReader(data[18:20]), binary.BigEndian, &ReservedLength)

	totalLength := int(methodLength + DBNameLength + KeysLength + ValuesLength + StartsLength + LimitsLength + PrefixesLength + SettingsLength + ReservedLength + uint16(requestDBEventPackageHeaderLength))
	if totalLength <= len(data) {
		return totalLength, data[:totalLength], nil
	}

	return -1, nil, errors.New(fmt.Sprintf("Wrong protocol version 1, data: %s", data))
}
