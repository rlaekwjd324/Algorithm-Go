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

   arr := make([]int, 4)

   for i := 0; i < 4; i++ {
      var a int
      fmt.Fscan(reader, &a)
      arr[i] = a
   }
   
   fmt.Fprintf(writer, "%v\n", arr[0] * arr[2])
}
