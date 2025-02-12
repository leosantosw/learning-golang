package main

import (
	"fmt"
	"io"
	"os"
	"net/http"
	"time"
	"bufio"
	"strings"
	"strconv"
	"io/ioutil"
)

const monitoringTimes = 5
const monitoringInterval = 5

func main() {
	showIntroduction()

	for {
		showMenu() 
		command := readCommandLine()

		switch command {
			case 1:
				startMonitoring()
			case 2:
				showLogs()
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
	sites := readSitesFromFile("sites.txt")
	
	for i := 0; i < monitoringTimes; i++ {
		for _, site := range sites {
			fmt.Println(healthCheck(site))
		}
		time.Sleep(time.Second * monitoringInterval)
		fmt.Println("")
	} 
	fmt.Println("")
}

func healthCheck(site string) string {
	res, err := http.Get(site)
	
	if err != nil {
        return "Error making request to " + site + ": " + err.Error()
    }

	if res.StatusCode >= 200 && res.StatusCode < 300 {
		registerLog(site, true)
        return "Site " + site + " is up!"
    } else {
		registerLog(site, false)
        return "Site " + site + " is down!"
    }
}

func readSitesFromFile(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
        os.Exit(1)
	}
	reader := bufio.NewReader(file)
	
	var sites []string
	
	for {
		line, err := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		sites = append(sites, line)
		if err == io.EOF {
			break
		}
	}

	return sites
}

func registerLog(site string, status bool) {
	file, err := os.OpenFile("log.txt", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("Could not open log file", err)
	}
	currentTime := time.Now().Format("02/01/2006 15:04:05")
	file.WriteString(currentTime + " - " + site + "- online: " + strconv.FormatBool(status) + "\n")
	file.Close()
}

func showLogs() {
	file, err := ioutil.ReadFile("log.txt")
	if err != nil {
		fmt.Println("Error reading log file:", err)
    }
	fmt.Println(string(file))
}