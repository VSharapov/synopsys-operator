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
	"crypto/x509"
	"encoding/pem"
	"fmt"

	horizonapi "github.com/blackducksoftware/horizon/pkg/api"
	"github.com/blackducksoftware/horizon/pkg/components"
	"github.com/blackducksoftware/synopsys-operator/pkg/util"
	"github.com/sirupsen/logrus"
)

// GetSecrets will return the secrets
func (c *Creater) GetSecrets(cert string, key string, sealKey []byte) []*components.Secret {
	var secrets []*components.Secret

	certificateSecret := components.NewSecret(horizonapi.SecretConfig{Namespace: c.hubSpec.Namespace, Name: util.GetResourceName(c.name, util.BlackDuckName, "webserver-certificate", c.config.IsClusterScoped), Type: horizonapi.SecretTypeOpaque})
	certificateSecret.AddData(map[string][]byte{"WEBSERVER_CUSTOM_CERT_FILE": []byte(cert), "WEBSERVER_CUSTOM_KEY_FILE": []byte(key)})
	certificateSecret.AddLabels(c.GetVersionLabel("secret"))
	secrets = append(secrets, certificateSecret)

	if len(c.hubSpec.ProxyCertificate) > 0 {
		cert, err := c.stringToCertificate(c.hubSpec.ProxyCertificate)
		if err != nil {
			logrus.Warnf("The proxy certificate provided is invalid")
		} else {
			logrus.Debugf("Adding Proxy certificate with SN: %x", cert.SerialNumber)
			proxyCertificateSecret := components.NewSecret(horizonapi.SecretConfig{Namespace: c.hubSpec.Namespace, Name: util.GetResourceName(c.name, util.BlackDuckName, "proxy-certificate", c.config.IsClusterScoped), Type: horizonapi.SecretTypeOpaque})
			proxyCertificateSecret.AddData(map[string][]byte{"HUB_PROXY_CERT_FILE": []byte(c.hubSpec.ProxyCertificate)})
			proxyCertificateSecret.AddLabels(c.GetVersionLabel("secret"))
			secrets = append(secrets, proxyCertificateSecret)
		}
	}

	if len(c.hubSpec.AuthCustomCA) > 0 {
		cert, err := c.stringToCertificate(c.hubSpec.AuthCustomCA)
		if err != nil {
			logrus.Warnf("The Auth Custom CA provided is invalid")
		} else {
			logrus.Debugf("Adding The Auth Custom CA with SN: %x", cert.SerialNumber)
			authCustomCASecret := components.NewSecret(horizonapi.SecretConfig{Namespace: c.hubSpec.Namespace, Name: util.GetResourceName(c.name, util.BlackDuckName, "auth-custom-ca", c.config.IsClusterScoped), Type: horizonapi.SecretTypeOpaque})
			authCustomCASecret.AddData(map[string][]byte{"AUTH_CUSTOM_CA": []byte(c.hubSpec.AuthCustomCA)})
			authCustomCASecret.AddLabels(c.GetVersionLabel("secret"))
			secrets = append(secrets, authCustomCASecret)
		}
	}

	// creating secrets to store the seal key
	uploadCacheSecret := components.NewSecret(horizonapi.SecretConfig{Namespace: c.hubSpec.Namespace, Name: util.GetResourceName(c.name, util.BlackDuckName, "upload-cache", c.config.IsClusterScoped), Type: horizonapi.SecretTypeOpaque})
	uploadCacheSecret.AddData(map[string][]byte{"SEAL_KEY": sealKey})
	uploadCacheSecret.AddLabels(c.GetVersionLabel("uploadcache"))
	secrets = append(secrets, uploadCacheSecret)

	return secrets
}

func (c *Creater) stringToCertificate(certificate string) (*x509.Certificate, error) {
	block, _ := pem.Decode([]byte(certificate))
	if block == nil {
		return nil, fmt.Errorf("failed to parse certificate PEM")
	}

	cert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		return nil, err
	}

	return cert, nil
}
