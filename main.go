package main

import (
	"example/bootstrap"
	"fmt"

	_ "example/routers"
)

func main() {
	fmt.Println("main")
	bootstrap.Setup()
}
