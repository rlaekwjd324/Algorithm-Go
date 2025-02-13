package main

import (
   "bufio"
   "fmt"
   "os"
   "strings"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func main() {
   defer writer.Flush()

   var t, s, h int
   fmt.Fscan(reader, &t, &s, &h)
   
   answer := ""
   for i := 0; i < t; i++ {
      answer += "*"
      answer += strings.Repeat(" ", s)
      answer += "*"
      answer += strings.Repeat(" ", s)
      answer += "*\n"
   }
   answer += strings.Repeat("*", 3+s+s)
   answer += "\n"
   for i := 0; i < h; i++ {
      answer += strings.Repeat(" ", s+1)
      answer += "*\n"
   }

   fmt.Fprint(writer, answer)
}
