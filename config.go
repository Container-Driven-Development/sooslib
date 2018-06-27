package sooslib

import (
	"io/ioutil"

	yaml "gopkg.in/yaml.v2"
)

// BuildConfig is build config
type BuildConfig struct {
	Build     string
	Image     string
	Ports     []string
	Volumes   []string
	Hashfiles []string
}

// Config get object
func Config(file string) BuildConfig {

	configData, readFileErr := ioutil.ReadFile(file)

	check(readFileErr)

	buildConfigData := BuildConfig{}

	parseYamlErr := yaml.UnmarshalStrict([]byte(configData), &buildConfigData)

	check(parseYamlErr)

	return buildConfigData
}
