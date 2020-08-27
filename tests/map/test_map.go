package main

import "fmt"

func main() {

	//frist initilize
	var mapp map[string]string
	mapp = make(map[string]string)
	mapp["name"] = "sssssss"
	fmt.Println(mapp)

	//secend initilize
	var mappp map[string]int = make(map[string]int)
	mappp["age"] = 19
	fmt.Println(mappp)

	//thrid initilize : auto derivation
	mapppp := map[string]string{"name": "sssss", "age": "19"}
	fmt.Println(mapppp)

	age_value, ok := mapppp["age"]

	if !ok {
		fmt.Println("这个key不存在")
	} else {
		fmt.Println(age_value)
	}

	for K, _ := range mapppp {
		fmt.Println(K)
	}

}
