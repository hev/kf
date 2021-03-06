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

package services

type getServiceConfig struct {
	// Namespace is the Kubernetes namespace to use.
	Namespace string
}

// GetServiceOption is a single option for configuring a getServiceConfig
type GetServiceOption func(*getServiceConfig)

// GetServiceOptions is a configuration set defining a getServiceConfig
type GetServiceOptions []GetServiceOption

// toConfig applies all the options to a new getServiceConfig and returns it.
func (opts GetServiceOptions) toConfig() getServiceConfig {
	cfg := getServiceConfig{}

	for _, v := range opts {
		v(&cfg)
	}

	return cfg
}

// Extend creates a new GetServiceOptions with the contents of other overriding
// the values set in this GetServiceOptions.
func (opts GetServiceOptions) Extend(other GetServiceOptions) GetServiceOptions {
	var out GetServiceOptions
	out = append(out, opts...)
	out = append(out, other...)
	return out
}

// Namespace returns the last set value for Namespace or the empty value
// if not set.
func (opts GetServiceOptions) Namespace() string {
	return opts.toConfig().Namespace
}

// WithGetServiceNamespace creates an Option that sets the Kubernetes namespace to use.
func WithGetServiceNamespace(val string) GetServiceOption {
	return func(cfg *getServiceConfig) {
		cfg.Namespace = val
	}
}

// GetServiceOptionDefaults gets the default values for GetService.
func GetServiceOptionDefaults() GetServiceOptions {
	return GetServiceOptions{
		WithGetServiceNamespace("default"),
	}
}
