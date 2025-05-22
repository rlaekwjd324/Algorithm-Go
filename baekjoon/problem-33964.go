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

   var x, y int
   fmt.Fscan(reader, &x, &y)

   xs := strings.Repeat("1", x)
   ys := strings.Repeat("1", y)

   xi, _ := strconv.Atoi(xs)
   yi, _ := strconv.Atoi(ys)

   sum := xi+yi
   
   fmt.Fprintf(writer, "%v\n", sum)
}
