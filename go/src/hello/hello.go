package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
	"io"
	"io/ioutil"
	"bufio"
	"strings"
	"strconv"
)

const amountMonitoring = 1
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
	fmt.Println("3- Start monitoring Files Sites")
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
		startMonitoring(1)
	case 2:
		printingLog()
	case 3: 
		startMonitoring(3)
	case 0:
		fmt.Println("Shutdown!")
		os.Exit(0)
	default:
		fmt.Println("Command unknown!")
		os.Exit(-1)
	}
}
func startMonitoring(option int) {
	fmt.Println("Monitoring...")
	var sites []string
	switch option {
		case 1:
			sites = getSites()
		case 3:
			sites = loadSitesFile()
	}
	for i := 0; i <= amountMonitoring; i++ {
		for _, site := range sites {
			testSite(site)
		}
		time.Sleep(delay * time.Second)
		fmt.Println("")
	}
}

func testSite(site string){
	resp, err := http.Get(site)
	catchError(err, "testSite-http.Get")
	if resp.StatusCode == 200 {
		fmt.Println("Site: ", " - ", site, "Status: UP")
		recordLog(site, true)
	} else {
		fmt.Println("Site: ", " - ", site, "Status: DOWN", "Status-Code:", resp.StatusCode)
		recordLog(site, false)
	}
}

func printingLog() {
	fmt.Println("Printing Log...")
	fileLog, err := ioutil.ReadFile("log.txt")
	catchError(err, "printLog-oiutil.ReadFile")
	fmt.Println(string(fileLog))	
}

func getSites() []string {
	sites := []string{"https://random-status-code.herokuapp.com/", "https://www.alura.com.br/", "https://uol.com.br/"}
	return sites
}

func readSitesFile() string{
	sitesFile, err := ioutil.ReadFile("sites.txt")
	catchError(err, "readSitesFile-ioutil.ReadFile")
	return string(sitesFile)
}

func loadSitesFile() []string {
	var sites []string
	sitesFile, err := os.Open("sites.txt")
	if err != nil {
		catchError(err, "loadSitesFile-os.Open")
	}
	reader := bufio.NewReader(sitesFile)	
	for{
		line, err := reader.ReadString('\n')
		if err != nil && err != io.EOF {
			catchError(err, "reader.ReadString")
		}
		line = strings.TrimSpace(line)
		fmt.Println("TEST PRINT LINE: ", line)
		sites = append(sites, line)
		if err == io.EOF {
			break
		}		
	}
	sitesFile.Close()
	return sites
}

func catchError(err error, operation string){
	if err != nil {
		fmt.Println("Fail Operation: [", operation, "] - ", err)
		os.Exit(-1)
	}	
}

func recordLog(site string, status bool){
	logFile, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0755)
	if err != nil{
		fmt.Println("Error on file log.txt")
	}
	logFile.WriteString(time.Now().Format("02/01/2006 15:04:05") +" | "+ site + " - Status: " + strconv.FormatBool(status) + "\n")
	logFile.Close()
}
