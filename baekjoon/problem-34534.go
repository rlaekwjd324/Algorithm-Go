package main

import (
	"fmt"
	"bufio"
	"os"
    "strings"
    "strconv"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func main() {
	defer writer.Flush()

   cards := ""
   for i := 0; i < 5; i++ {
      var a string
      fmt.Fscan(reader, &a)

      cards += (a+" ")
   }

   min := 5
   for i := 1; i <= 5; i++ {
      s := 0
      for j := 0; j < 5; j++ {
         if !strings.Contains(cards, strconv.Itoa(i+j)) {
            s++
         }
      }
      if s < min {
         min = s
      }
   }
   
   fmt.Fprintf(writer, "%v\n", min)
}
