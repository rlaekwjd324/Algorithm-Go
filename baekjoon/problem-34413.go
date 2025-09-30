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

   var t string
   fmt.Fscan(reader, &t)

   d := strings.Split(t, "d")
   x, _ := strconv.Atoi(d[0])
   y, _ := strconv.Atoi(strings.Split(d[1], "+")[0])
   z, _ := strconv.Atoi(strings.Split(d[1], "+")[1])
   
   min := x+z
   max := x*y+z
   
   fmt.Fprintf(writer, "%v\n", float64(min+max)/2.0)
}
