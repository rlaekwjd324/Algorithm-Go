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

   var t1, t2 string
   fmt.Fscan(reader, &t1, &t2)

   h1, _ := strconv.Atoi(strings.Split(t1, ":")[0])
   m1, _ := strconv.Atoi(strings.Split(t1, ":")[1])
   h2, _ := strconv.Atoi(strings.Split(t2, ":")[0])
   m2, _ := strconv.Atoi(strings.Split(t2, ":")[1])
   m := m2-m1
   if m < 0 {
      m += 60
      h2--
   }
   h := h2-h1
   h += 24

   ms := fmt.Sprintf("%v", m)
   if m < 10 {
      ms = "0"+ms
   }
   hs := fmt.Sprintf("%v", h)
   if h < 10 {
      hs = "0"+hs
   }
   fmt.Fprintf(writer, "%v:%v\n", hs, ms)
}
