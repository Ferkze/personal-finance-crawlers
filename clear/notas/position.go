package notas

func updatePositions(daytrades DayTradePositions, swingtrades SwingTradePositions) {
	for k, daytrade := range daytrades {
		if daytrade.Quant == 0 {
			continue
		}
		swing, ok := swingtrades[k]
		if !ok { // Abrir nova posição
			swing = Position{
				Start: daytrade.Start,
				Asset: daytrade.Asset,
				AssetType: daytrade.AssetType,
				FinancialVolume: daytrade.FinancialVolume,
				Price: daytrade.Price,
				Quant: daytrade.Quant,
				QuantityVolume: daytrade.QuantityVolume,
				ShortVolume: daytrade.ShortVolume,
				Total: daytrade.Total,
				Type: SwingTrade,
				Result: 0.0,
			}
		} else { // Aumentar, diminuir ou encerrar posição
			if (daytrade.Quant > 0 && swing.Quant < 0) || (daytrade.Quant < 0 && swing.Quant > 0) {
				swing.Result += calculateResult(swing.Price, daytrade.Price, daytrade.Quant)
			}
			swing.AssetType = daytrade.AssetType
			swing.Type = SwingTrade
			swing.Start = daytrade.Start
			swing.FinancialVolume += daytrade.FinancialVolume
			swing.ShortVolume += daytrade.ShortVolume
			swing.Total += daytrade.Total
			swing.Quant += daytrade.Quant
			swing.Price = calculateAvgPrice(swing.Price, daytrade.Price, swing.Quant, daytrade.Quant)
		}
		swingtrades[k] = swing
		
		daytrade.Quant = 0
		daytrade.Price = 0
		daytrade.Total = 0

		if daytrade.Result == 0 {
			delete(daytrades, k)
		}
	}
}