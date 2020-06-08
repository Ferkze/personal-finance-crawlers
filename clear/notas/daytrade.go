package notas

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func parseDayTradeFuturesOrders(pos AssetPosition, text string) (AssetPosition) {
	lines = strings.Split(text, "\n")

	win := &AssetPositionBalance{
		OperationType: "day_trade",
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
				win.Start = date
			}
		}
		
		if len(texts) < 2 {
			continue
		}

		if strings.HasPrefix(texts[1], "WIN") || strings.HasPrefix(texts[1], "IND") {
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
				win.Total -= total
			}
			if positionType == "V" {
				win.Total += total
			}
		}
	}

	pos["WIN"] = *win

	fmt.Println(pos)
	
	return pos
}