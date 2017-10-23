/*
Package twofer implements the "Two for one" string problem
See the package's README for further details
*/
package twofer

import "fmt"

// ShareWith returns the following string where X is the argument: "One for X, one for me."
func ShareWith(name string) string {
	if len(name) == 0 {
	  name = "you"
	}
	return fmt.Sprintf("One for %s, one for me.", name)
}
