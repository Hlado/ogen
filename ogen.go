// Package ogen implements OpenAPI v3 code generation.
package ogen

import (
	"sigs.k8s.io/yaml"
)

// Parse parses JSON/YAML into OpenAPI Spec.
func Parse(data []byte) (s *Spec, err error) {
	s = &Spec{}
	if err := yaml.Unmarshal(data, s); err != nil {
		return nil, err
	}
	s.Init()
	return s, nil
}
