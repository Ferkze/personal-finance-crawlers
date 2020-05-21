package types

import "time"

type Account struct {
	CPF string
	Password string
	DateOfBirth string
}

type Execution struct {
	Asset string
	AssetType string
	Quantity int64
	OrderType string
	Price float64
	Datetime time.Time	
}