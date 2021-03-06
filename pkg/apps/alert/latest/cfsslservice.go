/*
Copyright (C) 2019 Synopsys, Inc.

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
	horizonapi "github.com/blackducksoftware/horizon/pkg/api"
	"github.com/blackducksoftware/horizon/pkg/components"
	"github.com/blackducksoftware/synopsys-operator/pkg/util"
)

// getCfsslService returns a new Service for a Cffsl
func (a *SpecConfig) getCfsslService() *components.Service {
	service := components.NewService(horizonapi.ServiceConfig{
		Name:      util.GetResourceName(a.name, util.AlertName, "cfssl", a.isClusterScope),
		Namespace: a.config.Namespace,
		Type:      horizonapi.ServiceTypeServiceIP,
	})

	service.AddPort(horizonapi.ServicePortConfig{
		Port:       8888,
		TargetPort: "8888",
		Protocol:   horizonapi.ProtocolTCP,
		Name:       "8888-tcp",
	})

	service.AddSelectors(map[string]string{"app": util.AlertName, "name": a.name, "component": "cfssl"})
	service.AddLabels(map[string]string{"app": util.AlertName, "name": a.name, "component": "cfssl"})
	return service
}
