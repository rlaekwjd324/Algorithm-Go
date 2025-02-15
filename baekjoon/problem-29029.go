package main

import (
   "bufio"
   "fmt"
   "os"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func main() {
   defer writer.Flush()

   var n int
   var w string
   fmt.Fscan(reader, &n, &w)
   
   min := 400000
   arr := []string{"N", "W", "S", "E"}
   for i := 0; i < 4; i++ {
      c := 0
      for _, v := range w {
         if string(v) != arr[i] {
            c++
         }
      }
      if min > c {
         min = c
      }
   }

   fmt.Fprint(writer, min)
}
