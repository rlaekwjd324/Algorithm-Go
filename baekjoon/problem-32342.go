package main

import (
   "bufio"
	"fmt"
	"os"
   "strings"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func main() {
   defer writer.Flush()

   var n int
   fmt.Fscan(reader, &n)

   for i := 0; i < n; i++ {
      count := 0
      var str string
      fmt.Fscan(reader, &str)
      arr := strings.Split(str, "")
      for i := 0; i < len(str)-2; i++ {
         if arr[i] == "W" && arr[i+1] == "O" && arr[i+2] == "W" {
            count++
         }
      }
      fmt.Fprintln(writer, count)
   }
}
