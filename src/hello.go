package main

import (
	"fmt"
	"os"
	"net/http"
)

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
	// site := "https://leosantos.me/"

	sites := []string {"https://leosantos.me", "https://the-internet.herokuapp.com/status_codes/500"}
	
	for _, site := range sites {
		fmt.Println(healthCheck(site))
	}
	fmt.Println("")

}

func healthCheck(site string) string {
	res, _ := http.Get(site)
	if res.StatusCode >= 200 && res.StatusCode < 300 {
        return "Site " + site + " is up!"
    } else {
        return "Site " + site + " is down!"
    }
}