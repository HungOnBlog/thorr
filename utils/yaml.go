package utils

import "gopkg.in/yaml.v2"

func UnmarshalYaml(data []byte, v interface{}) error {
	err := yaml.Unmarshal(data, v)
	if err != nil {
		return err
	}

	return nil
}

func MarshalYaml(v interface{}) ([]byte, error) {
	data, err := yaml.Marshal(v)
	if err != nil {
		return nil, err
	}

	return data, nil
}
