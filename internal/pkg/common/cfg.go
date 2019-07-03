package common

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

// Cfg stores available settings
type Cfg struct {
	// Configurations of ENB
	MCC                   string `yaml:"mcc"`
	MNC                   string `yaml:"mnc"`
	EutraBand             string `yaml:"eutraBand"`
	DownlinkFrequency     string `yaml:"downlinkFrequency"`
	UplinkFrequencyOffset string `yaml:"uplinkFrequencyOffset"`
	FlexRAN               bool   `yaml:"flexRAN"`
	// Global setting
	ConfigurationPathofCN  string `yaml:"configurationPathofCN"`
	ConfigurationPathofRAN string `yaml:"configurationPathofRAN"`
	SnapBinaryPath         string `yaml:"snapBinaryPath"`
	DNS                    string `yaml:"dns"`
	MmeDomainName          string `yaml:"mmeDomainName"`
	SpgwDomainName         string `yaml:"SpgwDomainName"`
	FlexRANDomainName      string `yaml:"flexRANDomainName"`
	Test                   bool   `yaml:"test"` //test configuring without changing any file; No snap is installed
}

// GetConf : read yaml into struct
func (c *Cfg) GetConf(logger *log.Logger, path string) error {
	//Read yaml here
	yamlFile, err := ioutil.ReadFile(path)
	if err != nil {
		logger.Println(err.Error())
		return err
	}

	err = yaml.Unmarshal(yamlFile, c)

	if err != nil {
		logger.Println(err.Error())
		return err
	}

	return nil
}
