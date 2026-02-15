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
   fmt.Fscan(reader, &n)
  
   for i:=n; i>=4; i-- {
      if containNum(fmt.Sprintf("%v", i)) {
         fmt.Fprintf(writer, "%v\n", i)
         return
      }
   }
}

func containNum(t string) bool {
   for _, v:= range t {
       if string(v) != "4" && string(v) !=  "7" {
         return false
      }
   }
   return true
}
