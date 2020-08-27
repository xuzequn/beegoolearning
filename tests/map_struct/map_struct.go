package main

import "fmt"

type User struct {
	Id   int
	Name string
}

func main() {
	mapp := make(map[int]User)
	mapp[1] = User{Id: 1, Name: "sssss"}
	mapp[2] = User{Id: 2, Name: "sssss"}
	mapp[3] = User{Id: 3, Name: "sssss"}
	mapp[4] = User{Id: 4, Name: "sssss"}
	fmt.Println(mapp)

}
