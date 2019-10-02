package server

import (
	"bytes"
	"flag"
	"fmt"
	"github.com/Diode222/MarioDB/dbEventPackage/request"
	"github.com/Diode222/MarioDB/manager"
	"github.com/Diode222/MarioDB/response"
	"github.com/Diode222/MarioDB/response/responseErrors"
	"github.com/Diode222/MarioDB/response/responseOK"
	"github.com/panjf2000/gnet"
	"github.com/panjf2000/gnet/ringbuffer"
	"io"
	"log"
	"sync"
	"time"
)

// TODO This buffer should be managed by a map, key is client, value is buffer, when client closed, buffer should be removed.
var buffer *bytes.Buffer = bytes.NewBuffer(nil)

type server struct {
	IP        string
	Port      uint
	ReusePort bool
	Loops     uint
}

var l *server
var once sync.Once

func NewServer(ip string, port uint, reusePort bool, loops uint) *server {
	once.Do(func() {
		l = &server{
			IP:        ip,
			Port:      port,
			ReusePort: reusePort,
			Loops:     loops,
		}
	})
	return l
}

func (l *server) Init() {
	var ip string
	var port uint
	var reusePort bool
	var loops uint
	var transportLayerProtocol string = "tcp"

	flag.StringVar(&ip, "ip", l.IP, "NewServer ip")
	flag.UintVar(&port, "port", l.Port, "NewServer port")
	flag.BoolVar(&reusePort, "reusePort", l.ReusePort, "Reuse server port in cluster")
	flag.UintVar(&loops, "loops", l.Loops, "Loops number the server is using")
	flag.Parse()

	var tcpEventsServer gnet.Events
	tcpEventsServer.NumLoops = int(loops)

	tcpEventsServer.OnInitComplete = func(srv gnet.Server) (action gnet.Action) {
		log.Printf("MarioDB server started on tcp://%s.", srv.Addrs)
		return
	}

	tcpEventsServer.OnOpened = func(c gnet.Conn) (out []byte, opts gnet.Options, action gnet.Action) {
		log.Printf("Client started on tcp://%s.", c.RemoteAddr())
		err := manager.ClientManger.Add(c)
		if err != nil {
			action = gnet.Close
			out = responseErrors.ResponseErrorBinary([2]byte{'V', '1'}, []byte(err.Error()))
		} else {
			out = responseOK.ResponseOKBinary([2]byte{'V', '1'})
		}
		return
	}

	tcpEventsServer.OnClosed = func(c gnet.Conn, err error) (action gnet.Action) {
		log.Printf("Client(Address: %s) closed connection.", c.RemoteAddr())
		return
	}

	tcpEventsServer.OnDetached = func(c gnet.Conn, rwc io.ReadWriteCloser) (action gnet.Action) {
		log.Printf("NewServer detached connection of client(Address: %s).", c.RemoteAddr())
		_, err := rwc.Write([]byte(fmt.Sprintf("NewServer(Address: %s) detached connection.", c.LocalAddr())))
		if err != nil {
			log.Printf("NewServer detached info send failed. Client address: %s", c.RemoteAddr())
		}
		err = rwc.Close()
		if err != nil {
			log.Printf("rwc close failed when server detached connection. Client address: %s", c.RemoteAddr())
		}
		return
	}

	tcpEventsServer.React = func(c gnet.Conn, inBuf *ringbuffer.RingBuffer) (out []byte, action gnet.Action) {
		fmt.Println("remote addr: ", c.RemoteAddr())
		head, tail := inBuf.PreReadAll()
		dbEventSourceMessage := append(head, tail...)
		log.Printf("DB source request messages: %s, messages size: %d", string(dbEventSourceMessage), len(dbEventSourceMessage))
		inBuf.Reset()

		client, ok := manager.ClientManger.Get(c.RemoteAddr().String())
		if !ok {
			manager.ClientManger.Remove(c.RemoteAddr().String())
			out = responseErrors.ResponseErrorBinary([2]byte{'V', '1'}, []byte("Connection has closed."))
			action = gnet.Close
			return
		}

		err := client.Write(dbEventSourceMessage)
		if err != nil {
			client.ClearBuffer()
			log.Printf("Write dbEventSourceMessage to buffer failed, dbEventSourceMessage: %s, client addr: %s", dbEventSourceMessage, c.RemoteAddr())
		}

		packages, consumeBytesLength, err := request.RequestDBEventPackageParser().Parse(client.GetBuffer())
		if err != nil {
			log.Print(err)
		}
		err = client.DiscardConsumedBuffer(consumeBytesLength)
		if err != nil {
			client.ClearBuffer()
			log.Printf("Reset the length of buffer by consuming failed, consumeBytesLength: %d, buffer length: %d", consumeBytesLength, len(buffer.Bytes()))
		}

		for _, p := range packages {
			fmt.Println(string(p.Method))
			fmt.Println(string(p.Keys))
			fmt.Println(string(p.Values))
		}

		dataResponse := response.Response(packages)
		log.Printf("response: %s", string(dataResponse))
		out = dataResponse

		return
	}

	tcpEventsServer.Tick = func() (delay time.Duration, action gnet.Action) {
		return
	}

	err := gnet.Serve(tcpEventsServer, fmt.Sprintf("%s://%s:%d", transportLayerProtocol, ip, port))
	if err != nil {
		log.Fatalf("NewServer start failed, address: tcp://%s:%d", ip, port)
	}
}
