package states

import "fmt"

func CreateStateKey(name string, namespace string, resourceType string) string {
	return fmt.Sprintf("%s.%s.%s", name, namespace, resourceType)
}
