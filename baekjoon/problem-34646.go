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
   
   var n int
   fmt.Fscan(reader, &n)
   
   d := make([]int, n)
   for i := 0; i < n; i++ {
      fmt.Fscan(reader, &d[i])
   }
   
   a := d[0]/3
   c := d[n-1]/3
   b := d[1]-a*2

   fmt.Fprintf(writer, "%v %v %v\n", a, b, c)
}
