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
   
   answer := ""
   arr := []string{"", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}
   for {
      var n string
      fmt.Fscanln(reader, &n)
      if n == "halt" {
         break
      }
      count := 0
      preI := 0
      for _, input := range n {
         i, j := getPosition(arr, string(input))
         if preI == i {
            count += 2
         }
         preI = i
         count += j
      }
      answer += fmt.Sprintf("%v\n", count)
   }

   fmt.Fprintln(writer, answer)
}

func getPosition(arr []string, input string) (arrIndex int, strIndex int) {
   for i, str := range arr {
      if strings.Contains(str, input) {
         for j, c := range str {
            if string(c) == input {
               return i, j+1
            }
         }
      }
   }
   return 0, 0
}
