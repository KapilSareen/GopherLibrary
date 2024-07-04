package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	user := os.Getenv("MYSQL_USERNAME")
	passwd := os.Getenv("MYSQL_PASSWORD")
	url := os.Getenv("URL")
	dbname := os.Getenv("DATABASE")

	databaseURL := fmt.Sprintf("mysql://%s:%s@tcp(%s)/%s", user, passwd, url, dbname)

	argsWithoutProg := os.Args[1:]
	arg1 := argsWithoutProg[0]

	cmdArgs := []string{
		"-path", "migrations",
		"-database", databaseURL,
	}

	if arg1 == "up" {
		cmdArgs = append(cmdArgs, "up")
	} else if arg1 == "down" {
		cmdArgs = append(cmdArgs, "down")
		fmt.Print("Are you sure you want to apply all down migrations? [y/N]")
	} 
	cmd := exec.Command("migrate", cmdArgs...)

	cmd.Stdin = os.Stdin

	combinedOutput, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("Error executing command: %s\nCommand output:\n%s\n", err, string(combinedOutput))
	}
	fmt.Println(string(combinedOutput))
}
