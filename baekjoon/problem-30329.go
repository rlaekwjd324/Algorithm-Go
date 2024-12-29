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

   var str string
   fmt.Fscan(reader, &str)

   c := 0
   for i, s := range str {
      if s == 'k' && i <= len(str) - 4 && str[i+1] == 'i' && str[i+2] == 'c' && str[i+3] == 'k' {
         c++
      }
   }

   fmt.Fprintln(writer, c)
}
