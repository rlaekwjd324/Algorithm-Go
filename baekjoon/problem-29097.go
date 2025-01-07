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

   var n, m, k, a, b, c int
   fmt.Fscan(reader, &n, &m, &k, &a, &b, &c)

   max := n*a
   if max < m*b {
      max = m*b
   }
   if max < k*c {
      max = k*c
   }
   
   answer := ""
   if max == n*a {
      answer += "Joffrey "
   }
   if max == m*b {
      answer += "Robb "
   }
   if max == k*c {
      answer += "Stannis"
   }
   
   fmt.Fprintln(writer, answer)
}
