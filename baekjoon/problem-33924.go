package main

import (
   "bufio"
   "fmt"
   "os"
   "strings"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func main() {
   defer writer.Flush()

   var n, m, k int
   fmt.Fscan(reader, &n, &m, &k)
   
   var s string 
   fmt.Fscan(reader, &s)

   arr := []int{1, 2, 3, 4, 5, 6, 7, 8}
   for _, v := range s {
      switch string(v) {
      case "A":
         arr := []int{arr[4], arr[5], arr[6], arr[7], arr[0], arr[1], arr[2], arr[3]}
      case "B":
         arr := []int{arr[3], arr[2], arr[1], arr[0], arr[7], arr[6], arr[5], arr[4]}
      case "C":
         arr := []int{arr[7], arr[6], arr[5], arr[4], arr[3], arr[2], arr[1], arr[0]}
      case "D":
         arr := []int{arr[2], arr[0], arr[4], arr[1], arr[6], arr[3], arr[7], arr[5]}
      }
   }
   
   fmt.Fprintf(writer, "%v\n", )
}
