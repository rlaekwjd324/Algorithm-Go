package main

import (
   "fmt"
   "os"
   "bufio"
   "strings"
   "strconv"
)

func main() {
   reader := bufio.NewReader(os.Stdin)
	 input, _ := reader.ReadString('\n')
	 input = strings.TrimSpace(input)
   nums := strings.Split(input, " ")
   w, _ := strconv.Atoi(nums[0])
   h, _ := strconv.Atoi(nums[1])
   t := float64(w*h)/2.0
   fmt.Printf("%0.1f\n",t)
}
