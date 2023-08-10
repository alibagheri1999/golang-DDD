package test_test

import (
	"remote-task/interfaces/http/handler"
	"remote-task/utilities"
)

var (
	giftCartApp utilities.GiftCartAppInterface
	userApp     utilities.UserAppInterface

	g = handler.NewGiftCartHandler(&giftCartApp, &userApp)
)
