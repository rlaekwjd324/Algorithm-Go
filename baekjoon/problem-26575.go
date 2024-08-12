package main

import (
   "fmt"
   "bufio"
   "strings"
   "os"
   "strconv"
   "math"
)

func main() {
   var t int
   fmt.Scanf("%v", &t)

	reader := bufio.NewReader(os.Stdin)
   for i := 0; i < t; i++ {
      input, _ := reader.ReadString('\n')
      input = strings.TrimSpace(input)
      nums := strings.Split(input, " ")
      a := nums[0]
      b := nums[1]
      c := nums[2]
      if strings.Split(a, "")[0] == "." {
         a = fmt.Sprintf("0%v", a)
      }
      aNum, _ := strconv.ParseFloat(a, 64)
      if strings.Split(b, "")[0] == "." {
         b = fmt.Sprintf("0%v", b)
      }
      bNum, _ := strconv.ParseFloat(b, 64)
      if strings.Split(b, "")[0] == "." {
         c = fmt.Sprintf("0%v", c)
      }
      cNum, _ := strconv.ParseFloat(c, 64)
      var value float64
      value = aNum * bNum * cNum
      fmt.Printf("$%.2f\n", toFixed(value, 2))
   }
}

func round(num float64) int {
   return int(num + math.Copysign(0.5, num))
}

func toFixed(num float64, precision int) float64 {
   output := math.Pow(10, float64(precision))
   return float64(round(num * output)) / output
}
