package src2struct

import (
	"fmt"
	"io/ioutil"

	log "github.com/sirupsen/logrus"
	yaml "gopkg.in/yaml.v2"
)

// LoadFileYaml will load a Yaml file and unmarshal it into the provided interface
func LoadFileYaml(logger *log.Entry, yamlFile string, data interface{}) error {
	fileContent, err := ioutil.ReadFile(yamlFile)
	if err != nil {
		return fmt.Errorf("reading the file '%s': %v", yamlFile, err)
	}
	err = yaml.Unmarshal(fileContent, data)
	if err != nil {
		return fmt.Errorf("unmarshal the file '%s': %v", yamlFile, err)
	}
	return nil
}
