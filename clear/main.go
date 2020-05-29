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
// Main ...
func Main() {
	wb := support.WBInit()

	err := wb.Get(ClearLoginURL)
	if err != nil {
		log.Fatalf("Erro ao abrir página no chromedriver: %s\n", err.Error())
	}
	acc := types.Account{
		CPF: "48574314838",
		DateOfBirth: "16062000",
		Password: "091136",
	}
	err = login(&wb, acc)
	if err != nil {
		log.Fatalf("Erro ao realizar login: %s\n", err.Error())
	}

	pit := "main"

	err = selectPit(&wb, pit)
	if err != nil {
		log.Fatalf("Erro ao selecionar o pit %s: %s\n", pit, err.Error())
	}
	err = navigateToOrdersPage(&wb, pit)
	if err != nil {
		log.Fatalf("Erro no : %s", err.Error())
	}
	start, _ := time.Parse("02/01/2006", "01/01/2019")
	end, _ := time.Parse("02/01/2006", "31/12/2019")

	operationType := "day_trade"
	
	diff := end.Unix() - start.Unix()
	if diff < 0 {
		log.Fatalf("Erro de inconsistência no período de consulta: Início %v; Fim: %v\n", start, end)
	}

	end = end.AddDate(0,0,1)
	for start.Before(end) {
		log.Printf("Filtering orders for day %v\n", start)
		err = filterOrders(&wb, pit, start, start, operationType)
		if err != nil {
			log.Fatalf("Erro no filtro de ordens: %s\n", err.Error())
		}
		err = parseMainPitOrders(&wb, operationType)
		if err != nil {
			log.Fatalf("Erro na extração de ordens: %s\n", err.Error())
		}
		start = start.AddDate(0, 0, 1)
	}

	log.Printf("Operação finalizada\n")
}
