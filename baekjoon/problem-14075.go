package main

import (
   "bufio"
   "fmt"
   "os"
   "strconv"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func main() {
   defer writer.Flush()

   var s string
   fmt.Fscan(reader, &s)
   
   sum := 0
   for i, v := range s {
      n, _ := strconv.Atoi(string(v))
      if i % 2 != 0 {
         sum += n
         continue
      }
      n = n*2
      if n >= 10 {
         n = n/10+n%10
      }
      sum += n
   }
   if sum % 10 == 0 {
      fmt.Fprintf(writer, "%v\n", "DA")
      return
   }
   fmt.Fprintf(writer, "%v\n", "NE")
}
