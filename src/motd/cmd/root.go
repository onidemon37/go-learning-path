/*
Author:............ Edino Moniz
EMAIL:............. edinomoniz@gmail.com

*/

package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	// "github.com/onidemon37/golearningpath"
	"github.com/spf13/cobra"
)

// defining my vars
var name string
var greeting string
var preview bool
var prompt bool
var debug bool = false

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "motd",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		// Show usage if flags are invalid
		if prompt == false && (name == "" || greeting == "") {
			cmd.Usage()
			os.Exit(1)
		}

		// Optionally print flags and exit based on DEBUG environment variables
		if debug {
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
		message := buildMessage(name, greeting)

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
	},
}

func buildMessage(name, greeting string) string {
	return fmt.Sprintf("%s, %s", greeting, name)
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

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	// if err := rootCmd.Execute() ; err != nil {
	//   fmt.Prinln(err)
	//   os.Exit(1)
	// }

	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringVarP(&name, "name", "n", "", "Name to use in the message")
	rootCmd.Flags().StringVarP(&greeting, "greeting", "g", "", "Greeting message")
	rootCmd.Flags().BoolVarP(&preview, "preview", "v", false, "Preview message instead of writing to /etc/motd")

	if os.Getenv("DEBUG") != "" {
		debug = true
	}
}
