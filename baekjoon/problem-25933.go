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

   var n int
   fmt.Fscan(reader, &n)
   for i := 0; i < n; i++ {
      arr := make([]int, 6)
      for j := 0; j < 6; j++ {
         var a int
         fmt.Fscan(reader, &a)
         arr[j] = a
         fmt.Fprintf(writer, "%v ", a)
      }
      fmt.Fprint(writer, "\n")
      isCount := arr[0]+arr[1]+arr[2] > arr[3]+arr[4]+arr[5]
      isColor := isColor(arr)
      if isCount && isColor {
         fmt.Fprintf(writer, "%v\n", "both")
      } else if isCount {
         fmt.Fprintf(writer, "%v\n", "count")
      } else if isColor {
         fmt.Fprintf(writer, "%v\n", "color")
      } else {
         fmt.Fprintf(writer, "%v\n", "none")
      }
      fmt.Fprint(writer, "\n")
   }
}

func isColor(arr []int) bool {
   if arr[0] > arr[3] {
      return true
   }
   if arr[0] < arr[3] {
      return false
   }
   if arr[1] > arr[4] {
      return true
   }
   if arr[1] < arr[4] {
      return false
   }
   if arr[2] > arr[5] {
      return true
   }
   return false
}
