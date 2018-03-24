package modica

import "errors"

const (
	// constant error codes as returned by the API
	errCodeSendFailed     = "send_failed"
	errCodeInvalidJson    = "invalid_json"
	errCodeMissingAttrib  = "missing_attrib"
	errCodeInvalidAttrib  = "invalid_attrib"
	errCodeBroadcastLimit = "broadcast_limit"
	errCode400            = "400"
	errCode422            = "422"
)

// Modica Mobile Gateway Errors
var (
	// send_failed - Could not queue message due to an unknown error
	ErrMobileGatewaySendFailed = errors.New("could not queue message due to an unknown error")

	// invalid_json - Invalid JSON data in the request body
	ErrMobileGatewayInvalidJSON = errors.New("invalid json data in the request body")

	// missing_attrib - Missing a required attribute
	ErrMobileGatewayMissingAttribute = errors.New("missing a required attribute")

	// invalid_attrib - Invalid attribute value
	ErrMobileGatewayInvalidAttribute = errors.New("invalid attribute value")

	// broadcast_limit - Broadcast limit has been exceeded, please consult the error description for more detail.
	ErrMobileGatewayBroadcastLimit = errors.New("broadcast limit has been exceeded")

	// 400 - Invalid scheduled timestamp (must be RFC3339)
	ErrMobileGatewayInvalidTimestampFormat = errors.New("invalid scheduled timestamp (must be rfc3339)")

	// 422 - Invalid scheduled timestamp (must not be in the past)
	ErrMobileGatewayInvalidTimestamp = errors.New("invalid scheduled timestamp (must not be in the past)")

	// ErrMobileGatewayMessageIDNotFound is returned when a message id is not returned from the API, but the request to create a new message was successful.
	ErrMobileGatewayMessageIDNotFound = errors.New("message id not found")
)

var mobileGatewayErrorMap = map[string]error{
	errCodeSendFailed:     ErrMobileGatewaySendFailed,
	errCodeInvalidJson:    ErrMobileGatewayInvalidJSON,
	errCodeMissingAttrib:  ErrMobileGatewayMissingAttribute,
	errCodeInvalidAttrib:  ErrMobileGatewayInvalidAttribute,
	errCodeBroadcastLimit: ErrMobileGatewayBroadcastLimit,
	errCode400:            ErrMobileGatewayInvalidTimestampFormat,
	errCode422:            ErrMobileGatewayInvalidTimestamp,
}
