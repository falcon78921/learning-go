// Package twofer will take a name and output a certain string
package twofer

import "fmt"

// ShareWith will take a name input (string) and return another string based on the name
func ShareWith(name string) string {
    if name == "Alice" {
            return "One for Alice, one for me."
    } else if name == "Bob" {
            return "One for Bob, one for me."
    } else if name == "" {
            return "One for you, one for me."
    }

    return fmt.Sprintf("One for %s, one for me.", name)
}
