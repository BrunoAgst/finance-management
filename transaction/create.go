package transaction

import (
	"fmt"
	"os"
)

type Spent struct {
	name, date, worksheetName string
	value                     float64
}

func (s *Spent) Create(flag bool) {
	var value float64
	var valid bool

	if flag {
		fmt.Println("Digite o nome da sua planilha:")
		fmt.Scan(&s.worksheetName)
		fmt.Println("")

		fmt.Println("Digite o nome da transação:")
		fmt.Scan(&s.name)
		fmt.Println("")

		fmt.Println("Digite o valor da transação:")
		fmt.Scan(&value)
		fmt.Println("")

		valid = valueIsValid(value)

		if !valid {
			fmt.Println("Valor digitado é inválido")
			os.Exit(-1)
		}

		s.value = value

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
	fmt.Scan(&value)
	fmt.Println("")

	valid = valueIsValid(value)

	if !valid {
		fmt.Println("Valor digitado é inválido")
		os.Exit(-1)
	}

	s.value = value

	fmt.Println("Digite a data da transação:")
	fmt.Scan(&s.date)
	fmt.Println("")

}

func (s *Spent) GetName() string {
	return s.name
}

func (s *Spent) GetValue() float64 {
	return s.value
}

func (s *Spent) GetDate() string {
	return s.date
}

func (s *Spent) GetWorksheet() string {
	return s.worksheetName
}

func valueIsValid(value float64) bool {
	return value > 0
}
