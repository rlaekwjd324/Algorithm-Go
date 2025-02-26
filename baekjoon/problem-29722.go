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
   
   var c string
   var n int
   fmt.Fscan(reader, &c, &n)

   y, err := strconv.Atoi(strings.Split(c, "-")[0])
   if err != nil {
      return
   }
   m, err := strconv.Atoi(strings.Split(c, "-")[1])
   if err != nil {
      return
   }
   d, err := strconv.Atoi(strings.Split(c, "-")[2])
   if err != nil {
      return
   }

   d += n
   if d > 30 {
      m += ((d-1)/30)
      d = d % 30
      if d == 0 {
         d = 30
      }
   }
   if m > 12 {
      y += ((m-1)/12)
      m = m % 12
      if m == 0 {
         m = 12
      }
   }
   mf := fmt.Sprintf("%v", m)
   if m < 10 {
      mf = fmt.Sprintf("0%v", m)
   }
   df := fmt.Sprintf("%v", d)
   if d < 10 {
      df = fmt.Sprintf("0%v", d)
   }
   
   fmt.Fprintf(writer, "%v-%v-%v\n", y, mf, df)
}
