package main

import (
	"os"
)

func main() {

	if len(os.Args) > 1 {
		if os.Args[1] == "ls" {
			mainLs()
			os.Exit(0)
		}

		if os.Args[1] == "kube" {
			mainKubectl()
			os.Exit(0)
		}
	}

	cuiMain()
}
