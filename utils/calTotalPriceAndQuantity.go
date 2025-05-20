package utils


type Cart struct{
	Name     string `json:"name"`
	Price    float64 `json:"price"`
	Quantity float64 `json:"quantity"`
	Image    string `json:"image"`
	PizzaId  uint `json:"pizzaId"`
}

type Carts []Cart

func CalTotalPriceAndQuantity(carts Carts) (float64, float64) {
	totalQuantity := 0.0
	totalPrice := 0.0
	for _, value := range carts {
		totalQuantity += value.Quantity
		totalPrice += value.Quantity * value.Price
	}

	return totalQuantity, totalPrice
}