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
   for i := 0; i < n; i++ {
      year := 0
      input, _ := reader.ReadString('\n')
      arr := strings.Split(input, "")
      for j := 0; j < len(arr); j++ {
         switch arr[j] {
         case "I":
            year += 1
         case "V":
            year += 5
         case "X":
            year += 10
         case "L":
            year += 50
         case "C":
            year += 100
         case "D":
            year += 500
         case "M":
            year += 1000
         }
      }
      answer += fmt.Sprintf("%v\n", year)
   }
   fmt.Fprintln(writer, answer)
}
