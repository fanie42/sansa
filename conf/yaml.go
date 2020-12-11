package conf

import (
    "gopkg.in/yaml.v2"
)

// ParseYAML TODO
func ParseYAML(filename string, config interface{}) error {
    return yaml.Unmarshal(filename, config)
}
