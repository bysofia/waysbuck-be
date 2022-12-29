package models

type Cart struct {
	ID            int                 `json:"id" gorm:"primary_key:auto_increment"`
	QTY           int                 `json:"qty"`
	SubTotal      int                 `json:"subtotal"`
	ProductID     int                 `json:"product_id" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Product       ProductTransaction  `json:"product"`
	TransactionID int                 `json:"transaction_id"`
	Transaction   TransactionResponse `json:"transaction"`
	ToppingID     []int               `json:"topping_id" gorm:"-"`
	Topping       []Topping           `json:"topping" gorm:"many2many:cart_toppings;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Status        string              `json:"status"`
}

type CartResponse struct {
	ID            int                `json:"id"`
	Total         int                `json:"total"`
	TransactionID int                `json:"transaction_id"`
	ProductID     int                `json:"product_id"`
	Product       ProductTransaction `json:"product"`
	ToppingID     []int              `json:"topping_id" gorm:"-"`
	Topping       []Topping          `json:"topping" gorm:"many2many:cart_toppings"`
}

func (CartResponse) TableName() string {
	return "carts"
}
