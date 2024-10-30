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

   var t int
   fmt.Fscanln(reader, &t)
   
   answer := ""
   for i := 0; i < t; i++ {
      var a, b, c int
      fmt.Fscanln(reader, &a, &b, &c)
      answer += fmt.Sprintf("%v %v %v\n", a, b, c)

      a--
      count := a/b + a/c
      if b < c {
         temp := c
         c = b
         b = temp
      }
      minDup := getMinDup(b, c)
      count -= a/minDup
      
      answer += fmt.Sprintf("%v\n", count)
   }
   fmt.Fprintln(writer, answer)
}

func getMinDup(a int, b int) int {
   dup := a*b
   for {
      if b == 0 {
         break
      }
      r := a%b
      a = b
      b = r
   }
   return dup/a
}
