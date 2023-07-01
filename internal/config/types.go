package config

type Configs struct {
	Envs  EnvMap
	Specs SpecsMap // specs
}

type EnvMap struct {
	ConfigPath string // config file path
}

// for YAML
type ConfigMap struct {
	Specs SpecsMap `yaml:"specs"`
}

type SpecsMap struct {
	Resources ResourcesMap `yaml:"resources"`
}

type ResourcesMap struct {
	Deployments  []ResourceMap `yaml:"deployments"`
	Statefulsets []ResourceMap `yaml:"statefulsets"`
}
type ResourceMap struct {
	Name      string `yaml:"name"`
	Namespace string `yaml:"namespace"`
}
