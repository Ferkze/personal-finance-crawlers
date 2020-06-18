package notas

import "fmt"

func calculateAvgPrice(p1, p2 float64, q1, q2 int64)  float64 {
	if q1 - q2 == 0 {
		return 0
	}
	
	qf1 := float64(q1)
	qf2 := float64(q2)
	
	a1 := p1 * qf1
	a2 := p2 * qf2

	return (a1 + a2)/ (qf1 + qf2)
}

func calculateResult(p1, p2 float64, q int64) float64 {
	return (p2 * float64(q)) - (p1 * float64(q))
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
		fmt.Printf("-->[ %s ]<--\n", k)
		fmt.Printf("[%s]: %v\n", v.SwingTradeShares.AssetType, v.SwingTradeShares.Value)
		fmt.Printf("[%s]: %v\n", v.DayTradeShares.AssetType, v.DayTradeShares.Value)
		fmt.Printf("[%s]: %v\n", v.DayTradeFutures.AssetType, v.DayTradeFutures.Value)
	}
}