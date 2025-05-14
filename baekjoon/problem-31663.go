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
   
   arr := make([]string, n)
   maxL := 0
   for i := 0; i < n; i++ {
      var s string
      fmt.Fscan(reader, &s)

      arr[i] = s
      if maxL < len(s) {
         maxL = len(s)
      }
   }

   str := ""
   for i := 0; i < maxL; i++ {
      sum := 0
      c := 0
      for j := 0; j < n; j++ {
         if len(arr[j]) < i+1 {
            continue
         }
         sum += int(arr[j][i])
         c++
      }
      str += string(rune(sum/c))
   }

   fmt.Fprintf(writer, "%v\n", str)
}
