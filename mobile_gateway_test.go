package modica

import (
	"fmt"
	"net/http"
	"testing"
)

func TestMobileGatewayService_CreateMessage(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/messages", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		testHeader(t, r, "Authorization", expectedAuthHeader)
		testHeader(t, r, "Accept", mediaTypeV1)
		testBody(t, r, `{"destination":"+642123456789","content":"Hi, this is a test message to ensure you are texting correctly","source":"TEST","scheduled":"2017-05-05T10:00:00+12:00","reference":"alt-reference","class":"mt_message","sms_class":2}`+"\n")

		// Modica's REST api returns a non keyed raw array with a single int, representing the message ID if successful.
		fmt.Fprint(w, `[123]`)
	})

	payload := &Message{
		Destination: "+642123456789",
		Content:     "Hi, this is a test message to ensure you are texting correctly",
		Source:      "TEST",
		Scheduled:   "2017-05-05T10:00:00+12:00",
		Reference:   "alt-reference",
		Class:       "mt_message",
		Mask:        "",
		SMSClass:    2,
	}
	messageID, err := client.MobileGateway.CreateMessage(payload)
	if err != nil {
		t.Errorf("MobileGateway.CreateMessage returned error: %v", err)
	}

	want := 123
	if messageID != want {
		t.Errorf("MobileGateway.CreateMessage returned %+v, want %+v", messageID, want)
	}
}
