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

   word := ""
   for {
      var n string
      _, err := fmt.Fscan(reader, &n)
      if err != nil {
         break
      }
      word = n
   }
   if word == "eh?" {
      fmt.Fprintln(writer, "Canadian!")
      return
   }
   fmt.Fprintln(writer, "Imposter!")
}
