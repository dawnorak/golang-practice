package main

import (
	"fmt"
	"runtime"
	"time"
)

func sums() {
	sum := 1
	for ; sum < 1000; {
		sum += sum
	}
	fmt.Println("The first number post 1000 which is a power of 2 is: ")
	fmt.Println(sum)
}

func while() {
	sum := 1
	for sum < 90 {
		sum += sum
	}
	fmt.Println(sum)
}

func indent() {
	fmt.Print("My Go runs on: ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		fmt.Printf("%s.\n", os)
	}
}

func times() {
	fmt.Println("When's the Weekend?")
	today := time.Now().Weekday()
	switch time.Friday {
	case today + 0:
		fmt.Println("Today :D")
	case today + 1:
		fmt.Println("Tomorrow :)")
	case today + 2:
		fmt.Println("In two days :|")
	default:
		fmt.Println("Too far away :()")
	}
}

func greeting() {
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon!!")
	case t.Hour() < 21:
		fmt.Println("Good evening!!!")
	default:
		fmt.Println("Good night!!! Go to sleep!")
	}
}

func counting() {
	fmt.Println("Anyone can count in order. Let's count in reverse. Let me just think first...")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("Done.")
}

func main() {
	sum := 0
	fmt.Println("Wanna count from 0 to 10?")
	for i := 0; i < 10; i++ {
		fmt.Println(i)
		sum += i
	}
	fmt.Println("Bonus, here is their sum: ", sum)
	sums()
	while()
	indent()
	times()
	greeting()
	counting()
}
