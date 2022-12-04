package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var Uri string

func getName() string {
	fmt.Print("Enter username: ")
	reader := bufio.NewReader(os.Stdin)
	// ReadString will block until the delimiter is entered
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("An error occured while reading input. Please try again", err)
		return ""
	}

	// remove the delimeter from the string
	return strings.TrimRight(input, "\r\n")

}
func getPass() string {
	fmt.Print("Enter password: ")
	reader := bufio.NewReader(os.Stdin)
	// ReadString will block until the delimiter is entered
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("An error occured while reading input. Please try again", err)
		return ""
	}

	// remove the delimeter from the string
	return strings.TrimRight(input, "\r\n")

}

func Prompt() {
	// comment this 2 if you want to skip the prompt
	// username := getName()
	// password := getPass()
	// uncomment and write it here if you dont want the prompt
	username := "golang"
	password := "NlEag3qXsgpCCsrX"
	Uri = createURI(username, password)
}

func createURI(username, password string) string {
	return "mongodb+srv://" + username + ":" + password + "@cluster0.dqwwc.mongodb.net/test?retryWrites=true&w=majority"

}
