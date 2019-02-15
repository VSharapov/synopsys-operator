/*
Copyright (C) 2018 Synopsys, Inc.

Licensed to the Apache Software Foundation (ASF) under one
or more contributor license agreements. See the NOTICE file
distributed with this work for additional information
regarding copyright ownershia. The ASF licenses this file
to you under the Apache License, Version 2.0 (the
"License"); you may not use this file except in compliance
with the License. You may obtain a copy of the License at

http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing,
software distributed under the License is distributed on an
"AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
KIND, either express or implied. See the License for the
specific language governing permissions and limitations
under the License.
*/

package synopsysctl

import (
	"github.com/blackducksoftware/horizon/pkg/components"
	"github.com/blackducksoftware/synopsys-operator/pkg/api"
)

// SampleSpecConfig represents your Sample component
// Its methods include GetComponents() and any functions
// that create Kubernetes Resources for the Sample
type PrometheusSpecConfig struct {
	Namespace       string
	PrometheusImage string
}

// NewPrometheus will create a PromtheusSpecConfig type
func NewPrometheus(namespace, prometheusImage string) *PrometheusSpecConfig {
	return &PrometheusSpecConfig{
		Namespace:       namespace,
		PrometheusImage: prometheusImage,
	}
}

// GetComponents will return a ComponentList representing all
// Kubernetes Resources for Prometheus
func (specConfig *PrometheusSpecConfig) GetComponents() (*api.ComponentList, error) {
	components := &api.ComponentList{
		Deployments: []*components.Deployment{
			specConfig.GetPrometheusDeployment(),
		},
		Services: []*components.Service{
			specConfig.GetPrometheusService(),
		},
		ConfigMaps: []*components.ConfigMap{
			specConfig.GetPrometheusConfigMap(),
		},
	}
	return components, nil
}
