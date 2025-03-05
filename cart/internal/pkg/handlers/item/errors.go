package item

import "errors"

var (
	errorIncorrectUser  = errors.New("incorrect user")
	errorIncorrectSKU   = errors.New("incorrect SKU")
	errorIncorrectCount = errors.New("incorrect count")

	errorMethodNotAllowed = errors.New("method not allowed")
)
