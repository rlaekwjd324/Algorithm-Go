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
      t, _ := strconv.Atoi(input)
      if t == 0 {
         break
      }
      input, _ = reader.ReadString('\n')
      input = strings.TrimSpace(input)
      nums := strings.Split(input, " ")
      var arr = []int{}
      var max = 0
      for i := 0; i < t; i++ {
         a, _ := strconv.Atoi(nums[i])
         arr = append(arr, a)
         count := 0
         sum := 0
         for j := i; j >= 0; j-- {
            count++
            if count > 3 {
               break
            }
            sum += arr[j]
         }
         if max < sum {
            max = sum
         }
      }
      answer += fmt.Sprintf("%v\n", max)
   }
   
   fmt.Println(answer)
}
