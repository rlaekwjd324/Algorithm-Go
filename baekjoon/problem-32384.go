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

   var n int
   fmt.Fscanln(reader, &n)

   answer := ""
   for i := 0; i < n; i++ {
      answer += fmt.Sprintf("%s ", "LoveisKoreaUniversity")
   }
   
   fmt.Fprintln(writer, answer)
}
