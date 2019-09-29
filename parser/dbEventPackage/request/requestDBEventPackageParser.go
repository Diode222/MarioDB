package request

import (
	"bufio"
	"bytes"
	"sync"
)

type requestDBEventPackageParser struct{}

var parser *requestDBEventPackageParser
var once sync.Once

func RequestDBEventPackageParser() *requestDBEventPackageParser {
	once.Do(func() {
		parser = new(requestDBEventPackageParser)
	})
	return parser
}

func (p *requestDBEventPackageParser) Parse(buffer *bytes.Buffer) ([]*RequestDBEventPackage, int, error) {
	var err error
	packages := []*RequestDBEventPackage{}
	buf := bytes.NewBuffer(buffer.Bytes())
	scanner := bufio.NewScanner(buf)
	scanner.Split(ScannerSplit)
	consumeBytesCount := 0
	for scanner.Scan() {
		dbEventPack := new(RequestDBEventPackage)
		err = dbEventPack.Unpack(bytes.NewReader(scanner.Bytes()))
		if err != nil {
			continue
		}

		consumeBytesCount += dbEventPack.TotalLength()

		packages = append(packages, dbEventPack)
	}

	return packages, consumeBytesCount, nil
}
