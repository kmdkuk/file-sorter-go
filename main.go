package main

import (
	"github.com/kmdkuk/file-sorter-go/cmd"
	"github.com/kmdkuk/file-sorter-go/log"
	"github.com/spiegel-im-spiegel/logf"
)

func main(){
	log.SetMinLevel(logf.ERROR)
	log.Debug("start cmd")
	cmd.Execute()
	log.Debug("finish cmd")
}