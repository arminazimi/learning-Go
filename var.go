package main

import "fmt"

// type represents
//type float float64

// const represents
const (
	lightspeed int = 300000
)

// iota represents
const (
	j = iota
	r
	k
)

func main() {

	var a int = 3
	var txt string = "string"
	var b = -1.2
	c := 43
	lightspeed := lightspeed + 1

	fmt.Println(a)
	fmt.Println(txt)
	fmt.Println(b)
	fmt.Println(c)

	s, d, g := 1, 2, 3
	fmt.Println(s, d, g, lightspeed)

	var (
		x, f, l = 3, 2, 1.5
	)
	var z = x + int(l)
	var q = float64(x) + l

	fmt.Println(x, f, l, z, q)
	fmt.Println(j, r, k)

	//pointer represents
	ss, jj := 42, 2701

	p := &ss        // point to i
	fmt.Println(*p) // read i through the pointer     خوده مقدار
	fmt.Println(p)  // print memory address   ادرس مقدار
	*p = 21         // set i through the pointer
	fmt.Println(ss) // see the new value of i

	p = &jj      // point to j
	*p = *p / 37 // divide j through the pointer
	fmt.Println(j)

}
