package main

import (
	"github.com/urfave/cli"
	"io"
	"log"
	"net"
	"os"
	"sync"
	"time"
)

var (
	xmitBuf sync.Pool
)



func checkError(err error) {
	if err != nil {
		log.Printf("%+v\n", err)
		os.Exit(-1)
	}
}

func main() {

	// add more log flags for debugging
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	xmitBuf.New = func() interface{} {
		return make([]byte, 32768)
	}
	myApp := cli.NewApp()
	myApp.Name = "tcptun"
	myApp.Usage = "tcptun -l :2022  -t 192.168.1.200:80"
	myApp.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "targettcp, t",
			Value: "192.168.1.1:80",
			Usage: "target server address",
		},
		cli.StringFlag{
			Name:  "listentcp,l",
			Value: "127.0.0.1:2022",
			Usage: "local listen address",
		},
		cli.StringFlag{
			Name:  "c",
			Value: "", // when the value is not empty, the config path must exists
			Usage: "config from json file, which will override the command from shell",
		},
	}
	myApp.Action = func(c *cli.Context) error {
		xmitBuf.New = func() interface{} {
			return make([]byte, 32768)
		}

		config := Config{}
		config.ListenTcp = c.String("listentcp")
		config.TargetTcp = c.String("targettcp")
		//config.TargetTcp = "127.0.0.1:22"


		if c.String("c") != "" {
			err := parseJSONConfig(&config, c.String("c"))
			checkError(err)
		}


		chTCPConn := make(chan *net.TCPConn, 16)
		go tcpListener(chTCPConn, &config)


		tickerCheck := time.NewTicker(10*time.Second)
		defer tickerCheck.Stop()
		for {
			select {
			case p1 := <-chTCPConn:
				go handleLocalTcp(p1, &config)
			case <-tickerCheck.C:
				//log.Println("working")
			}
		}


		return nil
	}
	myApp.Run(os.Args)
}



func tcpListener(chTCPConn chan *net.TCPConn, config *Config){
	listenTcpAddr, err := net.ResolveTCPAddr("tcp4", config.ListenTcp)
	checkError(err)
	listener, err := net.ListenTCP("tcp4", listenTcpAddr)
	checkError(err)
	log.Println("tcp listening on:", listener.Addr())
	for{
		p1, err := listener.AcceptTCP()
		if err != nil {
			log.Fatalln(err)
		}
		chTCPConn <- p1
	}
}

func handleLocalTcp(p1 io.ReadWriteCloser, config *Config) {

	p2, err := net.DialTimeout("tcp", config.TargetTcp, 5*time.Second)
	if err != nil {
		p1.Close()
		log.Println(err)
		return
	}
	go func() {
		if true {
			log.Println("tcp client opened")
			defer log.Println("tcp client closed")
		}
		defer p1.Close()
		defer p2.Close()

		streamCopy := func(dst io.Writer, src io.ReadCloser) chan struct{} {
			die := make(chan struct{})
			go func() {
				buf := xmitBuf.Get().([]byte)
				CopyBuffer(dst, src, buf)
				//generic.Copy(dst, src)
				xmitBuf.Put(buf)
				close(die)
			}()
			return die
		}

		select {
		case <-streamCopy(p1, p2):
		case <-streamCopy(p2, p1):
		}
	}()
}

