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

	"github.com/blackducksoftware/synopsys-operator/pkg/util"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// dbMigrateCmd puts a Synopsys resource into the mode for database migration
var dbMigrateCmd = &cobra.Command{
	Use:     "db-migrate",
	Aliases: []string{"database-migrate"},
	Short:   "Put a Synopsys resource into the database-migrate state",
	RunE: func(cmd *cobra.Command, args []string) error {
		return fmt.Errorf("must specify a resource")
	},
}

// dbMigrateBlackDuckCmd puts a Black Duck instance into the db-migrate state
var dbMigrateBlackDuckCmd = &cobra.Command{
	Use:     "blackduck NAME",
	Example: "synopsysctl db-migrate blackduck <name>\nsynopsysctl db-migrate blackduck <name> -n <namespace>",
	Short:   "Put a Black Duck instance into database migration mode",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return fmt.Errorf("this command takes 1 argument")
		}
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		blackDuckName, blackDuckNamespace, _, err := getInstanceInfo(cmd, args[0], util.BlackDuckCRDName, "", namespace)
		if err != nil {
			log.Error(err)
			return nil
		}
		log.Infof("putting Black Duck '%s' in namespace '%s' into database migration mode...", blackDuckName, blackDuckNamespace)

		// Get the Black Duck
		currBlackDuck, err := util.GetHub(blackDuckClient, blackDuckNamespace, blackDuckName)
		if err != nil {
			log.Errorf("unable to get Black Duck '%s' in namespace '%s' due to %+v", blackDuckName, blackDuckNamespace, err)
			return nil
		}

		// Make changes to Spec
		currBlackDuck.Spec.DesiredState = "DbMigrate"
		// Update Black Duck
		_, err = util.UpdateBlackduck(blackDuckClient, currBlackDuck.Spec.Namespace, currBlackDuck)
		if err != nil {
			log.Errorf("error putting Black Duck '%s' in namespace '%s' into database migration mode due to %+v", blackDuckName, blackDuckNamespace, err)
			return nil
		}

		log.Infof("successfully put Black Duck '%s' in namespace '%s' into database migration mode", blackDuckName, blackDuckNamespace)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(dbMigrateCmd)

	dbMigrateBlackDuckCmd.Flags().StringVarP(&namespace, "namespace", "n", namespace, "Namespace of the instance(s)")
	dbMigrateCmd.AddCommand(dbMigrateBlackDuckCmd)
}
