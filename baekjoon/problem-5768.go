package main

import (
	"fmt"
	"bufio"
	"os"
   "math"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func main() {
	defer writer.Flush()

   for {
      var m, n int
      fmt.Fscan(reader, &m, &n)

      if m == 0 && n == 0 {
         break
      }

      answer := m
      maxCnt := 0
      for i:=m; i<=n; i++ {
         cnt := getDivisorCnt(i)
         if maxCnt <= cnt {
            maxCnt = cnt
            answer = i
         }
      }

      fmt.Fprintf(writer, "%v %v\n", answer, maxCnt)
   }
}

func getDivisorCnt(n int) int {
   cnt := 0
   for i:=1; i<=int(math.Sqrt(float64(n))); i++ {
      if n % i != 0 {
         continue
      }
      cnt++
      if i != n/i {
         cnt++
      }
   }
   return cnt
}
