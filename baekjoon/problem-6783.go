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
   
   tCount := 0
   sCount := 0
   for {
      var str string
      fmt.Fscanln(reader, &str)
      if str == "" {
         break
      }

      for _, input := range str {
         switch input {
         case 'T':
            tCount++
         case 't':
            tCount++
         case 'S':
            sCount++
         case 's':
            sCount++
         }
      }
   }

   if tCount > sCount {
      fmt.Fprintln(writer, "English")
      return
   }
   fmt.Fprintln(writer, "French")
}
