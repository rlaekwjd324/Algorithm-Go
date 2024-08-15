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
   b, _ := strconv.Atoi(nums[1])
   c, _ := strconv.Atoi(nums[2])
   d, _ := strconv.Atoi(nums[3])
   hour := c-a
   min := d-b
   if(min < 0) {
      min += 60
      hour = hour-1
   }
   if(hour < 0) {
      hour += 24
   }
   min = min + hour*60
   fmt.Printf("%v %v\n", min, min/30)
}
