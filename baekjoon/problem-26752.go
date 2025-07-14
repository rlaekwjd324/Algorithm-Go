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

   var h, m, s int
   fmt.Fscan(reader, &h, &m, &s)

   s++
   if s >= 60 {
      s-=60
      m++
   }
   if m >= 60 {
      m-=60
      h++
   }
   if h == 24 {
      h-=24
   }
   ss, sm, sh := fmt.Sprintf("%v", s), fmt.Sprintf("%v", m), fmt.Sprintf("%v", h)
   if s < 10 {
      ss = "0"+ss 
   }
   if m < 10 {
      sm = "0"+sm
   }
   if h < 10 {
      sh = "0"+sh
   }
   fmt.Fprintf(writer, "%v:%v:%v\n", sh, sm, ss)
}
