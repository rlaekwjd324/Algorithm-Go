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
   var p1, p2 string
   fmt.Fscan(reader, &n, &p1, &p2)
   
   c := 0
   arr1 := make([]int, 26)
   arr2 := make([]int, 26)
   for _, v := range p1 {
      arr1[int(v)-97]++
   }
   for _, v := range p2 {
      arr2[int(v)-97]++
   }
   for i := 0; i < 26; i++ {
      if arr1[i] > 0 && arr1[i] - arr2[i] > 0{
         c += (arr1[i] - arr2[i])
      }
   }
   
   fmt.Fprintln(writer, c)
}
