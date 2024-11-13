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
   fmt.Fscanln(reader, &n)
   
   sum := 0
   tempSum := 0
   for i := 0; i < n; i++ {
      var num int
      fmt.Fscan(reader, &num)

      if num == 1 {
         tempSum++
      } else {
         tempSum--
      }
      sum += tempSum
   }

   fmt.Fprintf(writer, strconv.Itoa(sum))
}
