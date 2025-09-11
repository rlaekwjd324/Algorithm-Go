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

   cnt := 0
   for {
      var s string
      if _, err := fmt.Fscan(reader, &s); err != nil {
         break
      }
      cnt += countCertainly(s)
   }
   
   fmt.Fprintf(writer, "%v\n", cnt)
}

func countCertainly(s string) int {
   cnt := 0
   
   for i := 0; i < len(s); i++ {
      if len(s) < i+9 {
         break
      }
      if s[i:i+9] == "certainly" {
         cnt++
         i += 8
      }
   }
   
   return cnt
}
