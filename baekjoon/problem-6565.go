package main

import (
   "bufio"
	"fmt"
	"os"
   "strings"
   "strconv"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func main() {
   defer writer.Flush()

   answer := ""
   for {
      var n string
      fmt.Fscanln(reader, &n)

      if n == "0+0=0" {
         answer += "True\n"
         break
      }

      a := strings.Split(n, "+")[0]
      a = Reverse(a)
      aa := strings.Split(n, "+")[1]
      b := strings.Split(aa, "=")[0]
      b = Reverse(b)
      c := strings.Split(aa, "=")[1]
      c = Reverse(c)
      
      ai, _ := strconv.Atoi(a)
      bi, _ := strconv.Atoi(b)
      ci, _ := strconv.Atoi(c)

      if ai+bi == ci {
         answer += "True\n"
      } else {
         answer += "False\n"
      }
   }
   
   fmt.Fprintln(writer, answer)
}

func Reverse(s string) string {
   runes := []rune(s)
   for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
       runes[i], runes[j] = runes[j], runes[i]
   }
   return string(runes)
}
