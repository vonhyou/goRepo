package main

import "fmt"

/*
func typeof(v interface{}) string {
  return fmt.Sprintf("%T", v)
}
*/

func main() {
  /*c := '+'
    d := "+"
    fmt.Println(typeof(c), typeof(d))*/
  m := make(map[string]int)
  m["one"], m["two"] = 1, 2

  fmt.Println("map:", m)

  v1 := m["one"]
  fmt.Println(v1, len(m))
  _, ok := m["three"]
  fmt.Println(ok)

  n := map[string]int{"foo": 1, "bar": 2}
  fmt.Println(n)
}
