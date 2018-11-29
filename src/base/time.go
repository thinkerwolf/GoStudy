package base

import (
	"time"
	"fmt"
)

func Tim() {
	t := time.Now().UTC().Local()
	fmt.Println(t)
}