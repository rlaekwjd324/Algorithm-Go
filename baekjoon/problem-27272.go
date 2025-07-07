package main

import (
   "bufio"
   "fmt"
   "os"
   "sort"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func main() {
   defer writer.Flush()

   arr := []int{}
   for i := 0; i < 4; i++ {
      var a int
      fmt.Fscan(reader, &a)
      
      arr = append(arr, a)
   }

   sort.Ints(arr)

   fmt.Fprintf(writer, "%v\n", arr[0]*arr[1]+arr[2]*arr[3])
}
