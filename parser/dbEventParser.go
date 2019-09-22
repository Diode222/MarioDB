package parser

import (
	"bufio"
	"bytes"
	"github.com/panjf2000/gnet/ringbuffer"
	"log"
	"sync"
)

type dbEventParser struct{}

var parser *dbEventParser
var once sync.Once

func DBEventParser() *dbEventParser {
	once.Do(func() {
		parser = new(dbEventParser)
	})
	return parser
}

func (p *dbEventParser) Parse(inBuf *ringbuffer.RingBuffer) ([]*RequestDBEventPackage, error) {
	var err error
	packages := []*RequestDBEventPackage{}
	scanner := bufio.NewScanner(inBuf)
	scanner.Split(scannerSplitVersion1)
	for scanner.Scan() {
		dbEventPack := new(RequestDBEventPackage)
		err = dbEventPack.Unpack(bytes.NewReader(scanner.Bytes()))
		if err != nil {
			head, tail := inBuf.PreReadAll()
			log.Printf("DBEventPackage parse failed, %s", string(append(head, tail...)))
			continue
		}
		//fmt.Println(string(dbEventPack.DBName))
		packages = append(packages, dbEventPack)
	}
	if err = scanner.Err(); err != nil {
		head, tail := inBuf.PreReadAll()
		log.Printf("DBEventPackage parse failed, %s", string(append(head, tail...)))
	}
	return packages, err
}
