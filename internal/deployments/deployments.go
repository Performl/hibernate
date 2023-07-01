package deployments

import (
	"context"
	"fmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	v1 "k8s.io/client-go/kubernetes/typed/apps/v1"
	"k8s.io/client-go/util/retry"

	"github.com/performl/hibernate/internal/utils"
)

// implements resources interface
type Deployments struct {
	clientset *kubernetes.Clientset
	namespace string
	name      string
}

func InitDeployment(clientset *kubernetes.Clientset, namespace string, name string) *Deployments {
	return &Deployments{
		clientset: clientset,
		namespace: namespace,
		name:      name,
	}
}

func (d *Deployments) getClient() v1.DeploymentInterface {
	return d.clientset.AppsV1().Deployments(d.namespace)
}

func (d *Deployments) Sleep() {
	d.tryUpdateDeploymentReplica(0)
}

func (d *Deployments) Wake() {

}

func (d *Deployments) tryUpdateDeploymentReplica(targetReplica int32) error {
	deploymentClient := d.getClient()

	retryErr := retry.RetryOnConflict(retry.DefaultRetry, func() error {
		// Retrieve the latest version of Deployment before attempting update
		// RetryOnConflict uses exponential backoff to avoid exhausting the apiserver
		result, getErr := deploymentClient.Get(context.TODO(), d.name, metav1.GetOptions{})
		if getErr != nil {
			panic(fmt.Errorf("Failed to get latest version of Deployment: %v", getErr))
		}

		if *result.Spec.Replicas == targetReplica {
			return nil // no need to update
		}

		result.Spec.Replicas = utils.Int32Ptr(targetReplica) // update replica count
		_, updateErr := deploymentClient.Update(context.TODO(), result, metav1.UpdateOptions{})
		return updateErr
	})
	if retryErr != nil {
		panic(fmt.Errorf("Update failed: %v", retryErr))
	}
	fmt.Printf("Updated deployment for %s\n", d.name)
	return nil
}