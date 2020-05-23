package clear

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/ferkze/personal-finance-crawlers/clear/types"
	"github.com/tebeka/selenium"
)

func parseOrders(d *selenium.WebDriver) (orders []types.Execution, err error) {
	ordersURL := "https://novopit.clear.com.br/Operacoes/Ordens"
	currentURL, err := (*d).CurrentURL()
	if err != nil { return }
	if currentURL != ordersURL {
		err = (*d).Get(ordersURL)
		if err != nil { return }
	}

	orderButtons, err := (*d).FindElements(selenium.ByCSSSelector, "#box-view-list-daytrade > li > section > div:nth-child(1) > nav > button")
	if err != nil { return }

	getExecutions := func(d *selenium.WebDriver) ([]types.Execution, error) {
		elements, _ := (*d).FindElements(selenium.ByCSSSelector, "#execution-list > li > div")
		orders := make([]types.Execution, len(elements))
		
		orderTypeEl, _ := (*d).FindElement(selenium.ByCSSSelector, "body > div.container.orders > div:nth-child(5) > section > div:nth-child(2) > div:nth-child(6) > label:nth-child(1) > span.order-side")
		orderType, _ := orderTypeEl.Text()

		assetEl, _ := (*d).FindElement(selenium.ByCSSSelector, "body > div.container.orders > div:nth-child(5) > section > div:nth-child(2) > p.order-symbol")
		asset, _ := assetEl.Text()

		assetTypeEl, _ := (*d).FindElement(selenium.ByCSSSelector, "body > div.container.orders > div:nth-child(5) > section > div:nth-child(2) > div:nth-child(7) > label:nth-child(4) > span.order-market")
		assetType, _ := assetTypeEl.Text()

		for i := range elements {
			selector := fmt.Sprintf("#execution-list > li > div > label:nth-child(%d) > span.execution-quantity", i+1)
			infoEl, _ := (*d).FindElement(selenium.ByCSSSelector, selector)
			quant, _ := infoEl.Text()
			orders[i].Quantity, _ = strconv.ParseInt(quant, 10, 64)

			selector = fmt.Sprintf("#execution-list > li > div > label:nth-child(%d) > span.execution-price", i+1)
			infoEl, _ = (*d).FindElement(selenium.ByCSSSelector, selector)
			price, _ := infoEl.Text()
			priceFl, _ := strconv.ParseFloat(strings.ReplaceAll(strings.TrimSuffix(price, "R$ "), ",", "."), 64)
			orders[i].Price = priceFl

			selector = fmt.Sprintf("#execution-list > li > div > label:nth-child(%d) > span.execution-date", i+1)
			infoEl, _ = (*d).FindElement(selenium.ByCSSSelector, selector)
			date, _ := infoEl.Text()
			orders[i].Datetime, _ = time.Parse("02/01/2006 15:04:05", date)
			
			orders[i].Asset = asset
			orders[i].AssetType = assetType
			orders[i].OrderType = orderType
		}

		return orders, nil
	}

	for _, btn := range orderButtons {
		var executions []types.Execution
		btn.Click()
		executions, err = getExecutions(d)
		if err != nil {
			return
		}
		orders = append(orders, executions...)
	}

	return
}