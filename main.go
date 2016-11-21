package main

import (
	"fmt"

	"github.com/coxley/gonsot/rest"
)

func printAll() {
	var rs = make([]rest.Resource, 8)
	rs = []rest.Resource{
		&rest.Site{},
		&rest.Attribute{},
		&rest.Device{},
		&rest.Circuit{},
		&rest.Interface{},
		&rest.Network{},
		&rest.User{},
	}
	for _, r := range rs {
		go func(res rest.Resource) {
			all, err := res.GetAll()
			if err != nil {
				panic(err)
			}
			fmt.Printf("Number of resources: %d", len(all))
			for val := range all {
				fmt.Printf("%+v\n", val)
			}
			fmt.Println()
		}(r)
	}
}

func main() {
	fmt.Println("Starting")
	printAll()
	fmt.Println("Done")

}
