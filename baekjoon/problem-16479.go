package main

import (
   "fmt"
   "bufio"
   "strings"
   "os"
   "strconv"
)

func main() {
   var k int
   fmt.Scanf("%v", &k)
   reader := bufio.NewReader(os.Stdin)
	 input, _ := reader.ReadString('\n')
	 input = strings.TrimSpace(input)
   nums := strings.Split(input, " ")
   d1, _ := strconv.Atoi(nums[0])
   r1 := float64(d1)/2.0
   d2, _ := strconv.Atoi(nums[1])
   r2 := float64(d2)/2.0
   fmt.Println(float64(k*k) - (r1-r2)*(r1-r2))
 }
