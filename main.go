package main

import (
	"fmt"
)

type addable interface {
	add(a addable) addable
}

func adder[T int | float64](a, b T) T {
	return a + b
}

type myInt struct {
	val int
}

func (i *myInt) add(i2 *myInt) myInt {
	i.val += i2.val
	return *i
}

type myFloat64 struct {
	val int
}

func (i *myFloat64) add(i2 *myFloat64) myFloat64 {
	i.val += i2.val
	return *i
}

func main() {
	i1 := 10
	i2 := 69
	adder(i1, i2)
	fmt.Printf("value 1 := %v\n", i1)
	f1 := 10.0
	f2 := 69.0
	adder(f1, f2)
	fmt.Printf("value 1 := %v\n", f1)
}
