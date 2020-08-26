package main

import "fmt"

type UserStruct struct {
	Id   int
	Name string
}

func main() {
	user := UserStruct{Id: 1, Name: "sssss"}
	fmt.Println(user)

	arr1 := [4]int{1, 2, 3, 5}
	fmt.Println(arr1)

	arr2 := [4]string{"2222", "dsafdas", "ssda", "sadas"}
	fmt.Println(arr2)

	arr_struct := [4]UserStruct{{Id: 1, Name: "22222"}, {Id: 1, Name: "22222"}, {Id: 1, Name: "22222"}, {Id: 1, Name: "22222"}}
	fmt.Println(arr_struct)

	for _, v := range arr_struct {
		fmt.Println(v.Id)
	}

}
