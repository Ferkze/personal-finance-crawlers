package notas

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func parseDayTradeIndexFuturesOrders(positions DayTradePositions, text string) (DayTradePositions) {
	lines = strings.Split(text, "\n")

	pos := Position{
		AssetType: IndFut,
		Asset: "WIN",
	}

	for _, line := range lines {
		texts := strings.Split(line, " ")

		if len(texts) == 3 {
			dateTxt := texts[2]

			if strings.Contains(dateTxt, "/") {
				date, err := time.Parse("02/01/2006", dateTxt)
				if err != nil {
					panic(err.Error())
				}
				pos.Start = date
			}
		}
		
		if len(texts) < 2 {
			continue
		}

		if strings.HasPrefix(strings.ToLower(texts[1]), "win") || strings.HasPrefix(strings.ToLower(texts[1]), "ind") {
			priceTxt := texts[5]
			fmt.Println(texts, priceTxt)
			price, _ := strconv.ParseFloat(strings.ReplaceAll(strings.ReplaceAll(priceTxt, ".", ""), ",", "."), 64)

			quantTxt := texts[4]
			fmt.Println(texts, quantTxt)
			quant, _ := strconv.ParseInt(strings.ReplaceAll(strings.ReplaceAll(quantTxt, ".", ""), ",", "."), 10, 64)

			positionType := texts[0]
			
			fmt.Println(price, quant, positionType)
			total := price * float64(quant) / 5

			if positionType == "C" {
				pos.Total -= total
			}
			if positionType == "V" {
				pos.Total += total
			}
		}
	}

	positions["WIN"] = pos

	fmt.Println(positions)
	
	return positions
}


func parseDayTradeDolarFuturesOrders(positions DayTradePositions, text string) (DayTradePositions) {
	lines = strings.Split(text, "\n")

	pos := Position{
		AssetType: DolFut,
		Asset: "WDO",
	}

	for _, line := range lines {
		texts := strings.Split(line, " ")
		
		if len(texts) < 2 {
			continue
		}

		if len(texts) == 3 {
			dateTxt := texts[2]

			if strings.Contains(dateTxt, "/") {
				date, err := time.Parse("02/01/2006", dateTxt)
				if err != nil {
					panic(err.Error())
				}
				pos.Start = date
			}
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
			}
			if positionType == "V" {
				pos.Total += total
			}
		}
	}

	positions["WDO"] = pos

	
	return positions
}