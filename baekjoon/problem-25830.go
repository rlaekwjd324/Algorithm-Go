package main

import (
   "bufio"
   "fmt"
   "os"
   "strconv"
   "strings"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func main() {
   defer writer.Flush()

   var n string
   fmt.Fscan(reader, &n)

   mm, _ := strconv.Atoi(strings.Split(n, ":")[0])
   ss, _ := strconv.Atoi(strings.Split(n, ":")[1])
   s := 0-ss
   m := ss-mm
   h := mm
   if s < 0 {
      s+=60
      m--
   }
   if m < 0 {
      m+=60
      h--
   }

   nh := fmt.Sprintf("%v", h)
   nm := fmt.Sprintf("%v", m)
   ns := fmt.Sprintf("%v", s)
   if h < 10 {
      nh = "0"+nh
   }
   if m < 10 {
      nm = "0"+nm
   }
   if s < 10 {
      ns = "0"+ns
   }
   fmt.Fprintf(writer, "%v:%v:%v\n", nh, nm, ns)
}
