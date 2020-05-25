package support

import (
	"fmt"

	"github.com/tebeka/selenium"
)

var driver selenium.WebDriver

// WBInit retorna uma instância do webdriver
func WBInit() selenium.WebDriver {
	var err error
	caps := selenium.Capabilities(map[string]interface{}{
		"browserName": "chrome",
	})
	driver, err = selenium.NewRemote(caps, "")
	if err != nil {
		panic(fmt.Errorf("Erro ao instanciar o driver: %s", err.Error()))
	}

	driver.ResizeWindow("note", 1280, 800)
	return driver
}