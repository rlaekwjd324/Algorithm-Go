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
   
   aArr := []string{"keys", "phone", "wallet"}
   bArr := make([]string, n)
   isReady := true
   for i := 0; i < n; i++ {
      var s string
      fmt.Fscan(reader, &s)
      bArr[i] = s
   }
   for i := 0; i < 3; i++ {
      if !isContains(bArr, aArr[i]) {
         fmt.Fprintf(writer, "%v\n", aArr[i])
         isReady = false
      }
   }
   if isReady {
      fmt.Fprintf(writer, "%v\n", "ready")
   }

}

func isContains(arr []string, target string) bool {
   for _, v := range arr {
      if v == target {
         return true
      }
   }
   return false
}
