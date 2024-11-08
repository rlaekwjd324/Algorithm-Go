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

   var k int
   fmt.Fscanln(reader, &k)
   var word string
   fmt.Fscanln(reader, &word)
   
   answer := ""
   for i, input := range word {
      s := 3*(i+1)+k
      new := int(input)-s
      if new < 65 {
         new += 26
      }
      answer += string(rune(new))
   }
   
   fmt.Fprintln(writer, answer)
}
