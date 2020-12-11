package conf

import (
    "fmt"
    "io"
    "io/ioutil"

    "gopkg.in/yaml.v2"
)

// ParseYAML TODO
func ParseYAML(file io.Reader, config interface{}) error {
    data, err := ioutil.ReadAll(file)
    if err != nil {
        return fmt.Errorf("could not read config file")
    }

    return yaml.Unmarshal(data, config)
}
