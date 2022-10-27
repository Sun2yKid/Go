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

The zro value of a slice is nil.

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

retrieve and element:
> elem = m[key]

delete an element:
> delete(m, key)

test that a key is present with a two-value assignment, if key is in m , ok is true, else ok is false and elem is the zero value for the map's element type. :
> elem, ok := m[key]

## Function closures
Go functions may be closures. A closure is a function value that references variables from outside its body. The function may access and assign to the referenced variables; in this sense the function is "bound" to the variables.

