package notas

import (
	"strconv"
	"strings"
)

func parseSharesOrders(results Results, positions SwingTradePositions, text string) (SwingTradePositions) {
	lines = strings.Split(text, "\n")

	res := Result{
		AssetType: Shares,
	}
	
	date, err := parsePageDate(text)
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
			
			pos  := Position{
				Asset: asset,
			}

			priceTxt := texts[len(texts)-3]
			price, _ := strconv.ParseFloat(strings.ReplaceAll(strings.ReplaceAll(priceTxt, ".", ""), ",", "."), 64)
			
			quantTxt := texts[len(texts)-4]
			quant, _ := strconv.ParseInt(strings.ReplaceAll(strings.ReplaceAll(quantTxt, ".", ""), ",", "."), 10, 64)
			
			positionType := texts[1]
			
			total := price * float64(quant)

			if positionType == "C" {
				pos.Price = calculateAvgPrice(pos.Price, price, pos.Quant, quant)
				pos.Quant += quant
				
				pos.Total -= total // A buy takes from total
			}
			if positionType == "V" {
				pos.Price = calculateAvgPrice(pos.Price, price, pos.Quant, -quant)
				pos.Quant -= quant

				pos.Total += total // A sell adds to total
				res.ShortVolume += total
			}
			res.QuantityVolume += quant
			res.FinancialVolume += total
			res.Date = date
			pos.Start = date
			
			positions[asset] = pos
		}
	}

	formatted := date.Format("2006-01-02")
	_, ok := results[formatted]
	if !ok {
		results[formatted] = make([]Result, 0)
	}
	results[formatted] = append(results[formatted], res)

	return positions
}

func calculateAvgPrice(p1, p2 float64, q1, q2 int64)  float64 {
	qf1 := float64(q1)
	qf2 := float64(q2)
	
	a1 := p1 * qf1
	a2 := p2 * qf2

	return (a1 + a2)/ (qf1 + qf2)
}