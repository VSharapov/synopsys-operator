/*
Copyright (C) 2018 Synopsys, Inc.

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

package util

import (
	"fmt"
	"strings"

	"github.com/blackducksoftware/synopsys-operator/pkg/api/hub/v1"
	hubClient "github.com/blackducksoftware/synopsys-operator/pkg/hub/client/clientset/versioned"
	"github.com/blackducksoftware/synopsys-operator/pkg/util"
	"github.com/sirupsen/logrus"
	"k8s.io/client-go/kubernetes"
)

// GetHubVersion will return the Hub version from the list of Hub environment variables
func GetHubVersion(environs []string) string {
	for _, value := range environs {
		if strings.Contains(value, "HUB_VERSION") {
			values := strings.SplitN(value, ":", 2)
			if len(values) == 2 {
				return strings.Trim(values[1], " ")
			}
			break
		}
	}
	return ""
}

// GetDefaultPasswords returns admin,user,postgres passwords for db maintainance tasks.  Should only be used during
// initialization, or for 'babysitting' ephemeral hub instances (which might have postgres restarts)
// MAKE SURE YOU SEND THE NAMESPACE OF THE SECRET SOURCE (operator), NOT OF THE new hub  THAT YOUR TRYING TO CREATE !
func GetDefaultPasswords(kubeClient *kubernetes.Clientset, nsOfSecretHolder string) (adminPassword string, userPassword string, postgresPassword string, err error) {
	blackduckSecret, err := util.GetSecret(kubeClient, nsOfSecretHolder, "blackduck-secret")
	if err != nil {
		logrus.Infof("warning: You need to first create a 'blackduck-secret' in this namespace with ADMIN_PASSWORD, USER_PASSWORD, POSTGRES_PASSWORD")
		return "", "", "", err
	}
	adminPassword = string(blackduckSecret.Data["ADMIN_PASSWORD"])
	userPassword = string(blackduckSecret.Data["USER_PASSWORD"])
	postgresPassword = string(blackduckSecret.Data["POSTGRES_PASSWORD"])

	// default named return
	return adminPassword, userPassword, postgresPassword, err
}

func updateHubObject(h *hubClient.Clientset, namespace string, obj *v1.Hub) (*v1.Hub, error) {
	return h.SynopsysV1().Hubs(namespace).Update(obj)
}

// UpdateState will be used to update the hub object
func UpdateState(h *hubClient.Clientset, namespace string, specState string, statusState string, err error, hub *v1.Hub) (*v1.Hub, error) {
	hub.Spec.State = specState
	hub.Status.State = statusState
	if err != nil {
		hub.Status.ErrorMessage = fmt.Sprintf("%+v", err)
	}
	hub, err = updateHubObject(h, namespace, hub)
	if err != nil {
		logrus.Errorf("couldn't update the state of hub object: %s", err.Error())
	}
	return hub, err
}
