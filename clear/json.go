package clear

import (
	"encoding/json"
	"time"

	jap "github.com/ferkze/jsonappender"
	"github.com/ferkze/personal-finance-crawlers/clear/types"
)

// WriteOrdersToJSONFile escreve os registros no arquivo YYYY-MM-DD-orders.json
func WriteOrdersToJSONFile(exs []types.Execution) (err error) {
	now := time.Now().Format("2006-01-02")
	f := now+"-orders.json"
	a, err := jap.JSONAppender(f)
	if err != nil {
		return err
	}

	if err = json.NewEncoder(a).Encode(&exs); err != nil {
		return err
	}

	if err = a.Close(); err != nil {
		return err
	}
	return
}