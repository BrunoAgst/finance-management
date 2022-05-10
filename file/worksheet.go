package file

import (
	"csv/transaction"
	"fmt"
	"os"
	"strconv"
)

func CreateWorksheet(transac transaction.Spent) {
	filename := transac.GetWorksheet() + ".csv"
	file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	file.WriteString("nome" + "," + "valor" + "," + "data" + "\n")
	file.WriteString(transac.GetName() + "," + strconv.FormatFloat(transac.GetValue(), 'f', 2, 64) + "," + transac.GetDate() + "\n")

	file.Close()

}

func UpdateWorksheet(transac transaction.Spent) {

	filename := transac.GetWorksheet() + ".csv"
	file, err := os.OpenFile(filename, os.O_RDWR|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	file.WriteString(transac.GetName() + "," + strconv.FormatFloat(transac.GetValue(), 'f', 2, 64) + "," + transac.GetDate() + "\n")

	file.Close()
}
