package main

import (
	"fmt"
	"github.com/MatthewHartstonge/go-modica"
)

func main() {
	client := modica.NewClient("clientID", "clientSecret", nil)

	newMessage := &modica.Message{
		Destination: "+642123456789",
		Content:     "Hello, this is a test message!",
	}
	msgID, err := client.MobileGateway.CreateMessage(newMessage)
	if err != nil {
		panic(err)
	}

	fmt.Printf("message: %#+v\n", msgID)
}
