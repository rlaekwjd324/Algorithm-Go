package main

import (
   "bufio"
   "fmt"
   "os"
   "strings"
   "strconv"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func main() {
   defer writer.Flush()

   var n int
   fmt.Fscan(reader, &n)
   
   h := 0
   m := 0
   s := 0
   for i := 0; i < n; i++ {
      var str string
      fmt.Fscan(reader, &str)
      tm, _ := strconv.Atoi(strings.Split(str, ":")[0])
      ts, _ := strconv.Atoi(strings.Split(str, ":")[1])
      m += tm
      s += ts
   }
   m += s/60
   s = s%60
   h += m/60
   m = m%60
   
   sh := fmt.Sprintf("%v", h)
   if h < 10 {
      sh = "0"+sh
   }
   sm := fmt.Sprintf("%v", m)
   if m < 10 {
      sm = "0"+sm
   }
   ss := fmt.Sprintf("%v", s)
   if s < 10 {
      ss = "0"+ss
   }

   fmt.Fprintf(writer, "%v:%v:%v\n", sh, sm, ss)
}
