package request

import (
	"bufio"
	"bytes"
	"github.com/panjf2000/gnet/ringbuffer"
	"log"
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

func (p *requestDBEventPackageParser) Parse(inBuf *ringbuffer.RingBuffer) ([]*RequestDBEventPackage, error) {
	var err error
	packages := []*RequestDBEventPackage{}
	scanner := bufio.NewScanner(inBuf)
	scanner.Split(ScannerSplit)
	for scanner.Scan() {
		dbEventPack := new(RequestDBEventPackage)
		err = dbEventPack.Unpack(bytes.NewReader(scanner.Bytes()))
		if err != nil {
			head, tail := inBuf.PreReadAll()
			log.Printf("requestDBEventPackageParser parse failed, %s", string(append(head, tail...)))
			continue
		}

		packages = append(packages, dbEventPack)
	}

	return packages, err
}
