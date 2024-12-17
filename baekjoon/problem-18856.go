package main

import (
   "bufio"
	"fmt"
	"os"
   "strings"
   "math"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func main() {
   defer writer.Flush()

   var n int
   fmt.Fscan(reader, &n)
   arr := make([]string, n)
   arr[0] = "1"
   arr[1] = "2"
   count := 2
   for i := 3; i < 1000; i++ {
      if isSosu(i) {
         arr[count] = fmt.Sprintf("%v", i)
         count++
         if count == n {
            break
         }
      }
   }
   
   fmt.Fprintf(writer, "%v\n%v\n", n, strings.Join(arr, " "))
}

func isSosu(num int) bool {
   for i := 2; float64(i) <= math.Sqrt(float64(num)); i++ {
      if num % i == 0 {
         return false
      }
   }
   return true
}
