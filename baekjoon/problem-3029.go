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
   
   var a, b string
   fmt.Fscan(reader, &a, &b)

   if a == b {
      fmt.Fprint(writer, "24:00:00")
      return
   }
   
   ah, _ := strconv.Atoi(strings.Split(a, ":")[0])
   am, _ := strconv.Atoi(strings.Split(a, ":")[1])
   as, _ := strconv.Atoi(strings.Split(a, ":")[2])
   bh, _ := strconv.Atoi(strings.Split(b, ":")[0])
   bm, _ := strconv.Atoi(strings.Split(b, ":")[1])
   bs, _ := strconv.Atoi(strings.Split(b, ":")[2])

   s := bs - as
   if s < 0 {
      s += 60
      bm--
   }
   m := bm - am
   if m < 0 {
      m += 60
      bh--
   }
   h := bh - ah
   if h < 0 {
      h += 24
   }
   if h < 10 {
      fmt.Fprint(writer, "0")
   }
   fmt.Fprint(writer, h)
   fmt.Fprint(writer, ":")
   if m < 10 {
      fmt.Fprint(writer, "0")
   }
   fmt.Fprint(writer, m)
   fmt.Fprint(writer, ":")
   if s < 10 {
      fmt.Fprint(writer, "0")
   }
   fmt.Fprint(writer, s)
}
