package main

import (
	"fmt"
	"math/rand"
	"math"
	"math/cmplx"
)

func sqrt() {
	fmt.Println("Now for a tale about problems...")
	fmt.Printf("I have %g problems.\n", math.Sqrt(10000))
	fmt.Println("But you are not one of them.")
}

func pi() {
	fmt.Println("You want to eat a pie? The only pi I know is: ", math.Pi)
}

func add(x, y int) int {
	return x + y
}

func nine() {
	fmt.Println("Do you know what 9 + 11 is?")
	fmt.Println("The answer is: ", add(10, 11))
	fmt.Println("Lol.")	
}

func swap(x, y string) (string, string) {
	return y, x
}

func swapper() {
	a, b := swap("Vaigarai", "Dawn")
	fmt.Println("Wanna know my name? It's Dawn. ", a, b)
}

func split(sum int) (x, y int) {
	x = sum * 4 / 8
	y = sum - x
	return
}

func splitter() {
	fmt.Println("Wanna split 20 between us? You will be surprised what you find.")
	fmt.Println(split(20))
}

var a, b, c bool

func bools() {
	var i int
	fmt.Println("Wanna know what I think of your proposition? ")
	fmt.Println(i,".", a,". ", b, ". ", c, ".")
}

func mambo() {
	var i, j int = 1, 2
	k := 3
	var x, y int = 4, 5
	a, b, c := true, false, "Mambo No. 5!!!"

	fmt.Println(i, j, k, x, y, a, b, c)
}

var (
	ToBe   bool       = true
	MaxInt uint64     = 1<<60 - 1
	Cmplx  complex128 = cmplx.Sqrt(-5 + 12i)
)

func types() {
	fmt.Printf("You want a '%T' value? Here: %v\n", ToBe, ToBe)
	fmt.Printf("You want a '%T' value? Here: %v\n", MaxInt, MaxInt)
	fmt.Printf("You want a '%T' value? Here: %v\n", Cmplx, Cmplx)
}

func conversion() {
	var x, y int = 13, 26
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	fmt.Println("Ever wonder what the square root of 13 square plus 26 square is? I don't promise decimal numbers.")
	fmt.Println(x, y, z)
}

func char() {
	const bye = "再见"
	fmt.Println("Every wonder how they say Goodbye in Chinese? Don't worry, I got you covered. Here it is: ", bye)

	const Truth = true
	fmt.Println("Am I the best?")
	fmt.Println("Engine says the above statement is 100 percent the ", Truth)
}

const (
	Big = 2 << 100
	Small = Big >> 99
)

func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 {
	return x * 0.1
}

func bigint() {
	fmt.Println("Want a small number? Here: ", needInt(Small))
	fmt.Println("Want a small float number? Here: ", needFloat(Small))
	fmt.Println("Want a big number? Here: ", needFloat(Big))
}

func main() { 
	fmt.Println("My first favorite number is: ", rand.Intn(20))
	fmt.Println("My second favorite number is: ", rand.Intn(10))
	sqrt()
	pi()
	nine()
	swapper()
	splitter()
	bools()
	mambo()
	types()
	conversion()
	char()
	bigint()
}
