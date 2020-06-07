package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

const amountMonitoring = 5
const delay = 5

func main() {
	for {
		showIntroduction()
		showMenu()
		choosingOperation(readCommand())
	}
}

func showIntroduction() {
	name := "Rodrigo"
	version := 0.1
	fmt.Println("Hello Mrs.", name)
	fmt.Println("This is program is in version ", version)
	fmt.Println("")
}
func showMenu() {
	fmt.Println("1- Start monitoring")
	fmt.Println("2- Show logs")
	fmt.Println("0- Shutdown")
	fmt.Println("")
}
func readCommand() int {
	var command int
	fmt.Scan(&command)
	fmt.Println("The option selected was: ", command)
	return command
}
func choosingOperation(operation int) {

	switch operation {
	case 1:
		startMonitoring()
	case 2:
		printingLog()
	case 0:
		fmt.Println("Shutdown!")
		os.Exit(0)
	default:
		fmt.Println("Command unknown!")
		os.Exit(-1)
	}
}
func startMonitoring() {
	fmt.Println("Monitoring...")
	sites := getSites()
	for i := 0; i <= amountMonitoring; i++ {
		for _, site := range sites {
			testSite(site)
		}
		time.Sleep(delay * time.Second)
		fmt.Println("")
	}
}

func testSite(site string){
	resp, _ := http.Get(site)
	if resp.StatusCode == 200 {
		fmt.Println("Site: ", " - ", site, "Status: UP")
	} else {
		fmt.Println("Site: ", " - ", site, "Status: DOWN", "Status-Code:", resp.StatusCode)
	}
}

func printingLog() {
	fmt.Println("Printing Log...")
}

func getSites() []string {
	sites := []string{"https://random-status-code.herokuapp.com/", "https://www.alura.com.br/", "https://uol.com.br/"}
	return sites
}
