package main

import "fmt"

type TestStruct struct {
	Id   int
	Name string
}

func main() {
	// frist initialize
	var test1 TestStruct = TestStruct{Id: 1, Name: "sssss"}
	fmt.Println(test1)

	// secend initialize
	var test2 TestStruct
	test2.Id = 2
	test2.Name = "ssss"
	fmt.Println(test2)

	// third initialize
	test3 := TestStruct{Id: 3, Name: "SSSSSSSS"}
	fmt.Println((test3))
}
