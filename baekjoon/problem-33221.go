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

   var t, p int
   fmt.Fscan(reader, &t, &p)
   
   answer := 0

   rC := p
   rS := 0
   for j := 0; j < p; j++ {
      var m string
      fmt.Fscan(reader, &m)
      if m == "X" {
         rC--
         continue
      }
      mi, _ := strconv.Atoi(m)
      rS += mi
   }
   for i := 1; i < t; i++ {
      count := p
      sum := 0
      for j := 0; j < p; j++ {
         var m string
         fmt.Fscan(reader, &m)
         if m == "X" {
            count--
            continue
         }
         mi, _ := strconv.Atoi(m)
         sum += mi
      }
      if rC < count {
         answer++
         continue
      }
      if rC == count && rS >= sum {
         answer++
      }
   }
   fmt.Fprintln(writer, answer)
}
