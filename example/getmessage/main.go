package main

import (
	"fmt"
	"github.com/MatthewHartstonge/go-modica"
)

func main() {
	client := modica.NewClient("clientID", "clientSecret", nil)
	message, err := client.MobileGateway.GetMessage(123456789)
	if err != nil {
		panic(err)
	}

	fmt.Printf("message: %#+v\n", message)
}
