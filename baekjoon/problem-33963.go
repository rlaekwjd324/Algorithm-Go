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

   var n int
   fmt.Fscan(reader, &n)

   c := -1
   l := 0
   for {
      s := strconv.Itoa(n)
      if l != 0 && len(s) != l {
         break
      }
      l = len(s)
      n = n * 2
      c++
   }

   fmt.Fprintf(writer, "%v\n", c)
}
