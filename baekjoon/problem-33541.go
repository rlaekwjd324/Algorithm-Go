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
   
   var t int
   fmt.Fscan(reader, &t)

   arr := []int{2025, 3025, 9801}
   
   if t >= arr[2] {
      fmt.Fprintln(writer, -1)
      return
   }
   if t >= arr[1] {
      fmt.Fprintln(writer, arr[2])
      return
   }
   if t >= arr[0] {
      fmt.Fprintln(writer, arr[1])
      return
   }
   fmt.Fprintln(writer, arr[0])
}
