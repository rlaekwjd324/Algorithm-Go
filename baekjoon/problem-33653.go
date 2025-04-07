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
   
   var s string
   var n int
   fmt.Fscan(reader, &s, &n)
   
   arr := make([]string, n)
   for i := 0; i < n; i++ {
      var a string
      fmt.Fscan(reader, &a)

      arr[i] = a
   }
   
   fmt.Fprintln(writer, countTarget(arr, s))
}

func countTarget(arr []string, target string) int {
   c := 0
   for _, v := range arr {
      if target == string(v) {
         c++
         continue
      }
      for i := 0; i <= len(v) - len(target); i++ {
         if v[i:len(target)+i] == target {
            c++
         }
      }
   }
   return c
}
