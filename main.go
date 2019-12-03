package main

import "fmt"

func timesTwo(n int) (int, int) {
	return n, n * 2
}

func main() {
	n, ns := timesTwo(8)
	fmt.Println(n, ns)
}
