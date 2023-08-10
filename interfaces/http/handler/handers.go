package handler

type Handlers struct {
	GiftCartService GiftCart
	GeneralService  General
}

// New create new handler instance
func New(GiftService GiftCart, GeneralService General) *Handlers {
	return &Handlers{
		GiftCartService: GiftService,
		GeneralService:  GeneralService,
	}
}
