package main

import "fmt"

func main() {
	a := []int{1,2,5,6,7}
	//b := []int{1,2,5,6}
	//b :=remove(a,2)
	a = append(a[:2],a[3:]...)
	fmt.Println(a)
}

func remove(a []int,i int)[]int  {
	copy(a[i:],a[i+1:])
	return a[:len(a)-1]

}