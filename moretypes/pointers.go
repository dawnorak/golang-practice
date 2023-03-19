package main

import (
	"fmt"
)

type Vertex struct {
	X int
	Y int
}

func vert() {
	v:= Vertex{2, 4}
	v.X = 20
	v.Y = 21
	fmt.Println("First number: ", v.X, ", Second number: ", v.Y)
}

func vert2() {
	v := Vertex{1, 2}
	p := &v
	p.X = 1e9
	fmt.Println(v)
}

var (
	v1 = Vertex{1, 2}
	v2 = Vertex{X: 1}
	v3 = Vertex{}
	p  = &Vertex{1, 2}
)

func vert3() {
	fmt.Println(v1, p, v2, v3)
}

func main() {
	i, j := 20, 200

	p := &i  
	fmt.Println(*p)
	*p = 21        
	fmt.Println(i)

	p = &j
	*p = *p / 20
	fmt.Println(j)

	vert()
	vert2()
	vert3()
}