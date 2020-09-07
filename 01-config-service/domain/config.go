package domain

import (
	"fmt"

	"gopkg.in/yaml.v2"
)

//Config is an abstraction around the map that holds the config values
type Config struct {
	config map[string]interface{}
}

//Get the serviced needed
func (c *Config) Get(serviceName string) (map[string]interface{}, error) {
	a, ok := c.config["base"].(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("Base does not exist, or is not a string")
	}

	if _, ok := c.config[serviceName]; !ok {
		return a, nil
	}

	b, ok := c.config[serviceName].(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("service is not a map")
	}

	config := make(map[string]interface{})
	for k, v := range a {
		config[k] = v
	}
	for k, v := range b {
		config[k] = v
	}

	return config, nil
}

//SetFromBytes sets the internal config based on the byte array of the YAML file
func (c *Config) SetFromBytes(data []byte) error {
	var rawConfig interface{}
	if err := yaml.Unmarshal(data, &rawConfig); err != nil {
		return err
	}

	unassertedConfig, ok := rawConfig.(map[interface{}]interface{})
	if !ok {
		return fmt.Errorf("rawConfig is not a map[interface{}].interface")
	}

	config, err := convertKeysToStrings(unassertedConfig)
	if err != nil {
		return err
	}

	c.config = config

	return nil
}

func convertKeysToStrings(m map[interface{}]interface{}) (map[string]interface{}, error) {
	n := make(map[string]interface{})

	for k, v := range m {
		str, ok := k.(string)
		if !ok {
			return nil, fmt.Errorf("key value is not a string")
		}

		if vMap, ok := v.(map[interface{}]interface{}); ok {
			var err error
			v, err = convertKeysToStrings(vMap)
			if err != nil {
				return nil, err
			}
		}
		n[str] = v
	}

	return n, nil

}
