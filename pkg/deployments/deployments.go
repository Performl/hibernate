package deployments

import (
	"context"
	"encoding/json"
	"fmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	v1 "k8s.io/client-go/kubernetes/typed/apps/v1"
	"k8s.io/client-go/util/retry"

	_states "github.com/performl/hibernate/pkg/states"
	"github.com/performl/hibernate/pkg/utils"
)

// implements resources interface
type Deployments struct {
	clientset    *kubernetes.Clientset
	namespace    string
	name         string
	replicas     int32
	resourceType string
}

func InitDeployment(clientset *kubernetes.Clientset, namespace string, name string) *Deployments {
	deployments := &Deployments{
		clientset:    clientset,
		namespace:    namespace,
		name:         name,
		resourceType: "deployment",
	}
	// pulls data from remote
	deployments.FetchRemoteInfo()

	return deployments
}

func (d *Deployments) getClient() v1.DeploymentInterface {
	return d.clientset.AppsV1().Deployments(d.namespace)
}

func (d *Deployments) Sleep() {
	d.tryUpdateDeploymentReplica(0)
}

func (d *Deployments) Wake() {
	globalState := _states.GetState()
	stateKey := _states.CreateStateKey(d.name, d.namespace, d.resourceType)
	state := map[string]interface{}{}
	json.Unmarshal([]byte(globalState[stateKey].(string)), &state)
	targetReplica := int32(state["replicas"].(float64))
	d.tryUpdateDeploymentReplica(targetReplica)
}

func (d *Deployments) GetState() map[string]interface{} {
	return map[string]interface{}{
		"name":         d.name,
		"namespace":    d.namespace,
		"replicas":     d.replicas,
		"resourceType": d.resourceType,
	}
}

func (d *Deployments) FetchRemoteInfo() {
	deploymentClient := d.getClient()
	result, getErr := deploymentClient.Get(context.TODO(), d.name, metav1.GetOptions{})
	if getErr != nil {
		panic(fmt.Errorf("Failed to get latest version of Deployment: %v", getErr))
	}

	d.replicas = *result.Spec.Replicas

}

func (d *Deployments) tryUpdateDeploymentReplica(targetReplica int32) error {
	deploymentClient := d.getClient()
	previousReplica := -1

	retryErr := retry.RetryOnConflict(retry.DefaultRetry, func() error {
		// Retrieve the latest version of Deployment before attempting update
		// RetryOnConflict uses exponential backoff to avoid exhausting the apiserver
		result, getErr := deploymentClient.Get(context.TODO(), d.name, metav1.GetOptions{})
		if getErr != nil {
			panic(fmt.Errorf("Failed to get latest version of Deployment: %v", getErr))
		}
		previousReplica = int(*result.Spec.Replicas)
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
	fmt.Printf("Updated deployment for %s: %d -> %d\n", d.name, previousReplica, targetReplica)
	return nil
}
