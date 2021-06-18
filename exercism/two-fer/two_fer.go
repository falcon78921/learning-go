// Package twofer will take a name and output a certain string
package twofer

import "fmt"

// ShareWith will take a name input (string) and return another string based on the name
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}
	return fmt.Sprintf("One for %s, one for me.", name)
}
