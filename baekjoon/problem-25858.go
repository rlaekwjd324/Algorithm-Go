package main

import (
   "fmt"
   "os"
   "bufio"
   "strings"
   "strconv"
)

func main() {
   reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
   nums := strings.Split(input, " ")
   x, _ := strconv.Atoi(nums[0])
   y, _ := strconv.Atoi(nums[1])
   sum := 0
   var array []int
   for i := 0; i < x; i++ {
      input, _ = reader.ReadString('\n')
	   input = strings.TrimSpace(input)
	   a, _ := strconv.Atoi(input)
      array = append(array, a)
      sum += a
   }
   for i := 0; i < x; i++ {
      value := y * array[i] / sum
      fmt.Println(value)
   }
}
