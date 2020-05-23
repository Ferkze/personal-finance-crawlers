package clear

import (
	"encoding/json"

	jap "github.com/ferkze/jsonappender"
	"github.com/ferkze/personal-finance-crawlers/clear/types"
)

// WriteRecords escreve os registros no arquivo records.json
func WriteRecords(exs []types.Execution) (err error) {
	f := "records.json"
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