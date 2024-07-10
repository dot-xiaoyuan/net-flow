package flow

import "fmt"

func Debug(s string, a ...interface{}) {
	if *debug {
		fmt.Printf(s, a)
	}
}
