package main

import (
	"fmt"
	"math"
	"math/rand"
)

func add(x int, y int) int {
	return x + y;
}

func sub(x, y int) int {
	return x - y;
}

func swap(x, y string) (string, string) {
	return y, x;
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9;
	y = sum - x;
	return;
}

var c, python, java bool;

func main() {
	fmt.Println("My favorite number is", rand.Intn(10))
	fmt.Printf("Now you have %g problems\n", math.Sqrt(7));
	fmt.Println(math.Pi);
	fmt.Println(add(5,4));
	fmt.Println(sub(4,2));
	a, b := swap("Hello", "World");
	fmt.Println(a, b);
	fmt.Println(split(17));

	var i int;
	fmt.Println(i, c, python, java);
}