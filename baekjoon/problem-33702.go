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
   
   var k int
   fmt.Fscan(reader, &k)
   
   arr := []int{8, 0, 8, 0, 8, 0, 8, 0, 8}
   
   fmt.Fprintf(writer, "%v\n", arr[k-1])
}
