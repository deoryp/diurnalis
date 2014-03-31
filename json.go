package main

import (
  "fmt"
  "log"
  "encoding/json"
)

type Message struct {
    Name string
    Body string
    Time int64
}

func main() {
  m := Message{"Alice", "Hello", 1294706395881547000}
  b, err := json.Marshal(m)
  if err != nil {
    log.Fatal(err)
  }

  var c = []byte(`{"Name":"Alice","Body":"Hello","Time":1294706395881547000}`)
  

  var n Message
  err2 := json.Unmarshal(c, &n)
  if err2 != nil {
    log.Fatal(err2)
  }

  fmt.Printf("m = %+v\nb = %s\nn = %+v\n", m, b, n)

}
