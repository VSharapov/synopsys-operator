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

package cmd

import (
	"flag"
	"fmt"
	"path/filepath"

	horizonapi "github.com/blackducksoftware/horizon/pkg/api"
	horizoncomponents "github.com/blackducksoftware/horizon/pkg/components"
	"github.com/blackducksoftware/horizon/pkg/deployer"
	blackduckv1 "github.com/blackducksoftware/synopsys-operator/pkg/api/blackduck/v1"
	blackduckclientset "github.com/blackducksoftware/synopsys-operator/pkg/blackduck/client/clientset/versioned"
	"github.com/blackducksoftware/synopsys-operator/pkg/util"
	"github.com/spf13/cobra"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/tools/clientcmd"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "create a synopsys resource (ex: blackduck, opssight)",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("create called")
	},
}

var blackduckCmd = &cobra.Command{
	Use:   "blackduck",
	Short: "create an instance of a Blackduck",
	Run: func(cmd *cobra.Command, args []string) {
		// Create kubernetes Clientset
		var kubeconfig *string
		if home := homeDir(); home != "" {
			kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
		} else {
			kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
		}
		flag.Parse()
		restconfig, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
		if err != nil {
			panic(err.Error())
		}

		// Create namespace for the Blackduck
		namespaceDeployer, err := deployer.NewDeployer(restconfig)
		ns := horizoncomponents.NewNamespace(horizonapi.NamespaceConfig{
			// APIVersion:  "string",
			// ClusterName: "string",
			Name:      namespace,
			Namespace: namespace,
		})
		namespaceDeployer.AddNamespace(ns)
		err = namespaceDeployer.Run()
		if err != nil {
			fmt.Printf("Error deploying namespace for the Blackduck with Horizon : %s\n", err)
			return
		}

		// Create Spec for a Blackduck CRD
		blackduck := &blackduckv1.Blackduck{}

		blackduckClient, err := blackduckclientset.NewForConfig(restconfig)
		// Get Namespace for the blackduck
		blackduckNamespace := blackduck.Spec.Namespace
		// Get hub_v2.Blackduck
		hubv2 := blackduckv1.Blackduck{
			ObjectMeta: metav1.ObjectMeta{
				Name: blackduck.Spec.Namespace,
			},
			Spec: blackduck.Spec,
		}

		// CreateHub(hubClientset *hubclientset.Clientset, namespace string, createHub *hub_v2.Blackduck)
		util.CreateHub(blackduckClient, blackduckNamespace, &hubv2)
	},
}

var opssightCmd = &cobra.Command{
	Use:   "opssight",
	Short: "create an instance of OpsSight",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Creating OpsSight\n")
	},
}

func init() {
	rootCmd.AddCommand(createCmd)
	blackduckCmd.Flags().IntVarP(&create_blackduck_size, "size", "s", create_blackduck_size, "blackduck size in GB")
	createCmd.AddCommand(blackduckCmd)
	createCmd.AddCommand(opssightCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
