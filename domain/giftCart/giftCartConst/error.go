package giftCartConst

import "errors"

var (
	ERR_CREATE_GIFT_CART = errors.New("failed to create a gift card")
	ERR_NOT_FOUND        = errors.New("gift card not found")
	ERR_UPDATE_GIFT_CART = errors.New("failed to update a gift card")
)
