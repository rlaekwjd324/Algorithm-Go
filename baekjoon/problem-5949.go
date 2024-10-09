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

   var n string
   fmt.Fscan(reader, &n)
   arr := strings.Split(n, "")

   answer := ""
   for i := 0; i < len(n); i++ {
      if i != 0 && i % 3 == 0 {
         answer = "," + answer
      }
      answer = arr[len(n)-i-1] + answer
   }
   fmt.Fprintln(writer, answer)
}
