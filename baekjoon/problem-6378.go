package main

import (
   "bufio"
   "fmt"
   "os"
   "strconv"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func main() {
   defer writer.Flush()
   
   for {
      var n string
      fmt.Fscan(reader, &n)

      if n == "0" {
         return
      }
      
      a := 0
      for {
         if a > 0 && a < 10 {
            fmt.Fprintln(writer, n)
            break
         }
         a = 0
         for _, v := range n {
            i, err := strconv.Atoi(string(v))
            if err != nil {
               panic(err)
            }
            a += i
         }
         n = strconv.Itoa(a)
      }
   }
}
