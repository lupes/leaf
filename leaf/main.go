package main

import (
	"leaf/cmd"
	"leaf/log"
)

func init() {
	log.Init()
}

func main() {
	cmd.Execute()
}
