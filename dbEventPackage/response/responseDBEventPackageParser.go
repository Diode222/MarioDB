package response

import (
	"bufio"
	"bytes"
	"github.com/panjf2000/gnet/ringbuffer"
	"log"
	"sync"
)

type responseDBEventPackageParser struct{}

var parser *responseDBEventPackageParser
var once sync.Once

func ResponseDBEventPackageParser() *responseDBEventPackageParser {
	once.Do(func() {
		parser = new(responseDBEventPackageParser)
	})
	return parser
}

func (p *responseDBEventPackageParser) Parse(inBuf *ringbuffer.RingBuffer) ([]*ResponseDBEventPackage, error) {
	var err error
	packages := []*ResponseDBEventPackage{}
	scanner := bufio.NewScanner(inBuf)
	scanner.Split(ScannerSplit)
	for scanner.Scan() {
		dbEventPack := new(ResponseDBEventPackage)
		err = dbEventPack.Unpack(bytes.NewReader(scanner.Bytes()))
		if err != nil {
			head, tail := inBuf.PreReadAll()
			log.Printf("responseDBEventPackageParser parse failed, %s", string(append(head, tail...)))
			continue
		}

		packages = append(packages, dbEventPack)
	}
	if err = scanner.Err(); err != nil {
		head, tail := inBuf.PreReadAll()
		log.Printf("responseDBEventPackageParser parse failed, %s", string(append(head, tail...)))
	}
	return packages, err
}
