package main

import (
	"csv/file"
	"csv/transaction"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Bem-vindo")

	for {
		fmt.Println("")
		fmt.Println("1 - Adicionar uma despesa em uma planilha existente")
		fmt.Println("2 - Criar uma nova planilha")
		fmt.Println("0 - Sair do programa")
		fmt.Println("")

		var command int
		fmt.Scan(&command)

		transac := transaction.Spent{}

		switch command {
		case 1:
			fmt.Println("")
			fmt.Println("Adicionando...")
			fmt.Println("")
			transac.Create(true)
			file.UpdateWorksheet(transac)
		case 2:
			fmt.Println("")
			fmt.Println("Criando uma nova planilha...")
			fmt.Println("")
			transac.Create(false)
			file.CreateWorksheet(transac)
		case 0:
			fmt.Println("Saindo do programa")
			os.Exit(0)
		default:
			fmt.Println("Comando inválido")
			os.Exit(-1)
		}

		var loop int
		fmt.Println("Deseja fazer uma nova operação?")
		fmt.Println("Digite 1 para sim e 0 para não:")
		fmt.Scan(&loop)
		if loop == 0 {
			os.Exit(0)
		}
	}
}
