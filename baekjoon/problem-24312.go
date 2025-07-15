package main

import (
   "bufio"
   "fmt"
   "os"
   "sort"
   "math"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func main() {
   defer writer.Flush()

   var arr []int
   for i := 0; i < 4; i++ {
      var a int
      fmt.Fscan(reader, &a)
      arr = append(arr, a)
   }
   sort.Ints(arr)

   min := math.Abs(float64(arr[0]+arr[3]-arr[1]-arr[2]))
   m := math.Abs(float64(arr[0]+arr[1]+arr[2]-arr[3]))
   if min > m {
      min = m
   }
   m = math.Abs(float64(arr[0]-arr[1]-arr[2]-arr[3]))
   if min > m {
      min = m
   }

   fmt.Fprintf(writer, "%v\n", min)
}
