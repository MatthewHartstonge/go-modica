package modica

import (
	"fmt"
	"net/http"
	"testing"
	"reflect"
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
	got, err := client.MobileGateway.CreateMessage(payload)
	if err != nil {
		t.Errorf("MobileGateway.CreateMessage returned error: %v", err)
	}

	want := 123
	if got != want {
		t.Errorf("MobileGateway.CreateMessage returned %+v, want %+v", got, want)
	}
}

func TestMobileGatewayService_GetMessage(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/messages/123", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Authorization", expectedAuthHeader)
		testHeader(t, r, "Accept", mediaTypeV1)

		fmt.Fprint(w, `{"id":123,"destination":"+642123456789","content":"Hi, this is a test message to ensure you are texting correctly","source":"TEST","reference":"alt-reference","operator":"2degrees","reply_to":"123"}`+"\n")
	})

	got, err := client.MobileGateway.GetMessage(123)
	if err != nil {
		t.Errorf("MobileGateway.CreateMessage returned error: %v", err)
	}

	want := &Message{
		ID:          123,
		Destination: "+642123456789",
		Content:     "Hi, this is a test message to ensure you are texting correctly",
		Source:      "TEST",
		Reference:   "alt-reference",
		ReplyTo:     "123",
		Operator:    "2degrees",
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("MobileGateway.GetMessage returned %+v, want %+v", got, want)
	}
}

func TestMobileGatewayService_GetMessage_NotFound(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/messages/321", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Authorization", expectedAuthHeader)
		testHeader(t, r, "Accept", mediaTypeV1)

		w.WriteHeader(404)
	})

	_, err := client.MobileGateway.GetMessage(321)
	if err == nil {
		t.Error("MobileGateway.GetMessage didn't return an error on not found")
	}

	if err != ErrNotFound {
		t.Errorf("MobileGateway.GetMessage returned the wrong error: got: %+v, want %+v", err, ErrNotFound)
	}
}