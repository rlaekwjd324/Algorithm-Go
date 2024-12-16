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

   dArr := []string{"social", "history", "language", "literacy"}
   pArr := []string{"bigdata", "public", "society"}

   for {
      var n string
      fmt.Fscan(reader, &n)

      if strings.Contains(strings.Join(dArr, " "), n) {
         fmt.Fprintln(writer, "digital humanities")
         return
      }
      if strings.Contains(strings.Join(pArr, " "), n) {
         fmt.Fprintln(writer, "public bigdata")
         return
      }
   }

}
