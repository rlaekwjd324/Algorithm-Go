package main

import (
   "bufio"
   "fmt"
   "os"
   "math"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func main() {
   defer writer.Flush()
   
   var n int
   fmt.Fscan(reader, &n)
   
   for i := 0; i < n; i++ {
      var a int
      fmt.Fscan(reader, &a)

      s := getSumDivNums(a)
      if s == a {
         fmt.Fprintf(writer, "%v is a perfect number.\n\n", a)
         continue
      }
      if s < a {
         fmt.Fprintf(writer, "%v is a deficient number.\n\n", a)
         continue
      }
      fmt.Fprintf(writer, "%v is an abundant number.\n\n", a)
   }
}

func getSumDivNums(t int) int {
   sum := 1
   for i := 2; i <= int(math.Sqrt(float64(t))); i++ {
      if t % i == 0 {
         sum += i
         if i != t/i {
            sum += (t/i)
         }
      }
   }
   return sum
}
