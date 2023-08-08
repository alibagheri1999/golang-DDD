package handler

import (
	"context"
	"remote-task/domain/giftCart/DTO"
)

type GiftCartService interface {
	SendGiftCart(c context.Context, req *DTO.SendGiftCartRequest) error
}
