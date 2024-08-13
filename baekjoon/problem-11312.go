package main

import (
   "fmt"
   "bufio"
   "strings"
   "os"
   "strconv"
)

func main() {
   var t int
   fmt.Scanf("%v", &t)
	 reader := bufio.NewReader(os.Stdin)
   for i := 0; i < t; i++ {
      input, _ := reader.ReadString('\n')
      input = strings.TrimSpace(input)
      nums := strings.Split(input, " ")
      a, _ := strconv.Atoi(nums[0])
      b, _ := strconv.Atoi(nums[1])
      num := a*a / (b*b)
      fmt.Printf("%v\n", num)
   }
}
