package main

import (
	"runtime"

	"github.com/coomsy/mmo/pkg/api/server"
	"github.com/coomsy/mmo/pkg/config"
	"github.com/coomsy/mmo/pkg/log"
)

func init() {
	if err := config.InitAppConfig(); err != nil {
		log.Fatal(err.Error())
	}

	log.Info("Config loaded")
	log.Debugf("Config Values %+v\n", config.AppConfig)
}

func main() {
	// might need with nginx, to get correct client metadata? https://github.com/gin-gonic/examples/tree/master/reverse-proxy
	numCPU := runtime.NumCPU()
	log.Infof("Running on %d CPU", numCPU)

	if runtime.NumCPU() > 2 {
		runtime.GOMAXPROCS(numCPU / 2)
	}

	server, err := server.NewServer()

	if err != nil {
		log.Error(err.Error())
	}
	if err := server.Run(); err != nil {
		log.Fatal(err.Error())
	}
}
