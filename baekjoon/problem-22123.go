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

   var n int
   fmt.Fscan(reader, &n)
   
   for i := 0; i < n; i++ {
      var s, f string
      var k int
      fmt.Fscan(reader, &s, &f, &k)

      h1, _ := strconv.Atoi(strings.Split(s, ":")[0])
      m1, _ := strconv.Atoi(strings.Split(s, ":")[1])
      s1, _ := strconv.Atoi(strings.Split(s, ":")[2])
      h2, _ := strconv.Atoi(strings.Split(f, ":")[0])
      m2, _ := strconv.Atoi(strings.Split(f, ":")[1])
      s2, _ := strconv.Atoi(strings.Split(f, ":")[2])

      ss := s2-s1
      if ss < 0 {
         ss += 60
         m2--
      }
      sm := m2-m1
      if sm < 0 {
         sm += 60
         h2--
      }
      sh := h2-h1
      if sh < 0 {
         sh += 24
      }
      sm += sh*60
      if sm == 0 && ss == 0 {
         sm = 24 * 60
      }
      if sm >= k {
         fmt.Fprintf(writer, "%v\n", "Perfect")
         continue
      }
      if sm+60 >= k {
         fmt.Fprintf(writer, "%v\n", "Test")
         continue
      }
      fmt.Fprintf(writer, "%v\n", "Fail")
   }
}
