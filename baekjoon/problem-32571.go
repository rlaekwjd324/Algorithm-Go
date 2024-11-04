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

   var n int
   fmt.Fscanln(reader, &n)
   
   answer := ""
   arr := make([]string, n)
   for i := 0; i < n; i++ {
      var str string
      fmt.Fscanln(reader, &str)
      if strings.Contains(str, "rest") {
         arr[i] = "rest"
         continue
      }
      if strings.Contains(str, "leg") {
         arr[i] = "leg"
         continue
      }
      arr[i] = "arm"
   }
   for i := 0; i < 31; i++ {
      if i%7 == 0 {
         answer += fmt.Sprintf("%v ", i/7+1)
      }
      switch arr[i%n] {
      case "rest":
         answer += "😎"
      case "leg":
         answer += "🦵"
      case "arm":
         answer += "💪"
      }
      if i % 7 == 6 {
         answer+= "\n"
      }
   }
   fmt.Fprintln(writer, answer)
}
