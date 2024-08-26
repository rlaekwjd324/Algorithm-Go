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
   answer := ""
   for i := 0; i < t; i++ {
      input, _ := reader.ReadString('\n')
      input = strings.TrimSpace(input)
      nums := strings.Split(input, " ")
      v, _ := strconv.Atoi(nums[0])
      e, _ := strconv.Atoi(nums[1])
      answer += fmt.Sprintf("%v\n", (2+e-v))
   }
   fmt.Println(answer)
}
