/*
Copyright (c) 2020 Red Hat, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

  http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package cluster

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/spf13/cobra"

	"github.com/openshift/moactl/pkg/aws"
	clusterprovider "github.com/openshift/moactl/pkg/cluster"
	"github.com/openshift/moactl/pkg/logging"
	"github.com/openshift/moactl/pkg/ocm"
	rprtr "github.com/openshift/moactl/pkg/reporter"
)

var args struct {
	count int
}

var Cmd = &cobra.Command{
	Use:     "clusters",
	Aliases: []string{"cluster"},
	Short:   "List clusters",
	Long:    "List clusters.",
	Example: `  # List all clusters
  rosa list clusters`,
	Run: run,
}

func init() {
	flags := Cmd.Flags()
	flags.SortFlags = false

	// Basic options
	flags.IntVar(
		&args.count,
		"count",
		100,
		"Number of clusters to display.",
	)
}

func run(_ *cobra.Command, argv []string) {
	reporter := rprtr.CreateReporterOrExit()
	logger := logging.CreateLoggerOrExit(reporter)

	// Check command line arguments:
	if len(argv) != 0 {
		reporter.Errorf("Expected exactly zero command line parameters")
		os.Exit(1)
	}

	// Create the AWS client:
	awsClient, err := aws.NewClient().
		Logger(logger).
		Build()
	if err != nil {
		reporter.Errorf("Failed to create AWS client: %v", err)
		os.Exit(1)
	}

	awsCreator, err := awsClient.GetCreator()
	if err != nil {
		reporter.Errorf("Failed to get AWS creator: %v", err)
		os.Exit(1)
	}

	// Create the client for the OCM API:
	ocmConnection, err := ocm.NewConnection().
		Logger(logger).
		Build()
	if err != nil {
		reporter.Errorf("Failed to create OCM connection: %v", err)
		os.Exit(1)
	}
	defer func() {
		err = ocmConnection.Close()
		if err != nil {
			reporter.Errorf("Failed to close OCM connection: %v", err)
		}
	}()

	// Retrieve the list of clusters:
	clustersCollection := ocmConnection.ClustersMgmt().V1().Clusters()
	clusters, err := clusterprovider.GetClusters(clustersCollection, awsCreator.ARN, args.count)
	if err != nil {
		reporter.Errorf("Failed to get clusters: %v", err)
		os.Exit(1)
	}

	if len(clusters) == 0 {
		reporter.Infof("No clusters available")
		os.Exit(0)
	}

	// Create the writer that will be used to print the tabulated results:
	writer := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	fmt.Fprintf(writer, "ID\tNAME\tSTATE\n")
	for _, cluster := range clusters {
		fmt.Fprintf(
			writer,
			"%s\t%s\t%s\n",
			cluster.ID(),
			cluster.Name(),
			cluster.State(),
		)
	}
	writer.Flush()
}
