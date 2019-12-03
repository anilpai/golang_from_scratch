package main

import "fmt"

func timesTwo(n int) int {
	return n * 2
}

func main() {
	n := timesTwo(8)
	fmt.Println(n)
}
