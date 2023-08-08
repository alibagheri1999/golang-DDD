package handler

type Handlers struct {
	GiftCartService GiftCart
	GeneralService  General
}

func New(GiftService GiftCart, GeneralService General) *Handlers {
	return &Handlers{
		GiftCartService: GiftService,
		GeneralService:  GeneralService,
	}
}
