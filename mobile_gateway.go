package modica

import (
	"fmt"
	"strconv"
)

const (
	baseMessagePath          = "messages"
	baseBroadcastMessagePath = baseMessagePath + "/broadcast"
)

// MobileGatewayService implements modica's Mobile Gateway HTTPS API v2
type MobileGatewayService service

// CreateMessage sends an (outbound) message to a single destination.
func (m MobileGatewayService) CreateMessage(newMessage *Message) (messageID int, err error) {
	req, err := m.client.newRequest(methodPost, baseMessagePath, newMessage)
	if err != nil {
		return
	}

	// Parse the message ID from the response body
	var resMessageID []int
	_, err = m.client.do(req, &resMessageID)

	// If a message ID exists, return it.
	if len(resMessageID) > 0 {
		return resMessageID[0], err
	}

	return messageID, ErrMobileGatewayMessageIDNotFound
}

// GetMessage retrieves a message
func (m MobileGatewayService) GetMessage(messageID int) (message *Message, err error) {
	uri := fmt.Sprintf("%s/%s", baseMessagePath, strconv.Itoa(messageID))
	req, err := m.client.newRequest(methodGet, uri, nil)
	if err != nil {
		return nil, err
	}

	_, err = m.client.do(req, &message)
	return message, err
}

// CreateBroadcastMessage sends an (outbound) message to multiple destinations
func (m MobileGatewayService) CreateBroadcastMessage(newMessage *BroadcastMessage) (message *BroadcastMessage, err error) {
	req, err := m.client.newRequest(methodPost, baseBroadcastMessagePath, newMessage)
	if err != nil {
		return nil, err
	}

	_, err = m.client.do(req, &message)
	return message, err
}

// Message provides the data model to unmarshal and marshal a single message for Modica's mobile gateway API.
type Message struct {
	// ID contains the message ID
	ID int `json:"id,omitempty"`

	// Required Attributes
	// Destination contains the destination mobile number the SMS message will be sent to.
	// Number format must be international format e.g. +64211234567 / +61414123456 / +18123456789.
	Destination string `json:"destination"`

	// Content contains the text to be sent verbatim to the mobile phone.
	Content string `json:"content"`

	// Optional Attributes.
	// Source allows you to select the source short code or mobile number where you have multiple available.
	// This parameter is optional but when it is used it must be used instead of the class parameter.
	Source string `json:"source,omitempty"`

	// Scheduled enables a message to be sent at a scheduled time.
	Scheduled string `json:"scheduled,omitempty"`

	// Reference is unknown.
	Reference string `json:"reference,omitempty"`

	// Class allows you to define the type of message.
	// Classes are used to determine how to route and bill the message, as well as provide useful reporting information.
	// Modica will supply you with set of valid classes available to your account. In most cases this should be mt_message.
	// If the source parameter is used then class must not be included.
	Class string `json:"class,omitempty"`

	// Mask is unknown.
	Mask string `json:"mask,omitempty"`

	// SMSClass is unknown, but must be between 1-3.
	SMSClass int `json:"sms_class,omitempty"`
}

// BroadcastMessage provides the data model to unmarshal and marshal multiple messages for Modica's mobile gateway API.
type BroadcastMessage struct {
	Destination []string `json:"destination"`

	Message
}
