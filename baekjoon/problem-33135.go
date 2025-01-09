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

   var n string
   fmt.Fscan(reader, &n)
   
   arr := make([]int, 26)
   for i := 0; i < len(n); i++ {
      if arr[int(n[len(n)-i-1])-65] > 0 {
         fmt.Fprintln(writer, len(n)-i)
         return
      }
      arr[int(n[len(n)-i-1])-65]++
   }
}
