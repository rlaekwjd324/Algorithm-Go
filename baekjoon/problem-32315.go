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

   var a string
   fmt.Fscan(reader, &a)

   arr := strings.Split(a, "")

   var numbers []int
   for i := 0; i < 10; i++ {
      numbers = append(numbers, 0)
   }
   
   for i := 0; i < len(arr); i++ {
      if arr[i] == "-" {
         continue
      }
      num, _ := strconv.Atoi(arr[i])
      numbers[num]++
   }

   count := 0
   for i := 0; i < 10; i++ {
      if numbers[i] > 0 {
         count++
      }
   }

   fmt.Fprintln(writer, count)
}
