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
   fmt.Fscanln(reader, &n)
   
   arr := []int{0, 0}
   count := 0

   for {
      if n == 0 {
         break
      }
      if count == 0 {
         count = 1
      } else {
         count = 0
      }
      if n % 2 == 0 {
         arr[count] += n/2
         n = n/2
      }else {
         arr[count] += (n/2 + 1)
         n = n/2
      }
   }
   
   fmt.Fprintln(writer, fmt.Sprintf("%v %v", arr[0], arr[1]))
}
