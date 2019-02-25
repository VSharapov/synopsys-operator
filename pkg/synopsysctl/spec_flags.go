// Copyright © 2019 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package synopsysctl

import (
	"encoding/json"

	alertv1 "github.com/blackducksoftware/synopsys-operator/pkg/api/alert/v1"
	blackduckv1 "github.com/blackducksoftware/synopsys-operator/pkg/api/blackduck/v1"
	opssightv1 "github.com/blackducksoftware/synopsys-operator/pkg/api/opssight/v1"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

// Gloabal Specs
var globalBlackduckSpec = &blackduckv1.BlackduckSpec{}
var globalOpsSightSpec = &opssightv1.OpsSightSpec{}
var globalAlertSpec = &alertv1.AlertSpec{}

// Blackduck Spec Flags
var blackduckSize = ""
var blackduckDbPrototype = ""
var blackduckExternalPostgresPostgresHost = ""
var blackduckExternalPostgresPostgresPort = 0
var blackduckExternalPostgresPostgresAdmin = ""
var blackduckExternalPostgresPostgresUser = ""
var blackduckExternalPostgresPostgresSsl = false
var blackduckExternalPostgresPostgresAdminPassword = ""
var blackduckExternalPostgresPostgresUserPassword = ""
var blackduckPvcStorageClass = ""
var blackduckLivenessProbes = false
var blackduckScanType = ""
var blackduckPersistentStorage = false
var blackduckPVCJSONSlice = []string{}
var blackduckCertificateName = ""
var blackduckCertificate = ""
var blackduckCertificateKey = ""
var blackduckProxyCertificate = ""
var blackduckType = ""
var blackduckDesiredState = ""
var blackduckEnvirons = []string{}
var blackduckImageRegistries = []string{}
var blackduckImageUIDMapJSONSlice = []string{}
var blackduckLicenseKey = ""

// OpsSight Spec Flags
var opssightPerceptorName = ""
var opssightPerceptorImage = ""
var opssightPerceptorPort = 0
var opssightPerceptorCheckForStalledScansPauseHours = 0
var opssightPerceptorStalledScanClientTimeoutHours = 0
var opssightPerceptorModelMetricsPauseSeconds = 0
var opssightPerceptorUnknownImagePauseMilliseconds = 0
var opssightPerceptorClientTimeoutMilliseconds = 0
var opssightScannerPodName = ""
var opssightScannerPodScannerName = ""
var opssightScannerPodScannerImage = ""
var opssightScannerPodScannerPort = 0
var opssightScannerPodScannerClientTimeoutSeconds = 0
var opssightScannerPodImageFacadeName = ""
var opssightScannerPodImageFacadeImage = ""
var opssightScannerPodImageFacadePort = 0
var opssightScannerPodImageFacadeInternalRegistriesJSONSlice = []string{}
var opssightScannerPodImageFacadeImagePullerType = ""
var opssightScannerPodImageFacadeServiceAccount = ""
var opssightScannerPodReplicaCount = 0
var opssightScannerPodImageDirectory = ""
var opssightPerceiverEnableImagePerceiver = false
var opssightPerceiverEnablePodPerceiver = false
var opssightPerceiverImagePerceiverName = ""
var opssightPerceiverImagePerceiverImage = ""
var opssightPerceiverPodPerceiverName = ""
var opssightPerceiverPodPerceiverImage = ""
var opssightPerceiverPodPerceiverNamespaceFilter = ""
var opssightPerceiverAnnotationIntervalSeconds = 0
var opssightPerceiverDumpIntervalMinutes = 0
var opssightPerceiverServiceAccount = ""
var opssightPerceiverPort = 0
var opssightPrometheusName = ""
var opssightPrometheusImage = ""
var opssightPrometheusPort = 0
var opssightEnableSkyfire = false
var opssightSkyfireName = ""
var opssightSkyfireImage = ""
var opssightSkyfirePort = 0
var opssightSkyfirePrometheusPort = 0
var opssightSkyfireServiceAccount = ""
var opssightSkyfireHubClientTimeoutSeconds = 0
var opssightSkyfireHubDumpPauseSeconds = 0
var opssightSkyfireKubeDumpIntervalSeconds = 0
var opssightSkyfirePerceptorDumpIntervalSeconds = 0
var opssightBlackduckHosts = []string{}
var opssightBlackduckUser = ""
var opssightBlackduckPort = 0
var opssightBlackduckConcurrentScanLimit = 0
var opssightBlackduckTotalScanLimit = 0
var opssightBlackduckPasswordEnvVar = ""
var opssightBlackduckInitialCount = 0
var opssightBlackduckMaxCount = 0
var opssightBlackduckDeleteHubThresholdPercentage = 0
var opssightEnableMetrics = false
var opssightDefaultCPU = ""
var opssightDefaultMem = ""
var opssightLogLevel = ""
var opssightConfigMapName = ""
var opssightSecretName = ""

// Create Alert Spec Flags
var alertRegistry = ""
var alertImagePath = ""
var alertAlertImageName = ""
var alertAlertImageVersion = ""
var alertCfsslImageName = ""
var alertCfsslImageVersion = ""
var alertBlackduckHost = ""
var alertBlackduckUser = ""
var alertBlackduckPort = 0
var alertPort = 0
var alertStandAlone = false
var alertAlertMemory = ""
var alertCfsslMemory = ""
var alertState = ""

func addBlackduckSpecFlags(cmd *cobra.Command) {
	cmd.Flags().StringVar(&blackduckSize, "size", blackduckSize, "Blackduck size - small, medium, large")
	cmd.Flags().StringVar(&blackduckDbPrototype, "db-prototype", blackduckDbPrototype, "TODO")
	cmd.Flags().StringVar(&blackduckExternalPostgresPostgresHost, "external-postgres-host", blackduckExternalPostgresPostgresHost, "TODO")
	cmd.Flags().IntVar(&blackduckExternalPostgresPostgresPort, "external-postgres-port", blackduckExternalPostgresPostgresPort, "TODO")
	cmd.Flags().StringVar(&blackduckExternalPostgresPostgresAdmin, "external-postgres-admin", blackduckExternalPostgresPostgresAdmin, "TODO")
	cmd.Flags().StringVar(&blackduckExternalPostgresPostgresUser, "external-postgres-user", blackduckExternalPostgresPostgresUser, "TODO")
	cmd.Flags().BoolVar(&blackduckExternalPostgresPostgresSsl, "external-postgres-ssl", blackduckExternalPostgresPostgresSsl, "TODO")
	cmd.Flags().StringVar(&blackduckExternalPostgresPostgresAdminPassword, "external-postgres-admin-password", blackduckExternalPostgresPostgresAdminPassword, "TODO")
	cmd.Flags().StringVar(&blackduckExternalPostgresPostgresUserPassword, "external-postgres-user-password", blackduckExternalPostgresPostgresUserPassword, "TODO")
	cmd.Flags().StringVar(&blackduckPvcStorageClass, "pvc-storage-class", blackduckPvcStorageClass, "TODO")
	cmd.Flags().BoolVar(&blackduckLivenessProbes, "liveness-probes", blackduckLivenessProbes, "Enable liveness probes")
	cmd.Flags().StringVar(&blackduckScanType, "scan-type", blackduckScanType, "TODO")
	cmd.Flags().BoolVar(&blackduckPersistentStorage, "persistent-storage", blackduckPersistentStorage, "Enable persistent storage")
	cmd.Flags().StringSliceVar(&blackduckPVCJSONSlice, "pvc", blackduckPVCJSONSlice, "TODO")
	cmd.Flags().StringVar(&blackduckCertificateName, "db-certificate-name", blackduckCertificateName, "TODO")
	cmd.Flags().StringVar(&blackduckCertificate, "certificate", blackduckCertificate, "TODO")
	cmd.Flags().StringVar(&blackduckCertificateKey, "certificate-key", blackduckCertificateKey, "TODO")
	cmd.Flags().StringVar(&blackduckProxyCertificate, "proxy-certificate", blackduckProxyCertificate, "TODO")
	cmd.Flags().StringVar(&blackduckType, "type", blackduckType, "TODO")
	cmd.Flags().StringVar(&blackduckDesiredState, "desired-state", blackduckDesiredState, "TODO")
	cmd.Flags().StringSliceVar(&blackduckEnvirons, "environs", blackduckEnvirons, "TODO")
	cmd.Flags().StringSliceVar(&blackduckImageRegistries, "image-registries", blackduckImageRegistries, "List of image registries")
	cmd.Flags().StringSliceVar(&blackduckImageUIDMapJSONSlice, "image-uid-map", blackduckImageUIDMapJSONSlice, "TODO")
	cmd.Flags().StringVar(&blackduckLicenseKey, "license-key", blackduckLicenseKey, "TODO")
}

func addOpsSightSpecFlags(cmd *cobra.Command) {
	cmd.Flags().StringVar(&opssightPerceptorName, "perceptor-name", opssightPerceptorName, "TODO")
	cmd.Flags().StringVar(&opssightPerceptorImage, "perceptor-image", opssightPerceptorImage, "TODO")
	cmd.Flags().IntVar(&opssightPerceptorPort, "perceptor-port", opssightPerceptorPort, "TODO")
	cmd.Flags().IntVar(&opssightPerceptorCheckForStalledScansPauseHours, "perceptor-check-scan-hours", opssightPerceptorCheckForStalledScansPauseHours, "TODO")
	cmd.Flags().IntVar(&opssightPerceptorStalledScanClientTimeoutHours, "perceptor-scan-client-timeout-hours", opssightPerceptorStalledScanClientTimeoutHours, "TODO")
	cmd.Flags().IntVar(&opssightPerceptorModelMetricsPauseSeconds, "perceptor-metrics-pause-seconds", opssightPerceptorModelMetricsPauseSeconds, "TODO")
	cmd.Flags().IntVar(&opssightPerceptorUnknownImagePauseMilliseconds, "perceptor-unknown-image-pause-milliseconds", opssightPerceptorUnknownImagePauseMilliseconds, "TODO")
	cmd.Flags().IntVar(&opssightPerceptorClientTimeoutMilliseconds, "perceptor-client-timeout-milliseconds", opssightPerceptorClientTimeoutMilliseconds, "TODO")
	cmd.Flags().StringVar(&opssightScannerPodName, "scannerpod-name", opssightScannerPodName, "TODO")
	cmd.Flags().StringVar(&opssightScannerPodScannerName, "scannerpod-scanner-name", opssightScannerPodScannerName, "TODO")
	cmd.Flags().StringVar(&opssightScannerPodScannerImage, "scannerpod-scanner-image", opssightScannerPodScannerImage, "TODO")
	cmd.Flags().IntVar(&opssightScannerPodScannerPort, "scannerpod-scanner-port", opssightScannerPodScannerPort, "TODO")
	cmd.Flags().IntVar(&opssightScannerPodScannerClientTimeoutSeconds, "scannerpod-scanner-client-timeout-seconds", opssightScannerPodScannerClientTimeoutSeconds, "TODO")
	cmd.Flags().StringVar(&opssightScannerPodImageFacadeName, "scannerpod-imagefacade-name", opssightScannerPodImageFacadeName, "TODO")
	cmd.Flags().StringVar(&opssightScannerPodImageFacadeImage, "scannerpod-imagefacade-image", opssightScannerPodImageFacadeImage, "TODO")
	cmd.Flags().IntVar(&opssightScannerPodImageFacadePort, "scannerpod-imagefacade-port", opssightScannerPodImageFacadePort, "TODO")
	cmd.Flags().StringSliceVar(&opssightScannerPodImageFacadeInternalRegistriesJSONSlice, "scannerpod-imagefacade-internal-registries", opssightScannerPodImageFacadeInternalRegistriesJSONSlice, "TODO")
	cmd.Flags().StringVar(&opssightScannerPodImageFacadeImagePullerType, "scannerpod-imagefacade-image-puller-type", opssightScannerPodImageFacadeImagePullerType, "TODO")
	cmd.Flags().StringVar(&opssightScannerPodImageFacadeServiceAccount, "scannerpod-imagefacade-service-account", opssightScannerPodImageFacadeServiceAccount, "TODO")
	cmd.Flags().IntVar(&opssightScannerPodReplicaCount, "scannerpod-replica-count", opssightScannerPodReplicaCount, "TODO")
	cmd.Flags().StringVar(&opssightScannerPodImageDirectory, "scannerpod-image-directory", opssightScannerPodImageDirectory, "TODO")
	cmd.Flags().BoolVar(&opssightPerceiverEnableImagePerceiver, "enable-image-perceiver", opssightPerceiverEnableImagePerceiver, "TODO")
	cmd.Flags().BoolVar(&opssightPerceiverEnablePodPerceiver, "enable-pod-perceiver", opssightPerceiverEnablePodPerceiver, "TODO")
	cmd.Flags().StringVar(&opssightPerceiverImagePerceiverName, "imageperceiver-name", opssightPerceiverImagePerceiverName, "TODO")
	cmd.Flags().StringVar(&opssightPerceiverImagePerceiverImage, "imageperceiver-image", opssightPerceiverImagePerceiverImage, "TODO")
	cmd.Flags().StringVar(&opssightPerceiverPodPerceiverName, "podperceiver-name", opssightPerceiverPodPerceiverName, "TODO")
	cmd.Flags().StringVar(&opssightPerceiverPodPerceiverImage, "podperceiver-image", opssightPerceiverPodPerceiverImage, "TODO")
	cmd.Flags().StringVar(&opssightPerceiverPodPerceiverNamespaceFilter, "podperceiver-namespace-filter", opssightPerceiverPodPerceiverNamespaceFilter, "TODO")
	cmd.Flags().IntVar(&opssightPerceiverAnnotationIntervalSeconds, "perceiver-annotation-interval-seconds", opssightPerceiverAnnotationIntervalSeconds, "TODO")
	cmd.Flags().IntVar(&opssightPerceiverDumpIntervalMinutes, "perceiver-dump-interval-minutes", opssightPerceiverDumpIntervalMinutes, "TODO")
	cmd.Flags().StringVar(&opssightPerceiverServiceAccount, "perceiver-service-account", opssightPerceiverServiceAccount, "TODO")
	cmd.Flags().IntVar(&opssightPerceiverPort, "perceiver-port", opssightPerceiverPort, "TODO")
	cmd.Flags().StringVar(&opssightPrometheusName, "prometheus-name", opssightPrometheusName, "TODO")
	cmd.Flags().StringVar(&opssightPrometheusName, "prometheus-image", opssightPrometheusName, "TODO")
	cmd.Flags().IntVar(&opssightPrometheusPort, "prometheus-port", opssightPrometheusPort, "TODO")
	cmd.Flags().BoolVar(&opssightEnableSkyfire, "enable-skyfire", opssightEnableSkyfire, "TODO")
	cmd.Flags().StringVar(&opssightSkyfireName, "skyfire-name", opssightSkyfireName, "TODO")
	cmd.Flags().StringVar(&opssightSkyfireImage, "skyfire-image", opssightSkyfireImage, "TODO")
	cmd.Flags().IntVar(&opssightSkyfirePort, "skyfire-port", opssightSkyfirePort, "TODO")
	cmd.Flags().IntVar(&opssightSkyfirePrometheusPort, "skyfire-prometheus-port", opssightSkyfirePrometheusPort, "TODO")
	cmd.Flags().StringVar(&opssightSkyfireServiceAccount, "skyfire-service-account", opssightSkyfireServiceAccount, "TODO")
	cmd.Flags().IntVar(&opssightSkyfireHubClientTimeoutSeconds, "skyfire-hub-client-timeout-seconds", opssightSkyfireHubClientTimeoutSeconds, "TODO")
	cmd.Flags().IntVar(&opssightSkyfireHubDumpPauseSeconds, "skyfire-hub-dump-pause-seconds", opssightSkyfireHubDumpPauseSeconds, "TODO")
	cmd.Flags().IntVar(&opssightSkyfireKubeDumpIntervalSeconds, "skyfire-kube-dump-interval-seconds", opssightSkyfireKubeDumpIntervalSeconds, "TODO")
	cmd.Flags().IntVar(&opssightSkyfirePerceptorDumpIntervalSeconds, "skyfire-perceptor-dump-interval-seconds", opssightSkyfirePerceptorDumpIntervalSeconds, "TODO")
	cmd.Flags().StringSliceVar(&opssightBlackduckHosts, "blackduck-hosts", opssightBlackduckHosts, "TODO")
	cmd.Flags().StringVar(&opssightBlackduckUser, "blackduck-user", opssightBlackduckUser, "TODO")
	cmd.Flags().IntVar(&opssightBlackduckPort, "blackduck-port", opssightBlackduckPort, "TODO")
	cmd.Flags().IntVar(&opssightBlackduckConcurrentScanLimit, "blackduck-concurrent-scan-limit", opssightBlackduckConcurrentScanLimit, "TODO")
	cmd.Flags().IntVar(&opssightBlackduckTotalScanLimit, "blackduck-total-scan-limit", opssightBlackduckTotalScanLimit, "TODO")
	cmd.Flags().StringVar(&opssightBlackduckPasswordEnvVar, "blackduck-password-environment-variable", opssightBlackduckPasswordEnvVar, "TODO")
	cmd.Flags().IntVar(&opssightBlackduckInitialCount, "blackduck-initial-count", opssightBlackduckInitialCount, "TODO")
	cmd.Flags().IntVar(&opssightBlackduckMaxCount, "blackduck-max-count", opssightBlackduckMaxCount, "TODO")
	cmd.Flags().IntVar(&opssightBlackduckDeleteHubThresholdPercentage, "blackduck-delete-blackduck-threshold-percentage", opssightBlackduckDeleteHubThresholdPercentage, "TODO")
	cmd.Flags().BoolVar(&opssightEnableMetrics, "enable-metrics", opssightEnableMetrics, "TODO")
	cmd.Flags().StringVar(&opssightDefaultCPU, "default-cpu", opssightDefaultCPU, "TODO")
	cmd.Flags().StringVar(&opssightDefaultMem, "default-mem", opssightDefaultMem, "TODO")
	cmd.Flags().StringVar(&opssightLogLevel, "log-level", opssightLogLevel, "TODO")
	cmd.Flags().StringVar(&opssightConfigMapName, "config-map-name", opssightConfigMapName, "TODO")
	cmd.Flags().StringVar(&opssightSecretName, "secret-name", opssightSecretName, "TODO")
}

func addAlertSpecFlags(cmd *cobra.Command) {
	cmd.Flags().StringVar(&alertRegistry, "alert-registry", alertRegistry, "TODO")
	cmd.Flags().StringVar(&alertImagePath, "image-path", alertImagePath, "TODO")
	cmd.Flags().StringVar(&alertAlertImageName, "alert-image-name", alertAlertImageName, "TODO")
	cmd.Flags().StringVar(&alertAlertImageVersion, "alert-image-version", alertAlertImageVersion, "TODO")
	cmd.Flags().StringVar(&alertCfsslImageName, "cfssl-image-name", alertCfsslImageName, "TODO")
	cmd.Flags().StringVar(&alertCfsslImageVersion, "cfssl-image-version", alertCfsslImageVersion, "TODO")
	cmd.Flags().StringVar(&alertBlackduckHost, "blackduck-host", alertBlackduckHost, "TODO")
	cmd.Flags().StringVar(&alertBlackduckUser, "blackduck-user", alertBlackduckUser, "TODO")
	cmd.Flags().IntVar(&alertBlackduckPort, "blackduck-port", alertBlackduckPort, "TODO")
	cmd.Flags().IntVar(&alertPort, "port", alertPort, "TODO")
	cmd.Flags().BoolVar(&alertStandAlone, "stand-alone", alertStandAlone, "TODO")
	cmd.Flags().StringVar(&alertAlertMemory, "alert-memory", alertAlertMemory, "TODO")
	cmd.Flags().StringVar(&alertCfsslMemory, "cfssl-memory", alertCfsslMemory, "TODO")
	cmd.Flags().StringVar(&alertState, "alert-state", alertState, "TODO")
}

func setBlackduckFlags(f *pflag.Flag) {
	if f.Changed {
		log.Debugf("Flag %s: CHANGED\n", f.Name)
		switch f.Name {
		case "size":
			globalBlackduckSpec.Size = blackduckSize
		case "db-prototype":
			globalBlackduckSpec.DbPrototype = blackduckDbPrototype
		case "external-postgres-host":
			if globalBlackduckSpec.ExternalPostgres == nil {
				globalBlackduckSpec.ExternalPostgres = &blackduckv1.PostgresExternalDBConfig{}
			}
			globalBlackduckSpec.ExternalPostgres.PostgresHost = blackduckExternalPostgresPostgresHost
		case "external-postgres-port":
			if globalBlackduckSpec.ExternalPostgres == nil {
				globalBlackduckSpec.ExternalPostgres = &blackduckv1.PostgresExternalDBConfig{}
			}
			globalBlackduckSpec.ExternalPostgres.PostgresPort = blackduckExternalPostgresPostgresPort
		case "external-postgres-admin":
			if globalBlackduckSpec.ExternalPostgres == nil {
				globalBlackduckSpec.ExternalPostgres = &blackduckv1.PostgresExternalDBConfig{}
			}
			globalBlackduckSpec.ExternalPostgres.PostgresAdmin = blackduckExternalPostgresPostgresAdmin
		case "external-postgres-user":
			if globalBlackduckSpec.ExternalPostgres == nil {
				globalBlackduckSpec.ExternalPostgres = &blackduckv1.PostgresExternalDBConfig{}
			}
			globalBlackduckSpec.ExternalPostgres.PostgresUser = blackduckExternalPostgresPostgresUser
		case "external-postgres-ssl":
			if globalBlackduckSpec.ExternalPostgres == nil {
				globalBlackduckSpec.ExternalPostgres = &blackduckv1.PostgresExternalDBConfig{}
			}
			globalBlackduckSpec.ExternalPostgres.PostgresSsl = blackduckExternalPostgresPostgresSsl
		case "external-postgres-admin-password":
			if globalBlackduckSpec.ExternalPostgres == nil {
				globalBlackduckSpec.ExternalPostgres = &blackduckv1.PostgresExternalDBConfig{}
			}
			globalBlackduckSpec.ExternalPostgres.PostgresAdminPassword = blackduckExternalPostgresPostgresAdminPassword
		case "external-postgres-user-password":
			if globalBlackduckSpec.ExternalPostgres == nil {
				globalBlackduckSpec.ExternalPostgres = &blackduckv1.PostgresExternalDBConfig{}
			}
			globalBlackduckSpec.ExternalPostgres.PostgresUserPassword = blackduckExternalPostgresPostgresUserPassword
		case "pvc-storage-class":
			if globalBlackduckSpec.ExternalPostgres == nil {
				globalBlackduckSpec.ExternalPostgres = &blackduckv1.PostgresExternalDBConfig{}
			}
			globalBlackduckSpec.PVCStorageClass = blackduckPvcStorageClass
		case "liveness-probes":
			globalBlackduckSpec.LivenessProbes = blackduckLivenessProbes
		case "scan-type":
			globalBlackduckSpec.ScanType = blackduckScanType
		case "persistent-storage":
			globalBlackduckSpec.PersistentStorage = blackduckPersistentStorage
		case "pvc":
			for _, pvcJSON := range blackduckPVCJSONSlice {
				pvc := &blackduckv1.PVC{}
				json.Unmarshal([]byte(pvcJSON), pvc)
				globalBlackduckSpec.PVC = append(globalBlackduckSpec.PVC, *pvc)
			}
		case "db-certificate-name":
			globalBlackduckSpec.CertificateName = blackduckCertificateName
		case "certificate":
			globalBlackduckSpec.Certificate = blackduckCertificate
		case "certificate-key":
			globalBlackduckSpec.CertificateKey = blackduckCertificateKey
		case "proxy-certificate":
			globalBlackduckSpec.ProxyCertificate = blackduckProxyCertificate
		case "type":
			globalBlackduckSpec.Type = blackduckType
		case "desired-state":
			globalBlackduckSpec.DesiredState = blackduckDesiredState
		case "environs":
			globalBlackduckSpec.Environs = blackduckEnvirons
		case "image-registries":
			globalBlackduckSpec.ImageRegistries = blackduckImageRegistries
		case "image-uid-map":
			type uid struct {
				Key   string `json:"key"`
				Value int64  `json:"value"`
			}
			globalBlackduckSpec.ImageUIDMap = make(map[string]int64)
			for _, uidJSON := range blackduckImageUIDMapJSONSlice {
				uidStruct := &uid{}
				json.Unmarshal([]byte(uidJSON), uidStruct)
				globalBlackduckSpec.ImageUIDMap[uidStruct.Key] = uidStruct.Value
			}
		case "license-key":
			globalBlackduckSpec.LicenseKey = blackduckLicenseKey
		default:
			log.Debugf("Flag %s: Not Found\n", f.Name)
		}
	}
	log.Debugf("Flag %s: UNCHANGED\n", f.Name)
}

func setOpsSightFlags(f *pflag.Flag) {
	if f.Changed {
		log.Debugf("Flag %s: CHANGED\n", f.Name)
		switch f.Name {
		case "perceptor-name":
			if globalOpsSightSpec.Perceptor == nil {
				globalOpsSightSpec.Perceptor = &opssightv1.Perceptor{}
			}
			globalOpsSightSpec.Perceptor.Name = opssightPerceptorName
		case "perceptor-image":
			if globalOpsSightSpec.Perceptor == nil {
				globalOpsSightSpec.Perceptor = &opssightv1.Perceptor{}
			}
			globalOpsSightSpec.Perceptor.Image = opssightPerceptorImage
		case "perceptor-port":
			if globalOpsSightSpec.Perceptor == nil {
				globalOpsSightSpec.Perceptor = &opssightv1.Perceptor{}
			}
			globalOpsSightSpec.Perceptor.Port = opssightPerceptorPort
		case "perceptor-check-scan-hours":
			if globalOpsSightSpec.Perceptor == nil {
				globalOpsSightSpec.Perceptor = &opssightv1.Perceptor{}
			}
			globalOpsSightSpec.Perceptor.CheckForStalledScansPauseHours = opssightPerceptorCheckForStalledScansPauseHours
		case "perceptor-scan-client-timeout-hours":
			if globalOpsSightSpec.Perceptor == nil {
				globalOpsSightSpec.Perceptor = &opssightv1.Perceptor{}
			}
			globalOpsSightSpec.Perceptor.StalledScanClientTimeoutHours = opssightPerceptorStalledScanClientTimeoutHours
		case "perceptor-metrics-pause-seconds":
			if globalOpsSightSpec.Perceptor == nil {
				globalOpsSightSpec.Perceptor = &opssightv1.Perceptor{}
			}
			globalOpsSightSpec.Perceptor.ModelMetricsPauseSeconds = opssightPerceptorModelMetricsPauseSeconds
		case "perceptor-unknown-image-pause-milliseconds":
			if globalOpsSightSpec.Perceptor == nil {
				globalOpsSightSpec.Perceptor = &opssightv1.Perceptor{}
			}
			globalOpsSightSpec.Perceptor.UnknownImagePauseMilliseconds = opssightPerceptorUnknownImagePauseMilliseconds
		case "perceptor-client-timeout-milliseconds":
			if globalOpsSightSpec.Perceptor == nil {
				globalOpsSightSpec.Perceptor = &opssightv1.Perceptor{}
			}
			globalOpsSightSpec.Perceptor.ClientTimeoutMilliseconds = opssightPerceptorClientTimeoutMilliseconds
		case "scannerpod-name":
			if globalOpsSightSpec.ScannerPod == nil {
				globalOpsSightSpec.ScannerPod = &opssightv1.ScannerPod{}
			}
			globalOpsSightSpec.ScannerPod.Name = opssightScannerPodName
		case "scannerpod-scanner-name":
			if globalOpsSightSpec.ScannerPod == nil {
				globalOpsSightSpec.ScannerPod = &opssightv1.ScannerPod{}
			}
			if globalOpsSightSpec.ScannerPod.Scanner == nil {
				globalOpsSightSpec.ScannerPod.Scanner = &opssightv1.Scanner{}
			}
			globalOpsSightSpec.ScannerPod.Scanner.Name = opssightScannerPodScannerName
		case "scannerpod-scanner-image":
			if globalOpsSightSpec.ScannerPod == nil {
				globalOpsSightSpec.ScannerPod = &opssightv1.ScannerPod{}
			}
			if globalOpsSightSpec.ScannerPod.Scanner == nil {
				globalOpsSightSpec.ScannerPod.Scanner = &opssightv1.Scanner{}
			}
			globalOpsSightSpec.ScannerPod.Scanner.Image = opssightScannerPodScannerImage
		case "scannerpod-scanner-port":
			if globalOpsSightSpec.ScannerPod == nil {
				globalOpsSightSpec.ScannerPod = &opssightv1.ScannerPod{}
			}
			if globalOpsSightSpec.ScannerPod.Scanner == nil {
				globalOpsSightSpec.ScannerPod.Scanner = &opssightv1.Scanner{}
			}
			globalOpsSightSpec.ScannerPod.Scanner.Port = opssightScannerPodScannerPort
		case "scannerpod-scanner-client-timeout-seconds":
			if globalOpsSightSpec.ScannerPod == nil {
				globalOpsSightSpec.ScannerPod = &opssightv1.ScannerPod{}
			}
			if globalOpsSightSpec.ScannerPod.Scanner == nil {
				globalOpsSightSpec.ScannerPod.Scanner = &opssightv1.Scanner{}
			}
			globalOpsSightSpec.ScannerPod.Scanner.ClientTimeoutSeconds = opssightScannerPodScannerClientTimeoutSeconds
		case "scannerpod-imagefacade-name":
			if globalOpsSightSpec.ScannerPod == nil {
				globalOpsSightSpec.ScannerPod = &opssightv1.ScannerPod{}
			}
			if globalOpsSightSpec.ScannerPod.ImageFacade == nil {
				globalOpsSightSpec.ScannerPod.ImageFacade = &opssightv1.ImageFacade{}
			}
			globalOpsSightSpec.ScannerPod.ImageFacade.Name = opssightScannerPodImageFacadeName
		case "scannerpod-imagefacade-image":
			if globalOpsSightSpec.ScannerPod == nil {
				globalOpsSightSpec.ScannerPod = &opssightv1.ScannerPod{}
			}
			if globalOpsSightSpec.ScannerPod.ImageFacade == nil {
				globalOpsSightSpec.ScannerPod.ImageFacade = &opssightv1.ImageFacade{}
			}
			globalOpsSightSpec.ScannerPod.ImageFacade.Image = opssightScannerPodImageFacadeImage
		case "scannerpod-imagefacade-port":
			if globalOpsSightSpec.ScannerPod == nil {
				globalOpsSightSpec.ScannerPod = &opssightv1.ScannerPod{}
			}
			if globalOpsSightSpec.ScannerPod.ImageFacade == nil {
				globalOpsSightSpec.ScannerPod.ImageFacade = &opssightv1.ImageFacade{}
			}
			globalOpsSightSpec.ScannerPod.ImageFacade.Port = opssightScannerPodImageFacadePort
		case "scannerpod-imagefacade-internal-registries":
			if globalOpsSightSpec.ScannerPod == nil {
				globalOpsSightSpec.ScannerPod = &opssightv1.ScannerPod{}
			}
			if globalOpsSightSpec.ScannerPod.ImageFacade == nil {
				globalOpsSightSpec.ScannerPod.ImageFacade = &opssightv1.ImageFacade{}
			}
			for _, registryJSON := range opssightScannerPodImageFacadeInternalRegistriesJSONSlice {
				registry := &opssightv1.RegistryAuth{}
				json.Unmarshal([]byte(registryJSON), registry)
				globalOpsSightSpec.ScannerPod.ImageFacade.InternalRegistries = append(globalOpsSightSpec.ScannerPod.ImageFacade.InternalRegistries, *registry)
			}
		case "scannerpod-imagefacade-image-puller-type":
			if globalOpsSightSpec.ScannerPod == nil {
				globalOpsSightSpec.ScannerPod = &opssightv1.ScannerPod{}
			}
			if globalOpsSightSpec.ScannerPod.ImageFacade == nil {
				globalOpsSightSpec.ScannerPod.ImageFacade = &opssightv1.ImageFacade{}
			}
			globalOpsSightSpec.ScannerPod.ImageFacade.ImagePullerType = opssightScannerPodImageFacadeImagePullerType
		case "scannerpod-imagefacade-service-account":
			if globalOpsSightSpec.ScannerPod == nil {
				globalOpsSightSpec.ScannerPod = &opssightv1.ScannerPod{}
			}
			if globalOpsSightSpec.ScannerPod.ImageFacade == nil {
				globalOpsSightSpec.ScannerPod.ImageFacade = &opssightv1.ImageFacade{}
			}
			globalOpsSightSpec.ScannerPod.ImageFacade.ServiceAccount = opssightScannerPodImageFacadeServiceAccount
		case "scannerpod-replica-count":
			if globalOpsSightSpec.ScannerPod == nil {
				globalOpsSightSpec.ScannerPod = &opssightv1.ScannerPod{}
			}
			globalOpsSightSpec.ScannerPod.ReplicaCount = opssightScannerPodReplicaCount
		case "scannerpod-image-directory":
			if globalOpsSightSpec.ScannerPod == nil {
				globalOpsSightSpec.ScannerPod = &opssightv1.ScannerPod{}
			}
			globalOpsSightSpec.ScannerPod.ImageDirectory = opssightScannerPodImageDirectory
		case "enable-image-perceiver":
			if globalOpsSightSpec.Perceiver == nil {
				globalOpsSightSpec.Perceiver = &opssightv1.Perceiver{}
			}
			globalOpsSightSpec.Perceiver.EnableImagePerceiver = opssightPerceiverEnableImagePerceiver
		case "enable-pod-perceiver":
			if globalOpsSightSpec.Perceiver == nil {
				globalOpsSightSpec.Perceiver = &opssightv1.Perceiver{}
			}
			globalOpsSightSpec.Perceiver.EnablePodPerceiver = opssightPerceiverEnablePodPerceiver
		case "imageperceiver-name":
			if globalOpsSightSpec.Perceiver == nil {
				globalOpsSightSpec.Perceiver = &opssightv1.Perceiver{}
			}
			if globalOpsSightSpec.Perceiver.ImagePerceiver == nil {
				globalOpsSightSpec.Perceiver.ImagePerceiver = &opssightv1.ImagePerceiver{}
			}
			globalOpsSightSpec.Perceiver.ImagePerceiver.Name = opssightPerceiverImagePerceiverName
		case "imageperceiver-image":
			if globalOpsSightSpec.Perceiver == nil {
				globalOpsSightSpec.Perceiver = &opssightv1.Perceiver{}
			}
			if globalOpsSightSpec.Perceiver.ImagePerceiver == nil {
				globalOpsSightSpec.Perceiver.ImagePerceiver = &opssightv1.ImagePerceiver{}
			}
			globalOpsSightSpec.Perceiver.ImagePerceiver.Image = opssightPerceiverImagePerceiverImage
		case "podperceiver-name":
			if globalOpsSightSpec.Perceiver == nil {
				globalOpsSightSpec.Perceiver = &opssightv1.Perceiver{}
			}
			if globalOpsSightSpec.Perceiver.PodPerceiver == nil {
				globalOpsSightSpec.Perceiver.PodPerceiver = &opssightv1.PodPerceiver{}
			}
			globalOpsSightSpec.Perceiver.PodPerceiver.Name = opssightPerceiverPodPerceiverName
		case "podperceiver-image":
			if globalOpsSightSpec.Perceiver == nil {
				globalOpsSightSpec.Perceiver = &opssightv1.Perceiver{}
			}
			if globalOpsSightSpec.Perceiver.PodPerceiver == nil {
				globalOpsSightSpec.Perceiver.PodPerceiver = &opssightv1.PodPerceiver{}
			}
			globalOpsSightSpec.Perceiver.PodPerceiver.Image = opssightPerceiverPodPerceiverImage
		case "podperceiver-namespace-filter":
			if globalOpsSightSpec.Perceiver == nil {
				globalOpsSightSpec.Perceiver = &opssightv1.Perceiver{}
			}
			if globalOpsSightSpec.Perceiver.PodPerceiver == nil {
				globalOpsSightSpec.Perceiver.PodPerceiver = &opssightv1.PodPerceiver{}
			}
			globalOpsSightSpec.Perceiver.PodPerceiver.NamespaceFilter = opssightPerceiverPodPerceiverNamespaceFilter
		case "perceiver-annotation-interval-seconds":
			if globalOpsSightSpec.Perceiver == nil {
				globalOpsSightSpec.Perceiver = &opssightv1.Perceiver{}
			}
			globalOpsSightSpec.Perceiver.AnnotationIntervalSeconds = opssightPerceiverAnnotationIntervalSeconds
		case "perceiver-dump-interval-minutes":
			if globalOpsSightSpec.Perceiver == nil {
				globalOpsSightSpec.Perceiver = &opssightv1.Perceiver{}
			}
			globalOpsSightSpec.Perceiver.DumpIntervalMinutes = opssightPerceiverDumpIntervalMinutes
		case "perceiver-service-account":
			if globalOpsSightSpec.Perceiver == nil {
				globalOpsSightSpec.Perceiver = &opssightv1.Perceiver{}
			}
			globalOpsSightSpec.Perceiver.ServiceAccount = opssightPerceiverServiceAccount
		case "perceiver-port":
			if globalOpsSightSpec.Perceiver == nil {
				globalOpsSightSpec.Perceiver = &opssightv1.Perceiver{}
			}
			globalOpsSightSpec.Perceiver.Port = opssightPerceiverPort
		case "prometheus-name":
			if globalOpsSightSpec.Prometheus == nil {
				globalOpsSightSpec.Prometheus = &opssightv1.Prometheus{}
			}
			globalOpsSightSpec.Prometheus.Name = opssightPrometheusName
		case "prometheus-image":
			if globalOpsSightSpec.Prometheus == nil {
				globalOpsSightSpec.Prometheus = &opssightv1.Prometheus{}
			}
			globalOpsSightSpec.Prometheus.Image = opssightPrometheusImage
		case "prometheus-port":
			if globalOpsSightSpec.Prometheus == nil {
				globalOpsSightSpec.Prometheus = &opssightv1.Prometheus{}
			}
			globalOpsSightSpec.Prometheus.Port = opssightPrometheusPort
		case "enable-skyfire":
			globalOpsSightSpec.EnableSkyfire = opssightEnableSkyfire
		case "skyfire-name":
			if globalOpsSightSpec.Skyfire == nil {
				globalOpsSightSpec.Skyfire = &opssightv1.Skyfire{}
			}
			globalOpsSightSpec.Skyfire.Name = opssightSkyfireName
		case "skyfire-image":
			if globalOpsSightSpec.Skyfire == nil {
				globalOpsSightSpec.Skyfire = &opssightv1.Skyfire{}
			}
			globalOpsSightSpec.Skyfire.Image = opssightSkyfireImage
		case "skyfire-port":
			if globalOpsSightSpec.Skyfire == nil {
				globalOpsSightSpec.Skyfire = &opssightv1.Skyfire{}
			}
			globalOpsSightSpec.Skyfire.Port = opssightSkyfirePort
		case "skyfire-prometheus-port":
			if globalOpsSightSpec.Skyfire == nil {
				globalOpsSightSpec.Skyfire = &opssightv1.Skyfire{}
			}
			globalOpsSightSpec.Skyfire.PrometheusPort = opssightSkyfirePrometheusPort
		case "skyfire-service-account":
			if globalOpsSightSpec.Skyfire == nil {
				globalOpsSightSpec.Skyfire = &opssightv1.Skyfire{}
			}
			globalOpsSightSpec.Skyfire.ServiceAccount = opssightSkyfireServiceAccount
		case "skyfire-hub-client-timeout-seconds":
			if globalOpsSightSpec.Skyfire == nil {
				globalOpsSightSpec.Skyfire = &opssightv1.Skyfire{}
			}
			globalOpsSightSpec.Skyfire.HubClientTimeoutSeconds = opssightSkyfireHubClientTimeoutSeconds
		case "skyfire-hub-dump-pause-seconds":
			if globalOpsSightSpec.Skyfire == nil {
				globalOpsSightSpec.Skyfire = &opssightv1.Skyfire{}
			}
			globalOpsSightSpec.Skyfire.HubDumpPauseSeconds = opssightSkyfireHubDumpPauseSeconds
		case "skyfire-kube-dump-interval-seconds":
			if globalOpsSightSpec.Skyfire == nil {
				globalOpsSightSpec.Skyfire = &opssightv1.Skyfire{}
			}
			globalOpsSightSpec.Skyfire.KubeDumpIntervalSeconds = opssightSkyfireKubeDumpIntervalSeconds
		case "skyfire-perceptor-dump-interval-seconds":
			if globalOpsSightSpec.Skyfire == nil {
				globalOpsSightSpec.Skyfire = &opssightv1.Skyfire{}
			}
			globalOpsSightSpec.Skyfire.PerceptorDumpIntervalSeconds = opssightSkyfirePerceptorDumpIntervalSeconds
		case "blackduck-hosts":
			if globalOpsSightSpec.Blackduck == nil {
				globalOpsSightSpec.Blackduck = &opssightv1.Blackduck{}
			}
			globalOpsSightSpec.Blackduck.Hosts = opssightBlackduckHosts
		case "blackduck-user":
			if globalOpsSightSpec.Blackduck == nil {
				globalOpsSightSpec.Blackduck = &opssightv1.Blackduck{}
			}
			globalOpsSightSpec.Blackduck.User = opssightBlackduckUser
		case "blackduck-port":
			if globalOpsSightSpec.Blackduck == nil {
				globalOpsSightSpec.Blackduck = &opssightv1.Blackduck{}
			}
			globalOpsSightSpec.Blackduck.Port = opssightBlackduckPort
		case "blackduck-concurrent-scan-limit":
			if globalOpsSightSpec.Blackduck == nil {
				globalOpsSightSpec.Blackduck = &opssightv1.Blackduck{}
			}
			globalOpsSightSpec.Blackduck.ConcurrentScanLimit = opssightBlackduckConcurrentScanLimit
		case "blackduck-total-scan-limit":
			if globalOpsSightSpec.Blackduck == nil {
				globalOpsSightSpec.Blackduck = &opssightv1.Blackduck{}
			}
			globalOpsSightSpec.Blackduck.TotalScanLimit = opssightBlackduckTotalScanLimit
		case "blackduck-password-environment-variable":
			if globalOpsSightSpec.Blackduck == nil {
				globalOpsSightSpec.Blackduck = &opssightv1.Blackduck{}
			}
			globalOpsSightSpec.Blackduck.PasswordEnvVar = opssightBlackduckPasswordEnvVar
		case "blackduck-initial-count":
			if globalOpsSightSpec.Blackduck == nil {
				globalOpsSightSpec.Blackduck = &opssightv1.Blackduck{}
			}
			globalOpsSightSpec.Blackduck.InitialCount = opssightBlackduckInitialCount
		case "blackduck-max-count":
			if globalOpsSightSpec.Blackduck == nil {
				globalOpsSightSpec.Blackduck = &opssightv1.Blackduck{}
			}
			globalOpsSightSpec.Blackduck.MaxCount = opssightBlackduckMaxCount
		case "blackduck-delete-blackduck-threshold-percentage":
			if globalOpsSightSpec.Blackduck == nil {
				globalOpsSightSpec.Blackduck = &opssightv1.Blackduck{}
			}
			globalOpsSightSpec.Blackduck.DeleteHubThresholdPercentage = opssightBlackduckDeleteHubThresholdPercentage
		case "enable-metrics":
			globalOpsSightSpec.EnableMetrics = opssightEnableMetrics
		case "default-cpu":
			globalOpsSightSpec.DefaultCPU = opssightDefaultCPU
		case "default-mem":
			globalOpsSightSpec.DefaultMem = opssightDefaultMem
		case "log-level":
			globalOpsSightSpec.LogLevel = opssightLogLevel
		case "config-map-name":
			globalOpsSightSpec.ConfigMapName = opssightConfigMapName
		case "secret-name":
			globalOpsSightSpec.SecretName = opssightSecretName
		default:
			log.Debugf("Flag %s: Not Found\n", f.Name)
		}
	}
	log.Debugf("Flag %s: UNCHANGED\n", f.Name)

}

func setAlertFlags(f *pflag.Flag) {
	if f.Changed {
		log.Debugf("Flag %s: CHANGED\n", f.Name)
		switch f.Name {
		case "alert-registry":
			globalAlertSpec.Registry = alertRegistry
		case "image-path":
			globalAlertSpec.ImagePath = alertImagePath
		case "alert-image-name":
			globalAlertSpec.AlertImageName = alertAlertImageName
		case "alert-image-version":
			globalAlertSpec.AlertImageVersion = alertAlertImageVersion
		case "cfssl-image-name":
			globalAlertSpec.CfsslImageName = alertCfsslImageName
		case "cfssl-image-version":
			globalAlertSpec.CfsslImageVersion = alertCfsslImageVersion
		case "blackduck-host":
			globalAlertSpec.BlackduckHost = alertBlackduckHost
		case "blackduck-user":
			globalAlertSpec.BlackduckUser = alertBlackduckUser
		case "blackduck-port":
			globalAlertSpec.BlackduckPort = &alertBlackduckPort
		case "port":
			globalAlertSpec.Port = &alertPort
		case "stand-alone":
			globalAlertSpec.StandAlone = &alertStandAlone
		case "alert-memory":
			globalAlertSpec.AlertMemory = alertAlertMemory
		case "cfssl-memory":
			globalAlertSpec.CfsslMemory = alertCfsslMemory
		default:
			log.Debugf("Flag %s: Not Found\n", f.Name)
		}
	}
	log.Debugf("Flag %s: UNCHANGED\n", f.Name)
}
