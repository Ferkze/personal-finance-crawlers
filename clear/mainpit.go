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

func filterMainPitOrders(d *selenium.WebDriver, start, end time.Time, operationType string) (err error) {
	switch operationType {
	case "day_trade":
		_, err = (*d).ExecuteScript("document.querySelector('#content_middle > div:nth-child(1) > div > label:nth-child(2) > input').checked = false", nil)
		_, err = (*d).ExecuteScript("document.querySelector('#content_middle > div:nth-child(1) > div > label:nth-child(3) > input').checked = true", nil)
	case "swing_trade":
		_, err = (*d).ExecuteScript("document.querySelector('#content_middle > div:nth-child(1) > div > label:nth-child(2) > input').checked = true", nil)
		_, err = (*d).ExecuteScript("document.querySelector('#content_middle > div:nth-child(1) > div > label:nth-child(3) > input').checked = false", nil)
	}

	_, err = (*d).ExecuteScript(fmt.Sprintf("document.querySelector('#datefilter').value = '%s'", start.Format("02/01/2006")), nil)
	if err != nil {
		fmt.Printf("Error injecting script: %v\n", err) 
	}
	_, err = (*d).ExecuteScript(fmt.Sprintf("document.querySelector('#datefilterend').value = '%s'", end.Format("02/01/2006")), nil)
	if err != nil {
		fmt.Printf("Error injecting script: %v\n", err) 
	}

	_, err = (*d).ExecuteScript(fmt.Sprintf("document.querySelector('#status').value = '%s'", "WithExecutions"), nil)
	if err != nil {
		fmt.Printf("Error injecting script: %v\n", err) 
	}

	submitButton, err := (*d).FindElement(selenium.ByID, "btnSearchForOrders")
	if err != nil {
		return 
	}

	err = submitButton.Click()

	time.Sleep(15 * time.Second)

	return 
}

func parseMainPitOrders(d *selenium.WebDriver, operationType string) (err error) {
	if !support.IsCurrentInURL(d, OldPitOrdersURL) {
		err = fmt.Errorf("A página de ordens não está carregada")
		return
	}

	var ordersSelector string
	switch operationType {
	case "day_trade":
		ordersSelector = `#content_middle > div.container_middle > ul > li:nth-child(1) > div[data-type="DayTrade"] > div > div.content_right_orders > div > div.orderlistcontainer > div.container_s_orders > ul > li`
	case "swing_trade":
		ordersSelector = `#content_middle > div.container_middle > ul > li:nth-child(1) > div[data-type="SwingTrade"] > div > div.content_right_orders > div > div.orderlistcontainer > div.container_s_orders > ul > li`
	}

	ordersEls, err := (*d).FindElements(selenium.ByCSSSelector, ordersSelector)
	if err != nil {
		return 
	}

	getExecutions := func(d *selenium.WebDriver, selector string, index int) (orders []types.Execution, err error) {
		baseSelector := fmt.Sprintf("%s:nth-child(%d)", selector, index)
		
		assetEl, err := (*d).FindElement(selenium.ByCSSSelector,  baseSelector+" > div > h5")
		if err != nil {
			return orders, err 
		}
		assetText, err := assetEl.Text()
		if err != nil {
			return orders, err 
		}
		asset := strings.TrimSpace(strings.Split(assetText, "\n")[0])

		var orderType string
		if strings.Contains(assetText, "Venda") || strings.Contains(assetText, "Vd") {
			orderType = "Venda"
		} else if strings.Contains(assetText, "Compra") {
			orderType = "Compra"
		}

		var assetType string
		if strings.Contains(assetText, "WIN") || strings.Contains(assetText, "IND") {
			assetType = "IndiceFuturo"
		} else if strings.Contains(assetText, "WDO") || strings.Contains(assetText, "DOL") {
			assetType = "DolarFuturo"
		} else {
			assetType = "Acoes"
		}

		detailsButton, _ := (*d).FindElement(selenium.ByCSSSelector, baseSelector + " > div > div > div.container-orders-status > a")
		if err := detailsButton.Click(); err != nil { return orders, err }

		time.Sleep(1500 * time.Millisecond)

		elements, err := (*d).FindElements(selenium.ByCSSSelector, "#orderDetails > div > div > div > div:nth-child(7) > div.container_overflow_02.orderscol_01 > table > tbody > tr")
		if err != nil {
			fmt.Printf("Erro ao obter detalhes da ordem: %v\n", err)
			return orders, err
		}

		// log.Printf("Extracting executions from asset %q orders", asset)
		for i := range elements {
			e := types.Execution{}
			baseExecSelector := fmt.Sprintf("#orderDetails > div > div > div > div:nth-child(7) > div.container_overflow_02.orderscol_01 > table > tbody > tr:nth-child(%d)", i+1)

			infoEl, _ := (*d).FindElement(selenium.ByCSSSelector, baseExecSelector+"> td.line_01")
			quant, _ := infoEl.Text()
			e.Quantity, _ = strconv.ParseInt(quant, 10, 64)

			priceEl, _ := (*d).FindElement(selenium.ByCSSSelector, baseExecSelector+"> td.line_02")
			priceText, _ := priceEl.Text()
			priceTextTransformed := strings.ReplaceAll(strings.ReplaceAll(strings.TrimPrefix(priceText, "R$ "), ",", ""), ".", "")
			priceFl, err := strconv.ParseFloat(priceTextTransformed, 64)
			if err != nil {
				fmt.Printf("Erro ao converter o texto do elemento de preço em float64: %s\nPara o seletor: %s\n", err.Error(), baseExecSelector+"td.line_02")
			}
			e.Price = priceFl / 100

			dateEl, _ := (*d).FindElement(selenium.ByCSSSelector, baseExecSelector+"> td.line_03")
			date, _ := dateEl.Text()
			e.Datetime, _ = time.Parse("02/01/2006 15:04:05", date)
			
			e.Asset = asset
			e.AssetType = assetType
			e.OrderType = orderType

			orders = append(orders, e)
			// log.Printf("Extracted %d out of %d order executions", i+1, len(elements))
		}

		closeButtons, _ := (*d).FindElements(selenium.ByCSSSelector, "#orderDetails > div > div > a.bt_details_fechar")
		for _, closeButton := range closeButtons {
			ok, err := closeButton.IsDisplayed()
			if err != nil || !ok {
				continue
			}
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
		if len(executions) == 0 {
			return fmt.Errorf("executions not found for selector %q at index %d", ordersSelector, i)
		}
		err = WriteOrdersToJSONFile(executions[0].Datetime.Format("20060102")+".json", executions)
		if err != nil {
			return
		}
	}
	return
}
