package clear

import (
	"encoding/json"
	"time"

	jap "github.com/ferkze/jsonappender"
	"github.com/ferkze/personal-finance-crawlers/clear/types"
)

// WriteOrdersToJSONFile escreve os registros no arquivo YYYY-MM-DD-orders.json
func WriteOrdersToJSONFile(fileName string, exs []types.Execution) (err error) {
	outDir := "out/"
	if fileName == "" {
		now := time.Now().Format("2006-01-02")
		fileName = now+"-orders.json"
	}
	
	a, err := jap.JSONAppender(outDir + fileName)
	if err != nil {
		return err
	}

	var b []byte
	b, err = json.Marshal(&exs)
  if err != nil {
    return
  }

  if _, err = a.Write(b); err != nil {
    return
  }

  if err = a.Close(); err != nil {
    return
  }
	return
}