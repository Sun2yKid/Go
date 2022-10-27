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