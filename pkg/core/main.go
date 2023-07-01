package main

import (
	"flag"
	"log"
	"os"

	"github.com/performl/hibernate/pkg/config"
	_kubeclient "github.com/performl/hibernate/pkg/kubernetes"
	_resources "github.com/performl/hibernate/pkg/resources"
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

	if action == "none" {
		log.Printf("Usage of %s:\n", os.Args[0])
		flag.PrintDefaults()
		os.Exit(1)
	}

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
