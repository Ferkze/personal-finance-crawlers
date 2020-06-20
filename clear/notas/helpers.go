package notas

import (
	"fmt"
	"math"
)

func calculateAvgPrice(p1, p2 float64, q1, q2 int64)  float64 {
	if q1 + q2 == 0 {
		return 0
	}
	
	qf1 := float64(q1)
	qf2 := float64(q2)
	
	a1 := p1 * qf1
	a2 := p2 * qf2

	return math.Round((a1 + a2)*100/ (qf1 + qf2)) /100
}

func calculateTotal(p float64, q int64) float64 {
	return math.Round(p * float64(q) *100)/100
}

func calculateResult(p1, p2 float64, q int64) float64 {
	return math.Round(((p2 * float64(q)) - (p1 * float64(q)))*100)/100
}

// func appendResult(results Results, res Result) Results {
// 	date := res.Date.Format("2006-01-02")
// 	_, ok := results[date]
// 	if !ok {
// 		results[date] = make([]Result, 0)
// 	}
// 	results[date] = append(results[date], res)
// 	return results
// }

func appendPosition(positions map[string]Position, pos Position) map[string]Position {
	positions[pos.Asset] = pos
	return positions
}


func printMap(data map[string]interface{}) {
	for k, v := range data {
		fmt.Printf("[%s]: %v\n", k, v)
	}
}

func printPositions(data DayTradePositions) {
	fmt.Println("Printing positions")
	for k, v := range data {
		fmt.Printf("[%s]: %v (%v x %v)\n", k, v.Total, v.Quant, v.Price)
	}
}

func printResults(data Results) {
	fmt.Println("Printing results")
	for k, v := range data {
		fmt.Printf("Data %s\n", k)
		if v.SwingTradeShares.AssetType != "" {
			fmt.Printf("[Operações Normais %s %s]: %v\n", v.SwingTradeShares.MarketType, v.SwingTradeShares.AssetType, v.SwingTradeShares.Value)
		}
		if v.DayTradeShares.AssetType != "" {
			fmt.Printf("[Day Trade %s %s]: %v\n", v.DayTradeShares.MarketType, v.DayTradeShares.AssetType, v.DayTradeShares.Value)
		}
		if v.DayTradeFuturesDolar.AssetType != "" {
			fmt.Printf("[Day Trade %s %s]: %v\n", v.DayTradeFuturesDolar.MarketType, v.DayTradeFuturesDolar.AssetType, v.DayTradeFuturesDolar.Value)
		}
		if v.DayTradeFuturesIndex.AssetType != "" {
			fmt.Printf("[Day Trade %s %s]: %v\n", v.DayTradeFuturesIndex.MarketType, v.DayTradeFuturesIndex.AssetType, v.DayTradeFuturesIndex.Value)
		}
	}
}