package main

import (
  "fmt"
  // "sort"
  "time"
)

func length(s string) time.Duration {
  d, err := time.ParseDuration(s)
  if err!= nil {
    panic(err)
  }
  return d
}

type Track struct {
  Title string
  Artist string
  Album string
  Yeat int
  Length time.Duration
}

var tracks = []Track{
  {"Go", "Delilah", "From the Roots Up", 2012, length("3m20s")},
  {"Go", "Delilah", "From the Roots Up", 2012, length("3m20s")},
  {"Go Ahead", "Delilah", "From the Roots Up", 2012, length("3m20s")},
  {"Go", "Delilah", "From the Roots Up", 2012, length("3m20s")},
}


func main() {
  fmt.Println(tracks)
}
