package main

import (
	_ "github.com/ferkze/personal-finance-crawlers/clear"
	"github.com/ferkze/personal-finance-crawlers/clear/notas"
)

func main() {
	// clear.Main()
	notas.ParsePDF()
}
