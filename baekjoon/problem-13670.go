package main

import (
	"fmt"
	"bufio"
	"os"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func main() {
	defer writer.Flush()

   for {
      var h1, m1, h2, m2 int
      fmt.Fscan(reader, &h1, &m1, &h2, &m2)

      if h1 == 0 && m1 == 0 && h2 == 0 && m2 == 0 {
         break
      }
      m := m2-m1
      if m < 0 {
         m += 60
         h2--
      }
      h := h2-h1
      if h < 0 {
         h += 24
      }
      m += h * 60

      fmt.Fprintf(writer, "%v\n", m)
   }
}
