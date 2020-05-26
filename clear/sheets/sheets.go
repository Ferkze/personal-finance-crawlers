package sheets

import (
	"fmt"
	"time"

	excelize "github.com/360EntSecGroup-Skylar/excelize/v2"
	"github.com/ferkze/personal-finance-crawlers/clear/types"
)

// OrdersToXLSX loads execution data to xlsx file
func OrdersToXLSX(data []types.Execution) (error) {
	f := excelize.NewFile()

	sn := "Sheet1"
	f.SetActiveSheet(f.NewSheet(sn))

	for i, e := range data {
		row := fmt.Sprintf("%d", i+1)
		
		date := fmt.Sprintf("=DATE(%d, %d, %d)", e.Datetime.Year(), e.Datetime.Month(), e.Datetime.Day())
		err := f.SetCellFormula(sn, "A"+row, date)
		checkErr(err)

		err = f.SetCellStr(sn, "B"+row, e.Asset)
		checkErr(err)

		err = f.SetCellStr(sn, "C"+row, e.AssetType)
		checkErr(err)

		err = f.SetCellStr(sn, "D"+row, e.OrderType)
		checkErr(err)

		err = f.SetCellFloat(sn, "E"+row, e.Price, 2, 32)
		checkErr(err)

		err = f.SetCellInt(sn, "F"+row, int(e.Quantity))
		checkErr(err)
	}

	err := f.SaveAs("Orders.xlsx")
	if err != nil {
		return err
	}

	return nil
}


func checkErr(err error) {
	if err != nil {
		panic(err.Error())
	}
}
