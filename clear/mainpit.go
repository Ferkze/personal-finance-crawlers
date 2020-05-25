package clear

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/ferkze/personal-finance-crawlers/clear/support"
	"github.com/ferkze/personal-finance-crawlers/clear/types"
	"github.com/tebeka/selenium"
)


func parseMainPitOrders(d *selenium.WebDriver, operationType string) (err error) {
	if !support.IsCurrentInURL(d, OldPitOrdersURL) {
		err = fmt.Errorf("A página de ordens não está carregada")
		return
	}

	var ordersSelector string
	switch operationType {
	case "day_trade":
		ordersSelector = "#content_middle > div.container_middle > ul > li:nth-child(1) > div:nth-child(5) > div > div.content_right_orders > div > div.orderlistcontainer > div.container_s_orders.Main > ul > li"
	case "swing_trade":
		ordersSelector = "#content_middle > div.container_middle > ul > li:nth-child(1) > div:nth-child(4) > div > div.content_right_orders > div > div.orderlistcontainer > div.container_s_orders.Main > ul > li"
	}

	ordersEls, err := (*d).FindElements(selenium.ByCSSSelector, "#content_middle > div.container_middle > ul > li:nth-child(1) > div.drop-shadow.lifted.orderbar.xls > div > div.content_right_orders > div > div.orderlistcontainer > div.container_s_orders.Main > ul > li")
	if err != nil { return }

	getExecutions := func(d *selenium.WebDriver, selector string, index int) (orders []types.Execution, err error) {
		baseSelector := fmt.Sprintf("%s:nth-child(%d)", selector, index)
		// o, _ := (*d).FindElement(selenium.ByCSSSelector, baseSelector)

		// _, err = (*d).ExecuteScript(fmt.Sprintf("var s = document.querySelector('%s > div > h5') %s", baseSelector, strings.Repeat(";s.removeChild(s.lastElementChild)", 2)), nil)
		// if err != nil {
		// 	fmt.Printf("Erro ao executar o script para deletar a sujeira do nome do ativo, continuando: %s", err.Error())
		// 	// return orders, err
		// }
		assetEl, err := (*d).FindElement(selenium.ByCSSSelector,  baseSelector+" > div > h5")
		if err != nil { return orders, err }
		assetText, err := assetEl.Text()
		if err != nil { return orders, err }
		asset := strings.Split(assetText, " ")[0]
		fmt.Printf("AssetText: %s\nAsset extracted: %s\n", assetText, asset)
		var orderType string
		if strings.Contains(assetText, "Venda") {
			orderType = "Venda"
		} else if strings.Contains(assetText, "Compra") {
			orderType = "Compra"
		}

		var assetType string
		if strings.Contains(assetText, "WIN") || strings.Contains(assetText, "WDO") || strings.Contains(assetText, "IND") || strings.Contains(assetText, "DOL") {
			assetType = "futuros"
		} else {
			assetType = "acoes"
		}

		detailsButton, _ := (*d).FindElement(selenium.ByCSSSelector, baseSelector + " > div > div > div.container-orders-status > a")
		if err := detailsButton.Click(); err != nil { return orders, err }

		time.Sleep(1 * time.Second)

		elements, err := (*d).FindElements(selenium.ByCSSSelector, "#orderDetails > div > div > div > div:nth-child(7) > div.container_overflow_02.orderscol_01 > table > tbody > tr")
		if err != nil {
			fmt.Printf("Erro ao obter detalhes da ordem: %v", err)
			return orders, err
		}


		// orderTypeEl, _ := (*d).FindElement(selenium.ByCSSSelector, baseSelector+"> div > h5 > span")
		// orderType, _ := orderTypeEl.Text()
		
		// _, err = (*d).ExecuteScript("var s = document.querySelector('#orderDetails > div > div > div > table > tbody > tr:nth-child(2) > td:nth-child(2) > span');s.removeChild(s.lastElementChild)", nil)
		// if err != nil {
		// 	fmt.Printf("Erro ao executar o script para , continuando: %s", err.Error())
		// 	// return orders, err
		// }
		// assetTypeEl, err := (*d).FindElement(selenium.ByCSSSelector, "#orderDetails > div > div > div > table > tbody > tr:nth-child(2) > td:nth-child(2) > span")
		// if err != nil { return orders, err }
		// assetType, err := assetTypeEl.Text()
		// if err != nil { return orders, err }

		fmt.Printf("Found %d order executions\n", len(elements))
		for i := range elements {
			e := types.Execution{}
			baseExecSelector := fmt.Sprintf("#orderDetails > div > div > div > div:nth-child(7) > div.container_overflow_02.orderscol_01 > table > tbody > tr:nth-child(%d)", i+1)

			infoEl, _ := (*d).FindElement(selenium.ByCSSSelector, baseExecSelector+"> td.line_01")
			quant, _ := infoEl.Text()
			e.Quantity, _ = strconv.ParseInt(quant, 10, 64)

			infoEl, _ = (*d).FindElement(selenium.ByCSSSelector, baseExecSelector+"> td.line_02")
			price, _ := infoEl.Text()
			priceFl, _ := strconv.ParseFloat(strings.ReplaceAll(strings.TrimSuffix(price, "R$ "), ",", "."), 64)
			e.Price = priceFl

			infoEl, _ = (*d).FindElement(selenium.ByCSSSelector, baseExecSelector+"> td.line_03")
			date, _ := infoEl.Text()
			e.Datetime, _ = time.Parse("02/01/2006 15:04:05", date)
			
			e.Asset = asset
			e.AssetType = assetType
			e.OrderType = orderType

			orders = append(orders, e)
		}

		closeButtons, _ := (*d).FindElements(selenium.ByCSSSelector, "#orderDetails > div > div > a.bt_details_fechar")
		for _, closeButton := range closeButtons {
			closeButton.Click()
		}

		return orders, nil
	}

	for i := range ordersEls {
		var executions []types.Execution
		executions, err = getExecutions(d, ordersSelector, i+1)
		if err != nil {
			return
		}
		WriteRecords(executions)
	}
	return
}
