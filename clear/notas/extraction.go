package notas

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func extract(text string, pos map[string]Position) {
	if strings.Contains(text, "WIN ") || strings.Contains(text, "IND ") {
		fmt.Println("Parsing Index Futures Day Trades")
		extractDayTradeIndexFuturesOrders(pos, text)
	}
	if strings.Contains(text, "WDO ") || strings.Contains(text, "DOL ") {
		fmt.Println("Parsing Dolar Futures Day Trades")
		extractDayTradeDolarFuturesOrders(pos, text)
	}
	if strings.Contains(text, "1-BOVESPA") {
		fmt.Println("Parsing Shares Swing Trades")
		extractSharesOrders(pos, text)
	}
	
	printPositions(pos)
}

func extractDayTradeIndexFuturesOrders(positions DayTradePositions, text string) (DayTradePositions) {
	lines = strings.Split(text, "\n")

	pos, ok := positions["WIN"]
	if !ok {
		pos = Position{
			AssetType: IndFut,
			Asset: "WIN",
		}
	}
	
	date, err := extractPageDate(text)
	if err != nil {
		panic(err.Error())
	}
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
				pos.Value -= total
			}
			if positionType == "V" {
				pos.Total += total
				pos.Value += total
				pos.ShortVolume += total
			}
			pos.QuantityVolume += quant
			pos.FinancialVolume += total
		}
	}

	positions = appendPosition(positions, pos)

	return positions
}

func extractDayTradeDolarFuturesOrders(positions DayTradePositions, text string) (DayTradePositions) {
	lines = strings.Split(text, "\n")

	pos, ok := positions["WDO"]
	if !ok {
		pos = Position{
			AssetType: DolFut,
			Asset: "WDO",
		}
	}
	
	date, err := extractPageDate(text)
	if err != nil {
		panic(err.Error())
	}
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
				pos.Value -= total
			}
			if positionType == "V" {
				pos.Total += total
				pos.Value += total
				pos.ShortVolume += total
			}
			pos.QuantityVolume += quant
			pos.FinancialVolume += total
		}
	}

	positions = appendPosition(positions, pos)
	
	return positions
}

func extractSharesOrders(positions DayTradePositions, text string) (DayTradePositions) {
	lines = strings.Split(text, "\n")

	date, err := extractPageDate(text)
	if err != nil {
		panic(err.Error())
	}

	for _, line := range lines {
		texts := strings.Split(line, " ")
		
		if len(texts) < 2 {
			continue
		}

		if strings.HasPrefix(strings.ToLower(texts[0]), "1-bovespa"){
			asset := strings.TrimSpace(strings.Join(texts[3:5], " "))
			
			pos, ok := positions[asset]
			if !ok {
				pos = Position{
					Asset: asset,
				}
			}

			priceTxt := texts[len(texts)-3]
			price, _ := strconv.ParseFloat(strings.ReplaceAll(strings.ReplaceAll(priceTxt, ".", ""), ",", "."), 64)
			
			quantTxt := texts[len(texts)-4]
			quant, _ := strconv.ParseInt(strings.ReplaceAll(strings.ReplaceAll(quantTxt, ".", ""), ",", "."), 10, 64)
			
			positionType := texts[1]
			
			total := price * float64(quant)

			if positionType == "C" {
				if pos.Quant < 0 {
					pos.Value += calculateResult(pos.Price, price, quant)
				}
				pos.Price = calculateAvgPrice(pos.Price, price, pos.Quant, quant)
				pos.Quant += quant
				
				pos.Total -= total // A buy takes from total
			}
			if positionType == "V" {
				if pos.Quant > 0 {
					pos.Value += calculateResult(pos.Price, price, quant)
				}
				pos.Price = calculateAvgPrice(pos.Price, price, pos.Quant, -quant)
				pos.Quant -= quant

				pos.Total += total // A sell adds to total
				pos.ShortVolume += total
			}
			pos.QuantityVolume += quant
			pos.FinancialVolume += total
			pos.Start = date
			
			positions[asset] = pos
		}
	}

	return positions
}

func extractPageDate(pageText string) (date time.Time, err error) {
	lines = strings.Split(pageText, "\n")

	for _, line := range lines {
		texts := strings.Split(line, " ")
		if len(texts) == 3 {
			dateTxt := texts[2]
			if strings.Contains(dateTxt, "/") {
				return time.Parse("02/01/2006", dateTxt)
			}
		}
	}
	fmt.Println("(parsePageDate) Could not find date in expected form")

	return
}