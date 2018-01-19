package twofer

import (
	"fmt"
)

// ShareWith shares with someone.
func ShareWith(someone string) string {
	whom := "you"
	if len(someone) > 0 {
		whom = someone
	}
	return fmt.Sprintf("One for %s, one for me.", whom)
}
