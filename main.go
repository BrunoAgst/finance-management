package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println("Bem-vindo")
	fmt.Println("1 - Adicionar uma despesa em uma planilha existente")
	fmt.Println("2 - Criar uma nova planilha")
	fmt.Println("0 - Sair do programa")

	var command int
	fmt.Scan(&command)

	switch command {
	case 1:
		fmt.Println("adicionando...")
		updateWorksheet()
	case 2:
		fmt.Println("criando uma nova planilha")
		name, amount, date := createTransaction()
		createWorksheet(name, amount, date)
	case 0:
		fmt.Println("saindo do programa")
		os.Exit(0)
	default:
		fmt.Println("comando inválido")
		os.Exit(-1)
	}
}

func createWorksheet(name string, amount float64, date string) {
	file, err := os.OpenFile("planilha.csv", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	file.WriteString("name" + "," + "valor" + "," + "data" + "\n")
	file.WriteString(name + "," + strconv.FormatFloat(amount, 'f', 2, 64) + "," + date + "\n")

	file.Close()

}

func createTransaction() (string, float64, string) {
	var name, date string
	var amount float64

	fmt.Println("Digite o nome da transação:")
	fmt.Scan(&name)
	fmt.Println("")

	fmt.Println("Digite o valor da transação:")
	fmt.Scan(&amount)
	fmt.Println("")

	fmt.Println("Digite a data da transação:")
	fmt.Scan(&date)
	fmt.Println("")

	return name, amount, date

}

func updateWorksheet() {
	var nameWorksheet, name, date string
	var amount float64

	fmt.Println("Digite o nome da planilha:")
	fmt.Scan(&nameWorksheet)
	fmt.Println("")

	fmt.Println("Digite o nome da transação:")
	fmt.Scan(&name)
	fmt.Println("")

	fmt.Println("Digite o valor da transação:")
	fmt.Scan(&amount)
	fmt.Println("")

	fmt.Println("Digite a data da transação:")
	fmt.Scan(&date)
	fmt.Println("")

	filename := nameWorksheet + ".csv"
	file, err := os.OpenFile(filename, os.O_RDWR|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	file.WriteString(name + "," + strconv.FormatFloat(amount, 'f', 2, 64) + "," + date + "\n")

	file.Close()
}
