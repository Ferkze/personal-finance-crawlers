package clear

import (
	"log"
	"time"

	"github.com/ferkze/personal-finance-crawlers/clear/support"
	"github.com/ferkze/personal-finance-crawlers/clear/types"
)

const (
	// ClearLoginURL URL de login padrão da Clear
	ClearLoginURL = "https://www.clear.com.br/pit/signin?controller=SignIn"
	// OldPitURL URL do pit antigo
	OldPitURL = "https://www.clear.com.br/pit"
	// OldPitOrdersURL URL de ordens do pit antigo
	OldPitOrdersURL = "https://www.clear.com.br/pit/Orders"
	// NewPitOrdersURL URL de ordens do pit novo
	NewPitOrdersURL = "https://novopit.clear.com.br/Operacoes/Ordens"
)

func main() {
	wb := support.WBInit()

	err := wb.Get(ClearLoginURL)
	if err != nil {
		log.Fatalf("Erro ao abrir página no chromedriver: %s", err.Error())
	}
	acc := types.Account{
		CPF: "48574314838",
		DateOfBirth: "16062000",
		Password: "091136",
	}
	err = login(&wb, acc)
	if err != nil {
		log.Fatalf("Erro ao realizar login: %s", err.Error())
	}

	pit := "main"

	err = selectPit(&wb, pit)
	if err != nil {
		log.Fatalf("Erro ao selecionar o pit %s: %s", pit, err.Error())
	}
	err = navigateToOrders(&wb, pit)
	if err != nil {
		log.Fatalf("Erro no : %s", err.Error())
	}
	start, _ := time.Parse("02/01/2006", "26/05/2020")
	end, _ := time.Parse("02/01/2006", "26/05/2020")

	operationType := "day_trade"
	
	err = filterOrders(&wb, pit, start, end, operationType)
	if err != nil {
		log.Fatalf("Erro no : %s", err.Error())
	}
	err = parseMainPitOrders(&wb, operationType)
	if err != nil {
		log.Fatalf("Erro no : %s", err.Error())
	}
	
}
	