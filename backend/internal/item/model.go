package item

// Item represents a product in the system
type Item struct {
	ID      int     `json:"id"`
	Name    string  `json:"name"`
	Price   float64 `json:"price"`
	IsOffer bool    `json:"is_offer"`
}
