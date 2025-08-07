package main

import (
	"fmt"
	"bufio"
	"os"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func main() {
	defer writer.Flush()

   var t int
   fmt.Fscan(reader, &t)
   
   for i := 0; i < t; i++ {
      var w string
      var k int
      fmt.Fscan(reader, &w, &k)

      max := -1
      min := 10001
      // a~z 반복문으로 돌며 k개 만족하는 문자열 찾기
      for j := 0; j < 26; j++ {
         positions := []int{}
         for idx, tmp := range w {
            if tmp == 'a'+rune(j) {
               positions = append(positions, idx)
            }
         }

         if len(positions) < k {
            continue
         }

         for p := 0; p <= len(positions)-k; p++ {
            start := p
            end := p+k-1
            len := positions[end] - positions[start] + 1

            if len < min {
               min = len
            } 
            if len > max {
               max = len
            }
         }
      }
      
      if max == -1 {
         fmt.Fprintf(writer, "%v\n", -1)
         continue
      }
      fmt.Fprintf(writer, "%v %v\n", min, max)
   }
}
