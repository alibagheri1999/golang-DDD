package handler

type Handlers struct {
	GiftCartService GiftCart
}

func New(GiftService GiftCart) *Handlers {
	return &Handlers{
		GiftCartService: GiftService,
	}
}
