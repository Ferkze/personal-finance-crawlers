package clear

import (
	"github.com/cucumber/godog"
	messages "github.com/cucumber/messages-go/v10"
	"github.com/ferkze/personal-finance-crawlers/clear/support"
	"github.com/tebeka/selenium"
)

var Driver selenium.WebDriver

func imAccessingTheLoginPage() error {
	
	return godog.ErrPending
}

func iFillTheLoginForm() error {
	return godog.ErrPending
}

func iGetRedirectedToThePitSelection() error {
	return godog.ErrPending
}

func FeatureContext(s *godog.Suite) {
	s.Step(`^I\'m accessing the login page$`, imAccessingTheLoginPage)
	s.Step(`^I fill the login form$`, iFillTheLoginForm)
	s.Step(`^I get redirected to the pit selection$`, iGetRedirectedToThePitSelection)

	s.BeforeScenario(func(*messages.Pickle) {
		Driver = support.WBInit()
	})
	
}