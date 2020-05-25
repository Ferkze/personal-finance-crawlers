package types

import "time"

// Account credenciais de login para conta da Clear
type Account struct {
	CPF string
	Password string
	DateOfBirth string
}

// Execution representa a execução de uma ordem, podendo existir várias execuções para uma ordem
type Execution struct {
	Asset string
	AssetType string
	Quantity int64
	OrderType string
	Price float64
	Datetime time.Time	
}

// OrderType representa o tipo de ordem de Compra ou de Venda
type OrderType string

// AssetType representa o tipo de ativo negociado, sendo DolarFuturo, IndiceFuturo ou Acoes
type AssetType string