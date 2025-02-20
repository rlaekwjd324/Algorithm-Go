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
   
   var n float64
   fmt.Fscan(reader, &n)
   
   a := n / math.Pi
   a = math.Sqrt(a)
   a = a*2*math.Pi
   a = math.Ceil(a*10)/10
   fmt.Fprintf(writer, "%.1f\n", a)
}
