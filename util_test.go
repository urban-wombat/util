package util

import (
	"fmt"
	"testing"
)

func TestGoFmtProgramString(t *testing.T) {

	var deprecationMsg string = `
	util.GoFmtProgramString() will be DEPRECATED

	Instead use go/format.Source() as follows:

		import "go/format"
	
		var myGoCode string	// code to be formatted
		var goCodeBytes []byte
		goCodeBytes, err = format.Source([]byte(myGoCode))
		myGoCode = string(goCodeBytes)
	`

	fmt.Println(deprecationMsg)
}
