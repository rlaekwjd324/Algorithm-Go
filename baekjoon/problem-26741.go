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
   
   var x, y int
   fmt.Fscan(reader, &x, &y)

   // a+b = x
   // a*2 + 4*b = y
   cow := (y - 2 * x)/2
   chicken := x - cow
   fmt.Fprintf(writer, "%v %v\n", chicken, cow)
}
