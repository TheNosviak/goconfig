package main

import (
	"fmt"
	"goconfig"
	"reflect"
)

func main() {
	config := goconfig.NewConfig()


	if err := config.Parse("assets"); err != nil {
		panic(err)
	}

	options, err := config.Options()
	if err != nil {
		panic(err)
	}

	fmt.Println(options.Get(reflect.String, "Hello"))
}