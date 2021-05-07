package display

import (
  "reflect"
  "fmt"
)

func Display(name string, x interface{}) string {
  fmt.Printf("Display %s (%T): \n", name, x)
  display(name, reflect.ValueOf(x))
}

func display(name string, x reflect.Value) string {
  switch v.Kind() {

  }
}
