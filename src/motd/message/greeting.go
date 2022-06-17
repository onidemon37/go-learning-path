package message

import "fmt"

// Greeting function
func Greeting(name string, message string) (salutation string) {
	salutation = fmt.Sprintf("%s, %s", message, name)
	return
}
