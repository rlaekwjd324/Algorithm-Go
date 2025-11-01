package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

func main() {
    reader := bufio.NewReader(os.Stdin)
    writer := bufio.NewWriter(os.Stdout)
    defer writer.Flush()

    var n int
    fmt.Fscan(reader, &n) // 개행 제거

    for i := 0; i < n; i++ {
        line, _ := reader.ReadString('\n')
        line = strings.TrimSpace(line)

        words := strings.Fields(line)
        count := make(map[rune]int)

        for _, w := range words {
            first := rune(w[0])
            if first >= 'a' && first <= 'z' {
                count[first]++
            }
        }

        maxLetter := 'z' + 1
        maxCount := 0
        for ch := 'a'; ch <= 'z'; ch++ {
            if count[ch] > maxCount || (count[ch] == maxCount && ch < maxLetter) {
                maxLetter = ch
                maxCount = count[ch]
            }
        }

        fmt.Fprintf(writer, "%c\n", maxLetter)
    }
}
