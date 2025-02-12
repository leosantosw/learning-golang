package main

import "fmt"
import "os"

func main() {
	showIntroduction()
	showMenu()

	command := readCommandLine()

	switch command {
		case 1:
			fmt.Println("Monitoring...")
			startMonitoring()
		case 2:
			fmt.Println("Showing logs...")
		case 3:
			fmt.Println("Exiting...")
			os.Exit(0)
		default:
			fmt.Println("Invalid command")
			os.Exit(-1)
	}
}

func showIntroduction() {
	name := "Leo"
	version := 1.0
	
	fmt.Println("Hi,", name)
	fmt.Println("This program is in version", version)
}

func showMenu() {
	fmt.Println("1 - Start monitoring")
	fmt.Println("2 - Show logs")
	fmt.Println("3 - Exit")
}

func readCommandLine() int {
	var command int 
	fmt.Scan(&command)
	return command
}

