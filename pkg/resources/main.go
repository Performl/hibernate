package resources

import (
	"github.com/performl/hibernate/pkg/config"
	d "github.com/performl/hibernate/pkg/deployments"
	"k8s.io/client-go/kubernetes"
)

func CreateResources(clientSet *kubernetes.Clientset, conf *config.Configs) []Resources {
	var resources []Resources
	for _, deployment := range conf.Specs.Resources.Deployments {
		resources = append(resources, d.InitDeployment(
			clientSet,
			deployment.Namespace,
			deployment.Name,
		))
	}
	// todo statefulset
	// for _, statefulset := range conf.Specs.Resources.Statefulsets {
	// 	statefulset.sleep()
	// }

	return resources
}

func SleepAll(resources []Resources) {
	for _, resource := range resources {
		resource.Sleep()
	}
}

func WakeAll(resources []Resources) {
	for _, resource := range resources {
		resource.Wake()
	}
}
