package clear

import (
	"fmt"
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
