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
   
	line, _ := reader.ReadString('\n')
   one := "abdegopq@ADOPQR"
   two := "B"
   c := 0
   for _, v := range line {
      if strings.Contains(two, string(v)) {
         c += 2
      } else if strings.Contains(one, string(v)) {
         c += 1
      }
   }
   fmt.Fprintln(writer, c)
}
