package conf

import (
    "io/ioutil"

    "gopkg.in/yaml.v2"
)

// ParseYAML TODO
func ParseYAML(filename string, config interface{}) error {
    data := ioutil.ReadFile(filename)

    return yaml.Unmarshal(data, config)
}
