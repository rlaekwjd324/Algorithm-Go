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
   
   var a, b int
   fmt.Fscan(reader, &a, &b)
   
   ax := (a-1)/4
   ay := (a-1)%4
   bx := (b-1)/4
   by := (b-1)%4

   d := 0
   if ax > bx {
      d += (ax - bx)
   } else {
      d += (bx - ax)
   }
   if ay > by {
      d += (ay - by)
   } else {
      d += (by - ay)
   }
   fmt.Fprintln(writer, d)
}
