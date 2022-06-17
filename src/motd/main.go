package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"motd/message"
	"os"
	"strings"
)

func main() {
	// define flags

	// Parse flags

	// Show usage if flags are invalid

	// Conditionally read from stdin

	// Optionally print flags and exit based on DEBUG environment variables

	// Generate message

	// Either preview message or write to file

	// Write content
}

func main() {
	f, err := os.OpenFile("/etc/motd", os.O_WRONLY, 644)

	if err != nil {
		fmt.Println("Error: Unable to open /etc/motd")
		os.Exit(1)
	}
	defer f.Close()

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Your Greeting: ")
	phrase, _ := reader.ReadString('\n')
	phrase = strings.TrimSpace(phrase)

	fmt.Print("Your Name: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	m := message.Greeting(name, phrase)
	err = ioutil.WriteFile("/etc/motd", []byte(m), 0644)

	if err != nil {
		fmt.Println("Unable to write /etc/motd")
		os.Exit(1)
	}
}
