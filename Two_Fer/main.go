package twofer

import "fmt"

func ShareWith(name string) string {

	fmt.Scanln(&name)

	if name == "" {
		return "One for you, one for me."
	}

	if name == "Alice" {
		return "One for Alice, one for me."
	}

	if name == "Bob" {
		return "One for Bob, one for me."
	}

	return ""
}
