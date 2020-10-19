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

package machines

import (
	"fmt"

	cmv1 "github.com/openshift-online/ocm-sdk-go/clustersmgmt/v1"
)

func GetMachineTypes(client *cmv1.Client) (machineTypes []*cmv1.MachineType, err error) {
	collection := client.MachineTypes()
	page := 1
	size := 100
	for {
		var response *cmv1.MachineTypesListResponse
		response, err = collection.List().
			Search("cloud_provider.id = 'aws'").
			Order("size desc").
			Page(page).
			Size(size).
			Send()
		if err != nil {
			return nil, fmt.Errorf(response.Error().Reason())
		}
		machineTypes = append(machineTypes, response.Items().Slice()...)
		if response.Size() < size {
			break
		}
		page++
	}
	return
}
