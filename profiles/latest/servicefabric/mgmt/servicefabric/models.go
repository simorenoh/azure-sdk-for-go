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

package servicefabric

import original "github.com/Azure/azure-sdk-for-go/services/servicefabric/mgmt/2016-09-01/servicefabric"

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type BaseClient = original.BaseClient

func New(subscriptionID string) BaseClient {
	return original.New(subscriptionID)
}
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}

type ClustersClient = original.ClustersClient

func NewClustersClient(subscriptionID string) ClustersClient {
	return original.NewClustersClient(subscriptionID)
}
func NewClustersClientWithBaseURI(baseURI string, subscriptionID string) ClustersClient {
	return original.NewClustersClientWithBaseURI(baseURI, subscriptionID)
}

type ClusterVersionsClient = original.ClusterVersionsClient

func NewClusterVersionsClient(subscriptionID string) ClusterVersionsClient {
	return original.NewClusterVersionsClient(subscriptionID)
}
func NewClusterVersionsClientWithBaseURI(baseURI string, subscriptionID string) ClusterVersionsClient {
	return original.NewClusterVersionsClientWithBaseURI(baseURI, subscriptionID)
}

type ClusterState = original.ClusterState

const (
	AutoScale                 ClusterState = original.AutoScale
	BaselineUpgrade           ClusterState = original.BaselineUpgrade
	Deploying                 ClusterState = original.Deploying
	EnforcingClusterVersion   ClusterState = original.EnforcingClusterVersion
	Ready                     ClusterState = original.Ready
	UpdatingInfrastructure    ClusterState = original.UpdatingInfrastructure
	UpdatingUserCertificate   ClusterState = original.UpdatingUserCertificate
	UpdatingUserConfiguration ClusterState = original.UpdatingUserConfiguration
	UpgradeServiceUnreachable ClusterState = original.UpgradeServiceUnreachable
	WaitingForNodes           ClusterState = original.WaitingForNodes
)

func PossibleClusterStateValues() [10]ClusterState {
	return original.PossibleClusterStateValues()
}

type DurabilityLevel = original.DurabilityLevel

const (
	Bronze DurabilityLevel = original.Bronze
	Gold   DurabilityLevel = original.Gold
	Silver DurabilityLevel = original.Silver
)

func PossibleDurabilityLevelValues() [3]DurabilityLevel {
	return original.PossibleDurabilityLevelValues()
}

type Environment = original.Environment

const (
	Linux   Environment = original.Linux
	Windows Environment = original.Windows
)

func PossibleEnvironmentValues() [2]Environment {
	return original.PossibleEnvironmentValues()
}

type ProvisioningState = original.ProvisioningState

const (
	Canceled  ProvisioningState = original.Canceled
	Failed    ProvisioningState = original.Failed
	Succeeded ProvisioningState = original.Succeeded
	Updating  ProvisioningState = original.Updating
)

func PossibleProvisioningStateValues() [4]ProvisioningState {
	return original.PossibleProvisioningStateValues()
}

type ReliabilityLevel = original.ReliabilityLevel

const (
	ReliabilityLevelBronze ReliabilityLevel = original.ReliabilityLevelBronze
	ReliabilityLevelGold   ReliabilityLevel = original.ReliabilityLevelGold
	ReliabilityLevelSilver ReliabilityLevel = original.ReliabilityLevelSilver
)

func PossibleReliabilityLevelValues() [3]ReliabilityLevel {
	return original.PossibleReliabilityLevelValues()
}

type ReliabilityLevel1 = original.ReliabilityLevel1

const (
	ReliabilityLevel1Bronze   ReliabilityLevel1 = original.ReliabilityLevel1Bronze
	ReliabilityLevel1Gold     ReliabilityLevel1 = original.ReliabilityLevel1Gold
	ReliabilityLevel1Platinum ReliabilityLevel1 = original.ReliabilityLevel1Platinum
	ReliabilityLevel1Silver   ReliabilityLevel1 = original.ReliabilityLevel1Silver
)

func PossibleReliabilityLevel1Values() [4]ReliabilityLevel1 {
	return original.PossibleReliabilityLevel1Values()
}

type UpgradeMode = original.UpgradeMode

const (
	Automatic UpgradeMode = original.Automatic
	Manual    UpgradeMode = original.Manual
)

func PossibleUpgradeModeValues() [2]UpgradeMode {
	return original.PossibleUpgradeModeValues()
}

type UpgradeMode1 = original.UpgradeMode1

const (
	UpgradeMode1Automatic UpgradeMode1 = original.UpgradeMode1Automatic
	UpgradeMode1Manual    UpgradeMode1 = original.UpgradeMode1Manual
)

func PossibleUpgradeMode1Values() [2]UpgradeMode1 {
	return original.PossibleUpgradeMode1Values()
}

type X509StoreName = original.X509StoreName

const (
	AddressBook          X509StoreName = original.AddressBook
	AuthRoot             X509StoreName = original.AuthRoot
	CertificateAuthority X509StoreName = original.CertificateAuthority
	Disallowed           X509StoreName = original.Disallowed
	My                   X509StoreName = original.My
	Root                 X509StoreName = original.Root
	TrustedPeople        X509StoreName = original.TrustedPeople
	TrustedPublisher     X509StoreName = original.TrustedPublisher
)

func PossibleX509StoreNameValues() [8]X509StoreName {
	return original.PossibleX509StoreNameValues()
}

type AvailableOperationDisplay = original.AvailableOperationDisplay
type AzureActiveDirectory = original.AzureActiveDirectory
type CertificateDescription = original.CertificateDescription
type ClientCertificateCommonName = original.ClientCertificateCommonName
type ClientCertificateThumbprint = original.ClientCertificateThumbprint
type Cluster = original.Cluster
type ClusterCodeVersionsListResult = original.ClusterCodeVersionsListResult
type ClusterCodeVersionsListResultIterator = original.ClusterCodeVersionsListResultIterator
type ClusterCodeVersionsListResultPage = original.ClusterCodeVersionsListResultPage
type ClusterCodeVersionsResult = original.ClusterCodeVersionsResult
type ClusterHealthPolicy = original.ClusterHealthPolicy
type ClusterListResult = original.ClusterListResult
type ClusterListResultIterator = original.ClusterListResultIterator
type ClusterListResultPage = original.ClusterListResultPage
type ClusterProperties = original.ClusterProperties
type ClusterPropertiesUpdateParameters = original.ClusterPropertiesUpdateParameters
type ClustersCreateFuture = original.ClustersCreateFuture
type ClustersUpdateFuture = original.ClustersUpdateFuture
type ClusterUpdateParameters = original.ClusterUpdateParameters
type ClusterUpgradeDeltaHealthPolicy = original.ClusterUpgradeDeltaHealthPolicy
type ClusterUpgradePolicy = original.ClusterUpgradePolicy
type ClusterVersionDetails = original.ClusterVersionDetails
type DiagnosticsStorageAccountConfig = original.DiagnosticsStorageAccountConfig
type EndpointRangeDescription = original.EndpointRangeDescription
type ErrorModel = original.ErrorModel
type ErrorModelError = original.ErrorModelError
type NodeTypeDescription = original.NodeTypeDescription
type OperationListResult = original.OperationListResult
type OperationListResultIterator = original.OperationListResultIterator
type OperationListResultPage = original.OperationListResultPage
type OperationResult = original.OperationResult
type Resource = original.Resource
type SettingsParameterDescription = original.SettingsParameterDescription
type SettingsSectionDescription = original.SettingsSectionDescription
type OperationsClient = original.OperationsClient

func NewOperationsClient(subscriptionID string) OperationsClient {
	return original.NewOperationsClient(subscriptionID)
}
func NewOperationsClientWithBaseURI(baseURI string, subscriptionID string) OperationsClient {
	return original.NewOperationsClientWithBaseURI(baseURI, subscriptionID)
}
func UserAgent() string {
	return original.UserAgent() + " profiles/latest"
}
func Version() string {
	return original.Version()
}
