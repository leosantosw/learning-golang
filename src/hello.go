package main

import "fmt"
import "os"

func main() {
	showIntroduction()
	showMenu()

	command := readCommandLine()

	switch command {
		case 1:
			fmt.Println("Monitorando...")
		case 2:
			fmt.Println("Exibindo logs...")
		case 3:
			fmt.Println("Saindo do programa...")
			os.Exit(0)
		default:
			fmt.Println("Comando inválido.")
			os.Exit(-1)
	}
}

func showIntroduction() {
	name := "Leo"
	version := 1.0
	
	fmt.Println("Olá,", name)
	fmt.Println("Este programa está na versão", version)
}

func showMenu() {
	fmt.Println("1 - Iniciar monitoramento")
	fmt.Println("2 - Exibir logs")
	fmt.Println("3 - Sair do programa")
}

func readCommandLine() int {
	var command int 
	fmt.Scan(&command)
	return command
}