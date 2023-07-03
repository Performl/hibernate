package states

import (
	"encoding/json"

	"k8s.io/client-go/kubernetes"
)

// todo use k8 api to get the statefile containing previous resource replica state

// cases
// 1. (resources exist, empty state) -> sleep
//     - create statefile
// 2. (resources exist, statefile exists) -> sleep
// 	   - replace statefile
// 3. (resources exist, empty state) -> wake
//     - create statefile
// 4. (resources exist, statefile exists) -> wake
//     a) current resource replicas > statefile replicas
//        - update statefile to current resource replicas

//     b) current resource replicas <= statefile replicas
//        - do nothing
//		  - delete statefile

// 5. (resources dont exist, empty state) -> sleep
//   - do nothing
//
// 6. (resources dont exist, statefile exists) -> sleep
//   - do nothing
//
// 7. (resources dont exist, empty state) -> wake
//   - do nothing
//
// 8. (resources dont exist, statefile exists) -> wake
//   - create resources using statefile
//   - delete statefile
var state = State{}

// default name and namespace
// TODO to be overriden by ENV VARS
func GetStateFileAttrs() (string, string) {
	name := "hibernate-state"
	namespace := "hibernate"
	return name, namespace
}

// load state from configmap
// does nothing if configmap does not exist
func LoadState(clientset *kubernetes.Clientset) State {
	name, namespace := GetStateFileAttrs()
	configMap, _ := GetConfigMap(clientset, name, namespace)
	// if configMap already exists, load it into state
	if configMap != nil {
		for key, data := range configMap.Data {
			state[key] = data // data is unstructured strings right now
		}
	}
	return state
} // returns unstructured state

func PersistState(clientset *kubernetes.Clientset) error {
	name, namespace := GetStateFileAttrs()

	data := make(map[string]string)
	for key, value := range state {
		b, _ := json.Marshal(value)
		data[key] = string(b)
	}

	err := CreateConfigMap(clientset, name, namespace, data)

	// if configMap already exists, replace it
	// TODO
	return err
}

func GetState() State {
	return state
}

func SetState(key string, value interface{}) {
	state[key] = value
}
