package main

import (
   "fmt"
   "bufio"
   "strings"
   "os"
   "strconv"
)

func main() {
   reader := bufio.NewReader(os.Stdin)
	 input, _ := reader.ReadString('\n')
	 input = strings.TrimSpace(input)
   nums := strings.Split(input, " ")
   a, _ := strconv.Atoi(nums[0])
   b, _ := strconv.Atoi(nums[2])
   c, _ := strconv.Atoi(nums[4])
   if a+b == c {
      fmt.Println("YES")
      return
   }
   fmt.Println("NO")
 }
