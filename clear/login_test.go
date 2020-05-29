package clear

import (
	"time"

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
func iCanAccessOrdersPage() error {
	return navigateToOrdersPage(&Driver, "main")
}
func iCanFilterOrdersFromTo(dateStart, dateEnd string) error {
	start, _ := time.Parse("02/01/2006", dateStart)
	end, _ := time.Parse("02/01/2006", dateEnd)
	
	return filterOrders(&Driver, "main", start, end, "day_trade")
}
func iCanExtractDayTradeOrders() error {
	return parseMainPitOrders(&Driver, "day_trade")
}

func FeatureContext(s *godog.Suite) {
	s.Step(`^I\'m accessing the login page$`, imAccessingTheLoginPage)
	s.Step(`^I fill the login form$`, iFillTheLoginForm)
	s.Step(`^I get redirected to the pit selection$`, iGetRedirectedToThePitSelection)
	s.Step(`^I can access orders page$`, iCanAccessOrdersPage)
	s.Step(`^I can filter orders from "([^"]*)" to "([^"]*)"$`, iCanFilterOrdersFromTo)
	s.Step(`^I can extract day trade orders$`, iCanExtractDayTradeOrders)

	s.BeforeScenario(func(*messages.Pickle) {
		Driver = support.WBInit()
	})
	
}