package models

import "time"

type Transaction struct {
	ID        int64     `json:"id"`
	UserID    int       `json:"user_id" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	User      User      `json:"user"`
	Status    string    `json:"status"`
	Total     int       `json:"total"`
	Carts     []Cart    `json:"carts"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

type TransactionResponse struct {
	ID     int64 `json:"id"`
	UserID int   `json:"user_id"`
}

func (TransactionResponse) TableName() string {
	return "transactions"
}
