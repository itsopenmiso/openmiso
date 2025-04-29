package jobspec

import "gopkg.in/yaml.v3"

type JobSpec struct {
}

type BuildSpec struct {
	PlugName    string
	BuildConfig *yaml.Node
}
