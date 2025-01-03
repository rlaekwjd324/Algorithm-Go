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

   arr := make([][]string, 10)
   for i := 0; i < 10; i++ {
      arr[i] = make([]string, 10)
   }
   for i := 0; i < 10; i++ { 
      for j := 0; j < 10; j++ {
         var n string
         fmt.Fscan(reader, &n)
         arr[i][j] = n
      }
   }

   for i := 0; i < 10; i++ {
      if isSameLine(arr[i]) {
         fmt.Fprintln(writer, 1)
         return
      }
      temp := make([]string, 10)
      for j := 0; j < 10; j++ {
         temp[j] = arr[j][i]
      }
      if isSameLine(temp) {
         fmt.Fprintln(writer, 1)
         return
      }
   }
   for i := 0; i < 10; i++ {
      if isSameLine(arr[i]) {
         fmt.Fprintln(writer, 1)
         return
      }
   }

   fmt.Fprintln(writer, 0)
}

func isSameLine(arr []string) bool {
   v := arr[0]
   for _, l := range arr {
      if v != l {
         return false
      }
   }
   return true
}
