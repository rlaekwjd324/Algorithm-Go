package main

import (
   "bufio"
   "fmt"
   "os"
   "math"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func main() {
   defer writer.Flush()

   var a string
   fmt.Fscan(reader, &a)
   
   sum := 0
   for i, v := range a {
      sum += (int(math.Pow(26, float64(len(a)-i-1)))*(int(v)-64))
   }
   fmt.Fprintf(writer, "%v\n", sum)
}
