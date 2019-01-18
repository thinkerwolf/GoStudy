package datarw

import (
	"fmt"
	"os"
	"strings"
)

func StartArgs() {
	args := os.Args

	name := ""

	if len(args) > 1 {
		name += strings.Join(args[1:], " ")
	}

	fmt.Println("Hello", name)

}
