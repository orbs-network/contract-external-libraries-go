package keys

import "strings"

func Key(values... string) []byte {
	return []byte(strings.Join(values, ""))
}
