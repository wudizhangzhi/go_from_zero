package main

import (
  "net"
  "log"
  "time"
  "io"
  "fmt"
)

func main() {
  listener, err := net.Listen("tcp", "192.168.56.1:8000")
  if err != nil {
    log.Fatal(err)
  }
  fmt.Println("开始!")
  for {
    conn, err := listener.Accept()
    if err != nil {
       log.Print(err)
       continue
    }
    go handleConn(conn)
  }
}


func handleConn(c net.Conn) {
  defer c.Close()
  for {
    _, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
    if err != nil {
      return
    }
    time.Sleep(1*time.Second)
  }
}
