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
   
   var a, b int
   fmt.Fscan(reader, &a, &b)

   aArray := getDivNumber(a)
   bArray := getDivNumber(b)
   
   for _, v := range aArray {
      for _, w := range bArray {
         fmt.Fprintf(writer, "%v %v\n", v, w)
      }
   } 
}

func getDivNumber(num int) []int {
   var arr []int
   for i := 1; i <= num; i++ {
      if num % i == 0 {
         arr = append(arr, i)
      }
   }
   return arr
}
