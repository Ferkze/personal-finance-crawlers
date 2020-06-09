package notas

import "time"

// AssetType tipo de ativo
type AssetType string

const (
	// IndFut Índice futuro
	IndFut AssetType = "IndFut"
	// DolFut Dólar futuro
	DolFut AssetType = "DolFut"
	// Shares Ações
	Shares AssetType = "Shares"
)

// OperationType tipo de ativo
type OperationType string

const (
	// DayTrade DayTrade
	DayTrade OperationType = "DayTrade"
	// SwingTrade SwingTrade
	SwingTrade OperationType = "SwingTrade"
)

// Position Position
type Position struct {
	Start     time.Time
	Type      OperationType
	Asset     string
	Quant     int64
	Price     float64
	Total     float64
	AssetType AssetType
}

// DayTradePositions Positions
type DayTradePositions map[string]Position

// SwingTradePositions Positions
type SwingTradePositions map[string]Position

// Result Result
type Result struct {
	Date            time.Time
	QuantityVolume  int64
	FinancialVolume float64
	ShortVolume     float64
	Value           float64
	AssetType       AssetType
}

// Results all results of closed trades
type Results map[string]Result

// Total calcula o preço * quantidade
// func (pos *Position) Total() float64 {
// 	switch pos.AssetType {
// 	case "Shares":
// 		return pos.Price * float64(pos.Quant)
// 	case "IndFut":
// 		return pos.Price * float64(pos.Quant) / 5
// 	case "DolFut":
// 		return pos.Price * float64(pos.Quant) * 10
// 	default:
// 		return 0
// 	}
// }
