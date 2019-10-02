package manager

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/Diode222/MarioDB/client"
	"github.com/Diode222/MarioDB/global"
	"github.com/panjf2000/gnet"
	"log"
	"sync"
)

type clientManger struct {
	clientMap sync.Map
	maxClientCount uint
}

var clientM *clientManger
var clientMangerOnce sync.Once

func NewClientManger(maxClientCount uint) *clientManger {
	clientMangerOnce.Do(func() {
		clientM = &clientManger{
			clientMap: sync.Map{},
			maxClientCount: maxClientCount,
		}
	})
	return clientM
}

func (m *clientManger) Add(c gnet.Conn) error {
	var err error
	if (m.ConnctionFull()) {
		err = errors.New(fmt.Sprintf("MAXCLIENTCOUNT. Over max client count, maxCount: %d", global.MAX_CLIENT_COUNT))
		log.Printf(err.Error())
		return err
	}
	address := c.RemoteAddr()
	newClient := &client.Client{
		Address: address.String(),
		Buffer:  bytes.NewBuffer(nil),
		Type:    client.NORMAL,
	}
	m.clientMap.Store(address.String(), newClient)
	return nil
}

func (m *clientManger) Remove(address string) {
	m.clientMap.Delete(address)
}

func (m *clientManger) Get(address string) (*client.Client, bool) {
	c, ok := m.clientMap.Load(address)
	fmt.Println("shaqingkuang: ", c)
	if !ok {
		return nil, false
	}
	return c.(*client.Client), true
}

func (m *clientManger) ConnctionFull() bool {
	var connectionCount uint = 0
	m.clientMap.Range(func(k, v interface{}) bool {
		connectionCount++
		return true
	})

	if connectionCount >= global.MAX_CLIENT_COUNT {
		return true
	} else {
		return false
	}
}