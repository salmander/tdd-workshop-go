package shop

type Till interface {
	createBasket() string
	addProduct(BasketId, ean string, quantity int) error
	removeProduct(BasketId, ean string, quantity int) error
	getBasket(BasketId string) (Basket, error)
}

type Basket struct {
}
