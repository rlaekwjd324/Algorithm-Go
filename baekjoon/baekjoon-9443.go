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
   
   arr := make([]int, 26)
   for i := 0; i < n; i++ {
      var s string
      fmt.Fscan(reader, &s)

      arr[int(s[0])-65]++
   }
   c := 0
   for i := 0; i < 26; i++ {
      if arr[i] == 0 {
         break
      }
      c++
   }
   
   fmt.Fprintf(writer, "%v\n", c)
}
