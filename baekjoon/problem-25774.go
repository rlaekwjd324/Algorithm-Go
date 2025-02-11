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

   var d, d1, m, m1, y, y1, n int
   fmt.Fscan(reader, &d, &m, &y, &n, &d1, &m1, &y1)
   
   sumD := d1-d
   if sumD < 0 {
      m1--
      sumD += 30
   }
   
   sumM := m1-m
   if sumM < 0 {
      y1--
      sumM += 12
   }

   sumY := y1-y
   sum := sumD + sumM*30 + sumY*360 + n
   
   a := sum % 7
   if a == 0 {
      a = 7
   }
   fmt.Fprintln(writer, a)
}
