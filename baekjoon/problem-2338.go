package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var aStr, bStr string
	fmt.Fscan(reader, &aStr)
	fmt.Fscan(reader, &bStr)

	a := new(big.Int)
	b := new(big.Int)

	a.SetString(aStr, 10)
	b.SetString(bStr, 10)

	// A + B
	sum := new(big.Int).Add(a, b)

	// A - B
	diff := new(big.Int).Sub(a, b)

	// A * B
	prod := new(big.Int).Mul(a, b)

	fmt.Println(sum)
	fmt.Println(diff)
	fmt.Println(prod)
}
