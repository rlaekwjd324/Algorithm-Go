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
   
   answer := ""
   for i := 0; i < n; i++ {
       var p, q float64
       fmt.Fscan(reader, &p, &q)
       rf := 1/p+1/q
       f := 1/rf
       f = math.Round(f*10)/10
       answer += fmt.Sprintf("f = %.1f\n", f) 
   }

   fmt.Fprintln(writer, answer)
}
