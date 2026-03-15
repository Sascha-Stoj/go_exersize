package main

import "fmt"

func main() {
	var s []int
	var t = []int{}
	var u = make([]int, 5)
	var v = make([]int, 0, 5)

	fmt.Printf("%d, %d, %T, %5T, %#[3]v\n", len(s), cap(s), s, s == nil)
	fmt.Printf("%d, %d, %T, %5T, %#[3]v\n", len(t), cap(t), t, t == nil)
	fmt.Printf("%d, %d, %T, %5T, %#[3]v\n", len(u), cap(u), u, u == nil)
	fmt.Printf("%d, %d, %T, %5T, %#[3]v\n", len(v), cap(v), v, v == nil)
}
