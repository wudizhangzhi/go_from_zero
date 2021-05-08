package main

import (
	"flag"
	"fmt"
	"io"
	"net"
	"time"
)

var port = flag.Int("port", 8000, "port number")

func main() {
	flag.Parse()
	fmt.Printf("start at port: %d\n", *port)
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		fmt.Errorf("%v", err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		handleConn(conn)
	}
}

func handleConn(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:05 -0700 MST\n"))
		if err != nil {
			return
		}
		time.Sleep(1 * time.Second)
	}

}
