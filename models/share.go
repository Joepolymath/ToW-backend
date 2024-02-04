package models

import "time"

type Share struct {
	ID		uint
	OwnerID	uint
	Price		float64
	ForSale	bool
	TransactionHistory []Transaction
	CreatedAt	time.Time
	UpdatedAt	time.Time
}

type Transaction struct {
	ID        uint `gorm:"primaryKey"`
	BuyerID   uint
	SellerID  uint
	ShareID   uint
	Price     float64
	Timestamp time.Time
}

type AdminSettings struct {
	ID                   uint `gorm:"primaryKey"`
	MinimumFollowSpend  float64
	ReserveFund          float64
	CreatedAt            time.Time
	UpdatedAt            time.Time
}
