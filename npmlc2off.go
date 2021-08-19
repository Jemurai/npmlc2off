package main

import (
	"os"

	cmd "github.com/jemurai/npmlc2off/cmd"
)

func main() {
	var npmlcjson string = os.Args[1]
	cmd.Convert(npmlcjson)
}
