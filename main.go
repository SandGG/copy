package main

import "fmt"

func main() {
	var s = make([]int, 3)
	var n = copy(s, []int{0, 1, 2, 3, 4})
	fmt.Println(n, s)
	s = append(s, 3, 4)
	var m = copy(s, []int{3, 4})
	fmt.Println(m, s)

	var b = make([]byte, 5)
	copy(b, "Insert text here")
	fmt.Println(b)
}
