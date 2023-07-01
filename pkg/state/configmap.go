package state

import (
	"context"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	v1go "k8s.io/client-go/kubernetes/typed/core/v1"
)

func getConfigMapClient(clientset *kubernetes.Clientset, namespace string) v1go.ConfigMapInterface {
	return clientset.CoreV1().ConfigMaps(namespace)
}

func GetConfigMap(
	clientset *kubernetes.Clientset,
	name string,
	namespace string,
) (*v1.ConfigMap, error) {
	configMapClient := getConfigMapClient(clientset, namespace)
	configmap, err := configMapClient.Get(context.TODO(), name, metav1.GetOptions{})
	return configmap, err
}

func CreateConfigMap(
	clientset *kubernetes.Clientset,
	name string,
	namespace string,
	data map[string]string,
) error {
	configMapClient := getConfigMapClient(clientset, namespace)
	configmap := &v1.ConfigMap{
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: namespace,
		},
		Data: data,
	}
	_, err := configMapClient.Create(context.TODO(), configmap, metav1.CreateOptions{})
	return err
}

func DeleteConfigMap(
	clientset *kubernetes.Clientset,
	name string,
	namespace string,
) error {
	configMapClient := getConfigMapClient(clientset, namespace)
	err := configMapClient.Delete(context.TODO(), name, metav1.DeleteOptions{})
	return err
}
