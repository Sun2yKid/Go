package main

import (
	"fmt"
	"math"
	"math/cmplx"
	"math/rand"
	"runtime"
	"strings"
	"time"
)

var c, python, java bool
var i, j int = 1, 3

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

type Vertex struct {
	X int
	Y int
}

var (
	v1 = Vertex{1, 2}     // has type Vertex
	v2 = Vertex{X: 1}     // Y:0 is implicit
	v3 = Vertex{}         // X:0 and Y:0
	p_v2 = &Vertex{1, 2}  // has type *Vertex
)

func main() {
	fmt.Println("My favorite number is", rand.Intn(10))
	fmt.Println(add(58, 85))

	// multiple results
	a, b := swap("hello", "world")
	fmt.Println(a, b)

	// named return values
	fmt.Println(split(19))

	// variables
	var i int
	fmt.Println(i, c, python, java)

	var c, python, java = true, false, "no!"
	fmt.Println(i, j, c, python, java)

	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)

	var i2 int
	var f float64
	var b2 bool
	var s string
	fmt.Printf("%v %v %v %q\n", i2, f, b2, s)

	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)

	// if statement
	fmt.Println(sqrt(2), sqrt(-4))

	// if with a short statement
	fmt.Println(pow(3, 2, 10))
	fmt.Println(pow(3, 3, 20))

	fmt.Println(Sqrt(3))

	// switch
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.", os)
	case "linux":
		fmt.Println("Linux.", os)
	default:
		fmt.Printf("%s.\n", os)
	}

	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days")
	default:
		fmt.Println("Too far away.")
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good moring!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}

	// defer
	defer fmt.Println("world")
	fmt.Println("Hello")

	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("done")

	// pointer
	i, j := 42, 2701
	p := &i
	fmt.Println(i, p, *p)
	*p = 21
	fmt.Println(i, p, *p)

	p2 := &j
	fmt.Println(j, p2, *p2)
	*p2 = *p2 / 37
	fmt.Println(j, p2, *p2)

	// struct
	fmt.Println(Vertex{1, 2})
	v := Vertex{1, 2}
	v.X = 4
	fmt.Println(v)

	p_v := &v
	fmt.Println(p_v, (*p_v).X, p_v.X)
	p_v.Y = 1e9
	fmt.Println(v)

	fmt.Println(v1, p_v2, v2, v3)

	// arrays
	var a1 [2]string
	a1[0] = "hello"
	a1[1] = "world"
	fmt.Println(a1, a1[0], a1[1])

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)

	// slices
	var s1 []int = primes[1:4]
	fmt.Println(s1)

	names := [4]string{
		"John",
		"Paul",
		"George H",
		"Ringo",
	}
	fmt.Println(names)

	a3 := names[0:2]
	b3 := names[1:3]
	fmt.Println(a3, b3)

	b3[0] = "XXX"
	fmt.Println(a3, b3, names)

	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)
	printSlice(q)

	// slice the slice to give it zero length
	q = q[:0]
	printSlice(q)

	// extend its length
	q = q[:4]
	printSlice(q)

	// Drop its first two values
	q = q[2:]
	printSlice(q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	s2 := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(s2)

	// nil slices
	var s3 []int
	printSlice(s3)
	if s3 == nil {
		fmt.Println("nil! slice")
	}

	// create slice with make
	printSlice(make([]int, 5))
	printSlice(make([]int, 0, 5))
	printSlice(make([]int, 0, 5)[:2])
	printSlice(make([]int, 0, 5)[2:5])

	//slice of slice
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i :=0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}

	// append to a slice
	var s_a []int
	fmt.Println(s_a == nil)
	printSlice(s_a)

	// append works on nil slices
	s_a = append(s_a, 0)
	printSlice(s_a)

	// slice grows as needed.
	s_a = append(s_a, 1)
	printSlice(s_a)

	s_a = append(s_a, 2, 3, 4)
	printSlice(s_a)

}

func add(x, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	return lim
}

func Sqrt(x float64) float64 {
	z := 1.0
	for i := 1; i <= 10; i++ {
		// Newton's method
		fmt.Println(i, z)
		z -= (z*z - x) / (2 * z)
	}
	return z
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
