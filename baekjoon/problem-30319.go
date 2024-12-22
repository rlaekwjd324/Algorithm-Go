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

   var n string
   fmt.Fscan(reader, &n)
   m, _ := strconv.Atoi(strings.Split(n, "-")[1])
   d, _ := strconv.Atoi(strings.Split(n, "-")[2])
   
   if m == 9 && d <= 16 {
      fmt.Fprintln(writer, "GOOD")
      return
   }
   if m < 9 {
      fmt.Fprintln(writer, "GOOD")
      return
   }

   fmt.Fprintln(writer, "TOO LATE")
}
