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
	"fmt"
	"strings"

	horizonapi "github.com/blackducksoftware/horizon/pkg/api"
	"github.com/blackducksoftware/horizon/pkg/components"
	"github.com/blackducksoftware/synopsys-operator/pkg/util"
)

// GetConfigmaps will return the configMaps
func (c *Creater) GetConfigmaps() []*components.ConfigMap {

	var configMaps []*components.ConfigMap

	hubConfig := components.NewConfigMap(horizonapi.ConfigMapConfig{Namespace: c.hubSpec.Namespace, Name: util.GetResourceName(c.name, util.BlackDuckName, "config", c.config.IsClusterScoped)})

	hubData := map[string]string{
		"RUN_SECRETS_DIR": "/tmp/secrets",
		"HUB_VERSION":     c.hubSpec.Version,
	}

	blackduckServiceData := map[string]string{
		"HUB_AUTHENTICATION_HOST": util.GetResourceName(c.name, util.BlackDuckName, "authentication", c.config.IsClusterScoped),
		"AUTHENTICATION_HOST":     fmt.Sprintf("%s:%d", util.GetResourceName(c.name, util.BlackDuckName, "authentication", c.config.IsClusterScoped), authenticationPort),
		"CLIENT_CERT_CN":          util.GetResourceName(c.name, util.BlackDuckName, "binaryscanner", c.config.IsClusterScoped),
		"CFSSL":                   fmt.Sprintf("%s:8888", util.GetResourceName(c.name, util.BlackDuckName, "cfssl", c.config.IsClusterScoped)),
		"HUB_CFSSL_HOST":          util.GetResourceName(c.name, util.BlackDuckName, "cfssl", c.config.IsClusterScoped),
		"BLACKDUCK_CFSSL_HOST":    util.GetResourceName(c.name, util.BlackDuckName, "cfssl", c.config.IsClusterScoped),
		"BLACKDUCK_CFSSL_PORT":    "8888",
		"HUB_DOC_HOST":            util.GetResourceName(c.name, util.BlackDuckName, "documentation", c.config.IsClusterScoped),
		"HUB_JOBRUNNER_HOST":      util.GetResourceName(c.name, util.BlackDuckName, "jobrunner", c.config.IsClusterScoped),
		"HUB_LOGSTASH_HOST":       util.GetResourceName(c.name, util.BlackDuckName, "logstash", c.config.IsClusterScoped),
		"RABBIT_MQ_HOST":          util.GetResourceName(c.name, util.BlackDuckName, "rabbitmq", c.config.IsClusterScoped),
		"BROKER_URL":              fmt.Sprintf("amqps://%s/protecodesc", util.GetResourceName(c.name, util.BlackDuckName, "rabbitmq", c.config.IsClusterScoped)),
		"HUB_REGISTRATION_HOST":   util.GetResourceName(c.name, util.BlackDuckName, "registration", c.config.IsClusterScoped),
		"HUB_SCAN_HOST":           util.GetResourceName(c.name, util.BlackDuckName, "scan", c.config.IsClusterScoped),
		"HUB_SOLR_HOST":           util.GetResourceName(c.name, util.BlackDuckName, "solr", c.config.IsClusterScoped),
		// TODO: commented the below 2 environs until the HUB-20412 is fixed. once it if fixed, uncomment them
		// "BLACKDUCK_UPLOAD_CACHE_HOST": util.GetResourceName(c.name, util.BlackDuckName, "uploadcache", c.config.IsClusterScoped),
		// "HUB_UPLOAD_CACHE_HOST":       util.GetResourceName(c.name, util.BlackDuckName, "uploadcache", c.config.IsClusterScoped),
		// TODO: commented the below environs until the HUB-20462 is fixed. once it if fixed, uncomment them
		// "HUB_WEBAPP_HOST":    util.GetResourceName(c.name, util.BlackDuckName, "webapp", c.config.IsClusterScoped),
		"HUB_WEBSERVER_HOST": util.GetResourceName(c.name, util.BlackDuckName, "webserver", c.config.IsClusterScoped),
		"HUB_ZOOKEEPER_HOST": util.GetResourceName(c.name, util.BlackDuckName, "zookeeper", c.config.IsClusterScoped),
	}
	hubData = util.MergeEnvMaps(blackduckServiceData, hubData)

	for _, value := range c.hubSpec.Environs {
		values := strings.SplitN(value, ":", 2)
		if len(values) == 2 {
			mapKey := strings.TrimSpace(values[0])
			mapValue := strings.TrimSpace(values[1])
			if len(mapKey) > 0 && len(mapValue) > 0 {
				hubData[mapKey] = mapValue
			}
		}
	}

	// merge default and input environs
	environs := GetHubKnobs()
	hubData = util.MergeEnvMaps(hubData, environs)

	hubConfig.AddData(hubData)
	hubConfig.AddLabels(c.GetVersionLabel("configmap"))
	configMaps = append(configMaps, hubConfig)

	return configMaps
}
