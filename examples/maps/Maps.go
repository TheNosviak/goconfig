package main

import (
	"fmt"
	"goconfig"
	"reflect"
)

func main() {
	pkg := goconfig.NewConfig()
	if err := pkg.Parse("."); err != nil {
		panic(err)
	}

	config, err := pkg.Options()
	if err != nil {
		panic(err)
	}

	Hello, err := config.Get(reflect.String, "Hello")
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s\n", Hello)
}