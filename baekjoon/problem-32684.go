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

   cho := make([]int, 6)
   choSum := 0.0
   han := make([]int, 6)
   hanSum := 1.5
   score := []int{13, 7, 5, 3, 3, 2}
   for i := 0; i < 12; i++ {
      var n int
      fmt.Fscan(reader, &n)
      if i < 6 {
         cho[i] = n
         choSum += float64(n*score[i%6])
      } else {
         han[i%6] = n
         hanSum += float64(n*score[i%6])
      }
   }
   if choSum > hanSum {
      fmt.Fprintf(writer, "cocjr0208")
      return
   }
   fmt.Fprintf(writer, "ekwoo")
}
