package main

import (
	"flag"
	"log"
	"os"

	"github.com/performl/hibernate/pkg/config"
	_kubeclient "github.com/performl/hibernate/pkg/kubernetes"
	_resources "github.com/performl/hibernate/pkg/resources"
	_states "github.com/performl/hibernate/pkg/states"
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
			// loading State
			// fetches from ConfigMap Object
			// loads into memory
			// if does not exist, do nothing
			// since it will create a new statefile
			_states.LoadState(clientset)

			_resources.SleepAll(resources)
			for _, resource := range resources {
				s := resource.GetState()
				name := s["name"].(string)
				namespace := s["namespace"].(string)
				resourceType := s["resourceType"].(string)
				_states.SetState(
					_states.CreateStateKey(name, namespace, resourceType),
					map[string]interface{}{
						"replicas": s["replicas"],
					},
				)
			}
			_states.PersistState(clientset)
		}
	case "wake":
		{
			// loading State
			// fetches from ConfigMap Object
			// loads into memory
			// if does not exist, Log Fatal since prior statefile must exist before it can be woken up
			_, err := _states.LoadState(clientset)
			if err != nil {
				log.Fatal(err)
			}
			_resources.WakeAll(resources)

			// should delete statefile
			stateFileName, stateFileNamespace := _states.GetStateFileAttrs()
			_states.DeleteConfigMap(clientset, stateFileName, stateFileNamespace)
		}
	default:
		{
			log.Fatal("invalid mode: please provide (sleep|wake)")
		}
	}

}
