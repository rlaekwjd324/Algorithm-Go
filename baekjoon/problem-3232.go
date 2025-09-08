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

   var t int
   fmt.Fscan(reader, &t)
   
   for i:=0; i<t; i++ {
      var p, q, r string
      fmt.Fscan(reader, &p, &q, &r)

      num := 0
      for j := 2; j <= 16; j ++ {
         if j < 10 {
            if !valid(r, j) || !valid(r, j) || !valid(r, j) {
               continue
            }
         }
         v := getRealNum(r, j)
         d := getRealNum(p, j)*getRealNum(q, j)
         if v == d {
            num = j
            break
         }
      }
      fmt.Fprintf(writer, "%v\n", num)
   }
}

func getRealNum(r string, j int) int {
   s := 0
   for i, v := range r {
      n := int(v - '0')
      s += (int(math.Pow(float64(j), float64(len(r)-i-1)))*n)
   }
   return s
}

func valid(r string, j int) bool {
   for _, v := range r {
      if int(v-'0') >= j {
         return false
      }
   }
   return true
}
