package main

import (
	"fmt"
	"sync"

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
	var wg sync.WaitGroup
	wg.Add(len(rs))
	for _, r := range rs {
		go func(res rest.Resource) {
			defer wg.Done()
			all, err := res.GetAll()
			if err != nil {
				panic(err)
			}
			fmt.Printf("Number of resources: %d\n", len(all))
			for _, val := range all {
				fmt.Printf("%+v\n", val)
			}
			fmt.Println()
		}(r)
	}
	wg.Wait()
}

func main() {
	fmt.Println("Starting")
	printAll()
	fmt.Println("Done")

}
