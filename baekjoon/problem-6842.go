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
   fmt.Fscanln(reader, &n)
   papers := []int{100, 500, 1000, 5000, 10000, 25000, 50000, 100000, 500000, 1000000}
   
   sum := 0
   for i := 0; i < len(papers); i++ {
      sum += papers[i]
   }
   for i := 0; i < n; i++ {
      var index int
      fmt.Fscanln(reader, &index)
      sum -= papers[index-1]
   }

   var banker int
   fmt.Fscanln(reader, &banker)

   if banker > sum/(10-n) {
      fmt.Fprintf(writer, "deal")
      return
   }
   fmt.Fprintf(writer, "no deal")
}
