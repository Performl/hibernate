package main

import (
	"flag"
	"log"

	"github.com/performl/hibernate/internal/config"
	_kubeclient "github.com/performl/hibernate/internal/kubernetes"
	_resources "github.com/performl/hibernate/internal/resources"
)

// flags
var (
	mode   string
	action string
)

func init() {
	// initialise flags
	flag.StringVar(&mode, "mode", "prod", "--mode=(prod|local)")
	flag.StringVar(&action, "action", "none", "--action=(sleep|wake)")
}

func main() {
	// parse flags
	flag.Parse()

	// kubernetes client
	clientset := _kubeclient.CreateEnvironmentAwareClientSet(mode)

	// config file
	cfg := config.InitConfig()

	// resources
	resources := _resources.CreateResources(clientset, cfg)

	// commands
	switch action {
	case "sleep":
		{
			_resources.SleepAll(resources)
		}
	case "wake":
		{
			_resources.WakeAll(resources)
		}
	default:
		{
			log.Fatal("invalid mode: please provide (sleep|wake)")
		}
	}

}
