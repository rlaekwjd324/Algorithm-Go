package main

import (
   "bufio"
	"fmt"
	"os"
   "strings"
   "math"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func main() {
   defer writer.Flush()

   var n1, n2, n3 float64
   var c1, c2 string
   fmt.Fscan(reader, &n1, &c1, &n2, &c2, &n3)

   answer := calculate(n1, n2, c1)
   answer = calculate(answer, n3, c2)
   var n float64
   if answer <= 0 {
      n = 0.0
   } else {
      n = math.Round(answer*1000)/1000.0
   }
   num := fmt.Sprintf("%.3f", n)

   fmt.Fprintln(writer, "=================\n|SASA CALCULATOR|")
   fmt.Fprintf(writer, "|%s%s|\n", strings.Repeat(" ", 15-len(num)), num)
   fmt.Fprintln(writer, "-----------------\n|               |\n| AC         /  |\n| 7  8  9    *  |\n| 4  5  6    -  |\n| 1  2  3    +  |\n|    0  .    =  |\n=================")
}

func calculate(n1, n2 float64, c string) float64 {
   switch c {
   case "+":
      return n1 + n2
   case "-":
      return n1 - n2
   case "*":
      return n1 * n2
   case "/":
      return n1 / n2
   }
   return 0.0
}
