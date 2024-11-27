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

   var n int
   fmt.Fscanln(reader, &n)
   
   for i := 0; i < n; i++ {
      var a string
      fmt.Fscanln(reader, &a)
      var b string
      fmt.Fscanln(reader, &b)
      bArr := strings.Split(b, "")
      for j := 0; j <len(bArr); j++ {
         if !strings.Contains(a, bArr[j]) {
            fmt.Fprintln(writer, "NO")
            break
         }
         if j == len(bArr)-1 {
            fmt.Fprintln(writer, "YES")
         }
      }
   }
}
