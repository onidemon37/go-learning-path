package main

import (
	"bufio"
	"flag"
	"fmt"
	"motd/message"
	"os"
	"strings"
)

func main() {
	// define flags
	var name string
	var greeting string
	var prompt bool
	var preview bool

	// Parse flags
	flag.StringVar(&name, "name", "", "name to use within the message")
	flag.StringVar(&greeting, "greeting", "", "greeting to use before the message")
	flag.BoolVar(&prompt, "prompt", false, "use prompt to input a name and greeting")
	flag.BoolVar(&preview, "preview", false, "Use preview to output message without writting to /etc/motd")

	flag.Parse()

	// Show usage if flags are invalid
	if prompt == false && (name == "" || greeting == "") {
		flag.Usage()
		os.Exit(1)
	}

	// Optionally print flags and exit based on DEBUG environment variables
	if os.Getenv("DEBUG") != "" {
		fmt.Println("Name:", name)
		fmt.Println("Greeting:", greeting)
		fmt.Println("Prompt:", prompt)
		fmt.Println("Preview:", preview)

		os.Exit(1)
	}
	// Conditionally read from stdin

	if prompt {
		name, greeting = renderPrompt()
	}
	// Generate message
	message := message.Greeting(name, greeting)

	// Either preview message or write to file
	if preview {
		// Write content
		f, err := os.OpenFile("./etc_motd", os.O_WRONLY, 644)

		if err != nil {
			fmt.Println("Error: Unable to open /etc/motd")
			os.Exit(1)
		}
		defer f.Close()

		// _, err = ioutil.WriteFile("./etc_motd", []byte(m), 0644)
		_, err = f.Write([]byte(message))

		if err != nil {
			fmt.Println("Unable to write /etc/motd")
			os.Exit(1)
		}
	}
}

func renderPrompt() (name, greeting string) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Your Greeting: ")
	greeting, _ = reader.ReadString('\n')
	greeting = strings.TrimSpace(greeting)

	fmt.Print("Your Name: ")
	name, _ = reader.ReadString('\n')
	name = strings.TrimSpace(name)
	return
}
