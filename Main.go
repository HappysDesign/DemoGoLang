package main

import "fmt"

func pow(x, n int) int {
	tatal := x
	for i := 1; i < n; i++ {
		tatal *= x
		fmt.Println(tatal)
	}
	return tatal
}

func main() {
	result := pow(2, 10)
	fmt.Println(result) // 2^10 = 1024
}
