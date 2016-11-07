package main

import (
	"fmt"
	"github.com/coxley/gonsot/conf"
)

func main() {
	c := conf.Config{}
	err := c.Load()
	if err != nil {
		// fmt.Printf("%v", err)
		panic(err)
	} else {
		fmt.Printf("%v", c)
	}
}
