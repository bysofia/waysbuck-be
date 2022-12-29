package transactiondto

import "waysbucks/models"

type CreateTransaction struct {
	ID     int64  `json:"id"`
	UserID int    `json:"user_id" form:"user_id"`
	Status string `json:"status"`
}

type UpdateTransaction struct {
	UserID int    `json:"user_id" form:"user_id"`
	Status string `json:"status"`
	Total  int    `json:"total"`
}

type TransactionResponse struct {
	UserID    int                       `json:"user_id" form:"user_id"`
	ProductID int                       `json:"product_id" form:"product_id"`
	ToppingID int                       `json:"topping_id" form:"topping_id"`
	Cart      []models.CartResponse     `json:"order"`
	Product   models.ProductTransaction `json:"product"`
}
