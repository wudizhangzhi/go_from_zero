package format

import (
  "reflect"
  "strconv"
)

func Any(value interface{}) {
  return formatAtom(reflect.ValueOf(value))
}

func formatAtom(v reflect.Value) string {
  switch v.Kind() {
  case reflect.Invalid:
    return "invalid"
  case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
    return strconv.FormatInt(v.Int(), 10)
  case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uintptr:
    return strconv.FormatUnit(v.Uint(), 10)
  // ...floating-point and complex cases
  case reflect.Bool:
    return strconv.FormatBool(v.Bool())
  case reflect.String:
    return strconv.Quote(v.String())
  case reflect.Chan, reflect.Ptr, reflect.Func, reflect.Slice, reflect.Map:
    return v.Type().String() + " 0x" + strconv.FormatUnit(unit64(v.Pointer()), 16)
  default: // reflect.Array, reflect.Struct, reflect.Interface 
    return v.Type().String() + " value"
  }
}
