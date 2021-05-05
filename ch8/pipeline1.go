package main

import "fmt"


func main() {
  naturals := make(chan int)
  squares := make(chan int)
  // counter
  go func() {
    for x :=0;;x++ {
      naturals <- x
    }
  }()

  // squarer
  go func() {
    // for {
    //  x := <-naturals
    //  squares <- x * x
    //}
    for x := range naturals {
      squares <- x * x
    }
    close(squares)
  }()

  for {
    fmt.Println(<-squares)
  }
}



