// Copyright 2019 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// This file was generated with option-builder.go, DO NOT EDIT IT.

package secrets

type deleteConfig struct {
	// Namespace is the Kubernetes namespace to use
	Namespace string
}

// DeleteOption is a single option for configuring a deleteConfig
type DeleteOption func(*deleteConfig)

// DeleteOptions is a configuration set defining a deleteConfig
type DeleteOptions []DeleteOption

// toConfig applies all the options to a new deleteConfig and returns it.
func (opts DeleteOptions) toConfig() deleteConfig {
	cfg := deleteConfig{}

	for _, v := range opts {
		v(&cfg)
	}

	return cfg
}

// Extend creates a new DeleteOptions with the contents of other overriding
// the values set in this DeleteOptions.
func (opts DeleteOptions) Extend(other DeleteOptions) DeleteOptions {
	var out DeleteOptions
	out = append(out, opts...)
	out = append(out, other...)
	return out
}

// Namespace returns the last set value for Namespace or the empty value
// if not set.
func (opts DeleteOptions) Namespace() string {
	return opts.toConfig().Namespace
}

// WithDeleteNamespace creates an Option that sets the Kubernetes namespace to use
func WithDeleteNamespace(val string) DeleteOption {
	return func(cfg *deleteConfig) {
		cfg.Namespace = val
	}
}

// DeleteOptionDefaults gets the default values for Delete.
func DeleteOptionDefaults() DeleteOptions {
	return DeleteOptions{
		WithDeleteNamespace("default"),
	}
}
