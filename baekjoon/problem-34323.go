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

   var n, m, s int64
   fmt.Fscan(reader, &n, &m, &s)
   
   dn := s*(m+1)*(100-n)/100
   dm := s*m

   if dn > dm {
      fmt.Fprintf(writer, "%v\n", dm)
      return
   }
   fmt.Fprintf(writer, "%v\n", dn)
}
