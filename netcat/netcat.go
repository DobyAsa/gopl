package main

import (
	"flag"
	"io"
	"log"
	"net"
	"os"
	"sync"
)

var port int
var hostname string
var inter bool

func init() {
	flag.IntVar(&port, "port", 8080, "The port you want to connect")
	flag.StringVar(&hostname, "hostname", "localhost", "The hostname you want to connect")
	flag.BoolVar(&inter, "i", true, "If you want to immediately close remote server")
	flag.Parse()
}
func main() {

	conn, err := net.DialTCP("tcp", &net.TCPAddr{IP: net.ParseIP("127.0.0.1")},
		&net.TCPAddr{IP: net.ParseIP(hostname), Port: port})
	if err != nil {
		log.Fatal(err)
	}
	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		io.Copy(os.Stdout, conn)
		log.Println("done")
		wg.Done()
	}()
	if inter {
		handle(conn, os.Stdin)
	}
	conn.CloseWrite()
	wg.Wait()
	conn.Close()
}

func handle(d io.Writer, s io.Reader) {
	_, err := io.Copy(d, s)
	if err != nil {
		log.Fatal(err)
	}
}
