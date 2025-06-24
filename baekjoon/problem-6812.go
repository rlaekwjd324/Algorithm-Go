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

   var t int
   fmt.Fscan(reader, &t)

   h := t/100
   m := t%100
   fmt.Fprintf(writer, "%v in Ottawa\n", t)
   h -= 3
   if h < 0 {
      h += 24
   }
   t = h*100+m
   fmt.Fprintf(writer, "%v in Victoria\n", t)
   h++
   if h >= 24 {
      h -= 24
   }
   t = h*100+m
   fmt.Fprintf(writer, "%v in Edmonton\n", t)
   h++
   if h >= 24 {
      h -= 24
   }
   t = h*100+m
   fmt.Fprintf(writer, "%v in Winnipeg\n", t)
   h++
   if h >= 24 {
      h -= 24
   }
   t = h*100+m
   fmt.Fprintf(writer, "%v in Toronto\n", t)
   h++
   if h >= 24 {
      h -= 24
   }
   t = h*100+m
   fmt.Fprintf(writer, "%v in Halifax\n", t)
   m += 30
   if m >= 60 {
      h++
      if h >= 24 {
         h -= 24
      }
      m -= 60
   }
   t = h*100+m
   fmt.Fprintf(writer, "%v in St. John's\n", t)
}
