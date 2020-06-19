package notas

import (
	"fmt"
	"math"
	"time"
)

func calculateResults(res Results, dayt DayTradePositions, swingt SwingTradePositions) {
	var date time.Time
	dr := DailyResult{}
	for k, v := range dayt {
		if v.Quant != 0 {
			fmt.Printf("Daytrade[%s] with Quant > 0: %#v", k, v)
			continue
		}
		date = v.Start
		
		switch v.AssetType {
		case Shares:
			dr.DayTradeShares = Result{
				AssetType: v.AssetType,
				MarketType: MercadoAVista,
				Date: v.Start,
				FinancialVolume: v.FinancialVolume + dr.DayTradeShares.FinancialVolume,
				ShortVolume: v.ShortVolume + dr.DayTradeShares.ShortVolume,
				Value: v.Result + dr.DayTradeShares.Value,
			}
		case IndFut:
			dr.DayTradeFutures = Result{
				AssetType: v.AssetType,
				MarketType: MercadoFuturo,
				Date: v.Start,
				FinancialVolume: v.FinancialVolume + dr.DayTradeFutures.FinancialVolume,
				ShortVolume: v.ShortVolume + dr.DayTradeFutures.ShortVolume,
				Value: v.Result + dr.DayTradeFutures.Value,
			}
		case DolFut:
			dr.DayTradeFutures = Result{
				AssetType: v.AssetType,
				MarketType: MercadoFuturo,
				Date: v.Start,
				FinancialVolume: v.FinancialVolume + dr.DayTradeFutures.FinancialVolume,
				ShortVolume: v.ShortVolume + dr.DayTradeFutures.ShortVolume,
				Value: v.Result + dr.DayTradeFutures.Value,
			}
		}
	}
	for k, v := range swingt {
		currentVol := (math.Round(v.Price * float64(v.Quant)*100)/100)
		r := Result{
			AssetType: v.AssetType,
			MarketType: MercadoAVista,
			Date: v.Start,
			FinancialVolume: v.FinancialVolume + dr.SwingTradeShares.FinancialVolume - currentVol,
			ShortVolume: v.ShortVolume + dr.SwingTradeShares.ShortVolume,
			Value: v.Result + dr.SwingTradeShares.Value,
		}
		dr.SwingTradeShares = r

		if v.Quant != 0 {
			v.FinancialVolume = currentVol
			v.ShortVolume = 0
			v.Result = 0
		} else {
			delete(swingt, k)
		}
	}
	res[date.Format("2006-01-02")] = dr
}