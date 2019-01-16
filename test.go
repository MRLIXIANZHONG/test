package main

import "fmt"

func main() {
	a := []int{1, 2, 5, 6, 7}
	//b := []int{1,2,5,6}
	//b :=remove(a,2)
	a = append(a[:2], a[3:]...)
	b := append(a[:2], a[3:]...)

	fmt.Println(a, b)
}
