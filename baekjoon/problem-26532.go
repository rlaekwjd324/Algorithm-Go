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
   reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
   nums := strings.Split(input, " ")
   a, _ := strconv.Atoi(nums[0])
   b, _ := strconv.Atoi(nums[1])
   answer := float64(a*b)/5.0/4840.0
   fmt.Println(math.Ceil(answer))
 }
