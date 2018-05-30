package main

import (
	"fmt"
	"strconv"

	"github.com/matthewhartstonge/go-modica"
)

func main() {
	client := modica.NewClient("clientID", "clientSecret", nil)

	newBroadcastMessage := &modica.BroadcastMessage{
		Destinations: []string{"+642123456789", "+64987654321"},
		Message: modica.Message{
			Content: "Hello, this is a test message!",
		},
	}
	messageStatuses, err := client.MobileGateway.CreateBroadcastMessage(newBroadcastMessage)
	if err != nil {
		panic(err)
	}

	for _, status := range messageStatuses {
		switch status.Status {
		case modica.MessageStatusSubmitted:
			fmt.Printf("Yass! Message %s has been %s", strconv.Itoa(status.ID), status.Status)
		case modica.MessageStatusSent:
			fmt.Printf("Okay, we got to wait, but at least Message %s has been %s", strconv.Itoa(status.ID), status.Status)
		case modica.MessageStatusReceived:
			fmt.Printf("Message %s has been %s", strconv.Itoa(status.ID), status.Status)
		case modica.MessageStatusFrozen:
			fmt.Printf("Oh no! Message %s is stuck in terrible disney movie, also known as %s", strconv.Itoa(status.ID), status.Status)
		case modica.MessageStatusRejected, modica.MessageStatusFailed, modica.MessageStatusDead, modica.MessageStatusExpired:
			fmt.Printf("Oh no! Message %s has %s", strconv.Itoa(status.ID), status.Status)
		default:
			fmt.Printf("Well, this is awkward.. Message %s's status is unknown, but we at least got this: %s", strconv.Itoa(status.ID), status.Status)
		}
	}

	fmt.Printf("message: %#+v\n", messageStatuses)
}
