package notas

import (
	"strconv"
	"strings"
)

func parseDayTradeIndexFuturesOrders(results Results, positions DayTradePositions, text string) (DayTradePositions) {
	lines = strings.Split(text, "\n")

	res := Result{
		AssetType: IndFut,
	}
	pos, ok := positions["WIN"]
	if !ok {
		pos = Position{
			AssetType: IndFut,
			Asset: "WIN",
		}
	}
	
	date, err := parsePageDate(text)
	if err != nil {
		panic(err.Error())
	}
	res.Date = date
	pos.Start = date

	for _, line := range lines {
		texts := strings.Split(line, " ")
		
		if len(texts) < 2 {
			continue
		}

		if strings.HasPrefix(strings.ToLower(texts[1]), "win") || strings.HasPrefix(strings.ToLower(texts[1]), "ind") {
			priceTxt := texts[5]
			price, _ := strconv.ParseFloat(strings.ReplaceAll(strings.ReplaceAll(priceTxt, ".", ""), ",", "."), 64)

			quantTxt := texts[4]
			quant, _ := strconv.ParseInt(strings.ReplaceAll(strings.ReplaceAll(quantTxt, ".", ""), ",", "."), 10, 64)

			positionType := texts[0]
			
			total := price * float64(quant) / 5

			if positionType == "C" {
				pos.Total -= total
				res.Value -= total
			}
			if positionType == "V" {
				pos.Total += total
				res.Value += total
				res.ShortVolume += total
			}
			res.QuantityVolume += quant
			res.FinancialVolume += total
		}
	}

	results = appendResult(results, res)
	positions = appendPosition(positions, pos)

	return positions
}


func parseDayTradeDolarFuturesOrders(results Results, positions DayTradePositions, text string) (DayTradePositions) {
	lines = strings.Split(text, "\n")

	res := Result{
		AssetType: DolFut,
	}
	pos, ok := positions["WDO"]
	if !ok {
		pos = Position{
			AssetType: DolFut,
			Asset: "WDO",
		}
	}
	
	date, err := parsePageDate(text)
	if err != nil {
		panic(err.Error())
	}
	res.Date = date
	pos.Start = date

	for _, line := range lines {
		texts := strings.Split(line, " ")
		
		if len(texts) < 2 {
			continue
		}

		if strings.HasPrefix(strings.ToLower(texts[1]), "wdo") || strings.HasPrefix(strings.ToLower(texts[1]), "dol") {
			priceTxt := texts[5]
			price, _ := strconv.ParseFloat(strings.ReplaceAll(strings.ReplaceAll(priceTxt, ".", ""), ",", "."), 64)

			quantTxt := texts[4]
			quant, _ := strconv.ParseInt(strings.ReplaceAll(strings.ReplaceAll(quantTxt, ".", ""), ",", "."), 10, 64)

			positionType := texts[0]
			
			total := price * float64(quant) * 10

			if positionType == "C" {
				pos.Total -= total
				res.Value -= total
			}
			if positionType == "V" {
				pos.Total += total
				res.Value += total
				res.ShortVolume += total
			}
			res.QuantityVolume += quant
			res.FinancialVolume += total
		}
	}

	results = appendResult(results, res)
	positions = appendPosition(positions, pos)
	
	return positions
}

func appendResult(results Results, res Result) Results {
	date := res.Date.Format("2006-01-02")
	_, ok := results[date]
	if !ok {
		results[date] = make([]Result, 0)
	}
	results[date] = append(results[date], res)
	return results
}

func appendPosition(positions map[string]Position, pos Position) map[string]Position {
	positions[pos.Asset] = pos
	return positions
}