package main

import (
   "bufio"
	"fmt"
	"os"
   "math"
   "strings"
   "strconv"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func main() {
   defer writer.Flush()

   var n1, n2, n3 int
   fmt.Fscan(reader, &n1, &n2, &n3)
   
   binaryNumber := getBinaryNumber(n1, 4)+getBinaryNumber(n2, 4)+getBinaryNumber(n3, 4)
   decimalNumber := getDecimalNumber(binaryNumber)
   
   fmt.Fprintf(writer, "%s%v\n", strings.Repeat("0", 4-len(strconv.Itoa(decimalNumber))),decimalNumber)
}

func getBinaryNumber(n int, count int) string {
   num := ""
   for i := 0; i < count; i++ {
      num = fmt.Sprintf("%v%s", n%2, num)
      n = n/2
   }
   return num
}

func getDecimalNumber(n string) int {
   num := float64(0)
   arr := strings.Split(n, "")
   for i := len(n)-1; i >= 0; i-- {
      ni, _ := strconv.Atoi(arr[len(n)-i-1])
      num += float64(ni)*math.Pow(2, float64(i))
   }
   return int(num)
}
