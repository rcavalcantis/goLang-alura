package main

import (
	"fmt"
	"os"
	"net/http"
)

func main() {
	showIntroduction()
	showMenu()
	choosingOperation(readCommand())

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
func startMonitoring(){
	fmt.Println("Monitoring...")
	site := "https://www.alura.com.br"
	http.Get(site)
}

func printingLog(){
	fmt.Println("Printing Log...")
}