package main

import (
   "bufio"
	"fmt"
	"os"
   "strings"
)

func main() {
   scanner := bufio.NewScanner(os.Stdin)
   for scanner.Scan() {
      line := scanner.Text()
      if line == "" {
         break
      }

      ungarbled := strings.ReplaceAll(line, "iiing", "th")

		fmt.Println(ungarbled)
   }
}
