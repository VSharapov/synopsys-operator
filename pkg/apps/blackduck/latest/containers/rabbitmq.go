/*
Copyright (C) 2019 Synopsys, Inc.

Licensed to the Apache Software Foundation (ASF) under one
or more contributor license agreements. See the NOTICE file
distributed with this work for additional information
regarding copyright ownership. The ASF licenses this file
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

package containers

import (
	horizonapi "github.com/blackducksoftware/horizon/pkg/api"
	"github.com/blackducksoftware/horizon/pkg/components"
	"github.com/blackducksoftware/synopsys-operator/pkg/util"
)

// GetRabbitmqDeployment will return the rabbitmq deployment
func (c *Creater) GetRabbitmqDeployment(imageName string) (*components.ReplicationController, error) {
	volumeMounts := c.getRabbitmqVolumeMounts()

	rabbitmqContainerConfig := &util.Container{
		ContainerConfig: &horizonapi.ContainerConfig{Name: "rabbitmq", Image: imageName,
			PullPolicy: horizonapi.PullAlways, MinMem: c.hubContainerFlavor.RabbitmqMemoryLimit, MaxMem: c.hubContainerFlavor.RabbitmqMemoryLimit,
			MinCPU: "", MaxCPU: ""},
		EnvConfigs:   []*horizonapi.EnvConfig{c.getHubConfigEnv()},
		VolumeMounts: volumeMounts,
		PortConfig:   []*horizonapi.PortConfig{{ContainerPort: rabbitmqPort, Protocol: horizonapi.ProtocolTCP}},
	}

	var initContainers []*util.Container
	if c.hubSpec.PersistentStorage {
		initContainerConfig := &util.Container{
			ContainerConfig: &horizonapi.ContainerConfig{Name: "alpine", Image: "alpine", Command: []string{"sh", "-c", "chmod -c 777 /var/lib/rabbitmq"}},
			VolumeMounts:    volumeMounts,
		}
		initContainers = append(initContainers, initContainerConfig)
	}

	c.PostEditContainer(rabbitmqContainerConfig)

	return util.CreateReplicationControllerFromContainer(
		&horizonapi.ReplicationControllerConfig{Namespace: c.hubSpec.Namespace, Name: util.GetResourceName(c.name, util.BlackDuckName, "rabbitmq", c.config.IsClusterScoped), Replicas: util.IntToInt32(1)},
		&util.PodConfig{
			Volumes:             c.getRabbitmqVolumes(),
			Containers:          []*util.Container{rabbitmqContainerConfig},
			InitContainers:      initContainers,
			ImagePullSecrets:    c.hubSpec.RegistryConfiguration.PullSecrets,
			Labels:              c.GetVersionLabel("rabbitmq"),
			NodeAffinityConfigs: c.GetNodeAffinityConfigs("rabbitmq"),
		}, c.GetLabel("rabbitmq"))
}

// getRabbitmqVolumes will return the rabbitmq volumes
func (c *Creater) getRabbitmqVolumes() []*components.Volume {
	rabbitmqSecurityEmptyDir, _ := util.CreateEmptyDirVolumeWithoutSizeLimit("dir-rabbitmq-security")
	var rabbitmqDataEmptyDir *components.Volume
	if c.hubSpec.PersistentStorage {
		rabbitmqDataEmptyDir, _ = util.CreatePersistentVolumeClaimVolume("dir-rabbitmq-data", util.GetResourceName(c.name, util.BlackDuckName, "rabbitmq", c.config.IsClusterScoped))
	} else {
		rabbitmqDataEmptyDir, _ = util.CreateEmptyDirVolumeWithoutSizeLimit("dir-rabbitmq-data")
	}
	volumes := []*components.Volume{rabbitmqSecurityEmptyDir, rabbitmqDataEmptyDir}
	return volumes
}

// getRabbitmqVolumeMounts will return the rabbitmq volume mounts
func (c *Creater) getRabbitmqVolumeMounts() []*horizonapi.VolumeMountConfig {
	volumesMounts := []*horizonapi.VolumeMountConfig{
		{Name: "dir-rabbitmq-security", MountPath: "/opt/blackduck/rabbitmq/security"},
		{Name: "dir-rabbitmq-data", MountPath: "/var/lib/rabbitmq"},
	}
	return volumesMounts
}

// GetRabbitmqService will return the rabbitmq service
func (c *Creater) GetRabbitmqService() *components.Service {
	return util.CreateService(util.GetResourceName(c.name, util.BlackDuckName, "rabbitmq", c.config.IsClusterScoped), c.GetVersionLabel("rabbitmq"), c.hubSpec.Namespace, rabbitmqPort, rabbitmqPort, horizonapi.ServiceTypeServiceIP, c.GetVersionLabel("rabbitmq"))
}
