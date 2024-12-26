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

   var t, n int
   var x string
   fmt.Fscan(reader, &t, &x, &n)

   for i := 0; i < n; i++ {
      var k int
      fmt.Fscan(reader, &k)
      arr := make([]string, k)
      for j := 0; j < k; j++ {
         var ki string
         fmt.Fscan(reader, &ki)
         arr[j] = ki
      }
      if !strings.Contains(strings.Join(arr, " "), x) {
         fmt.Fprintln(writer, "NO")
         return
      }
   }

   fmt.Fprintln(writer, "YES")
}
