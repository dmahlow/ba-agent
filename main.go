package main

import (
	"fmt"

	"github.com/dmahlow/go-bytearena/client"
)

func main() {
	client, err := client.New()
	if err != nil {
		panic(err)
	}

	err = client.Start(&Agent{})
	if err != nil {
		panic(err)
	}

	fmt.Println("End")
}
