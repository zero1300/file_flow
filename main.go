package main

import (
	"file_flow/cmd"
)

func main() {

	defer cmd.Clean()
	cmd.Start()

}
