package states

// example state
// {name}.{namespace}.(deployment|statefulset) = map[string]string
// i.e. unstructured json data
type State map[string]interface{}
