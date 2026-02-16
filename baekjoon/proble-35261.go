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
   var s string
   fmt.Fscan(reader, &n, &s)
  
   t := "eagle"
    min := 5
     for j:=0; j<=n-5; j++ {
         ts:=s[j:]
         c := 0
         for i, _ := range t {
            if t[i] != ts[i] {
               c++
            }
         }
         if min > c {
            min=c
         }
      }

   fmt.Fprintf(writer, "%v\n", min)
}
