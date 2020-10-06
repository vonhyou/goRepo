package main

import "fmt"

func main() {
  s := make([]string, 3)
  fmt.Println("empty", s)

  s[0], s[1], s[2] = "a", "b", "c"
  fmt.Println("len:", len(s))
  fmt.Println(s)
  s = append(s, "d")
  fmt.Println(s)
  s = append(s, "e", "f")
  fmt.Println("final:", s)

  c := make([]string, len(s))
  copy(c, s)
  fmt.Println("copy:", c)

  fmt.Println(append(s, c...))
}
