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
	"fmt"

	alertv1 "github.com/blackducksoftware/synopsys-operator/pkg/api/alert/v1"
	blackduckv1 "github.com/blackducksoftware/synopsys-operator/pkg/api/blackduck/v1"
	opssightv1 "github.com/blackducksoftware/synopsys-operator/pkg/api/opssight/v1"
	crddefaults "github.com/blackducksoftware/synopsys-operator/pkg/util"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// Create Flags
var createBlackduckSpecType = "persistentStorage"
var createOpsSightSpecType = "disabledBlackduck"
var createAlertSpecType = "spec1"

// Default Specs
var defaultBlackduckSpec = &blackduckv1.BlackduckSpec{}
var defaultOpsSightSpec = &opssightv1.OpsSightSpec{}
var defaultAlertSpec = &alertv1.AlertSpec{}

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a Synopsys Resource in your cluster",
	Args: func(cmd *cobra.Command, args []string) error {
		numArgs := 1
		if len(args) < numArgs {
			return fmt.Errorf("Not enough arguments")
		}
		return nil
	},
	PreRunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 1 && args[0] == "--help" {
			return fmt.Errorf("Help Called")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		log.Debugf("Creating a Non-Synopsys Resource\n")
		kubeCmdArgs := append([]string{"create"}, args...)
		out, err := RunKubeCmd(kubeCmdArgs...)
		if err != nil {
			fmt.Printf("Error Creating the Resource with KubeCmd: %s\n", err)
		} else {
			fmt.Printf("%+v", out)
		}
	},
}

// createCmd represents the create command for Blackduck
var createBlackduckCmd = &cobra.Command{
	Use:   "blackduck",
	Short: "Create an instance of a Blackduck",
	Args: func(cmd *cobra.Command, args []string) error {
		// Check Number of Arguments
		if len(args) > 1 {
			return fmt.Errorf("This command only accepts up to 1 argument - NAME")
		}
		// Check the Spec Type
		switch createBlackduckSpecType {
		case "empty":
			globalBlackduckSpec = &blackduckv1.BlackduckSpec{}
		case "persistentStorage":
			globalBlackduckSpec = crddefaults.GetHubDefaultPersistentStorage()
		case "default":
			globalBlackduckSpec = crddefaults.GetHubDefaultValue()
		default:
			return fmt.Errorf("Blackduck Spec Type %s does not match: empty, persistentStorage, default", createBlackduckSpecType)
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		log.Debugf("Creating a Blackduck\n")
		// Read Commandline Parameters
		blackduckName := "blackduck"
		if len(args) == 1 {
			blackduckName = args[0]
		}

		// Create namespace for the Blackduck
		DeployCRDNamespace(restconfig, blackduckName)

		// Read Flags Into Default Blackduck Spec
		flagset := cmd.Flags()
		flagset.VisitAll(setBlackduckFlags)

		// Set Namespace in Spec
		globalBlackduckSpec.Namespace = blackduckName

		// Create and Deploy Blackduck CRD
		blackduck := &blackduckv1.Blackduck{
			ObjectMeta: metav1.ObjectMeta{
				Name: blackduckName,
			},
			Spec: *globalBlackduckSpec,
		}
		log.Debugf("%+v\n", blackduck)
		_, err := blackduckClient.SynopsysV1().Blackducks(blackduckName).Create(blackduck)
		if err != nil {
			fmt.Printf("Error creating the Blackduck : %s\n", err)
			return
		}
	},
}

// createCmd represents the create command for OpsSight
var createOpsSightCmd = &cobra.Command{
	Use:   "opssight",
	Short: "Create an instance of OpsSight",
	Args: func(cmd *cobra.Command, args []string) error {
		// Check Number of Arguments
		if len(args) > 1 {
			return fmt.Errorf("This command only accepts up to 1 argument - NAME")
		}
		// Check the Spec Type
		switch createOpsSightSpecType {
		case "empty":
			globalOpsSightSpec = &opssightv1.OpsSightSpec{}
		case "disabledBlackduck":
			globalOpsSightSpec = crddefaults.GetOpsSightDefaultValueWithDisabledHub()
		case "default":
			globalOpsSightSpec = crddefaults.GetOpsSightDefaultValue()
		default:
			return fmt.Errorf("OpsSight Spec Type %s does not match: empty, disabledBlackduck, default", createOpsSightSpecType)
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		log.Debugf("Creating an OpsSight\n")
		// Read Commandline Parameters
		opsSightName := "opssight"
		if len(args) == 1 {
			opsSightName = args[0]
		}

		// Create namespace for the OpsSight
		DeployCRDNamespace(restconfig, opsSightName)

		// Read Flags Into Default OpsSight Spec
		flagset := cmd.Flags()
		flagset.VisitAll(setOpsSightFlags)

		// Set Namespace in Spec
		globalOpsSightSpec.Namespace = opsSightName

		// Create and Deploy OpsSight CRD
		opssight := &opssightv1.OpsSight{
			ObjectMeta: metav1.ObjectMeta{
				Name: opsSightName,
			},
			Spec: *globalOpsSightSpec,
		}
		log.Debugf("%+v\n", opssight)
		_, err := opssightClient.SynopsysV1().OpsSights(opsSightName).Create(opssight)
		if err != nil {
			fmt.Printf("Error creating the OpsSight : %s\n", err)
			return
		}
	},
}

// createCmd represents the create command for Alert
var createAlertCmd = &cobra.Command{
	Use:   "alert",
	Short: "Create an instance of Alert",
	Args: func(cmd *cobra.Command, args []string) error {
		// Check Number of Arguments
		if len(args) > 1 {
			return fmt.Errorf("This command only accepts up to 1 argument - NAME")
		}
		// Check the Spec Type
		switch createAlertSpecType {
		case "empty":
			globalAlertSpec = &alertv1.AlertSpec{}
		case "spec1":
			globalAlertSpec = crddefaults.GetAlertDefaultValue()
		case "spec2":
			globalAlertSpec = crddefaults.GetAlertDefaultValue2()
		default:
			return fmt.Errorf("Alert Spec Type %s does not match: empty, spec1, spec2", createAlertSpecType)
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		log.Debugf("Creating an Alert\n")
		// Read Commandline Parameters
		alertName := "alert"
		if len(args) == 1 {
			alertName = args[0]
		}

		// Create namespace for the Alert
		DeployCRDNamespace(restconfig, alertName)

		// Read Flags Into Default Alert Spec
		flagset := cmd.Flags()
		flagset.VisitAll(setAlertFlags)

		// Set Namespace in Spec
		globalAlertSpec.Namespace = alertName

		// Create and Deploy Alert CRD
		alert := &alertv1.Alert{
			ObjectMeta: metav1.ObjectMeta{
				Name: alertName,
			},
			Spec: *globalAlertSpec,
		}
		log.Debugf("%+v\n", alert)
		_, err := alertClient.SynopsysV1().Alerts(alertName).Create(alert)
		if err != nil {
			fmt.Printf("Error creating the Alert : %s\n", err)
			return
		}
	},
}

func init() {
	createCmd.DisableFlagParsing = true
	rootCmd.AddCommand(createCmd)

	// Add Blackduck Command Flags
	createBlackduckCmd.Flags().StringVar(&createBlackduckSpecType, "spec", createBlackduckSpecType, "TODO")
	// Add Blackduck Spec Flags
	addBlackduckSpecFlags(createBlackduckCmd)
	createCmd.AddCommand(createBlackduckCmd)

	// Add OpsSight Command Flags
	createOpsSightCmd.Flags().StringVar(&createOpsSightSpecType, "spec", createOpsSightSpecType, "TODO")
	// Add OpsSight Spec Flags
	addOpsSightSpecFlags(createOpsSightCmd)
	createCmd.AddCommand(createOpsSightCmd)

	// Add Alert Command Flags
	createAlertCmd.Flags().StringVar(&createAlertSpecType, "spec", createAlertSpecType, "TODO")
	// Add Alert Spec Flags
	addAlertSpecFlags(createAlertCmd)
	createCmd.AddCommand(createAlertCmd)
}
