package main

import "fmt"

func main() {
	//frist initilize
	var arr1 [5]int = [5]int{1, 2, 3, 4}
	fmt.Println(arr1)

	//secend initilize
	arr2 := [5]int{2, 3, 4, 5, 6}
	fmt.Println(arr2)

	//thrid initilize
	arr3 := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	fmt.Println(arr3)

	// frist for
	var i int
	for i = 0; i < len(arr3); i++ {
		fmt.Println(arr3[i])
	}

	for i, v := range arr2 {
		fmt.Println(i)
		fmt.Println(v)
	}

	// for {
	// 	fmt.Println(1)
	// }
}
