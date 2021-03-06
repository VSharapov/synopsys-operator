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

package synopsysctl

import (
	"fmt"
	"strings"

	"github.com/blackducksoftware/synopsys-operator/pkg/util"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	rbacv1 "k8s.io/api/rbac/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// destroyCmd removes Synopsys Operator from the cluster
var destroyCmd = &cobra.Command{
	Use:     "destroy [NAMESPACE...]",
	Example: "synopsysctl destroy\nsynopsysctl destroy <namespace>\nsynopsysctl destroy <namespace1> <namespace2>",
	Short:   "Remove one or many Synopsys Operator instances and their associated CRDs",
	Args: func(cmd *cobra.Command, args []string) error {
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		// Read Commandline Parameters
		if len(args) > 0 {
			for _, operatorNamespace := range args {
				destroy(operatorNamespace)
			}
		} else {
			operatorNamespace := DefaultOperatorNamespace
			var err error
			isClusterScoped := util.GetClusterScope(apiExtensionClient)
			if isClusterScoped {
				namespace, err = util.GetOperatorNamespace(kubeClient, metav1.NamespaceAll)
				if err != nil {
					log.Error(err)
					return nil
				}
				if metav1.NamespaceAll != namespace {
					operatorNamespace = namespace
				}
			}
			destroy(operatorNamespace)
		}
		return nil
	},
}

func destroy(namespace string) {
	log.Infof("destroying Synopsys Operator in namespace '%s'...", namespace)
	crds := []string{util.AlertCRDName, util.BlackDuckCRDName, util.OpsSightCRDName}

	// delete namespace
	err := util.DeleteResourceNamespace(restconfig, kubeClient, strings.Join(crds, ","), namespace, true)
	if err != nil {
		log.Warn(err)
	}

	// delete crds
	deleteCrd(crds)

	// delete cluster role bindings
	clusterRoleBindings, roleBindings, err := util.GetOperatorRoleBindings(kubeClient, namespace)
	if err != nil {
		log.Errorf("error getting role binding or cluster role binding due to %+v", err)
	}

	for _, clusterRoleBinding := range clusterRoleBindings {
		crb, err := util.GetClusterRoleBinding(kubeClient, clusterRoleBinding)
		if err != nil {
			log.Errorf("unable to get cluster role binding '%s' due to %+v", clusterRoleBinding, err)
		}
		// check whether any subject present for other namespace before deleting them
		newSubjects := []rbacv1.Subject{}
		for _, subject := range crb.Subjects {
			isExist := util.IsSubjectExistForOtherNamespace(subject, namespace)
			if isExist {
				newSubjects = append(newSubjects, subject)
			}
		}
		if len(newSubjects) > 0 {
			crb.Subjects = newSubjects
			// update the cluster role binding to remove the old cluster role binding subject
			_, err = util.UpdateClusterRoleBinding(kubeClient, crb)
			if err != nil {
				log.Errorf("unable to update cluster role binding '%s' due to %+v", clusterRoleBinding, err)
			}
		} else {
			log.Infof("deleting cluster role binding '%s'", clusterRoleBinding)
			err := util.DeleteClusterRoleBinding(kubeClient, clusterRoleBinding)
			if err != nil {
				log.Errorf("unable to delete cluster role binding '%s' due to %+v", clusterRoleBinding, err)
			}
		}
	}

	// delete role bindings
	for _, roleBinding := range roleBindings {
		log.Infof("deleting role binding '%s'", roleBinding)
		err = util.DeleteRoleBinding(kubeClient, namespace, roleBinding)

		if err != nil {
			log.Errorf("unable to delete role binding '%s' due to %+v", roleBinding, err)
		}
	}

	// delete cluster roles
	clusterRoles, roles, err := util.GetOperatorRoles(kubeClient, namespace)
	if err != nil {
		log.Errorf("error getting role or cluster role due to %+v", err)
	}

	crbs, err := util.ListClusterRoleBindings(kubeClient, "app in (opssight)")

	for _, clusterRole := range clusterRoles {
		isExist := false
		// check whether the cluster role is referenced in any cluster role binding
		for _, crb := range crbs.Items {
			if util.IsClusterRoleRefExistForOtherNamespace(crb.RoleRef, clusterRole, namespace, crb.Subjects) {
				isExist = true
			}
		}
		if !isExist {
			log.Infof("deleting cluster role '%s'", clusterRole)
			err := util.DeleteClusterRole(kubeClient, clusterRole)
			if err != nil {
				log.Errorf("unable to delete cluster role '%s' due to %+v", clusterRole, err)
			}
		}
	}

	// delete roles
	for _, role := range roles {
		log.Infof("deleting role '%s'", role)
		err := util.DeleteRole(kubeClient, namespace, role)
		if err != nil {
			log.Errorf("unable to delete role '%s' due to %+v", role, err)
		}
	}

	log.Infof("finished destroying Synopsys Operator in namespace '%s'", namespace)
}

func isDeleteCrd(crd string) error {
	switch crd {
	case util.AlertCRDName:
		alerts, err := util.ListAlerts(alertClient, metav1.NamespaceAll)
		if err != nil {
			return fmt.Errorf("unable to list Alert instances due to %+v", err)
		}

		for _, alert := range alerts.Items {
			if alert.Namespace != namespace {
				return fmt.Errorf("already Alert instances exist in other namespaces. Please delete them before deleting the custom resources")
			}
		}
	case util.BlackDuckCRDName:
		blackducks, err := util.ListHubs(blackDuckClient, metav1.NamespaceAll)
		if err != nil {
			return fmt.Errorf("unable to list Black Duck instances due to %+v", err)
		}

		for _, blackduck := range blackducks.Items {
			if blackduck.Namespace != namespace {
				return fmt.Errorf("already Black Duck instances exist in other namespaces. Please delete them before deleting the custom resources")
			}
		}
	case util.OpsSightCRDName:
		opssights, err := util.ListOpsSights(opsSightClient, metav1.NamespaceAll)
		if err != nil {
			return fmt.Errorf("unable to list OpsSight instances due to %+v", err)
		}

		for _, opssight := range opssights.Items {
			if opssight.Namespace != namespace {
				return fmt.Errorf("already OpsSight instances exist in other namespaces. Please delete them before deleting the custom resources")
			}
		}
	}
	return nil
}

func deleteCrd(crds []string) {
	for _, crd := range crds {
		err := isDeleteCrd(crd)
		if err != nil {
			log.Warn(err)
		} else {
			log.Infof("deleting Custom Resource Definition '%s'", crd)
			err := util.DeleteCustomResourceDefinition(apiExtensionClient, crd)
			if err != nil {
				log.Errorf("unable to delete Custom Resource Definition '%s' due to %+v", crd, err)
			}
		}
	}
}

func init() {
	rootCmd.AddCommand(destroyCmd)
}
