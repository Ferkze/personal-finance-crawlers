package clear

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/tebeka/selenium"
)

// logins using cpf, password and date_of_birth
func login(d *selenium.WebDriver, cpf, password, dob string) (err error) {
	err = (*d).Get("https://www.clear.com.br/pit/signin?controller=SignIn")
	if err != nil { return }

	cpfField, err := (*d).FindElement(selenium.ByCSSSelector, "#identificationNumber")
	if err != nil { return }
	cpfField.SendKeys(cpf)

	passwordField, err := (*d).FindElement(selenium.ByCSSSelector, "#password")
	if err != nil { return }
	passwordField.SendKeys(password)

	dobField, err := (*d).FindElement(selenium.ByCSSSelector, "#dob")
	if err != nil { return }
	dobField.SendKeys(dob)

	submitButton, err := (*d).FindElement(selenium.ByCSSSelector, "#form_id > input.bt_signin")
	if err != nil { return }
	submitButton.Click()

	return nil
}

func selectPit(d *selenium.WebDriver) (err error) {
	pitSelectionURL := "https://www.clear.com.br/pit/Selector"
	currentURL, err := (*d).CurrentURL()
	if err != nil { return }
	if currentURL != pitSelectionURL {
		err = (*d).Get(pitSelectionURL)
		if err != nil { return }
	}
	novopitLink, err := (*d).FindElement(selenium.ByCSSSelector, "#content_middle > div.middle > div.left > a")
	if err != nil { return }

	return novopitLink.Click()
}

func navigateToOrders(d *selenium.WebDriver) (err error) {
	ordersURL := "https://novopit.clear.com.br/Operacoes/Ordens"
	currentURL, err := (*d).CurrentURL()
	if err != nil { return }

	if currentURL != ordersURL {
		err = (*d).Get(ordersURL)
		if err == nil {
			return nil
		}
		fmt.Println("Erro ao navegar para a tela de ordens:", err.Error())
	}

	ordersLink, err := (*d).FindElement(selenium.ByCSSSelector, "body > div > div > nav > ul:nth-child(3) > li:nth-child(4) > a")
	if err != nil { return }

	return ordersLink.Click()
}

func filterOrders(d *selenium.WebDriver, start, end time.Time, operationType string) (err error) {
	ordersURL := "https://novopit.clear.com.br/Operacoes/Ordens"
	currentURL, err := (*d).CurrentURL()
	if err != nil { return }
	if currentURL != ordersURL {
		err = (*d).Get(ordersURL)
		if err != nil {
			return
		}
	}

	var checkboxes []selenium.WebElement

	switch operationType {
	case "day_trade":
		var cbDayTrade, cbPlatforms selenium.WebElement
		cbDayTrade, err = (*d).FindElement(selenium.ByCSSSelector, "body > div.container.orders > div.container_left_ords > div > div.container_filter > div > div:nth-child(2) > div > div:nth-child(1) > label:nth-child(3) > input[type=checkbox]")
		cbPlatforms, err = (*d).FindElement(selenium.ByCSSSelector, "body > div.container.orders > div.container_left_ords > div > div.container_filter > div > div:nth-child(2) > div > div:nth-child(2) > label:nth-child(3) > input[type=checkbox]")
		checkboxes = append(checkboxes, cbDayTrade, cbPlatforms)
	case "swing_trade":
		var checkbox selenium.WebElement
		checkbox, err = (*d).FindElement(selenium.ByCSSSelector, "body > div.container.orders > div.container_left_ords > div > div.container_filter > div > div:nth-child(2) > div > div:nth-child(1) > label:nth-child(2) > input[type=checkbox]")
		checkboxes = append(checkboxes, checkbox)
	}
	if err != nil { return }
	for _, checkbox := range checkboxes {
		checkbox.Click()
	}

	(*d).ExecuteScript("document.querySelector('#input-date-start').removeAttribute('readonly')")
	dateStartField, err := (*d).FindElement(selenium.ByID, "input-date-start")
	if err != nil { return }
	dateStartField.SendKeys(start.Format("02/01/2006"))

	(*d).ExecuteScript("document.querySelector('#input-date-end').removeAttribute('readonly')")
	dateEndField, err := (*d).FindElement(selenium.ByID, "input-date-end")
	if err != nil { return }
	dateEndField.SendKeys(end.Format("02/01/2006"))

	statusSelection, err := (*d).FindElement(selenium.ByID, "select-status")
	if err != nil { return }
	statusSelection.SendKeys("4")
	statusSelection.SendKeys(selenium.EnterKey)

	submitButton, err := (*d).FindElement(selenium.ByID, "body > div.container.orders > div.container_left_ords > div > div.container_filter > div > div:nth-child(5) > button")
	if err != nil { return }

	return submitButton.Click()
}

// Order structure
type Order struct {
	Type string
	Quantity int64
	Price float64
	Datetime time.Time	
}

func parseOrders(d *selenium.WebDriver) (orders []Order, err error) {
	ordersURL := "https://novopit.clear.com.br/Operacoes/Ordens"
	currentURL, err := (*d).CurrentURL()
	if err != nil { return }
	if currentURL != ordersURL {
		err = (*d).Get(ordersURL)
		if err != nil { return }
	}

	orderButtons, err := (*d).FindElements(selenium.ByCSSSelector, "#box-view-list-daytrade > li > section > div:nth-child(1) > nav > button")
	if err != nil { return }

	getExecutions := func(d *selenium.WebDriver) ([]Order, error) {
		elements, _ := (*d).FindElements(selenium.ByCSSSelector, "#execution-list > li > div")
		orders := make([]Order, len(elements))
		
		orderTypeEl, _ := (*d).FindElement(selenium.ByCSSSelector, "body > div.container.orders > div:nth-child(5) > section > div:nth-child(2) > div:nth-child(6) > label:nth-child(1) > span.order-side")
		orderType, _ := orderTypeEl.Text()

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
			
			orders[i].Type = orderType
		}

		return orders, nil
	}

	for _, btn := range orderButtons {
		var executions []Order
		btn.Click()
		executions, err = getExecutions(d)
		if err != nil {
			return
		}
		orders = append(orders, executions...)
	}

	return
}