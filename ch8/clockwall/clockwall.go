package main

import (
	"fmt"
	"log"
	"net"
	"bufio"
	"sort"

	//"io"
	"os"
	//"flag"
	"strings"
	//"time"
)

// name:host
type ZoneTime struct {
	name string
	t    string
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
	for msg := range message {
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
		resp, _ := bufio.NewReader(conn).ReadString(' ')
		//fmt.Println(resp)
		message <- ZoneTime{name: name, t: resp}
	}
}

func printTable(t ZoneTime) {
	var names []string
	var line string
	wall[t.name] = t.t
	//nw := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', tabwriter.AlignRight|tabwriter.Debug)
	for key := range wall {
		names = append(names, key)
	}
	sort.Strings(names)
	for _, name := range names {
		line += fmt.Sprintf("%s 时间: %s", name, wall[name])
	}
	//line += "\n"
	fmt.Printf("%s\r", line)
	//nw.Flush()
}
