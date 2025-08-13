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

   var c, m, l int
   fmt.Fscan(reader, &c, &m, &l)
   
   cnt := 0
   n := 0
   for {
      cnt++
      n += c
      if n >= l {
         break
      }
      n-=m
   }

   fmt.Fprintf(writer, "%v\n", cnt)
}
