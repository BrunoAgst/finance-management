package transaction

import (
	"fmt"
	"os"
)

type Spent struct {
	name, date, worksheetName string
	amount                    float64
}

func (s *Spent) Create(flag bool) {
	var amount float64
	var valid bool

	if flag {
		fmt.Println("Digite o nome da sua planilha:")
		fmt.Scan(&s.worksheetName)
		fmt.Println("")

		fmt.Println("Digite o nome da transação:")
		fmt.Scan(&s.name)
		fmt.Println("")

		fmt.Println("Digite o valor da transação:")
		fmt.Scan(&amount)
		fmt.Println("")

		valid = amountIsValid(amount)

		if !valid {
			fmt.Println("Valor digitado é inválido")
			os.Exit(-1)
		}

		s.amount = amount

		fmt.Println("Digite a data da transação:")
		fmt.Scan(&s.date)
		fmt.Println("")
		return

	}
	fmt.Println("Digite o nome da sua nova planilha(sem espaços):")
	fmt.Scan(&s.worksheetName)
	fmt.Println("")

	fmt.Println("Digite o nome da transação:")
	fmt.Scan(&s.name)
	fmt.Println("")

	fmt.Println("Digite o valor da transação:")
	fmt.Scan(&amount)
	fmt.Println("")

	valid = amountIsValid(amount)

	if !valid {
		fmt.Println("Valor digitado é inválido")
		os.Exit(-1)
	}

	s.amount = amount

	fmt.Println("Digite a data da transação:")
	fmt.Scan(&s.date)
	fmt.Println("")

}

func (s *Spent) GetName() string {
	return s.name
}

func (s *Spent) GetAmount() float64 {
	return s.amount
}

func (s *Spent) GetDate() string {
	return s.date
}

func (s *Spent) GetWorksheet() string {
	return s.worksheetName
}

func amountIsValid(amount float64) bool {
	return amount > 0
}
