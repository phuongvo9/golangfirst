// declare greetings package to collect related functions
package greetings

import "fmt"

// Hello returns a greeting to the name
func Hello (name string) string {
	message:= fmt.Sprint("Hi, %v. Welcome!", name)
	return message
}