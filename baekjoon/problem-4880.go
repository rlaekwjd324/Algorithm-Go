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
   answer := ""
   for {
      input, _ := reader.ReadString('\n')
      input = strings.TrimSpace(input)
      nums := strings.Split(input, " ")
      if input == "0 0 0" {
         break
      }
      a, _ := strconv.Atoi(nums[0])
      b, _ := strconv.Atoi(nums[1])
      c, _ := strconv.Atoi(nums[2])
      if b-a == c-b {
         value := c+c-b
         answer += fmt.Sprintf("AP %v\n", value)
      }else {
         value := c+(c-b)/(b-a)*(c-b)
         answer += fmt.Sprintf("GP %v\n", value)
      }
   } 
   fmt.Println(answer)
}
