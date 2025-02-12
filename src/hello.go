package main

import "fmt"
import "os"
import "net/http"

func main() {
	showIntroduction()

	for {
		showMenu() 
		command := readCommandLine()

		switch command {
			case 1:
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

func startMonitoring() {
	fmt.Println("Monitoring...")
	site := "https://leosantos.me/"
	fmt.Println("Monitoring site:", site)
	response, error := http.Get(site)
	fmt.Println("Error:", error)

	if response.StatusCode >= 200 && response.StatusCode < 300 {
		fmt.Println("Site is up!")
	} else {
		fmt.Println("Site is down!")
	}
}