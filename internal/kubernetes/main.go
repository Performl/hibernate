package kubernetes

import (
	"flag"
	"log"
	"path/filepath"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

func CreateEnvironmentAwareClientSet(mode string) *kubernetes.Clientset {
	// kubernetes client
	var clientset *kubernetes.Clientset

	switch mode {
	case "prod":
		{
			clientset = CreateInClusterClientSet()
		}
	case "local":
		{
			clientset = CreateOutOfClusterClientSet()
		}
	default:
		{
			log.Fatal("invalid mode: please provide (prod|local)")
		}
	}

	return clientset
}

// https://github.com/kubernetes/client-go/blob/master/examples/in-cluster-client-configuration/main.go
func CreateInClusterClientSet() *kubernetes.Clientset {
	config, err := rest.InClusterConfig()
	if err != nil {
		panic(err.Error())
	}

	// creates the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	return clientset
}

// https://github.com/kubernetes/client-go/blob/master/examples/out-of-cluster-client-configuration/main.go
func CreateOutOfClusterClientSet() *kubernetes.Clientset {
	var kubeconfig *string
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()

	// use the current context in kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err.Error())
	}

	// create the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	return clientset
}
