package main

import "fmt"

func main() {
	var m map[string]int
	m = make(map[string]int)
	//add
	m["k1"] = 7
	m["k2"] = 1
	m["k3"] = 6
	m["k4"] = 4
	fmt.Println("map : ", m)
	//access
	v1 := m["k1"]
	fmt.Println("v1 : ", v1)
	//len
	fmt.Println("lenM : ", len(m))
	//check
	_, ok := m["k1"]
	fmt.Println(ok)
	//iterate
	for k, v := range m {
		fmt.Println("key : ", k, "value : ", v)
	}
	//delete
	delete(m, "k1")
	fmt.Println("map : ", m)

}
