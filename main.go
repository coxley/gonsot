package main

import (
	"fmt"
	"strings"
	"github.com/coxley/gonsot/conf"
)

func main() {
	c := conf.Config{}
	err := c.Load()
	if err != nil {
		panic(err)
	} else {
		msg := "THIS IS TEMPORARY BINARY FOR DEBUGGING"
		line := strings.Repeat("=", len(msg))
		fmt.Println(msg + "\n" + line + "\n")
		fmt.Printf("Config: %+v\n", c)
	}

	c.Dump("/tmp/testconfig")
}
