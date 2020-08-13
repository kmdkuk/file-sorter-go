package main

import (
	"github.com/kmdkuk/gorter/cmd"
	"github.com/kmdkuk/gorter/log"
)

func main() {
	log.Debug("start cmd")
	cmd.Execute()
	log.Debug("finish cmd")
}
