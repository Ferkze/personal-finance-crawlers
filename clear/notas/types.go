package notas


// AssetPositionBalance AssetPositionBalance
type AssetPositionBalance struct {
	Start         time.Time
	Total         float64
	OperationType string
}

// AssetPosition
type AssetPosition map[string]AssetPositionBalance
