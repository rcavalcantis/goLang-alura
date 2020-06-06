package main

import (
	"fmt"
	"net/http"
	"os"
)

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
	var sites [4]string
	sites[0] = "https://random-status-code.herokuapp.com/"
	sites[1] = "https://www.alura.com.br/"
	sites[2] = "https://uol.com.br/"

	for i := 0; i < len(sites); i++ {
		if sites[i] == "" {
			break
		}
		resp, _ := http.Get(sites[i])
		if resp.StatusCode == 200 {
			fmt.Println("Site: ", sites[i], "Status: UP")
		} else {
			fmt.Println("Site: ", sites[i], "Status: DOWN", "Status-Code:", resp.StatusCode)
		}
	}
}

func printingLog() {
	fmt.Println("Printing Log...")
}

func getSites() [4]string {
	var sites [4]string
	sites[0] = "https://random-status-code.herokuapp.com/"
	sites[1] = "https://www.alura.com.br/"
	sites[2] = "https://uol.com.br/"
	return sites
}
