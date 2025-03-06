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
   
   var a, b, c, d int
   fmt.Fscan(reader, &a, &b, &c, &d)

   c += d
   b += (c / 60)
   c = c % 60
   a += (b / 60)
   b = b % 60
   a = a % 24
   
   fmt.Fprintf(writer, "%v %v %v\n", a, b, c)
}
