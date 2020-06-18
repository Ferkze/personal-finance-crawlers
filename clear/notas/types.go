package notas

import "time"

// AssetPositionBalance AssetPositionBalance
type AssetPositionBalance struct {
	Start         time.Time
	Total         float64
	OperationType string
}

// AssetsPosition balanço de posições em diferentes ativos
type AssetsPosition map[string]AssetPositionBalance

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
	Start           time.Time
	Type            OperationType
	Asset           string
	Quant           int64
	Price           float64
	Total           float64
	AssetType       AssetType
	Value           float64
	ShortVolume     float64
	QuantityVolume  int64
	FinancialVolume float64
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
type Results map[string]DailyResult

// DailyResult report diário de resultado
type DailyResult struct {
	SwingTradeShares Result
	DayTradeShares   Result
	DayTradeFutures  Result
}
