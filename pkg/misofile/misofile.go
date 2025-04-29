package misofile

import "gopkg.in/yaml.v3"

type Jobfile struct {
	Project   string      `yaml:"project,omitempty"`
	Apps      []*App      `yaml:"app,omitempty"`
	Variables []*Variable `yaml:"variable,omitempty"`
	Secrets   []*Variable `yaml:"secret,omitempty"`
}

type App struct {
	Name          string          `yaml:"name"`
	BuildBlocks   []*BuildBlock   `yaml:"build,omitempty"`
	DeployBlocks  []*DeployBlock  `yaml:"deploy,omitempty"`
	ReleaseBlocks []*ReleaseBlock `yaml:"release,omitempty"`
}

type BuildBlock struct {
	Use        *UseBlock  `yaml:"use"`
	StepConfig *yaml.Node `yaml:",inline"` // <- kluczowe
}

type DeployBlock struct {
	Use        *UseBlock  `yaml:"use"`
	StepConfig *yaml.Node `yaml:",inline"` // <- kluczowe
}

type ReleaseBlock struct {
	Use        *UseBlock  `yaml:"use"`
	StepConfig *yaml.Node `yaml:",inline"` // <- kluczowe
}

type UseBlock struct {
	PluginName string `yaml:"name"`
}

type Variable struct {
	Name      string     `yaml:"name"`
	Value     string     `yaml:"value"`
	Remaining *yaml.Node `yaml:",inline"` // <- pozostałe rzeczy
}
