# Basics

## Packages
Every Go program is made up of packages.

Programs start running in package main.

## Imports
groups imports into a parenthesized, "factored" import statement

## Exported names
In Go, a name is exported if it begins with a capital letter.

When importing a package, you can refer only to its exported names. Any "unexported" names are not accessible from outside the package.

## Basic types
Go's basic types are

```
bool

string

int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64 uintptr

byte // alias for uint8

rune // alias for int32
     // represents a Unicode code point

float32 float64

complex64 complex128
```

## Type conversions
The expression T(v) converts the value v to the type T.

## Constants
Constants are declared like variables, but with the const keyword.

## For loop
Go has only one looping construct, the for loop.

The basic for loop has three components separated by semicolons:

the init statement: executed before the first iteration
the condition expression: evaluated before every iteration
the post statement: executed at the end of every iteration
The init statement will often be a short variable declaration, and the variables declared there are visible only in the scope of the for statement.

The init and post statements are optional.

For is Go's "while".

Forever: If you omit the loop condition it loops forever, so an infinite loop is compactly expressed.
```
func main() {
    for {
    }
}
```

## Switch
Switch cases evaluate cases from top to bottom, stopping when a case succeeds.

## Defer
A defer statement defers the execution of a function until the surrounding function returns.

The deferred call's arguments are evaluated immediately, but the function call is not executed until the surrounding function returns.

## Pointers
Go has pointers. A pointer holds the memory address of a value.

The type `*T` is a pointer to a `T` value. Its zero value is nil.
> var p *int
The `&` operator generates a pointer to its operand.
> i := 42
> p = &i
The `*` operator denotes the pointer's underlying value.
> fmt.Println(*p) // read i through the pointer p
> *p = 21         // set i through the pointer p

This is known as "dereferencing" or "indirecting".
Unlike C, Go has no pointer arithmetic.

## Structs
A struct is a collection of fields.

Struct fields are accessed using a dot.

Struct fields can be accessed through a struct pointer whth`(*p).X`. to simplify, just write `p.X`, without th explicit dereferece.

## Arrays
the type `[n]T` is an array of `n` values of type T.
> var a [10] int   // declares a variable `a` as an array of ten ingegers.
An array's length is part of its type, so arrays cannot be resized.

## Slices
A slice is a dynamically-sized.
The type `[]T` is a slice with elemeents of type T.

changing the elements of a slice modifies the corresponding elements of its underlying array.

slice literals:
> []bool{true, true, false}
This creates the arrary `[3]bool{true, true, false}`, then builds a slice that references it.

A slice has both a length and a capacity.

The length of a slice is the number of elements it contains.

The capacity of a slice is the number of elements in the underlying array, counting from the first element in the slice.

The zero value of a slice is nil.

Slices can be created with the built-in mak function; this is how you create dynamically-sized arrays.
> b := make([]int, 0, 5)  // len(b)=0, cap(b)=5

Go provides a built-in `append` function.
> func append(s []T, vs ...T) []T

## Range
the range form of the for loop iterates over a slice or map.


## Maps
A map maps keys to values.

The zero value of a map is nil. 

Insert or update an element in map `m`:
> m[key] = elem

retrieve an element:
> elem = m[key]

delete an element:
> delete(m, key)

test that a key is present with a two-value assignment, if key is in m , ok is true, else ok is false and elem is the zero value for the map's element type. :
> elem, ok := m[key]

## Function closures
Go functions may be closures. A closure is a function value that references variables from outside its body. The function may access and assign to the referenced variables; in this sense the function is "bound" to the variables.


## Methods
Go does not have classes. Howerver, you can define methods on types.

A method is a function with a special `receiver` argument.

The receiver appears in its own argument list between the `func` keyword and the method name.

```
package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
}
```

Remember: a method is just a function with a receiver argument.

You can declare a method on non-struct types, too.

You can only declare a method with a receiver whose type is defined in the same package as the method. You cannot declare a method with a receiver whose type is defined in another package (which includes the built-in types such as int).

### pointer receiver
two reasons to use a pointer receiver:
* the method can modify the value that its receiver points to.
* to avoid copying the value on each method call. 


## Interfaces
An `interface type` is defined as a set of method signatures.

A value of interface type can hold any value that implements these methods.

Interfaces are implemented implicitly. A type implements an interface by implementing its methods. There is no explicit declaration of intent.


## Reader
The io package specifies the io.Reader interface, which represents the read end of a stream of data.

The `io.Reader` interface has a Read method:
> func (T) Read(b []byte) (n int, err error)

## Type parameters
Go functions can be written to work on multiple types using type parameters. The type parameters of a function appear between brackets, before the function's arguments.

> func Index[T comparable](s []T, x T) int

This declaration means that s is a slice of any type T that fulfills the built-in constraint comparable. x is also a value of the same type.

comparable is a useful constraint that makes it possible to use the == and != operators on values of the type. In this example, we use it to compare a value to all slice elements until a match is found. This Index function works for any type that supports comparison.


## Goroutines
A goroutine is a lightweight thread managed by the Go runtime.
> go f(x, y, z)
start a new goroutine running

## Channels
Channels are a typed conduit through which you can send and receive values with the channel operator, `<-`. // communication among goroutines.
```
ch <- v    // Send v to channel ch.
v := <- ch // Receive from ch, and assign value to v.
```

channels must be created before use:
> ch := make(chan int)

By default, sends and receives block until the other side is ready. This allows goroutines to synchronize without explict locks or condition variables.

### Buffered Channels
> ch := make(chan int, 100)
Sends to a buffered channel block only when the buffer is full. Receives blocks when the buffer is empty.

### Range and close
A sender can close a channel to indicate that not more values will be sent. Receivers can test whether a channel has been closed by assigning a second parameter to the receive expression
> v, ok := <- ch

ok is false if there are no more values to receive and the channel is closed.

The loop `for i := range c` receives values from the channel repeatedly until it is closed.


## Select
The select statement lets a goroutine wait on multiple communication operations.

A select blocks until one of its case can run, then it executes that case. It choosees one at random if multiple are ready.