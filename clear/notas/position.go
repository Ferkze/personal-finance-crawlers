package notas

func updatePositions(daytrades DayTradePositions, swingtrades SwingTradePositions) {
	for k, daytrade := range daytrades {
		if daytrade.Quant == 0 {
			if daytrade.Type == "" {
				daytrade.Type = DayTrade
			}
			daytrades[k] = daytrade
			continue
		}
		// fmt.Printf("Position [%s]: %v x %v = %v; Result: %v\n", k, daytrade.Price, daytrade.Quant, daytrade.Total, daytrade.Result)
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
				Type: SwingTrade,
				Result: 0.0,
			}
			swing.Total = calculateTotal(swing.Price, swing.Quant)

		} else { // Aumentar, diminuir ou encerrar posição
			if (daytrade.Quant > 0 && swing.Quant < 0) || (daytrade.Quant < 0 && swing.Quant > 0) {
				swing.Result += calculateResult(swing.Price, daytrade.Price, daytrade.Quant)
			}
			swing.AssetType = daytrade.AssetType
			swing.Type = SwingTrade
			swing.Start = daytrade.Start
			swing.FinancialVolume += daytrade.FinancialVolume
			swing.ShortVolume += daytrade.ShortVolume
			swing.Price = calculateAvgPrice(swing.Price, daytrade.Price, swing.Quant, daytrade.Quant)
			swing.Quant += daytrade.Quant
			swing.Total = calculateTotal(swing.Price, swing.Quant)
		}

		// fmt.Printf("Swing [%s]: %v x %v = %v\n", k, swing.Price, swing.Quant, swing.Total)

		swingtrades[k] = swing
		
		daytrade.Quant = 0
		daytrade.Price = 0
		daytrade.Total = 0

		if daytrade.Result == 0 {
			delete(daytrades, k)
		} else {
			daytrade.Type = DayTrade
			daytrades[k] = daytrade
		}
	}
}