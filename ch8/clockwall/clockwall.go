package main


import (
	//"io"
	"os"
	"net"
	"fmt"
	"log"
	//"flag"
	"strings"
	"text/tabwriter"
	"bufio"
	//"time"
)

// name:host
type ZoneTime struct {
	name string
	t string
}
var wall = make(map[string]string)
var message = make(chan ZoneTime)

func main() {
	args := os.Args[1:]
	// connection
	for _, arg := range args {
		arg_splited := strings.Split(arg, "=")
		// hosts[arg_splited[0]] = arg_splited[1]
		go netcat(arg_splited[1], arg_splited[0])
	}
	//fmt.Println(hosts)
	// display
	// go netcat()
	for msg := range message{
		printTable(msg)
	}
	defer close(message)
}


func netcat(host string, name string) {
	conn, err := net.Dial("tcp", host)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
        // if _, err := io.WriteString(os.Stdout, conn); err != nil {
        //		log.Fatal(err)
        // }
	for {
		resp, _ := bufio.NewReader(conn).ReadString('\n')
		//fmt.Println(resp)
		message <- ZoneTime{name: name, t: resp}
	}
}


func printTable(t ZoneTime) {
	wall[t.name] = t.t
	nw := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', tabwriter.AlignRight|tabwriter.Debug)
	for key := range wall {
		//
	}
}
