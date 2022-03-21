package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strings"

	"github.com/go-cmd/cmd"
)

type LicenseResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Data    string `json:"data"`
}

func configurationQuestion(reader *bufio.Reader, description string, defaultValue string) (value string) {
	colorBlue := "\033[34m"
	colorWhite := "\033[37m"
	fmt.Print(string(colorBlue), description+" ("+defaultValue+") ")
	value, _ = reader.ReadString('\n')
	if value != "\n" {
		fmt.Println(string(colorWhite), description, value)
		return strings.TrimSuffix(value, "\n")
	}
	fmt.Println(string(colorWhite), description, defaultValue)
	return defaultValue
}

func commandExecutor(application string, args ...string) {
	cmd := exec.Command(application, args...)
	exec_error := cmd.Run()
	if exec_error != nil {
		fmt.Println("Error: "+": ", exec_error)
	}
}

func RunCMD(path string, args []string, debug bool) (out string, err error) {
	cmd := exec.Command(path, args...)
	var b []byte
	b, err = cmd.Output()
	out = string(b)
	if debug {
		fmt.Println(strings.Join(cmd.Args[:], " "))
		if err != nil {
			fmt.Println("RunCMD ERROR", err)
			fmt.Println(out)
		}
	}
	return
}

func RunCMD2(name string, args ...string) (err error, stdout, stderr []string) {
	c := cmd.NewCmd(name, args...)
	s := <-c.Start()
	stdout = s.Stdout
	stderr = s.Stderr
	return
}

func main() {
	colorRed := "\033[31m"
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(string(colorRed), "\n\n   Welcome to App Creator Installation   ")
	fmt.Println(string(colorRed), "-----------------------------------------")
	// Configs to be collected from user
	var APP_NAME = configurationQuestion(reader, "Application Name", "App-Creator")
	var WEB_APP_PORT = "4200"
	var API_APP_PORT = "8000"
	var DOMAIN_NAME = configurationQuestion(reader, "Web Application Domain name", "appcreator.com")

	// ---------------------------
	// Check License
	// var LICENSE_EMAIL = configurationQuestion(reader, "Enter your email address", "example@email.com")
	// if LICENSE_EMAIL == "example@email.com" && true {
	// 	println("Whoops! invalid email addres")
	// 	return
	// }
	var LICENSE_KEY = configurationQuestion(reader, "Enter your License Key", "XXXX-XXXX-XXXX-XXXX")
	if LICENSE_KEY == "XXXX-XXXX-XXXX-XXXX" {
		fmt.Print(string(colorRed), "Whoops! invalid License Key used")
		fmt.Println("")
		fmt.Println(string(colorRed), "***********************************")
		return
	}
	requestForm := map[string]string{"product_id": "A16D1689", "license_code": LICENSE_KEY, "client_name": DOMAIN_NAME}
	jsonValue, _ := json.Marshal(requestForm)
	resp, err := http.Post("http://127.0.0.1:8000/api/licenses/verify", "application/json", bytes.NewBuffer(jsonValue))
	if err != nil {
		fmt.Println("Whoops This!: ", err)
		return
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	var licenseResp LicenseResponse
	err = json.Unmarshal(body, &licenseResp)
	if err != nil {
		fmt.Println("Whoops Here!: ", err.Error())
		return
	}
	if !licenseResp.Status {
		fmt.Print(string(colorRed), licenseResp.Message)
		fmt.Println("")
		fmt.Println(string(colorRed), "***********************************")
		return
	}

	// Create New Environment File
	filename := ".env"
	fileStat, err := os.Stat(filename)
	if fileStat != nil {
		os.Remove(filename)
	}
	file, err := os.OpenFile(filename, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0660)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	defer file.Close()
	fmt.Fprintf(file, "APP_NAME=%s\n", APP_NAME)
	fmt.Fprintf(file, "WEB_APP_PORT=%s\n", WEB_APP_PORT)
	fmt.Fprintf(file, "API_APP_PORT=%s\n", API_APP_PORT)
	fmt.Fprintf(file, "DOMAIN_NAME=%s\n", DOMAIN_NAME)
	fmt.Fprintf(file, "DB_DATABASE=%s\n", "app_creator_db")
	fmt.Fprintf(file, "DB_PASSWORD=%s\n", "password123")
	fmt.Fprintf(file, "DB_USERNAME=%s\n", "admin")
	fmt.Fprintf(file, "LICENSE_NAME=%s\n", DOMAIN_NAME)
	fmt.Fprintf(file, "LICENSE_KEY=%s\n", LICENSE_KEY)
	// fmt.Println(string(colorRed), "-----------------------------------------")
	// fmt.Println(string(colorRed), "Getting Started . . . ")
	// commandExecutor( "docker-compose", "down")
	// time.Sleep(1)
	// fmt.Println(string(colorRed), "Starting Application")
	// commandExecutor( "docker-compose", "up", "-d")
	// time.Sleep(2)
	// fmt.Println(string(colorRed), "Check Database")
	// commandExecutor( "docker-compose", "exec", "app_creator_api php artisan migrate")
	// time.Sleep(5)
	// fmt.Println(string(colorRed), "Check Configs")
	// commandExecutor( "docker-compose", "exec", "app_creator_api php artisan config:cache")
	// time.Sleep(1)
	// fmt.Println(string(colorRed), "Seeding")
	// commandExecutor( "docker-compose", "exec", "app_creator_api php artisan db:seed")
	// time.Sleep(5)
	fmt.Println(string(colorRed), "Finished successfully")
	fmt.Println(string(colorRed), "-----------------------------------------")
}
