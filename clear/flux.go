package clear

import (
	"fmt"
	"strings"
	"time"

	"github.com/ferkze/personal-finance-crawlers/clear/types"
	"github.com/tebeka/selenium"
)

func checkoutPageOrRedirect(d *selenium.WebDriver, url string) (err error) {
	current, err := (*d).CurrentURL()
	if err != nil {
		return
	}

	if !strings.Contains(current, url) {
		fmt.Printf("Redirecionando à URL: %v\n", url)
		time.Sleep(1 * time.Second)
		err = (*d).Get(url)
		if err != nil { return }
		current, err = (*d).CurrentURL()
		if err != nil { return }
		if !strings.Contains(current, url) {
			return fmt.Errorf("Não foi possível redirecionar para a url %s", url)
		}
	}
	return
}

// logins using cpf, password and date_of_birth
func login(d *selenium.WebDriver, acc types.Account) (err error) {
	if err = checkoutPageOrRedirect(d, "https://www.clear.com.br/pit/signin?controller=SignIn"); err != nil { return }

	cpfField, err := (*d).FindElement(selenium.ByCSSSelector, "#identificationNumber")
	if err != nil { return }
	cpfField.SendKeys(acc.CPF)

	passwordField, err := (*d).FindElement(selenium.ByCSSSelector, "#password")
	if err != nil { return }
	passwordField.SendKeys(acc.Password)

	dobField, err := (*d).FindElement(selenium.ByCSSSelector, "#dob")
	if err != nil { return }
	dobField.SendKeys(acc.DateOfBirth)

	submitButton, err := (*d).FindElement(selenium.ByCSSSelector, "#form_id > input.bt_signin")
	if err != nil { return }
	submitButton.Click()

	return nil
}

func selectPit(d *selenium.WebDriver, pit string) (err error) {
	if err = checkoutPageOrRedirect(d, "https://www.clear.com.br/pit/Selector"); err != nil { return }

	if pit == "" {
		pit = "main"
	}
	
	selector := "#content_middle > div.middle > div.right > a"
	if pit == "novo" {
		selector = "#content_middle > div.middle > div.left > a"
	}

	time.Sleep(3 * time.Second)

	_, err = (*d).ExecuteScript(fmt.Sprintf("window.location.href = document.querySelector('%s').getAttribute('href')", selector), nil)
	if err != nil {
		fmt.Printf("Error injecting script: %v\n", err)
	}

	current, err := (*d).CurrentURL()
	if err != nil { return }

	if strings.Contains(current, "SignIn") {
		acc := types.Account{
			CPF: "48574314838",
			DateOfBirth: "16062000",
			Password: "091136",
		}
		err = login(d, acc)
		if err != nil { return }
		return checkoutPageOrRedirect(d, OldPitURL)
	}
	if pit == "novo" && !strings.HasPrefix(current, "https://novopit") {
		return fmt.Errorf("A seleção do pit não redirecionou para o Novo Pit")
	} else if pit == "main" && current != "https://www.clear.com.br/pit" {
		return fmt.Errorf("A seleção do pit não redirecionou para o Pit Main")
	}
	return 
}

func navigateToOrdersPage(d *selenium.WebDriver, pit string) (err error) {
	url := OldPitOrdersURL
	selector := "#nav-menu > li:nth-child(2) > ul > li.Orders > a"
	if pit == "novo" {
		url = NewPitOrdersURL
		selector = "body > div > div > nav > ul:nth-child(3) > li:nth-child(4) > a"
	}
	if err = checkoutPageOrRedirect(d, url); err == nil {
		if err := checkAndClosePopup(d, pit); err != nil {
			fmt.Println("Erro ao fechar o popup:", err.Error())
		}
		return
	}
	fmt.Printf("Erro ao redirecionar à página de ordens: %v\n", err.Error())

	ordersLink, err := (*d).FindElement(selenium.ByCSSSelector, selector)
	if err != nil { return }

	return ordersLink.Click()
}

func filterOrders(d *selenium.WebDriver, pit string, start, end time.Time, operationType string) (err error) {
	ordersURL := OldPitOrdersURL
	if pit == "novo" {
		ordersURL = NewPitOrdersURL
	}
	if err = checkoutPageOrRedirect(d, ordersURL); err != nil { return }

	if pit == "novo" {
		return filterNewPitOrders(d, start, end, operationType)
	}
	return filterMainPitOrders(d, start, end, operationType)
}

func filterNewPitOrders(d *selenium.WebDriver, start, end time.Time, operationType string) (err error) {

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

	(*d).ExecuteScript("document.querySelector('#input-date-start').removeAttribute('readonly')", nil)
	dateStartField, err := (*d).FindElement(selenium.ByID, "input-date-start")
	if err != nil { return }
	err = dateStartField.Clear()
	if err != nil { return }
	err = dateStartField.SendKeys(start.Format("02/01/2006"))
	if err != nil { return }

	(*d).ExecuteScript("document.querySelector('#input-date-end').removeAttribute('readonly')", nil)
	dateEndField, err := (*d).FindElement(selenium.ByID, "input-date-end")
	if err != nil { return }
	err = dateEndField.Clear()
	if err != nil { return }
	err = dateEndField.SendKeys(end.Format("02/01/2006"))
	if err != nil { return }

	statusSelection, err := (*d).FindElement(selenium.ByID, "select-status")
	if err != nil { return }
	statusSelection.SendKeys("4")
	statusSelection.SendKeys(selenium.EnterKey)

	submitButton, err := (*d).FindElement(selenium.ByID, "body > div.container.orders > div.container_left_ords > div > div.container_filter > div > div:nth-child(5) > button")
	if err != nil { return }

	return submitButton.Click()
}

func checkAndClosePopup(d *selenium.WebDriver, pit string) (err error) {
	selector := "#ipo_close"
	if pit == "novo" {
		selector = "#disclaimer_tyform > div"
	}
	el, err := (*d).FindElement(selenium.ByCSSSelector, selector)
	
	
	if err != nil { return }
	isDisplayed, err := el.IsDisplayed()
	if err != nil { return }
	if !isDisplayed {
		return
	}
	return el.Click()
}