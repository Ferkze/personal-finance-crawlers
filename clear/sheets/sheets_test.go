package sheets

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"github.com/ferkze/personal-finance-crawlers/clear/types"
)

func TestDataToXLSX(t *testing.T) {
	orders, err := readOrders()
	if err != nil {
		t.Errorf("Could not read orders in executions.json: %v", err.Error())
	}

	OrdersToXLSX(orders)

	_, err = os.Stat("Orders.xlsx")
	if err != nil {
		t.Errorf("Did not generate Orders.xlsx: %v", err.Error())
		return
	}
}


func readOrders() (data []types.Execution, err error){
	file, err := ioutil.ReadFile("executions.json")
	if err != nil {
		return
	}

	b := bytes.NewBuffer(file)

	err = json.NewDecoder(b).Decode(&data)
	if err != nil {
		return
	}

	return
}