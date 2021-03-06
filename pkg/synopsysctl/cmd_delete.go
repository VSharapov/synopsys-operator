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

	util "github.com/blackducksoftware/synopsys-operator/pkg/util"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// deleteCmd deletes a resource from the cluster
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Remove Synopsys resources from your cluster",
	RunE: func(cmd *cobra.Command, args []string) error {
		return fmt.Errorf("must specify a sub-command")
	},
}

// deleteAlertCmd deletes Alert instances from the cluster
var deleteAlertCmd = &cobra.Command{
	Use:     "alert NAME...",
	Example: "synopsysctl delete alert <name>\nsynopsysctl delete alert <name1> <name2> <name3>\nsynopsysctl delete alert <name> -n <namespace>\nsynopsysctl delete alert <name1> <name2> <name3> -n <namespace>",
	Short:   "Delete one or many Alert instances",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return fmt.Errorf("this command takes 1 or more arguments")
		}
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		for _, alertName := range args {
			alertName, alertNamespace, _, err := getInstanceInfo(cmd, alertName, util.AlertCRDName, util.AlertName, namespace)
			if err != nil {
				log.Error(err)
				return nil
			}
			log.Infof("deleting Alert '%s' in namespace '%s'...", alertName, alertNamespace)
			err = alertClient.SynopsysV1().Alerts(alertNamespace).Delete(alertName, &metav1.DeleteOptions{})
			if err != nil {
				log.Errorf("error deleting Alert '%s' in namespace '%s' due to %+v", alertName, alertNamespace, err)
				return nil
			}
			log.Infof("successfully deleted Alert '%s' in namespace '%s'", alertName, alertNamespace)
		}
		return nil
	},
}

// deleteBlackDuckCmd deletes Black Duck instances from the cluster
var deleteBlackDuckCmd = &cobra.Command{
	Use:     "blackduck NAME...",
	Example: "synopsysctl delete blackduck <name>\nsynopsysctl delete blackduck <name1> <name2> <name3>\nsynopsysctl delete blackduck <name> -n <namespace>\nsynopsysctl delete blackduck <name1> <name2> <name3> -n <namespace>",
	Short:   "Delete one or many Black Duck instances",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return fmt.Errorf("this command takes 1 or more arguments")
		}
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		for _, blackDuckName := range args {
			blackDuckName, blackDuckNamespace, _, err := getInstanceInfo(cmd, blackDuckName, util.BlackDuckCRDName, util.BlackDuckName, namespace)
			if err != nil {
				log.Error(err)
				return nil
			}
			log.Infof("deleting Black Duck '%s' in namespace '%s'...", blackDuckName, blackDuckNamespace)
			err = blackDuckClient.SynopsysV1().Blackducks(blackDuckNamespace).Delete(blackDuckName, &metav1.DeleteOptions{})
			if err != nil {
				log.Errorf("error deleting Black Duck '%s' in namespace '%s' due to '%s'", blackDuckName, blackDuckNamespace, err)
				return nil
			}
			log.Infof("successfully deleted Black Duck '%s' in namespace '%s'", blackDuckName, blackDuckNamespace)
		}
		return nil
	},
}

// deleteOpsSightCmd deletes OpsSight instances from the cluster
var deleteOpsSightCmd = &cobra.Command{
	Use:     "opssight NAME...",
	Example: "synopsysctl delete opssight <name>\nsynopsysctl delete opssight <name1> <name2> <name3>\nsynopsysctl delete opssight <name> -n <namespace>\nsynopsysctl delete opssight <name1> <name2> <name3> -n <namespace>",
	Short:   "Delete one or many OpsSight instances",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return fmt.Errorf("this command takes 1 or more arguments")
		}
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		for _, opsSightName := range args {
			opsSightName, opsSightNamespace, _, err := getInstanceInfo(cmd, opsSightName, util.OpsSightCRDName, util.OpsSightName, namespace)
			if err != nil {
				log.Error(err)
				return nil
			}
			log.Infof("deleting OpsSight '%s' in namespace '%s'...", opsSightName, opsSightNamespace)
			err = opsSightClient.SynopsysV1().OpsSights(opsSightNamespace).Delete(opsSightName, &metav1.DeleteOptions{})
			if err != nil {
				log.Errorf("error deleting OpsSight '%s' in namespace '%s' due to '%s'", opsSightName, opsSightNamespace, err)
				return nil
			}
			log.Infof("successfully deleted OpsSight '%s' in namespace '%s'", opsSightName, opsSightNamespace)
		}
		return nil
	},
}

func init() {
	//(PassCmd) deleteCmd.DisableFlagParsing = true // lets deleteCmd pass flags to kube/oc
	rootCmd.AddCommand(deleteCmd)

	// Add Delete Alert Command
	deleteAlertCmd.Flags().StringVarP(&namespace, "namespace", "n", namespace, "Namespace of the instance(s)")
	deleteCmd.AddCommand(deleteAlertCmd)

	// Add Delete Black Duck Command
	deleteBlackDuckCmd.Flags().StringVarP(&namespace, "namespace", "n", namespace, "Namespace of the instance(s)")
	deleteCmd.AddCommand(deleteBlackDuckCmd)

	// Add Delete OpsSight Command
	deleteOpsSightCmd.Flags().StringVarP(&namespace, "namespace", "n", namespace, "Namespace of the instance(s)")
	deleteCmd.AddCommand(deleteOpsSightCmd)
}
