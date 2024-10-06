package main

import (
   "bufio"
   "fmt"
   "os"
   "strings"
   "strconv"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func main() {
   defer writer.Flush()

   arr := []string{}
   for i := 0; i < 5; i++ {
      var n string
      fmt.Fscan(reader, &n)
      
      if strings.Contains(n, "FBI") {
         index := strconv.Itoa(i+1)
         arr = append(arr, index)
      }
   }
   if len(arr) == 0 {
      fmt.Fprintln(writer, "HE GOT AWAY!")
      return
   }
   fmt.Fprintln(writer, strings.Join(arr, " "))
}
