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
