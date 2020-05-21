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
	return selectPit(&Driver, "main")
}
func iCanAccessTheOrdersPage() error {
	return navigateToOrders(&Driver, "main")
}
func iCanFilterMyOrders() error {
	start, _ := time.Parse("02/01/2006", "13/05/2020")
	end, _ := time.Parse("02/01/2006", "18/05/2020")
	
	return filterOrders(&Driver, "main", start, end, "day_trade")
}

func FeatureContext(s *godog.Suite) {
	s.Step(`^I\'m accessing the login page$`, imAccessingTheLoginPage)
	s.Step(`^I fill the login form$`, iFillTheLoginForm)
	s.Step(`^I get redirected to the pit selection$`, iGetRedirectedToThePitSelection)
	s.Step(`^I can access the orders page$`, iCanAccessTheOrdersPage)
	s.Step(`^I can filter my orders$`, iCanFilterMyOrders)

	s.BeforeScenario(func(*messages.Pickle) {
		Driver = support.WBInit()
	})
	
}