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

   var n string
   fmt.Fscanln(reader, &n)

   if strings.Contains(n, "ss") {
      fmt.Fprintln(writer, "hiss")
      return
   }
   fmt.Fprintln(writer, "no hiss")
}
