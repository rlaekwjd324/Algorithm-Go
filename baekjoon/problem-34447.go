package main

import (
	"fmt"
	"bufio"
	"os"
   "strconv"
   "math"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func main() {
	defer writer.Flush()

   var t int
   fmt.Fscan(reader, &t)
   
   for i := 0; i < t; i++ {
      var n, s int
      fmt.Fscan(reader, &n, &s)
      
      l := len(strconv.Itoa(s))
      c := int(math.Pow(10, float64(l)))
      for j := 0; j < l; j++ {
         c = c/10
         v := (s/c+n)%10
         fmt.Fprintf(writer, "%v", v)
      }
      fmt.Fprint(writer, "\n")
   }
}
