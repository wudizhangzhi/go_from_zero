package main

import (
	"io"
	"net"
	"os"
	"fmt"
	"log"
	"flag"
)


var port = flag.Int("port", 8000, "remote port")

func main() {
	flag.Parse()
	fmt.Printf("connecting %d...\n", *port)
	conn, err := net.Dial("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	mustCopy(os.Stdout, conn)

}


func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err!=nil {
		log.Fatal(err)
	}
}
