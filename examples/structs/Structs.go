package main

import (
	"fmt"
	"goconfig"
)

// NOTICE: when using this module you must specify that even when parsing objects from toml & json they must have the `json:""` part https://pkg.go.dev/encoding/json#Unmarshal
type Config struct {
	Hello string `json:"Hello"`
}

func main() {
	pkg := goconfig.NewConfig()
	if err := pkg.Parse("."); err != nil {
		panic(err)
	}

	config, err := pkg.Options()
	if err != nil {
		panic(err)
	}

	item := new(Config)
	if err := config.MarshalEntire(&item); err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s\n", item.Hello)
}