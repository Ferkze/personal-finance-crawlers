package clear

import (
	"github.com/cucumber/godog"
	messages "github.com/cucumber/messages-go/v10"
	"github.com/ferkze/personal-finance-crawlers/clear/support"
	"github.com/ferkze/personal-finance-crawlers/clear/types"
	"github.com/tebeka/selenium"
)

var Driver selenium.WebDriver

func imAccessingTheLoginPage() error {
	return Driver.Get("https://www.clear.com.br/pit/signin?controller=SignIn")
}

func iFillTheLoginForm() error {
	acc := types.Account{
}
	return login(&Driver, acc)
}

func iGetRedirectedToThePitSelection() (err error) {
	return selectPit(&Driver)
}
func iCanAccessTheOrdersPage() error {
	return navigateToOrders(&Driver)
}

func FeatureContext(s *godog.Suite) {
	s.Step(`^I\'m accessing the login page$`, imAccessingTheLoginPage)
	s.Step(`^I fill the login form$`, iFillTheLoginForm)
	s.Step(`^I get redirected to the pit selection$`, iGetRedirectedToThePitSelection)
	s.Step(`^I can access the orders page$`, iCanAccessTheOrdersPage)

	s.BeforeScenario(func(*messages.Pickle) {
		Driver = support.WBInit()
	})
	
}