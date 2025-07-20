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

   var g, s, c int
   fmt.Fscan(reader, &g, &s, &c)

   sum := g*3 + s*2 + c
   if sum >= 8 {
      fmt.Fprintf(writer, "%v or %v\n", "Province", "Gold")
      return
   }
   if sum >= 6 {
      fmt.Fprintf(writer, "%v or %v\n", "Duchy", "Gold")
      return
   }
   if sum >= 5 {
      fmt.Fprintf(writer, "%v or %v\n", "Duchy", "Silver")
      return
   }
   if sum >= 3 {
      fmt.Fprintf(writer, "%v or %v\n", "Estate", "Silver")
      return
   }
   if sum >= 2 {
      fmt.Fprintf(writer, "%v or %v\n", "Estate", "Copper")
      return
   }
   fmt.Fprintf(writer, "%v\n", "Copper")
}
