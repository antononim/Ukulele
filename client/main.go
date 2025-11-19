package main

import "fmt"

type thing struct {
	Name string
	Age  int
}

func main() {

	thingy1 := thing{
		Name: "hi",
		Age:  0,
	}
	for i := 0; i < 10; i++ {
		fmt.Println(thingy1)
	}
}
