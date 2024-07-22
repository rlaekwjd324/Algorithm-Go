package main

import (
   "bufio"
   "fmt"
   "os"
   "math"
   "strconv"
   "strings"
)

func main() {
   reader := bufio.NewReader(os.Stdin)
   for {
      input, _ := reader.ReadString('\n')
      input = strings.TrimSpace(input)
      if input == "0 0" {
         break
      }
      nums := strings.Split(input, " ")
      b, _ := strconv.Atoi(nums[0])
      n, _ := strconv.Atoi(nums[1])
      num := math.Pow(float64(b), 1.0/float64(n))
      min := math.Ceil(num)
      max := math.Floor(num)
      if math.Abs(float64(b)-math.Pow(float64(min), float64(n))) > math.Abs(float64(b)-math.Pow(float64(max), float64(n))) {
         fmt.Printf("%v\n", int(max))
      }else {
         fmt.Printf("%v\n", int(min))
      }
   }
}
