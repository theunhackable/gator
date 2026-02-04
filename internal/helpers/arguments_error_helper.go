package helpers

import "fmt"

func ExpectedRequired(exp int, req int) error {
	return fmt.Errorf("Required arguments: %d\n Found arguments: %d", exp, req)
}
