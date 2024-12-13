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

   var n string
   var m int
   fmt.Fscan(reader, &n, &m)

   var sum string
   if m <= 26 {
      sum = string(rune(m+64))
   } else {
      tmp := m-27
      sum = string(rune(tmp/26+97))
      sum += string(rune(tmp%26+97))
   }

   fmt.Fprintf(writer, "SN %s%v\n", n, sum)
}
