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
   
   var s string
   fmt.Fscan(reader, &s)
   
   t := "SciComLove"
   c := 0
   for i := 0; i < len(s); i++ {
      if t[i] != s[i] {
         c++
      }
   }

   fmt.Fprintf(writer, "%v\n", c)
}
