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

   var r, c int
   fmt.Fscan(reader, &r, &c)

   isLeft := true
   for i := 0; i < r; i++ {
      var word string
      fmt.Fscan(reader, &word)

      var answer string
      if (c-len(word))% 2 == 0 {
         answer = fmt.Sprintf("%s%s%s", strings.Repeat(".", (c-len(word))/2), word, strings.Repeat(".", (c-len(word))/2))
      } else {
         if isLeft {
            answer = fmt.Sprintf("%s%s%s", strings.Repeat(".", (c-len(word))/2), word, strings.Repeat(".", (c-len(word))/2+1))
            isLeft = false
         } else {
            answer = fmt.Sprintf("%s%s%s", strings.Repeat(".", (c-len(word))/2+1), word, strings.Repeat(".", (c-len(word))/2))
            isLeft = true
         }
      }
      fmt.Fprintln(writer, answer)
   }
}
