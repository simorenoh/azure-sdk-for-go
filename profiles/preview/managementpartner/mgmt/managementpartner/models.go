// +build go1.9

// Copyright 2018 Microsoft Corporation
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

// This code was auto-generated by:
// github.com/Azure/azure-sdk-for-go/tools/profileBuilder

package managementpartner

import original "github.com/Azure/azure-sdk-for-go/services/managementpartner/mgmt/2018-02-01/managementpartner"

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type BaseClient = original.BaseClient

func New() BaseClient {
	return original.New()
}
func NewWithBaseURI(baseURI string) BaseClient {
	return original.NewWithBaseURI(baseURI)
}

type Code = original.Code

const (
	BadRequest Code = original.BadRequest
	Conflict   Code = original.Conflict
	NotFound   Code = original.NotFound
)

func PossibleCodeValues() [3]Code {
	return original.PossibleCodeValues()
}

type State = original.State

const (
	Active  State = original.Active
	Deleted State = original.Deleted
)

func PossibleStateValues() [2]State {
	return original.PossibleStateValues()
}

type Error = original.Error
type ExtendedErrorInfo = original.ExtendedErrorInfo
type OperationDisplay = original.OperationDisplay
type OperationList = original.OperationList
type OperationListIterator = original.OperationListIterator
type OperationListPage = original.OperationListPage
type OperationResponse = original.OperationResponse
type PartnerProperties = original.PartnerProperties
type PartnerResponse = original.PartnerResponse
type OperationClient = original.OperationClient

func NewOperationClient() OperationClient {
	return original.NewOperationClient()
}
func NewOperationClientWithBaseURI(baseURI string) OperationClient {
	return original.NewOperationClientWithBaseURI(baseURI)
}

type PartnerClient = original.PartnerClient

func NewPartnerClient() PartnerClient {
	return original.NewPartnerClient()
}
func NewPartnerClientWithBaseURI(baseURI string) PartnerClient {
	return original.NewPartnerClientWithBaseURI(baseURI)
}
func UserAgent() string {
	return original.UserAgent() + " profiles/preview"
}
func Version() string {
	return original.Version()
}
