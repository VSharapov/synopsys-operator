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

package alert

import (
	"fmt"

	"github.com/blackducksoftware/synopsys-operator/pkg/api"
	alertapi "github.com/blackducksoftware/synopsys-operator/pkg/api/alert/v1"
	log "github.com/sirupsen/logrus"
)

// SpecConfig will contain the specification to create the
// components of an Alert
type SpecConfig struct {
	name           string
	config         *alertapi.AlertSpec
	isClusterScope bool
}

// NewSpecConfig will create the Alert SpecConfig
func NewSpecConfig(name string, config *alertapi.AlertSpec, isClusterScope bool) *SpecConfig {
	return &SpecConfig{name: name, config: config, isClusterScope: isClusterScope}
}

// GetComponents will return the list of components for alert
func (a *SpecConfig) GetComponents() (*api.ComponentList, error) {
	log.Infof("Getting Alert Components")
	components := &api.ComponentList{}

	// Add alert components
	components.ConfigMaps = append(components.ConfigMaps, a.getAlertConfigMap())

	rc, err := a.getAlertReplicationController()
	if err != nil {
		return nil, fmt.Errorf("failed to create Alert Replication Controller: %s", err)
	}
	components.ReplicationControllers = append(components.ReplicationControllers, rc)

	service := a.getAlertService()
	components.Services = append(components.Services, service)

	switch a.config.ExposeService {
	case "NODEPORT":
		log.Debugf("case %s: Adding NodePort Service to ComponentList for Alert", a.config.ExposeService)
		components.Services = append(components.Services, a.getAlertServiceNodePort())
	case "LOADBALANCER":
		log.Debugf("case %s: Adding LoadBalancer Service to ComponentList for Alert", a.config.ExposeService)
		components.Services = append(components.Services, a.getAlertServiceLoadBalancer())
	default:
		log.Debugf("not adding a Kubernetes Service to ComponentList for Alert")
	}

	sec, err := a.GetAlertSecret()
	if err != nil {
		return nil, fmt.Errorf("failed to create Alert Secret: %s", err)
	}
	components.Secrets = append(components.Secrets, sec)

	if a.config.PersistentStorage {
		pvc, err := a.getAlertPersistentVolumeClaim()
		if err != nil {
			return nil, fmt.Errorf("failed to create Alert's PVC: %s", err)
		}
		components.PersistentVolumeClaims = append(components.PersistentVolumeClaims, pvc)
	}

	// Add cfssl if running in stand alone mode
	if *a.config.StandAlone {
		rc, err := a.getCfsslReplicationController()
		if err != nil {
			return nil, fmt.Errorf("failed to create Cfssl Replication Controller: %v", err)
		}
		components.ReplicationControllers = append(components.ReplicationControllers, rc)
		components.Services = append(components.Services, a.getCfsslService())
	}

	// Add routes for OpenShift
	route := a.GetOpenShiftRoute()
	if route != nil {
		components.Routes = []*api.Route{route}
	}

	return components, nil
}
