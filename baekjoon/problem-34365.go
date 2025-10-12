package main

import (
	"fmt"
	"bufio"
	"os"
   "strings"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func main() {
	defer writer.Flush()

   var n int
   fmt.Fscan(reader, &n)
   
   d := []string{"abcABC", "defDEF", "ghiGHI", "jklJKL", "mnoMNO", "pqrsPQRS", "tuvTUV", "wxyzWXYZ"}
   for i := 0; i < n; i++ {
      var s string
      fmt.Fscan(reader, &s)
      
      answer := ""
      for _, v := range s {
         for j, dt := range d {
            if strings.Contains(dt, string(v)) {
               answer = fmt.Sprintf("%v%v", answer, j+2)
               break
            }
         }
      }
      fmt.Fprintf(writer, "%v\n", answer)
   }
}
