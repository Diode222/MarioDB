package server

import (
	"flag"
	"fmt"
	"github.com/Diode222/MarioDB/parser/dbEventPackage/request"
	"github.com/panjf2000/gnet"
	"github.com/panjf2000/gnet/ringbuffer"
	"io"
	"log"
	"sync"
	"time"
)

type listener struct {
	IP        string
	Port      uint64
	ReusePort bool
	Loops     uint64
}

var l *listener
var once sync.Once

func Listener(ip string, port uint64, reusePort bool, loops uint64) *listener {
	once.Do(func() {
		l = &listener{
			IP:        ip,
			Port:      port,
			ReusePort: reusePort,
			Loops:     loops,
		}
	})
	return l
}

func (l *listener) Init() {
	var ip string
	var port uint64
	var reusePort bool
	var loops uint64
	var transportLayerProtocol string = "tcp"

	flag.StringVar(&ip, "ip", l.IP, "Server ip")
	flag.Uint64Var(&port, "port", l.Port, "Server port")
	flag.BoolVar(&reusePort, "reusePort", l.ReusePort, "Reuse listener port in cluster")
	flag.Uint64Var(&loops, "loops", l.Loops, "Loops number the server is using")
	flag.Parse()

	var dbEventsListener gnet.Events
	dbEventsListener.NumLoops = int(loops)

	dbEventsListener.OnInitComplete = func(srv gnet.Server) (action gnet.Action) {
		log.Printf("MarioDB server started on tcp://%s.", srv.Addrs)
		return
	}

	dbEventsListener.OnOpened = func(c gnet.Conn) (out []byte, opts gnet.Options, action gnet.Action) {
		log.Printf("Client started on tcp://%s.", c.RemoteAddr())
		out = []byte("TCP has connected.")
		return
	}

	dbEventsListener.OnClosed = func(c gnet.Conn, err error) (action gnet.Action) {
		log.Printf("Client(Address: %s) closed connection.", c.RemoteAddr())
		return
	}

	dbEventsListener.OnDetached = func(c gnet.Conn, rwc io.ReadWriteCloser) (action gnet.Action) {
		log.Printf("Server detached connection of client(Address: %s).", c.RemoteAddr())
		_, err := rwc.Write([]byte(fmt.Sprintf("Server(Address: %s) detached connection.", c.LocalAddr())))
		if err != nil {
			log.Printf("Server detached info send failed. Client address: %s", c.RemoteAddr())
		}
		err = rwc.Close()
		if err != nil {
			log.Printf("rwc close failed when server detached connection. Client address: %s", c.RemoteAddr())
		}
		return
	}

	dbEventsListener.React = func(c gnet.Conn, inBuf *ringbuffer.RingBuffer) (out []byte, action gnet.Action) {
		head, tail := inBuf.PreReadAll()
		dbEventSourceMessage := append(head, tail...)
		log.Printf("DB source request messages: %s, messages size: %d", string(dbEventSourceMessage), len(dbEventSourceMessage))

		packages, err := request.RequestDBEventPackageParser().Parse(inBuf)
		if err != nil {
			log.Print(err)
		}

		for _, p := range packages {
			fmt.Println(string(p.Version[0]) + string(p.Version[1]))
			fmt.Println(string(p.Keys))
		}

		out = response(packages)

		return
	}

	dbEventsListener.Tick = func() (delay time.Duration, action gnet.Action) {
		return
	}

	err := gnet.Serve(dbEventsListener, fmt.Sprintf("%s://%s:%d", transportLayerProtocol, ip, port))
	if err != nil {
		log.Fatalf("Server start failed, address: tcp://%s:%d", ip, port)
	}
}
