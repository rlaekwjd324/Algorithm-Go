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
   fmt.Fscan(reader, &n)
   
   pre := 1
   for i := 0; i < n; i++ {
      var a, b int
      var c string
      fmt.Fscan(reader, &a, &c, &b)
      
      switch c {
      case "+":
         pre = a+b-pre
      case "-":
         pre = (a-b)*pre
      case "*":
         pre = a*b*a*b
      case "/":
         if a % 2 == 0 {
            pre = a/2
         } else {
            pre = (a+1)/2
         }
      }
      fmt.Fprintln(writer, pre)
   }   
}
