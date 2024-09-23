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
   fmt.Fscan(reader, &n)
   var str string
   fmt.Fscan(reader, &str)

   if strings.Contains(str, "gori") {
      fmt.Fprintln(writer, "YES")
      return
   }
   fmt.Fprintln(writer, "NO")
}
