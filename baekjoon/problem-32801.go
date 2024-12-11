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

   var n, a, b int
   fmt.Fscan(reader, &n, &a, &b)

   arr := []int{0, 0, 0}
   for i := 1; i <= n; i++ {
      if i % a == 0 && i % b == 0 {
         arr[2]++
      } else if i % a == 0 {
         arr[0]++
      } else if i % b == 0 {
         arr[1]++
      }
   }
   
   fmt.Fprintf(writer, "%v %v %v\n", arr[0], arr[1], arr[2])
}
