package main

import "fmt"

func main() {
	showIntroduction()
	fmt.Println("1 - Iniciar monitoramento")
	fmt.Println("2 - Exibir logs")
	fmt.Println("3 - Sair do programa")

	var command int
	fmt.Scan(&command)
	fmt.Println("O comando escolhido foi", command)

	switch command {
		case 1:
			fmt.Println("Monitorando...")
		case 2:
			fmt.Println("Exibindo logs...")
		case 3:
			fmt.Println("Saindo do programa...")
		default:
			fmt.Println("Comando inválido.")
	}
}

func showIntroduction() {
	name := "Leo"
	version := 1.0
	
	fmt.Println("Olá,", name)
	fmt.Println("Este programa está na versão", version)
}