package main

import (
	"fmt"
	"goconfig"
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

	type Hello struct {
		World string `json:"Hello"`
	}

	var i *Hello = new(Hello)
	if err := options.MarshalFromPath(i, "test"); err != nil {
		panic(err)
	}

	fmt.Println(i)
}