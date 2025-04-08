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
   
   var fs, cs, es, bs int
   fmt.Fscan(reader, &fs, &cs, &es, &bs)
   var fn, cn, en, bn int
   fmt.Fscan(reader, &fn, &cn, &en, &bn)
   var q int
   fmt.Fscan(reader, &q)
   
   c := 0
   for i := 0; i < q; i++ {
      var a, i int
      fmt.Fscan(reader, &a, &i)

      switch(a) {
         case 1:
            if fs >= i*fn && cs >= i*cn && es >= i*en && bs >= i*bn {
               fs -= i*fn
               cs -= i*cn
               es -= i*en
               bs -= i*bn
               c += i
               fmt.Fprintln(writer, c)
            } else {
               fmt.Fprintln(writer, "Hello, siumii")
            }
         case 2:
            fs += i
            fmt.Fprintln(writer, fs)
         case 3:
            cs += i
            fmt.Fprintln(writer, cs)
         case 4:
            es += i
            fmt.Fprintln(writer, es)
         case 5:
            bs += i
            fmt.Fprintln(writer, bs)
      }
   }
}
