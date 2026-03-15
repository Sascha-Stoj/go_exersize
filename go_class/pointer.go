package main

import "fmt"

// https://www.youtube.com/watch?v=2XEQsJLsLN0
func main() {
	i, j := 42, 2701
	fmt.Printf("i=%d, j=%d\n", i, j)   //=> normal value
	fmt.Printf("i=%d, j=%d\n", &i, &j) //=> pointer value

	w := *&i
	fmt.Printf("w=%d\n", w) //=> normal value

	p := &i
	fmt.Printf("p=%d\n", p)

	r := *p
	fmt.Printf("r=%d\n", r) //=> normal value

	a := 4
	squareVal(a)
	fmt.Println(&a, a) //=> normal value
	squareAddr(&a)
	fmt.Println(&a, a) //=> normal value

}

func squareVal(a int) {
	a *= a
	fmt.Println(&a, a)
}

// a is a pointer to an int, so we need to dereference it to get the value
// we want to modify a here so it is also modified in the caller
func squareAddr(a *int) {
	*a *= *a
	fmt.Println(&a, a)
}
