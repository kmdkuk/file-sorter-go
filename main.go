package main

import (
	"github.com/kmdkuk/gorter/cmd"
	"github.com/kmdkuk/gorter/log"
	"github.com/spiegel-im-spiegel/logf"
)

func main(){
	log.SetMinLevel(logf.ERROR)
	log.Debug("start cmd")
	cmd.Execute()
	log.Debug("finish cmd")
}