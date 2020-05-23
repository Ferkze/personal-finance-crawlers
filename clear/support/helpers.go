package support

import (
	"strings"

	"github.com/tebeka/selenium"
)

// IsCurrentInURL checks if the current active url in selenium is equal to the url passed
func IsCurrentInURL(d *selenium.WebDriver, url string) (bool) {
	current, err := (*d).CurrentURL()
	if err != nil {
		return false
	}
	return strings.Contains(current, url)
}