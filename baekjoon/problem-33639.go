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
   
   var n, m int
   fmt.Fscan(reader, &n, &m)
   
   arr := make([]string, n*2)
   for i := 0; i < n; i++ {
      var f, s string
      fmt.Fscan(reader, &f, &s)
      arr[i*2] = f
      arr[i*2+1] = s
   }
   for i := 0; i < m; i++ {
      var x string
      fmt.Fscan(reader, &x)
      c, j := countOfArray(arr, x)
      if c == 0 {
         fmt.Fprintln(writer, "nobody")
         continue
      }
      if c > 1 {
         fmt.Fprintln(writer, "ambiguous")
         continue
      }
      fmt.Fprintf(writer, "%s %s\n", arr[j*2], arr[j*2+1])
   }
}
func countOfArray(arr []string, target string) (count, index int) {
   c := 0
   i := -1
   for j := 0; j < len(arr)/2; j++ {
      f := arr[j*2]
      s := arr[j*2+1]

      new := string(f[0])+string(s[0])
      if new == target {
         c++
         i = j
      }
   }
   return c, i
}
