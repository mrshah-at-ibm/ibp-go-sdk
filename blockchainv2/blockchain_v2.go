/**
 * (C) Copyright IBM Corp. 2020.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

// Package blockchainv2 : Operations and models for the BlockchainV2 service
package blockchainv2

import (
	"encoding/json"
	"fmt"
	"github.com/IBM/go-sdk-core/v4/core"
	common "github.com/IBM-Blockchain/ibp-go-sdk/blockchainv2"
	"reflect"
)

// BlockchainV2 : This doc lists APIs that you can use to interact with your IBM Blockchain Platform console (IBP
// console)
//
// Version: 2.5.0
// See: http://swagger.io
type BlockchainV2 struct {
	Service *core.BaseService
}

// DefaultServiceName is the default key used to find external configuration information.
const DefaultServiceName = "blockchain"

// BlockchainV2Options : Service options
type BlockchainV2Options struct {
	ServiceName   string
	URL           string
	Authenticator core.Authenticator
}

// NewBlockchainV2UsingExternalConfig : constructs an instance of BlockchainV2 with passed in options and external configuration.
func NewBlockchainV2UsingExternalConfig(options *BlockchainV2Options) (blockchain *BlockchainV2, err error) {
	if options.ServiceName == "" {
		options.ServiceName = DefaultServiceName
	}

	if options.Authenticator == nil {
		options.Authenticator, err = core.GetAuthenticatorFromEnvironment(options.ServiceName)
		if err != nil {
			return
		}
	}

	blockchain, err = NewBlockchainV2(options)
	if err != nil {
		return
	}

	err = blockchain.Service.ConfigureService(options.ServiceName)
	if err != nil {
		return
	}

	if options.URL != "" {
		err = blockchain.Service.SetServiceURL(options.URL)
	}
	return
}

// NewBlockchainV2 : constructs an instance of BlockchainV2 with passed in options.
func NewBlockchainV2(options *BlockchainV2Options) (service *BlockchainV2, err error) {
	serviceOptions := &core.ServiceOptions{
		Authenticator: options.Authenticator,
	}

	baseService, err := core.NewBaseService(serviceOptions)
	if err != nil {
		return
	}

	if options.URL != "" {
		err = baseService.SetServiceURL(options.URL)
		if err != nil {
			return
		}
	}

	service = &BlockchainV2{
		Service: baseService,
	}

	return
}

// SetServiceURL sets the service URL
func (blockchain *BlockchainV2) SetServiceURL(url string) error {
	return blockchain.Service.SetServiceURL(url)
}

// GetComponent : Get component data
// Get the IBP console's data on a component (peer, CA, orderer, or MSP). The component might be imported or created.
func (blockchain *BlockchainV2) GetComponent(getComponentOptions *GetComponentOptions) (result *GenericComponentResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getComponentOptions, "getComponentOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getComponentOptions, "getComponentOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"ak/api/v2/components"}
	pathParameters := []string{*getComponentOptions.ID}

	builder := core.NewRequestBuilder(core.GET)
	_, err = builder.ConstructHTTPURL(blockchain.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range getComponentOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("blockchain", "V2", "GetComponent")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	if getComponentOptions.DeploymentAttrs != nil {
		builder.AddQuery("deployment_attrs", fmt.Sprint(*getComponentOptions.DeploymentAttrs))
	}
	if getComponentOptions.ParsedCerts != nil {
		builder.AddQuery("parsed_certs", fmt.Sprint(*getComponentOptions.ParsedCerts))
	}
	if getComponentOptions.Cache != nil {
		builder.AddQuery("cache", fmt.Sprint(*getComponentOptions.Cache))
	}
	if getComponentOptions.CaAttrs != nil {
		builder.AddQuery("ca_attrs", fmt.Sprint(*getComponentOptions.CaAttrs))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = blockchain.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalGenericComponentResponse)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// RemoveComponent : Remove imported component
// Remove a single component from the IBP console.
// - Using this api on an **imported** component removes it from the IBP console.
// - Using this api on a **created** component removes it from the IBP console **but** it will **not** delete the
// component from the Kubernetes cluster where it resides. Thus it orphans the Kubernetes deployment (if it exists).
// Instead use the [Delete component](#delete-component) API to delete the Kubernetes deployment and the IBP console
// data at once.
func (blockchain *BlockchainV2) RemoveComponent(removeComponentOptions *RemoveComponentOptions) (result *DeleteComponentResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(removeComponentOptions, "removeComponentOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(removeComponentOptions, "removeComponentOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"ak/api/v2/components"}
	pathParameters := []string{*removeComponentOptions.ID}

	builder := core.NewRequestBuilder(core.DELETE)
	_, err = builder.ConstructHTTPURL(blockchain.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range removeComponentOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("blockchain", "V2", "RemoveComponent")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = blockchain.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalDeleteComponentResponse)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// DeleteComponent : Delete component
// Removes a single component from the IBP console **and** it deletes the Kubernetes deployment.
// - Using this api on an **imported** component will *error out* since its Kubernetes deployment is unknown and cannot
// be removed. Instead use the [Remove imported component](#remove-imported-component) API to remove imported
// components.
// - Using this api on a **created** component removes it from the IBP console **and** it will delete the component from
// the Kubernetes cluster where it resides. The Kubernetes delete must succeed before the component will be removed from
// the IBP console.
func (blockchain *BlockchainV2) DeleteComponent(deleteComponentOptions *DeleteComponentOptions) (result *DeleteComponentResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteComponentOptions, "deleteComponentOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(deleteComponentOptions, "deleteComponentOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"ak/api/v2/kubernetes/components"}
	pathParameters := []string{*deleteComponentOptions.ID}

	builder := core.NewRequestBuilder(core.DELETE)
	_, err = builder.ConstructHTTPURL(blockchain.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range deleteComponentOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("blockchain", "V2", "DeleteComponent")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = blockchain.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalDeleteComponentResponse)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// CreateCa : Create a CA
// Create a Hyperledger Fabric Certificate Authority (CA) in your Kubernetes cluster.
func (blockchain *BlockchainV2) CreateCa(createCaOptions *CreateCaOptions) (result *CaResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(createCaOptions, "createCaOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(createCaOptions, "createCaOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"ak/api/v2/kubernetes/components/fabric-ca"}
	pathParameters := []string{}

	builder := core.NewRequestBuilder(core.POST)
	_, err = builder.ConstructHTTPURL(blockchain.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range createCaOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("blockchain", "V2", "CreateCa")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if createCaOptions.DisplayName != nil {
		body["display_name"] = createCaOptions.DisplayName
	}
	if createCaOptions.ConfigOverride != nil {
		body["config_override"] = createCaOptions.ConfigOverride
	}
	if createCaOptions.Resources != nil {
		body["resources"] = createCaOptions.Resources
	}
	if createCaOptions.Storage != nil {
		body["storage"] = createCaOptions.Storage
	}
	if createCaOptions.Zone != nil {
		body["zone"] = createCaOptions.Zone
	}
	if createCaOptions.Replicas != nil {
		body["replicas"] = createCaOptions.Replicas
	}
	if createCaOptions.Tags != nil {
		body["tags"] = createCaOptions.Tags
	}
	if createCaOptions.Hsm != nil {
		body["hsm"] = createCaOptions.Hsm
	}
	if createCaOptions.Region != nil {
		body["region"] = createCaOptions.Region
	}
	if createCaOptions.Version != nil {
		body["version"] = createCaOptions.Version
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = blockchain.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalCaResponse)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// ImportCa : Import a CA
// Import an existing Certificate Authority (CA) to your IBP console. It is recommended to only import components that
// were created by this or another IBP console.
func (blockchain *BlockchainV2) ImportCa(importCaOptions *ImportCaOptions) (result *CaResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(importCaOptions, "importCaOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(importCaOptions, "importCaOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"ak/api/v2/components/fabric-ca"}
	pathParameters := []string{}

	builder := core.NewRequestBuilder(core.POST)
	_, err = builder.ConstructHTTPURL(blockchain.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range importCaOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("blockchain", "V2", "ImportCa")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if importCaOptions.DisplayName != nil {
		body["display_name"] = importCaOptions.DisplayName
	}
	if importCaOptions.ApiURL != nil {
		body["api_url"] = importCaOptions.ApiURL
	}
	if importCaOptions.CaName != nil {
		body["ca_name"] = importCaOptions.CaName
	}
	if importCaOptions.TlscaName != nil {
		body["tlsca_name"] = importCaOptions.TlscaName
	}
	if importCaOptions.TlsCert != nil {
		body["tls_cert"] = importCaOptions.TlsCert
	}
	if importCaOptions.Location != nil {
		body["location"] = importCaOptions.Location
	}
	if importCaOptions.OperationsURL != nil {
		body["operations_url"] = importCaOptions.OperationsURL
	}
	if importCaOptions.Tags != nil {
		body["tags"] = importCaOptions.Tags
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = blockchain.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalCaResponse)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// UpdateCa : Update a CA
// Update Kubernetes deployment attributes of a Hyperledger Fabric Certificate Authority (CA) in your cluster.
func (blockchain *BlockchainV2) UpdateCa(updateCaOptions *UpdateCaOptions) (result *CaResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(updateCaOptions, "updateCaOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(updateCaOptions, "updateCaOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"ak/api/v2/kubernetes/components/fabric-ca"}
	pathParameters := []string{*updateCaOptions.ID}

	builder := core.NewRequestBuilder(core.PUT)
	_, err = builder.ConstructHTTPURL(blockchain.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range updateCaOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("blockchain", "V2", "UpdateCa")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if updateCaOptions.Resources != nil {
		body["resources"] = updateCaOptions.Resources
	}
	if updateCaOptions.Zone != nil {
		body["zone"] = updateCaOptions.Zone
	}
	if updateCaOptions.ConfigOverride != nil {
		body["config_override"] = updateCaOptions.ConfigOverride
	}
	if updateCaOptions.Replicas != nil {
		body["replicas"] = updateCaOptions.Replicas
	}
	if updateCaOptions.Version != nil {
		body["version"] = updateCaOptions.Version
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = blockchain.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalCaResponse)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// EditCa : Edit data about a CA
// Modify local metadata fields of a Certificate Authority (CA). For example, the "display_name" field. This API will
// **not** change any Kubernetes deployment attributes for the CA.
func (blockchain *BlockchainV2) EditCa(editCaOptions *EditCaOptions) (result *CaResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(editCaOptions, "editCaOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(editCaOptions, "editCaOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"ak/api/v2/components/fabric-ca"}
	pathParameters := []string{*editCaOptions.ID}

	builder := core.NewRequestBuilder(core.PUT)
	_, err = builder.ConstructHTTPURL(blockchain.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range editCaOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("blockchain", "V2", "EditCa")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if editCaOptions.DisplayName != nil {
		body["display_name"] = editCaOptions.DisplayName
	}
	if editCaOptions.ApiURL != nil {
		body["api_url"] = editCaOptions.ApiURL
	}
	if editCaOptions.OperationsURL != nil {
		body["operations_url"] = editCaOptions.OperationsURL
	}
	if editCaOptions.CaName != nil {
		body["ca_name"] = editCaOptions.CaName
	}
	if editCaOptions.Location != nil {
		body["location"] = editCaOptions.Location
	}
	if editCaOptions.Tags != nil {
		body["tags"] = editCaOptions.Tags
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = blockchain.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalCaResponse)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// CreatePeer : Create a peer
// Create a Hyperledger Fabric peer in your Kubernetes cluster.
func (blockchain *BlockchainV2) CreatePeer(createPeerOptions *CreatePeerOptions) (result *PeerResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(createPeerOptions, "createPeerOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(createPeerOptions, "createPeerOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"ak/api/v2/kubernetes/components/fabric-peer"}
	pathParameters := []string{}

	builder := core.NewRequestBuilder(core.POST)
	_, err = builder.ConstructHTTPURL(blockchain.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range createPeerOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("blockchain", "V2", "CreatePeer")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if createPeerOptions.MspID != nil {
		body["msp_id"] = createPeerOptions.MspID
	}
	if createPeerOptions.DisplayName != nil {
		body["display_name"] = createPeerOptions.DisplayName
	}
	if createPeerOptions.Config != nil {
		body["config"] = createPeerOptions.Config
	}
	if createPeerOptions.ConfigOverride != nil {
		body["config_override"] = createPeerOptions.ConfigOverride
	}
	if createPeerOptions.Resources != nil {
		body["resources"] = createPeerOptions.Resources
	}
	if createPeerOptions.Storage != nil {
		body["storage"] = createPeerOptions.Storage
	}
	if createPeerOptions.Zone != nil {
		body["zone"] = createPeerOptions.Zone
	}
	if createPeerOptions.StateDb != nil {
		body["state_db"] = createPeerOptions.StateDb
	}
	if createPeerOptions.Tags != nil {
		body["tags"] = createPeerOptions.Tags
	}
	if createPeerOptions.Hsm != nil {
		body["hsm"] = createPeerOptions.Hsm
	}
	if createPeerOptions.Region != nil {
		body["region"] = createPeerOptions.Region
	}
	if createPeerOptions.Version != nil {
		body["version"] = createPeerOptions.Version
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = blockchain.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalPeerResponse)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// ImportPeer : Import a peer
// Import an existing peer into your IBP console. It is recommended to only import components that were created by this
// or another IBP console.
func (blockchain *BlockchainV2) ImportPeer(importPeerOptions *ImportPeerOptions) (result *PeerResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(importPeerOptions, "importPeerOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(importPeerOptions, "importPeerOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"ak/api/v2/components/fabric-peer"}
	pathParameters := []string{}

	builder := core.NewRequestBuilder(core.POST)
	_, err = builder.ConstructHTTPURL(blockchain.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range importPeerOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("blockchain", "V2", "ImportPeer")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if importPeerOptions.DisplayName != nil {
		body["display_name"] = importPeerOptions.DisplayName
	}
	if importPeerOptions.MspID != nil {
		body["msp_id"] = importPeerOptions.MspID
	}
	if importPeerOptions.GrpcwpURL != nil {
		body["grpcwp_url"] = importPeerOptions.GrpcwpURL
	}
	if importPeerOptions.TlsCaRootCert != nil {
		body["tls_ca_root_cert"] = importPeerOptions.TlsCaRootCert
	}
	if importPeerOptions.Location != nil {
		body["location"] = importPeerOptions.Location
	}
	if importPeerOptions.ApiURL != nil {
		body["api_url"] = importPeerOptions.ApiURL
	}
	if importPeerOptions.OperationsURL != nil {
		body["operations_url"] = importPeerOptions.OperationsURL
	}
	if importPeerOptions.Tags != nil {
		body["tags"] = importPeerOptions.Tags
	}
	if importPeerOptions.TlsCert != nil {
		body["tls_cert"] = importPeerOptions.TlsCert
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = blockchain.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalPeerResponse)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// EditPeer : Edit data about a peer
// Modify local metadata fields of a peer. For example, the "display_name" field. This API will **not** change any
// Kubernetes deployment attributes for the peer.
func (blockchain *BlockchainV2) EditPeer(editPeerOptions *EditPeerOptions) (result *PeerResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(editPeerOptions, "editPeerOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(editPeerOptions, "editPeerOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"ak/api/v2/components/fabric-peer"}
	pathParameters := []string{*editPeerOptions.ID}

	builder := core.NewRequestBuilder(core.PUT)
	_, err = builder.ConstructHTTPURL(blockchain.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range editPeerOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("blockchain", "V2", "EditPeer")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if editPeerOptions.DisplayName != nil {
		body["display_name"] = editPeerOptions.DisplayName
	}
	if editPeerOptions.ApiURL != nil {
		body["api_url"] = editPeerOptions.ApiURL
	}
	if editPeerOptions.OperationsURL != nil {
		body["operations_url"] = editPeerOptions.OperationsURL
	}
	if editPeerOptions.GrpcwpURL != nil {
		body["grpcwp_url"] = editPeerOptions.GrpcwpURL
	}
	if editPeerOptions.MspID != nil {
		body["msp_id"] = editPeerOptions.MspID
	}
	if editPeerOptions.Location != nil {
		body["location"] = editPeerOptions.Location
	}
	if editPeerOptions.Tags != nil {
		body["tags"] = editPeerOptions.Tags
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = blockchain.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalPeerResponse)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// UpdatePeer : Update a peer
// Update Kubernetes deployment attributes of a Hyperledger Fabric Peer node.
func (blockchain *BlockchainV2) UpdatePeer(updatePeerOptions *UpdatePeerOptions) (result *PeerResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(updatePeerOptions, "updatePeerOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(updatePeerOptions, "updatePeerOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"ak/api/v2/kubernetes/components/fabric-peer"}
	pathParameters := []string{*updatePeerOptions.ID}

	builder := core.NewRequestBuilder(core.PUT)
	_, err = builder.ConstructHTTPURL(blockchain.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range updatePeerOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("blockchain", "V2", "UpdatePeer")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if updatePeerOptions.ConfigOverride != nil {
		body["config_override"] = updatePeerOptions.ConfigOverride
	}
	if updatePeerOptions.Resources != nil {
		body["resources"] = updatePeerOptions.Resources
	}
	if updatePeerOptions.Zone != nil {
		body["zone"] = updatePeerOptions.Zone
	}
	if updatePeerOptions.Version != nil {
		body["version"] = updatePeerOptions.Version
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = blockchain.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalPeerResponse)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// CreateOrderer : Create an ordering service
// Create a Hyperledger Ordering Service (OS) in your Kubernetes cluster. Currently, only raft ordering nodes are
// supported.
func (blockchain *BlockchainV2) CreateOrderer(createOrdererOptions *CreateOrdererOptions) (result *OrdererResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(createOrdererOptions, "createOrdererOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(createOrdererOptions, "createOrdererOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"ak/api/v2/kubernetes/components/fabric-orderer"}
	pathParameters := []string{}

	builder := core.NewRequestBuilder(core.POST)
	_, err = builder.ConstructHTTPURL(blockchain.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range createOrdererOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("blockchain", "V2", "CreateOrderer")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if createOrdererOptions.OrdererType != nil {
		body["orderer_type"] = createOrdererOptions.OrdererType
	}
	if createOrdererOptions.MspID != nil {
		body["msp_id"] = createOrdererOptions.MspID
	}
	if createOrdererOptions.DisplayName != nil {
		body["display_name"] = createOrdererOptions.DisplayName
	}
	if createOrdererOptions.Config != nil {
		body["config"] = createOrdererOptions.Config
	}
	if createOrdererOptions.ClusterName != nil {
		body["cluster_name"] = createOrdererOptions.ClusterName
	}
	if createOrdererOptions.ClusterID != nil {
		body["cluster_id"] = createOrdererOptions.ClusterID
	}
	if createOrdererOptions.ExternalAppend != nil {
		body["external_append"] = createOrdererOptions.ExternalAppend
	}
	if createOrdererOptions.ConfigOverride != nil {
		body["config_override"] = createOrdererOptions.ConfigOverride
	}
	if createOrdererOptions.Resources != nil {
		body["resources"] = createOrdererOptions.Resources
	}
	if createOrdererOptions.Storage != nil {
		body["storage"] = createOrdererOptions.Storage
	}
	if createOrdererOptions.SystemChannelID != nil {
		body["system_channel_id"] = createOrdererOptions.SystemChannelID
	}
	if createOrdererOptions.Zone != nil {
		body["zone"] = createOrdererOptions.Zone
	}
	if createOrdererOptions.Tags != nil {
		body["tags"] = createOrdererOptions.Tags
	}
	if createOrdererOptions.Region != nil {
		body["region"] = createOrdererOptions.Region
	}
	if createOrdererOptions.Hsm != nil {
		body["hsm"] = createOrdererOptions.Hsm
	}
	if createOrdererOptions.Version != nil {
		body["version"] = createOrdererOptions.Version
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = blockchain.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalOrdererResponse)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// ImportOrderer : Import an ordering service
// Import an existing Ordering Service (OS) to your IBP console. It is recommended to only import components that were
// created by this or another IBP console.
func (blockchain *BlockchainV2) ImportOrderer(importOrdererOptions *ImportOrdererOptions) (result *OrdererResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(importOrdererOptions, "importOrdererOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(importOrdererOptions, "importOrdererOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"ak/api/v2/components/fabric-orderer"}
	pathParameters := []string{}

	builder := core.NewRequestBuilder(core.POST)
	_, err = builder.ConstructHTTPURL(blockchain.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range importOrdererOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("blockchain", "V2", "ImportOrderer")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if importOrdererOptions.ClusterName != nil {
		body["cluster_name"] = importOrdererOptions.ClusterName
	}
	if importOrdererOptions.DisplayName != nil {
		body["display_name"] = importOrdererOptions.DisplayName
	}
	if importOrdererOptions.MspID != nil {
		body["msp_id"] = importOrdererOptions.MspID
	}
	if importOrdererOptions.GrpcwpURL != nil {
		body["grpcwp_url"] = importOrdererOptions.GrpcwpURL
	}
	if importOrdererOptions.TlsCaRootCert != nil {
		body["tls_ca_root_cert"] = importOrdererOptions.TlsCaRootCert
	}
	if importOrdererOptions.Location != nil {
		body["location"] = importOrdererOptions.Location
	}
	if importOrdererOptions.ApiURL != nil {
		body["api_url"] = importOrdererOptions.ApiURL
	}
	if importOrdererOptions.OperationsURL != nil {
		body["operations_url"] = importOrdererOptions.OperationsURL
	}
	if importOrdererOptions.SystemChannelID != nil {
		body["system_channel_id"] = importOrdererOptions.SystemChannelID
	}
	if importOrdererOptions.Tags != nil {
		body["tags"] = importOrdererOptions.Tags
	}
	if importOrdererOptions.TlsCert != nil {
		body["tls_cert"] = importOrdererOptions.TlsCert
	}
	if importOrdererOptions.ServerTlsCert != nil {
		body["server_tls_cert"] = importOrdererOptions.ServerTlsCert
	}
	if importOrdererOptions.ClientTlsCert != nil {
		body["client_tls_cert"] = importOrdererOptions.ClientTlsCert
	}
	if importOrdererOptions.ClusterID != nil {
		body["cluster_id"] = importOrdererOptions.ClusterID
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = blockchain.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalOrdererResponse)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// EditOrderer : Edit data about an orderer
// Modify local metadata fields of a single node in an Ordering Service (OS). For example, the "display_name" field.
// This API will **not** change any Kubernetes deployment attributes for the node.
func (blockchain *BlockchainV2) EditOrderer(editOrdererOptions *EditOrdererOptions) (result *OrdererResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(editOrdererOptions, "editOrdererOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(editOrdererOptions, "editOrdererOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"ak/api/v2/components/fabric-orderer"}
	pathParameters := []string{*editOrdererOptions.ID}

	builder := core.NewRequestBuilder(core.PUT)
	_, err = builder.ConstructHTTPURL(blockchain.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range editOrdererOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("blockchain", "V2", "EditOrderer")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if editOrdererOptions.ClusterName != nil {
		body["cluster_name"] = editOrdererOptions.ClusterName
	}
	if editOrdererOptions.DisplayName != nil {
		body["display_name"] = editOrdererOptions.DisplayName
	}
	if editOrdererOptions.ApiURL != nil {
		body["api_url"] = editOrdererOptions.ApiURL
	}
	if editOrdererOptions.OperationsURL != nil {
		body["operations_url"] = editOrdererOptions.OperationsURL
	}
	if editOrdererOptions.GrpcwpURL != nil {
		body["grpcwp_url"] = editOrdererOptions.GrpcwpURL
	}
	if editOrdererOptions.MspID != nil {
		body["msp_id"] = editOrdererOptions.MspID
	}
	if editOrdererOptions.ConsenterProposalFin != nil {
		body["consenter_proposal_fin"] = editOrdererOptions.ConsenterProposalFin
	}
	if editOrdererOptions.Location != nil {
		body["location"] = editOrdererOptions.Location
	}
	if editOrdererOptions.SystemChannelID != nil {
		body["system_channel_id"] = editOrdererOptions.SystemChannelID
	}
	if editOrdererOptions.Tags != nil {
		body["tags"] = editOrdererOptions.Tags
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = blockchain.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalOrdererResponse)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// UpdateOrderer : Update an orderer node
// Update Kubernetes deployment attributes of a Hyperledger Fabric Ordering node.
func (blockchain *BlockchainV2) UpdateOrderer(updateOrdererOptions *UpdateOrdererOptions) (result *OrdererResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(updateOrdererOptions, "updateOrdererOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(updateOrdererOptions, "updateOrdererOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"ak/api/v2/kubernetes/components/fabric-orderer"}
	pathParameters := []string{*updateOrdererOptions.ID}

	builder := core.NewRequestBuilder(core.PUT)
	_, err = builder.ConstructHTTPURL(blockchain.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range updateOrdererOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("blockchain", "V2", "UpdateOrderer")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if updateOrdererOptions.ConfigOverride != nil {
		body["config_override"] = updateOrdererOptions.ConfigOverride
	}
	if updateOrdererOptions.Resources != nil {
		body["resources"] = updateOrdererOptions.Resources
	}
	if updateOrdererOptions.Zone != nil {
		body["zone"] = updateOrdererOptions.Zone
	}
	if updateOrdererOptions.Version != nil {
		body["version"] = updateOrdererOptions.Version
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = blockchain.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalOrdererResponse)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// SubmitBlock : Submit config block to orderer
// Send a config block (or genesis block) to a pre-created raft orderer node. Use this api to finish the raft-append
// flow and finalize a pre-created orderer. This is the final step to append a node to a raft cluster. The orderer will
// restart, load this block, and connect to the other orderers listed in said block.
//
// The full flow to append a raft node:
//   1. Pre-create the orderer with the [Create an ordering service](#create-an-ordering-service) API (setting
// `cluster_id` is how you turn the normal create-orderer api into a pre-create-orderer api).
//   2. Retrieve the pre-created node's tls cert with the [Get component data](#get-component-data) API (set the
// `deployment_attrs=included` parameter).
//   3. Get the latest config block for the system-channel by using the Fabric API (use a Fabric CLI or another Fabric
// tool).
//   4. Edit the config block for the system-channel and add the pre-created orderer's tls cert and api url as a
// consenter.
//   5. Create and marshal a Fabric
// [ConfigUpdate](https://github.com/hyperledger/fabric/blob/release-1.4/protos/common/configtx.proto#L78) proposal with
// [configtxlator](https://hyperledger-fabric.readthedocs.io/en/release-1.4/commands/configtxlator.html#configtxlator-compute-update)
// using the old and new block.
//   6. Sign the `ConfigUpdate` proposal and create a
// [ConfigSignature](https://github.com/hyperledger/fabric/blob/release-1.4/protos/common/configtx.proto#L111). Create a
// set of signatures that will satisfy the system channel's update policy.
//   7. Build a [SignedProposal](https://github.com/hyperledger/fabric/blob/release-1.4/protos/peer/proposal.proto#L105)
// out of the `ConfigUpdate` & `ConfigSignature`. Submit the `SignedProposal` to an existing ordering node (do not use
// the pre-created node).
//   8. After the `SignedProposal` transaction is committed to a block, pull the latest config block (for the
// system-channel) from an existing ordering node (use a Fabric CLI or another Fabric tool).
//   9. Submit the latest config block to your pre-created node with the 'Submit config block to orderer' API (which is
// this api!)
//   10. Use the [Edit data about an orderer](#edit-data-about-an-orderer) API to change the pre-created node's field
// `consenter_proposal_fin` to `true`. This changes the status icon on the IBP console.
func (blockchain *BlockchainV2) SubmitBlock(submitBlockOptions *SubmitBlockOptions) (result *GenericComponentResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(submitBlockOptions, "submitBlockOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(submitBlockOptions, "submitBlockOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"ak/api/v2/kubernetes/components", "config"}
	pathParameters := []string{*submitBlockOptions.ID}

	builder := core.NewRequestBuilder(core.PUT)
	_, err = builder.ConstructHTTPURL(blockchain.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range submitBlockOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("blockchain", "V2", "SubmitBlock")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "text/plain")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if submitBlockOptions.B64Block != nil {
		body["b64_block"] = submitBlockOptions.B64Block
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = blockchain.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalGenericComponentResponse)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// ImportMsp : Import an MSP
// Create or import a Membership Service Provider (MSP) definition into your IBP console. This definition represents an
// organization that controls a peer or OS (Ordering Service).
func (blockchain *BlockchainV2) ImportMsp(importMspOptions *ImportMspOptions) (result *MspResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(importMspOptions, "importMspOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(importMspOptions, "importMspOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"ak/api/v2/components/msp"}
	pathParameters := []string{}

	builder := core.NewRequestBuilder(core.POST)
	_, err = builder.ConstructHTTPURL(blockchain.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range importMspOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("blockchain", "V2", "ImportMsp")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if importMspOptions.MspID != nil {
		body["msp_id"] = importMspOptions.MspID
	}
	if importMspOptions.DisplayName != nil {
		body["display_name"] = importMspOptions.DisplayName
	}
	if importMspOptions.RootCerts != nil {
		body["root_certs"] = importMspOptions.RootCerts
	}
	if importMspOptions.IntermediateCerts != nil {
		body["intermediate_certs"] = importMspOptions.IntermediateCerts
	}
	if importMspOptions.Admins != nil {
		body["admins"] = importMspOptions.Admins
	}
	if importMspOptions.TlsRootCerts != nil {
		body["tls_root_certs"] = importMspOptions.TlsRootCerts
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = blockchain.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalMspResponse)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// EditMsp : Edit an MSP
// Modify local metadata fields of a Membership Service Provider (MSP) definition. For example, the "display_name"
// property.
func (blockchain *BlockchainV2) EditMsp(editMspOptions *EditMspOptions) (result *MspResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(editMspOptions, "editMspOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(editMspOptions, "editMspOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"ak/api/v2/components/msp"}
	pathParameters := []string{*editMspOptions.ID}

	builder := core.NewRequestBuilder(core.PUT)
	_, err = builder.ConstructHTTPURL(blockchain.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range editMspOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("blockchain", "V2", "EditMsp")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if editMspOptions.MspID != nil {
		body["msp_id"] = editMspOptions.MspID
	}
	if editMspOptions.DisplayName != nil {
		body["display_name"] = editMspOptions.DisplayName
	}
	if editMspOptions.RootCerts != nil {
		body["root_certs"] = editMspOptions.RootCerts
	}
	if editMspOptions.IntermediateCerts != nil {
		body["intermediate_certs"] = editMspOptions.IntermediateCerts
	}
	if editMspOptions.Admins != nil {
		body["admins"] = editMspOptions.Admins
	}
	if editMspOptions.TlsRootCerts != nil {
		body["tls_root_certs"] = editMspOptions.TlsRootCerts
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = blockchain.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalMspResponse)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// GetMspCertificate : Get MSP's public certificates
// External IBP consoles can use this API to get the public certificate for your given MSP id.
func (blockchain *BlockchainV2) GetMspCertificate(getMspCertificateOptions *GetMspCertificateOptions) (result *GetMSPCertificateResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getMspCertificateOptions, "getMspCertificateOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getMspCertificateOptions, "getMspCertificateOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"ak/api/v2/components/msps"}
	pathParameters := []string{*getMspCertificateOptions.MspID}

	builder := core.NewRequestBuilder(core.GET)
	_, err = builder.ConstructHTTPURL(blockchain.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range getMspCertificateOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("blockchain", "V2", "GetMspCertificate")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	if getMspCertificateOptions.Cache != nil {
		builder.AddQuery("cache", fmt.Sprint(*getMspCertificateOptions.Cache))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = blockchain.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalGetMSPCertificateResponse)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// EditAdminCerts : Edit admin certs on a component
// This api will append or remove admin certs to the components' file system. Certificates will be parsed. If invalid
// they will be skipped. Duplicate certificates will also be skipped. To view existing admin certificate use the [Get
// component data](#get-component-data) API with the query parameters: `?deployment_attrs=included&cache=skip`.
//
// **This API will not work on *imported* components.**.
func (blockchain *BlockchainV2) EditAdminCerts(editAdminCertsOptions *EditAdminCertsOptions) (result *EditAdminCertsResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(editAdminCertsOptions, "editAdminCertsOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(editAdminCertsOptions, "editAdminCertsOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"ak/api/v2/kubernetes/components", "certs"}
	pathParameters := []string{*editAdminCertsOptions.ID}

	builder := core.NewRequestBuilder(core.PUT)
	_, err = builder.ConstructHTTPURL(blockchain.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range editAdminCertsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("blockchain", "V2", "EditAdminCerts")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if editAdminCertsOptions.AppendAdminCerts != nil {
		body["append_admin_certs"] = editAdminCertsOptions.AppendAdminCerts
	}
	if editAdminCertsOptions.RemoveAdminCerts != nil {
		body["remove_admin_certs"] = editAdminCertsOptions.RemoveAdminCerts
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = blockchain.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalEditAdminCertsResponse)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// ListComponents : Get all components
// Get the IBP console's data on all components (peers, CAs, orderers, and MSPs). The component might be imported or
// created.
func (blockchain *BlockchainV2) ListComponents(listComponentsOptions *ListComponentsOptions) (result *GetMultiComponentsResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(listComponentsOptions, "listComponentsOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"ak/api/v2/components"}
	pathParameters := []string{}

	builder := core.NewRequestBuilder(core.GET)
	_, err = builder.ConstructHTTPURL(blockchain.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range listComponentsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("blockchain", "V2", "ListComponents")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	if listComponentsOptions.DeploymentAttrs != nil {
		builder.AddQuery("deployment_attrs", fmt.Sprint(*listComponentsOptions.DeploymentAttrs))
	}
	if listComponentsOptions.ParsedCerts != nil {
		builder.AddQuery("parsed_certs", fmt.Sprint(*listComponentsOptions.ParsedCerts))
	}
	if listComponentsOptions.Cache != nil {
		builder.AddQuery("cache", fmt.Sprint(*listComponentsOptions.Cache))
	}
	if listComponentsOptions.CaAttrs != nil {
		builder.AddQuery("ca_attrs", fmt.Sprint(*listComponentsOptions.CaAttrs))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = blockchain.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalGetMultiComponentsResponse)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// GetComponentsByType : Get components of a type
// Get the IBP console's data on components that are a specific type. The component might be imported or created.
func (blockchain *BlockchainV2) GetComponentsByType(getComponentsByTypeOptions *GetComponentsByTypeOptions) (result *GetMultiComponentsResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getComponentsByTypeOptions, "getComponentsByTypeOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getComponentsByTypeOptions, "getComponentsByTypeOptions")
	if err != nil {
		return
	}

	pathSegments := []string{}
	pathParameters := []string{*getComponentsByTypeOptions.ComponentType}

	builder := core.NewRequestBuilder(core.GET)
	_, err = builder.ConstructHTTPURL(blockchain.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range getComponentsByTypeOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("blockchain", "V2", "GetComponentsByType")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	if getComponentsByTypeOptions.DeploymentAttrs != nil {
		builder.AddQuery("deployment_attrs", fmt.Sprint(*getComponentsByTypeOptions.DeploymentAttrs))
	}
	if getComponentsByTypeOptions.ParsedCerts != nil {
		builder.AddQuery("parsed_certs", fmt.Sprint(*getComponentsByTypeOptions.ParsedCerts))
	}
	if getComponentsByTypeOptions.Cache != nil {
		builder.AddQuery("cache", fmt.Sprint(*getComponentsByTypeOptions.Cache))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = blockchain.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalGetMultiComponentsResponse)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// GetComponentByTag : Get components with tag
// Get the IBP console's data on components that have a specific tag. The component might be imported or created. Tags
// are not case-sensitive.
func (blockchain *BlockchainV2) GetComponentByTag(getComponentByTagOptions *GetComponentByTagOptions) (result *GetMultiComponentsResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getComponentByTagOptions, "getComponentByTagOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getComponentByTagOptions, "getComponentByTagOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"ak/api/v2/components/tags"}
	pathParameters := []string{*getComponentByTagOptions.Tag}

	builder := core.NewRequestBuilder(core.GET)
	_, err = builder.ConstructHTTPURL(blockchain.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range getComponentByTagOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("blockchain", "V2", "GetComponentByTag")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	if getComponentByTagOptions.DeploymentAttrs != nil {
		builder.AddQuery("deployment_attrs", fmt.Sprint(*getComponentByTagOptions.DeploymentAttrs))
	}
	if getComponentByTagOptions.ParsedCerts != nil {
		builder.AddQuery("parsed_certs", fmt.Sprint(*getComponentByTagOptions.ParsedCerts))
	}
	if getComponentByTagOptions.Cache != nil {
		builder.AddQuery("cache", fmt.Sprint(*getComponentByTagOptions.Cache))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = blockchain.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalGetMultiComponentsResponse)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// RemoveComponentsByTag : Remove components with tag
// Removes components with the matching tag from the IBP console. Tags are not case-sensitive.
// - Using this api on **imported** components removes them from the IBP console.
// - Using this api on **created** components removes them from the IBP console **but** it will **not** delete the
// components from the Kubernetes cluster where they reside. Thus it orphans the Kubernetes deployments (if it exists).
// Instead use the [Delete components with tag](#delete-component-s-with-tag) API to delete the Kubernetes deployment
// and the IBP console data at once.
func (blockchain *BlockchainV2) RemoveComponentsByTag(removeComponentsByTagOptions *RemoveComponentsByTagOptions) (result *RemoveMultiComponentsResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(removeComponentsByTagOptions, "removeComponentsByTagOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(removeComponentsByTagOptions, "removeComponentsByTagOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"ak/api/v2/components/tags"}
	pathParameters := []string{*removeComponentsByTagOptions.Tag}

	builder := core.NewRequestBuilder(core.DELETE)
	_, err = builder.ConstructHTTPURL(blockchain.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range removeComponentsByTagOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("blockchain", "V2", "RemoveComponentsByTag")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = blockchain.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalRemoveMultiComponentsResponse)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// DeleteComponentsByTag : Delete components with tag
// Removes components with the matching tag from the IBP console **and** it deletes the Kubernetes deployment. Tags are
// not case-sensitive.
// - Using this api on **imported** components will be skipped over since their Kubernetes deployment is unknown and
// cannot be removed. Instead use the [Remove components with tag](#remove-component-s-with-tag) API to remove imported
// components with a tag.
// - Using this api on **created** components removes them from the IBP console **and** it will delete the components
// from the Kubernetes cluster where they reside. The Kubernetes delete must succeed before the component will be
// removed from the IBP console.
func (blockchain *BlockchainV2) DeleteComponentsByTag(deleteComponentsByTagOptions *DeleteComponentsByTagOptions) (result *DeleteMultiComponentsResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteComponentsByTagOptions, "deleteComponentsByTagOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(deleteComponentsByTagOptions, "deleteComponentsByTagOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"ak/api/v2/kubernetes/components/tags"}
	pathParameters := []string{*deleteComponentsByTagOptions.Tag}

	builder := core.NewRequestBuilder(core.DELETE)
	_, err = builder.ConstructHTTPURL(blockchain.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range deleteComponentsByTagOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("blockchain", "V2", "DeleteComponentsByTag")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = blockchain.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalDeleteMultiComponentsResponse)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// DeleteAllComponents : Delete all components
// Removes all components from the IBP console **and** their Kubernetes deployments (if applicable). Works on imported
// and created components (peers, CAs, orderers, MSPs, and signature collection transactions). This api attempts to
// effectively reset the IBP console to its initial (empty) state (except for logs & notifications, those will remain).
func (blockchain *BlockchainV2) DeleteAllComponents(deleteAllComponentsOptions *DeleteAllComponentsOptions) (result *DeleteMultiComponentsResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(deleteAllComponentsOptions, "deleteAllComponentsOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"ak/api/v2/kubernetes/components/purge"}
	pathParameters := []string{}

	builder := core.NewRequestBuilder(core.DELETE)
	_, err = builder.ConstructHTTPURL(blockchain.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range deleteAllComponentsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("blockchain", "V2", "DeleteAllComponents")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = blockchain.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalDeleteMultiComponentsResponse)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// GetSettings : Get public IBP console settings
// Retrieve all public (non-sensitive) settings for the IBP console. Use this API for debugging purposes. It shows what
// behavior to expect and confirms whether the desired settings are active.
func (blockchain *BlockchainV2) GetSettings(getSettingsOptions *GetSettingsOptions) (result *GetPublicSettingsResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(getSettingsOptions, "getSettingsOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"ak/api/v2/settings"}
	pathParameters := []string{}

	builder := core.NewRequestBuilder(core.GET)
	_, err = builder.ConstructHTTPURL(blockchain.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range getSettingsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("blockchain", "V2", "GetSettings")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = blockchain.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalGetPublicSettingsResponse)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// EditSettings : Change IBP console settings
// Edit a few IBP console settings (such as the rate limit and timeout settings). **Some edits will trigger an automatic
// server restart.**.
func (blockchain *BlockchainV2) EditSettings(editSettingsOptions *EditSettingsOptions) (result *GetPublicSettingsResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(editSettingsOptions, "editSettingsOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(editSettingsOptions, "editSettingsOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"ak/api/v2/settings"}
	pathParameters := []string{}

	builder := core.NewRequestBuilder(core.PUT)
	_, err = builder.ConstructHTTPURL(blockchain.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range editSettingsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("blockchain", "V2", "EditSettings")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if editSettingsOptions.InactivityTimeouts != nil {
		body["inactivity_timeouts"] = editSettingsOptions.InactivityTimeouts
	}
	if editSettingsOptions.FileLogging != nil {
		body["file_logging"] = editSettingsOptions.FileLogging
	}
	if editSettingsOptions.MaxReqPerMin != nil {
		body["max_req_per_min"] = editSettingsOptions.MaxReqPerMin
	}
	if editSettingsOptions.MaxReqPerMinAk != nil {
		body["max_req_per_min_ak"] = editSettingsOptions.MaxReqPerMinAk
	}
	if editSettingsOptions.FabricGetBlockTimeoutMs != nil {
		body["fabric_get_block_timeout_ms"] = editSettingsOptions.FabricGetBlockTimeoutMs
	}
	if editSettingsOptions.FabricInstantiateTimeoutMs != nil {
		body["fabric_instantiate_timeout_ms"] = editSettingsOptions.FabricInstantiateTimeoutMs
	}
	if editSettingsOptions.FabricJoinChannelTimeoutMs != nil {
		body["fabric_join_channel_timeout_ms"] = editSettingsOptions.FabricJoinChannelTimeoutMs
	}
	if editSettingsOptions.FabricInstallCcTimeoutMs != nil {
		body["fabric_install_cc_timeout_ms"] = editSettingsOptions.FabricInstallCcTimeoutMs
	}
	if editSettingsOptions.FabricLcInstallCcTimeoutMs != nil {
		body["fabric_lc_install_cc_timeout_ms"] = editSettingsOptions.FabricLcInstallCcTimeoutMs
	}
	if editSettingsOptions.FabricLcGetCcTimeoutMs != nil {
		body["fabric_lc_get_cc_timeout_ms"] = editSettingsOptions.FabricLcGetCcTimeoutMs
	}
	if editSettingsOptions.FabricGeneralTimeoutMs != nil {
		body["fabric_general_timeout_ms"] = editSettingsOptions.FabricGeneralTimeoutMs
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = blockchain.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalGetPublicSettingsResponse)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// GetFabVersions : Get supported Fabric versions
// Get list of supported Fabric versions by each component type. These are the Fabric versions your IBP console can use
// when creating or upgrading components.
func (blockchain *BlockchainV2) GetFabVersions(getFabVersionsOptions *GetFabVersionsOptions) (result *GetFabricVersionsResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(getFabVersionsOptions, "getFabVersionsOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"ak/api/v2/kubernetes/fabric/versions"}
	pathParameters := []string{}

	builder := core.NewRequestBuilder(core.GET)
	_, err = builder.ConstructHTTPURL(blockchain.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range getFabVersionsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("blockchain", "V2", "GetFabVersions")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	if getFabVersionsOptions.Cache != nil {
		builder.AddQuery("cache", fmt.Sprint(*getFabVersionsOptions.Cache))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = blockchain.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalGetFabricVersionsResponse)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// GetHealth : Get IBP console health stats
// See statistics of the IBP console process such as memory usage, CPU usage, up time, cache, and operating system
// stats.
func (blockchain *BlockchainV2) GetHealth(getHealthOptions *GetHealthOptions) (result *GetAthenaHealthStatsResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(getHealthOptions, "getHealthOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"ak/api/v2/health"}
	pathParameters := []string{}

	builder := core.NewRequestBuilder(core.GET)
	_, err = builder.ConstructHTTPURL(blockchain.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range getHealthOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("blockchain", "V2", "GetHealth")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = blockchain.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalGetAthenaHealthStatsResponse)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// ListNotifications : Get all notifications
// Retrieve all notifications. This API supports pagination through the query parameters. Notifications are generated
// from actions such as creating a component, deleting a component, server restart, and so on.
func (blockchain *BlockchainV2) ListNotifications(listNotificationsOptions *ListNotificationsOptions) (result *GetNotificationsResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(listNotificationsOptions, "listNotificationsOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"ak/api/v2/notifications"}
	pathParameters := []string{}

	builder := core.NewRequestBuilder(core.GET)
	_, err = builder.ConstructHTTPURL(blockchain.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range listNotificationsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("blockchain", "V2", "ListNotifications")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	if listNotificationsOptions.Limit != nil {
		builder.AddQuery("limit", fmt.Sprint(*listNotificationsOptions.Limit))
	}
	if listNotificationsOptions.Skip != nil {
		builder.AddQuery("skip", fmt.Sprint(*listNotificationsOptions.Skip))
	}
	if listNotificationsOptions.ComponentID != nil {
		builder.AddQuery("component_id", fmt.Sprint(*listNotificationsOptions.ComponentID))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = blockchain.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalGetNotificationsResponse)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// DeleteSigTx : Delete a signature collection tx
// Delete a signature collection transaction. These transactions involve creating or editing Fabric channels & chaincode
// approvals. This request is not distributed to external IBP consoles, thus the signature collection transaction is
// only deleted locally.
func (blockchain *BlockchainV2) DeleteSigTx(deleteSigTxOptions *DeleteSigTxOptions) (result *DeleteSignatureCollectionResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteSigTxOptions, "deleteSigTxOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(deleteSigTxOptions, "deleteSigTxOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"ak/api/v2/signature_collections"}
	pathParameters := []string{*deleteSigTxOptions.ID}

	builder := core.NewRequestBuilder(core.DELETE)
	_, err = builder.ConstructHTTPURL(blockchain.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range deleteSigTxOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("blockchain", "V2", "DeleteSigTx")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = blockchain.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalDeleteSignatureCollectionResponse)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// ArchiveNotifications : Archive notifications
// Archive 1 or more notifications. Archived notifications will no longer appear in the default [Get all
// notifications](#get-all-notifications) API.
func (blockchain *BlockchainV2) ArchiveNotifications(archiveNotificationsOptions *ArchiveNotificationsOptions) (result *ArchiveResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(archiveNotificationsOptions, "archiveNotificationsOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(archiveNotificationsOptions, "archiveNotificationsOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"ak/api/v2/notifications/bulk"}
	pathParameters := []string{}

	builder := core.NewRequestBuilder(core.POST)
	_, err = builder.ConstructHTTPURL(blockchain.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range archiveNotificationsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("blockchain", "V2", "ArchiveNotifications")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if archiveNotificationsOptions.NotificationIds != nil {
		body["notification_ids"] = archiveNotificationsOptions.NotificationIds
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = blockchain.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalArchiveResponse)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// Restart : Restart the IBP console
// Restart IBP console processes. This causes a small outage (10 - 30 seconds) which is possibly disruptive to active
// user sessions.
func (blockchain *BlockchainV2) Restart(restartOptions *RestartOptions) (result *RestartAthenaResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(restartOptions, "restartOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"ak/api/v2/restart"}
	pathParameters := []string{}

	builder := core.NewRequestBuilder(core.POST)
	_, err = builder.ConstructHTTPURL(blockchain.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range restartOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("blockchain", "V2", "Restart")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = blockchain.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalRestartAthenaResponse)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// DeleteAllSessions : Delete all IBP console sessions
// Delete all client sessions in IBP console. Use this API to clear any active logins and force everyone to log in
// again. This API is useful for debugging purposes and when changing roles of a user. It forces any role changes to
// take effect immediately. Otherwise, permission or role changes will take effect during the user's next login or
// session expiration.
func (blockchain *BlockchainV2) DeleteAllSessions(deleteAllSessionsOptions *DeleteAllSessionsOptions) (result *DeleteAllSessionsResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(deleteAllSessionsOptions, "deleteAllSessionsOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"ak/api/v2/sessions"}
	pathParameters := []string{}

	builder := core.NewRequestBuilder(core.DELETE)
	_, err = builder.ConstructHTTPURL(blockchain.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range deleteAllSessionsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("blockchain", "V2", "DeleteAllSessions")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = blockchain.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalDeleteAllSessionsResponse)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// DeleteAllNotifications : Delete all notifications
// Delete all notifications. This API is intended for administration.
func (blockchain *BlockchainV2) DeleteAllNotifications(deleteAllNotificationsOptions *DeleteAllNotificationsOptions) (result *DeleteAllNotificationsResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(deleteAllNotificationsOptions, "deleteAllNotificationsOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"ak/api/v2/notifications/purge"}
	pathParameters := []string{}

	builder := core.NewRequestBuilder(core.DELETE)
	_, err = builder.ConstructHTTPURL(blockchain.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range deleteAllNotificationsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("blockchain", "V2", "DeleteAllNotifications")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = blockchain.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalDeleteAllNotificationsResponse)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// ClearCaches : Clear IBP console caches
// Clear the in-memory caches across all IBP console server processes. No effect on caches that are currently disabled.
func (blockchain *BlockchainV2) ClearCaches(clearCachesOptions *ClearCachesOptions) (result *CacheFlushResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(clearCachesOptions, "clearCachesOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"ak/api/v2/cache"}
	pathParameters := []string{}

	builder := core.NewRequestBuilder(core.DELETE)
	_, err = builder.ConstructHTTPURL(blockchain.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range clearCachesOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("blockchain", "V2", "ClearCaches")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = blockchain.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalCacheFlushResponse)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// GetPostman : Generate Postman collection
// Generate and download a Postman API Collection. The JSON contains all the APIs available in the IBP console. It can
// be imported to the [Postman](https://www.postman.com/downloads) desktop application. **The examples in the collection
// will be pre-populated with authorization credentials.** The authorization credentials to use must be provided to this
// API. See the query parameters for available options.
//
// Choose an auth strategy that matches your environment & concerns:
//
// - **IAM Bearer Auth** - *[Available on IBM Cloud]* - This is the recommended auth strategy. The same bearer token
// used to authenticate this request will be copied into the Postman collection examples. Since the bearer token expires
// the auth embedded in the collection will also expire. At that point the collection might be deleted & regenerated, or
// manually edited to refresh the authorization header values. To use this strategy set `auth_type` to `bearer`.
// - **IAM Api Key Auth** - *[Available on IBM Cloud]* - The IAM api key will be copied into the Postman collection
// examples. This means the auth embedded in the collection will never expire. To use this strategy set `auth_type` to
// `api_key`.
// - **Basic Auth** - *[Available on OpenShift & IBM Cloud Private]* - A basic auth username and password will be copied
// into the Postman collection examples. This is **not** available for an IBP SaaS instance on IBM Cloud. To use this
// strategy set `auth_type` to `basic`.
func (blockchain *BlockchainV2) GetPostman(getPostmanOptions *GetPostmanOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getPostmanOptions, "getPostmanOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getPostmanOptions, "getPostmanOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"ak/api/v2/postman"}
	pathParameters := []string{}

	builder := core.NewRequestBuilder(core.GET)
	_, err = builder.ConstructHTTPURL(blockchain.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range getPostmanOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("blockchain", "V2", "GetPostman")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	builder.AddQuery("auth_type", fmt.Sprint(*getPostmanOptions.AuthType))
	if getPostmanOptions.Token != nil {
		builder.AddQuery("token", fmt.Sprint(*getPostmanOptions.Token))
	}
	if getPostmanOptions.ApiKey != nil {
		builder.AddQuery("api_key", fmt.Sprint(*getPostmanOptions.ApiKey))
	}
	if getPostmanOptions.Username != nil {
		builder.AddQuery("username", fmt.Sprint(*getPostmanOptions.Username))
	}
	if getPostmanOptions.Password != nil {
		builder.AddQuery("password", fmt.Sprint(*getPostmanOptions.Password))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = blockchain.Service.Request(request, nil)

	return
}

// GetSwagger : Download OpenAPI file
// Download the [OpenAPI](https://swagger.io/specification/) specification file (aka swagger file) for the IBP console.
// This is the same file that was used to generate the APIs on this page. This file documents APIs offered by the IBP
// console.
func (blockchain *BlockchainV2) GetSwagger(getSwaggerOptions *GetSwaggerOptions) (result *string, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(getSwaggerOptions, "getSwaggerOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"ak/api/v2/openapi"}
	pathParameters := []string{}

	builder := core.NewRequestBuilder(core.GET)
	_, err = builder.ConstructHTTPURL(blockchain.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range getSwaggerOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("blockchain", "V2", "GetSwagger")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "text/plain")

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = blockchain.Service.Request(request, &result)

	return
}

// ArchiveNotificationsOptions : The ArchiveNotifications options.
type ArchiveNotificationsOptions struct {
	// Array of notification IDs to archive.
	NotificationIds []string `json:"notification_ids" validate:"required"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewArchiveNotificationsOptions : Instantiate ArchiveNotificationsOptions
func (*BlockchainV2) NewArchiveNotificationsOptions(notificationIds []string) *ArchiveNotificationsOptions {
	return &ArchiveNotificationsOptions{
		NotificationIds: notificationIds,
	}
}

// SetNotificationIds : Allow user to set NotificationIds
func (options *ArchiveNotificationsOptions) SetNotificationIds(notificationIds []string) *ArchiveNotificationsOptions {
	options.NotificationIds = notificationIds
	return options
}

// SetHeaders : Allow user to set Headers
func (options *ArchiveNotificationsOptions) SetHeaders(param map[string]string) *ArchiveNotificationsOptions {
	options.Headers = param
	return options
}

// ArchiveResponse : ArchiveResponse struct
type ArchiveResponse struct {
	// Response message. "ok" indicates the api completed successfully.
	Message *string `json:"message,omitempty"`

	// Text with the number of notifications that were archived.
	Details *string `json:"details,omitempty"`
}


// UnmarshalArchiveResponse unmarshals an instance of ArchiveResponse from the specified map of raw messages.
func UnmarshalArchiveResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ArchiveResponse)
	err = core.UnmarshalPrimitive(m, "message", &obj.Message)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "details", &obj.Details)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Bccsp : Configures the Blockchain Crypto Service Providers (bccsp).
type Bccsp struct {
	// The name of the crypto library implementation to use for the BlockChain Crypto Service Provider (bccsp). Defaults to
	// `SW`.
	Default *string `json:"Default,omitempty"`

	// Software based blockchain crypto provider.
	SW *BccspSW `json:"SW,omitempty"`

	// Hardware-based blockchain crypto provider.
	PKCS11 *BccspPKCS11 `json:"PKCS11,omitempty"`
}

// Constants associated with the Bccsp.Default property.
// The name of the crypto library implementation to use for the BlockChain Crypto Service Provider (bccsp). Defaults to
// `SW`.
const (
	Bccsp_Default_Pkcs11 = "PKCS11"
	Bccsp_Default_Sw = "SW"
)


// UnmarshalBccsp unmarshals an instance of Bccsp from the specified map of raw messages.
func UnmarshalBccsp(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Bccsp)
	err = core.UnmarshalPrimitive(m, "Default", &obj.Default)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "SW", &obj.SW, UnmarshalBccspSW)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "PKCS11", &obj.PKCS11, UnmarshalBccspPKCS11)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// BccspPKCS11 : Hardware-based blockchain crypto provider.
type BccspPKCS11 struct {
	// Token Label.
	Label *string `json:"Label" validate:"required"`

	// The user PIN.
	Pin *string `json:"Pin" validate:"required"`

	// The hash family to use for the BlockChain Crypto Service Provider (bccsp).
	Hash *string `json:"Hash,omitempty"`

	// The length of hash to use for the BlockChain Crypto Service Provider (bccsp).
	Security *float64 `json:"Security,omitempty"`
}


// NewBccspPKCS11 : Instantiate BccspPKCS11 (Generic Model Constructor)
func (*BlockchainV2) NewBccspPKCS11(label string, pin string) (model *BccspPKCS11, err error) {
	model = &BccspPKCS11{
		Label: core.StringPtr(label),
		Pin: core.StringPtr(pin),
	}
	err = core.ValidateStruct(model, "required parameters")
	return
}

// UnmarshalBccspPKCS11 unmarshals an instance of BccspPKCS11 from the specified map of raw messages.
func UnmarshalBccspPKCS11(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(BccspPKCS11)
	err = core.UnmarshalPrimitive(m, "Label", &obj.Label)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "Pin", &obj.Pin)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "Hash", &obj.Hash)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "Security", &obj.Security)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// BccspSW : Software based blockchain crypto provider.
type BccspSW struct {
	// The hash family to use for the BlockChain Crypto Service Provider (bccsp).
	Hash *string `json:"Hash" validate:"required"`

	// The length of hash to use for the BlockChain Crypto Service Provider (bccsp).
	Security *float64 `json:"Security" validate:"required"`
}


// NewBccspSW : Instantiate BccspSW (Generic Model Constructor)
func (*BlockchainV2) NewBccspSW(hash string, security float64) (model *BccspSW, err error) {
	model = &BccspSW{
		Hash: core.StringPtr(hash),
		Security: core.Float64Ptr(security),
	}
	err = core.ValidateStruct(model, "required parameters")
	return
}

// UnmarshalBccspSW unmarshals an instance of BccspSW from the specified map of raw messages.
func UnmarshalBccspSW(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(BccspSW)
	err = core.UnmarshalPrimitive(m, "Hash", &obj.Hash)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "Security", &obj.Security)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// CaResponse : Contains the details of a CA.
type CaResponse struct {
	// The unique identifier of this component.
	ID *string `json:"id,omitempty"`

	// The unique id for the component in Kubernetes. Not available if component was imported.
	DepComponentID *string `json:"dep_component_id,omitempty"`

	// The CA's "CAName" attribute.
	CaName *string `json:"ca_name,omitempty"`

	// A descriptive name for this CA. The IBP console tile displays this name.
	DisplayName *string `json:"display_name,omitempty"`

	// The gRPC URL for the peer. Typically, client applications would send requests to this URL. Include the protocol,
	// hostname/ip and port.
	ApiURL *string `json:"api_url,omitempty"`

	// The operations URL for the CA. Include the protocol, hostname/ip and port.
	OperationsURL *string `json:"operations_url,omitempty"`

	// The **cached** configuration override that was set for the Kubernetes deployment. Field does not exist if an
	// override was not set of if the component was imported.
	ConfigOverride interface{} `json:"config_override,omitempty"`

	// Indicates where the component is running.
	Location *string `json:"location,omitempty"`

	// UTC UNIX timestamp of component onboarding to the UI. In milliseconds.
	Timestamp *float64 `json:"timestamp,omitempty"`

	// The **cached** Kubernetes resource attributes for this component. Not available if CA was imported.
	Resources *CaResponseResources `json:"resources,omitempty"`

	// The versioning of the IBP console format of this JSON.
	SchemeVersion *string `json:"scheme_version,omitempty"`

	// The **cached** Kubernetes storage attributes for this component. Not available if CA was imported.
	Storage *CaResponseStorage `json:"storage,omitempty"`

	Tags []string `json:"tags,omitempty"`

	// The TLS certificate as base 64 encoded PEM. Certificate is used to secure/validate a TLS connection.
	TlsCert *string `json:"tls_cert,omitempty"`

	// The cached Hyperledger Fabric release version.
	Version *string `json:"version,omitempty"`

	// Specify the Kubernetes zone for the deployment. The deployment will use a k8s node in this zone. Find the list of
	// possible zones by retrieving your Kubernetes node labels: `kubectl get nodes --show-labels`. [More
	// information](https://kubernetes.io/docs/setup/best-practices/multiple-zones).
	Zone *string `json:"zone,omitempty"`
}


// UnmarshalCaResponse unmarshals an instance of CaResponse from the specified map of raw messages.
func UnmarshalCaResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(CaResponse)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "dep_component_id", &obj.DepComponentID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "ca_name", &obj.CaName)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "display_name", &obj.DisplayName)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "api_url", &obj.ApiURL)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "operations_url", &obj.OperationsURL)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "config_override", &obj.ConfigOverride)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "location", &obj.Location)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "timestamp", &obj.Timestamp)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "resources", &obj.Resources, UnmarshalCaResponseResources)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "scheme_version", &obj.SchemeVersion)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "storage", &obj.Storage, UnmarshalCaResponseStorage)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "tags", &obj.Tags)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "tls_cert", &obj.TlsCert)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "version", &obj.Version)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "zone", &obj.Zone)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// CaResponseResources : The **cached** Kubernetes resource attributes for this component. Not available if CA was imported.
type CaResponseResources struct {
	Ca *GenericResources `json:"ca,omitempty"`
}


// UnmarshalCaResponseResources unmarshals an instance of CaResponseResources from the specified map of raw messages.
func UnmarshalCaResponseResources(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(CaResponseResources)
	err = core.UnmarshalModel(m, "ca", &obj.Ca, UnmarshalGenericResources)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// CaResponseStorage : The **cached** Kubernetes storage attributes for this component. Not available if CA was imported.
type CaResponseStorage struct {
	Ca *StorageObject `json:"ca,omitempty"`
}


// UnmarshalCaResponseStorage unmarshals an instance of CaResponseStorage from the specified map of raw messages.
func UnmarshalCaResponseStorage(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(CaResponseStorage)
	err = core.UnmarshalModel(m, "ca", &obj.Ca, UnmarshalStorageObject)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// CacheData : CacheData struct
type CacheData struct {
	// Number of cache hits.
	Hits *float64 `json:"hits,omitempty"`

	// Number of cache misses.
	Misses *float64 `json:"misses,omitempty"`

	// Number of entries in the cache.
	Keys *float64 `json:"keys,omitempty"`

	// Approximate size of the in memory cache.
	CacheSize *string `json:"cache_size,omitempty"`
}


// UnmarshalCacheData unmarshals an instance of CacheData from the specified map of raw messages.
func UnmarshalCacheData(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(CacheData)
	err = core.UnmarshalPrimitive(m, "hits", &obj.Hits)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "misses", &obj.Misses)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "keys", &obj.Keys)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "cache_size", &obj.CacheSize)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// CacheFlushResponse : CacheFlushResponse struct
type CacheFlushResponse struct {
	// Response message. "ok" indicates the api completed successfully.
	Message *string `json:"message,omitempty"`

	// The name of the caches that were cleared.
	Flushed []string `json:"flushed,omitempty"`
}

// Constants associated with the CacheFlushResponse.Flushed property.
const (
	CacheFlushResponse_Flushed_CouchCache = "couch_cache"
	CacheFlushResponse_Flushed_IamCache = "iam_cache"
	CacheFlushResponse_Flushed_ProxyCache = "proxy_cache"
	CacheFlushResponse_Flushed_SessionCache = "session_cache"
)


// UnmarshalCacheFlushResponse unmarshals an instance of CacheFlushResponse from the specified map of raw messages.
func UnmarshalCacheFlushResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(CacheFlushResponse)
	err = core.UnmarshalPrimitive(m, "message", &obj.Message)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "flushed", &obj.Flushed)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ClearCachesOptions : The ClearCaches options.
type ClearCachesOptions struct {

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewClearCachesOptions : Instantiate ClearCachesOptions
func (*BlockchainV2) NewClearCachesOptions() *ClearCachesOptions {
	return &ClearCachesOptions{}
}

// SetHeaders : Allow user to set Headers
func (options *ClearCachesOptions) SetHeaders(param map[string]string) *ClearCachesOptions {
	options.Headers = param
	return options
}

// ConfigCACfgIdentities : ConfigCACfgIdentities struct
type ConfigCACfgIdentities struct {
	// The maximum number of incorrect password attempts allowed per identity.
	Passwordattempts *float64 `json:"passwordattempts" validate:"required"`
}


// NewConfigCACfgIdentities : Instantiate ConfigCACfgIdentities (Generic Model Constructor)
func (*BlockchainV2) NewConfigCACfgIdentities(passwordattempts float64) (model *ConfigCACfgIdentities, err error) {
	model = &ConfigCACfgIdentities{
		Passwordattempts: core.Float64Ptr(passwordattempts),
	}
	err = core.ValidateStruct(model, "required parameters")
	return
}

// UnmarshalConfigCACfgIdentities unmarshals an instance of ConfigCACfgIdentities from the specified map of raw messages.
func UnmarshalConfigCACfgIdentities(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ConfigCACfgIdentities)
	err = core.UnmarshalPrimitive(m, "passwordattempts", &obj.Passwordattempts)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ConfigCACreate : ConfigCACreate struct
type ConfigCACreate struct {
	Cors *ConfigCACors `json:"cors,omitempty"`

	// Enable debug to debug the CA.
	Debug *bool `json:"debug,omitempty"`

	// Max size of an acceptable CRL in bytes.
	Crlsizelimit *float64 `json:"crlsizelimit,omitempty"`

	Tls *ConfigCATls `json:"tls,omitempty"`

	Ca *ConfigCACa `json:"ca,omitempty"`

	Crl *ConfigCACrl `json:"crl,omitempty"`

	Registry *ConfigCARegistry `json:"registry" validate:"required"`

	Db *ConfigCADb `json:"db,omitempty"`

	// Set the keys to the desired affiliation parent names. The keys 'org1' and 'org2' are examples.
	Affiliations *ConfigCAAffiliations `json:"affiliations,omitempty"`

	Csr *ConfigCACsr `json:"csr,omitempty"`

	Idemix *ConfigCAIdemix `json:"idemix,omitempty"`

	// Configures the Blockchain Crypto Service Providers (bccsp).
	BCCSP *Bccsp `json:"BCCSP,omitempty"`

	Intermediate *ConfigCAIntermediate `json:"intermediate,omitempty"`

	Cfg *ConfigCACfg `json:"cfg,omitempty"`

	Metrics *Metrics `json:"metrics,omitempty"`
}


// NewConfigCACreate : Instantiate ConfigCACreate (Generic Model Constructor)
func (*BlockchainV2) NewConfigCACreate(registry *ConfigCARegistry) (model *ConfigCACreate, err error) {
	model = &ConfigCACreate{
		Registry: registry,
	}
	err = core.ValidateStruct(model, "required parameters")
	return
}

// UnmarshalConfigCACreate unmarshals an instance of ConfigCACreate from the specified map of raw messages.
func UnmarshalConfigCACreate(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ConfigCACreate)
	err = core.UnmarshalModel(m, "cors", &obj.Cors, UnmarshalConfigCACors)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "debug", &obj.Debug)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "crlsizelimit", &obj.Crlsizelimit)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "tls", &obj.Tls, UnmarshalConfigCATls)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "ca", &obj.Ca, UnmarshalConfigCACa)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "crl", &obj.Crl, UnmarshalConfigCACrl)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "registry", &obj.Registry, UnmarshalConfigCARegistry)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "db", &obj.Db, UnmarshalConfigCADb)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "affiliations", &obj.Affiliations, UnmarshalConfigCAAffiliations)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "csr", &obj.Csr, UnmarshalConfigCACsr)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "idemix", &obj.Idemix, UnmarshalConfigCAIdemix)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "BCCSP", &obj.BCCSP, UnmarshalBccsp)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "intermediate", &obj.Intermediate, UnmarshalConfigCAIntermediate)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "cfg", &obj.Cfg, UnmarshalConfigCACfg)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "metrics", &obj.Metrics, UnmarshalMetrics)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ConfigCACsrCa : ConfigCACsrCa struct
type ConfigCACsrCa struct {
	// The expiration for the root CA certificate.
	Expiry *string `json:"expiry,omitempty"`

	// The pathlength field is used to limit CA certificate hierarchy. 0 means that the CA cannot issue CA certs, only
	// entity certificates. 1 means that the CA can issue both.
	Pathlength *float64 `json:"pathlength,omitempty"`
}


// UnmarshalConfigCACsrCa unmarshals an instance of ConfigCACsrCa from the specified map of raw messages.
func UnmarshalConfigCACsrCa(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ConfigCACsrCa)
	err = core.UnmarshalPrimitive(m, "expiry", &obj.Expiry)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "pathlength", &obj.Pathlength)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ConfigCACsrKeyrequest : ConfigCACsrKeyrequest struct
type ConfigCACsrKeyrequest struct {
	// The algorithm to use for CSRs.
	Algo *string `json:"algo" validate:"required"`

	// The size of the key for CSRs.
	Size *float64 `json:"size" validate:"required"`
}


// NewConfigCACsrKeyrequest : Instantiate ConfigCACsrKeyrequest (Generic Model Constructor)
func (*BlockchainV2) NewConfigCACsrKeyrequest(algo string, size float64) (model *ConfigCACsrKeyrequest, err error) {
	model = &ConfigCACsrKeyrequest{
		Algo: core.StringPtr(algo),
		Size: core.Float64Ptr(size),
	}
	err = core.ValidateStruct(model, "required parameters")
	return
}

// UnmarshalConfigCACsrKeyrequest unmarshals an instance of ConfigCACsrKeyrequest from the specified map of raw messages.
func UnmarshalConfigCACsrKeyrequest(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ConfigCACsrKeyrequest)
	err = core.UnmarshalPrimitive(m, "algo", &obj.Algo)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "size", &obj.Size)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ConfigCACsrNamesItem : ConfigCACsrNamesItem struct
type ConfigCACsrNamesItem struct {
	C *string `json:"C" validate:"required"`

	ST *string `json:"ST" validate:"required"`

	L *string `json:"L,omitempty"`

	O *string `json:"O" validate:"required"`

	OU *string `json:"OU,omitempty"`
}


// NewConfigCACsrNamesItem : Instantiate ConfigCACsrNamesItem (Generic Model Constructor)
func (*BlockchainV2) NewConfigCACsrNamesItem(c string, sT string, o string) (model *ConfigCACsrNamesItem, err error) {
	model = &ConfigCACsrNamesItem{
		C: core.StringPtr(c),
		ST: core.StringPtr(sT),
		O: core.StringPtr(o),
	}
	err = core.ValidateStruct(model, "required parameters")
	return
}

// UnmarshalConfigCACsrNamesItem unmarshals an instance of ConfigCACsrNamesItem from the specified map of raw messages.
func UnmarshalConfigCACsrNamesItem(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ConfigCACsrNamesItem)
	err = core.UnmarshalPrimitive(m, "C", &obj.C)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "ST", &obj.ST)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "L", &obj.L)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "O", &obj.O)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "OU", &obj.OU)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ConfigCADbTls : ConfigCADbTls struct
type ConfigCADbTls struct {
	Certfiles []string `json:"certfiles,omitempty"`

	Client *ConfigCADbTlsClient `json:"client,omitempty"`

	// Set to true if TLS is to be used between the CA and its database, else false.
	Enabled *bool `json:"enabled,omitempty"`
}


// UnmarshalConfigCADbTls unmarshals an instance of ConfigCADbTls from the specified map of raw messages.
func UnmarshalConfigCADbTls(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ConfigCADbTls)
	err = core.UnmarshalPrimitive(m, "certfiles", &obj.Certfiles)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "client", &obj.Client, UnmarshalConfigCADbTlsClient)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "enabled", &obj.Enabled)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ConfigCADbTlsClient : ConfigCADbTlsClient struct
type ConfigCADbTlsClient struct {
	// The TLS certificate for client TLS as base 64 encoded PEM.
	Certfile *string `json:"certfile" validate:"required"`

	// The TLS private key for client TLS as base 64 encoded PEM.
	Keyfile *string `json:"keyfile" validate:"required"`
}


// NewConfigCADbTlsClient : Instantiate ConfigCADbTlsClient (Generic Model Constructor)
func (*BlockchainV2) NewConfigCADbTlsClient(certfile string, keyfile string) (model *ConfigCADbTlsClient, err error) {
	model = &ConfigCADbTlsClient{
		Certfile: core.StringPtr(certfile),
		Keyfile: core.StringPtr(keyfile),
	}
	err = core.ValidateStruct(model, "required parameters")
	return
}

// UnmarshalConfigCADbTlsClient unmarshals an instance of ConfigCADbTlsClient from the specified map of raw messages.
func UnmarshalConfigCADbTlsClient(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ConfigCADbTlsClient)
	err = core.UnmarshalPrimitive(m, "certfile", &obj.Certfile)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "keyfile", &obj.Keyfile)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ConfigCAIntermediateEnrollment : ConfigCAIntermediateEnrollment struct
type ConfigCAIntermediateEnrollment struct {
	// Hosts to set when issuing the certificate.
	Hosts *string `json:"hosts" validate:"required"`

	// Name of the signing profile to use when issuing the certificate.
	Profile *string `json:"profile" validate:"required"`

	// Label to use in HSM operations.
	Label *string `json:"label" validate:"required"`
}


// NewConfigCAIntermediateEnrollment : Instantiate ConfigCAIntermediateEnrollment (Generic Model Constructor)
func (*BlockchainV2) NewConfigCAIntermediateEnrollment(hosts string, profile string, label string) (model *ConfigCAIntermediateEnrollment, err error) {
	model = &ConfigCAIntermediateEnrollment{
		Hosts: core.StringPtr(hosts),
		Profile: core.StringPtr(profile),
		Label: core.StringPtr(label),
	}
	err = core.ValidateStruct(model, "required parameters")
	return
}

// UnmarshalConfigCAIntermediateEnrollment unmarshals an instance of ConfigCAIntermediateEnrollment from the specified map of raw messages.
func UnmarshalConfigCAIntermediateEnrollment(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ConfigCAIntermediateEnrollment)
	err = core.UnmarshalPrimitive(m, "hosts", &obj.Hosts)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "profile", &obj.Profile)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "label", &obj.Label)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ConfigCAIntermediateParentserver : ConfigCAIntermediateParentserver struct
type ConfigCAIntermediateParentserver struct {
	// The url of the parent server. Include the protocol, hostname/ip and port.
	URL *string `json:"url" validate:"required"`

	// The name of the CA to enroll within the server.
	Caname *string `json:"caname" validate:"required"`
}


// NewConfigCAIntermediateParentserver : Instantiate ConfigCAIntermediateParentserver (Generic Model Constructor)
func (*BlockchainV2) NewConfigCAIntermediateParentserver(url string, caname string) (model *ConfigCAIntermediateParentserver, err error) {
	model = &ConfigCAIntermediateParentserver{
		URL: core.StringPtr(url),
		Caname: core.StringPtr(caname),
	}
	err = core.ValidateStruct(model, "required parameters")
	return
}

// UnmarshalConfigCAIntermediateParentserver unmarshals an instance of ConfigCAIntermediateParentserver from the specified map of raw messages.
func UnmarshalConfigCAIntermediateParentserver(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ConfigCAIntermediateParentserver)
	err = core.UnmarshalPrimitive(m, "url", &obj.URL)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "caname", &obj.Caname)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ConfigCAIntermediateTls : ConfigCAIntermediateTls struct
type ConfigCAIntermediateTls struct {
	Certfiles []string `json:"certfiles" validate:"required"`

	Client *ConfigCAIntermediateTlsClient `json:"client,omitempty"`
}


// NewConfigCAIntermediateTls : Instantiate ConfigCAIntermediateTls (Generic Model Constructor)
func (*BlockchainV2) NewConfigCAIntermediateTls(certfiles []string) (model *ConfigCAIntermediateTls, err error) {
	model = &ConfigCAIntermediateTls{
		Certfiles: certfiles,
	}
	err = core.ValidateStruct(model, "required parameters")
	return
}

// UnmarshalConfigCAIntermediateTls unmarshals an instance of ConfigCAIntermediateTls from the specified map of raw messages.
func UnmarshalConfigCAIntermediateTls(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ConfigCAIntermediateTls)
	err = core.UnmarshalPrimitive(m, "certfiles", &obj.Certfiles)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "client", &obj.Client, UnmarshalConfigCAIntermediateTlsClient)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ConfigCAIntermediateTlsClient : ConfigCAIntermediateTlsClient struct
type ConfigCAIntermediateTlsClient struct {
	// The TLS certificate for client TLS as base 64 encoded PEM.
	Certfile *string `json:"certfile" validate:"required"`

	// The TLS private key for client TLS as base 64 encoded PEM.
	Keyfile *string `json:"keyfile" validate:"required"`
}


// NewConfigCAIntermediateTlsClient : Instantiate ConfigCAIntermediateTlsClient (Generic Model Constructor)
func (*BlockchainV2) NewConfigCAIntermediateTlsClient(certfile string, keyfile string) (model *ConfigCAIntermediateTlsClient, err error) {
	model = &ConfigCAIntermediateTlsClient{
		Certfile: core.StringPtr(certfile),
		Keyfile: core.StringPtr(keyfile),
	}
	err = core.ValidateStruct(model, "required parameters")
	return
}

// UnmarshalConfigCAIntermediateTlsClient unmarshals an instance of ConfigCAIntermediateTlsClient from the specified map of raw messages.
func UnmarshalConfigCAIntermediateTlsClient(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ConfigCAIntermediateTlsClient)
	err = core.UnmarshalPrimitive(m, "certfile", &obj.Certfile)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "keyfile", &obj.Keyfile)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ConfigCARegistryIdentitiesItem : ConfigCARegistryIdentitiesItem struct
type ConfigCARegistryIdentitiesItem struct {
	// The ID for the identity, aka enroll id.
	Name *string `json:"name" validate:"required"`

	// The password for the identity, aka enroll secret.
	Pass *string `json:"pass" validate:"required"`

	// The type of identity.
	Type *string `json:"type" validate:"required"`

	// Maximum number of enrollments for id. Set -1 for infinite.
	Maxenrollments *float64 `json:"maxenrollments,omitempty"`

	// The affiliation data for the identity.
	Affiliation *string `json:"affiliation,omitempty"`

	Attrs *IdentityAttrs `json:"attrs,omitempty"`
}

// Constants associated with the ConfigCARegistryIdentitiesItem.Type property.
// The type of identity.
const (
	ConfigCARegistryIdentitiesItem_Type_Admin = "admin"
	ConfigCARegistryIdentitiesItem_Type_Client = "client"
	ConfigCARegistryIdentitiesItem_Type_Orderer = "orderer"
	ConfigCARegistryIdentitiesItem_Type_Peer = "peer"
	ConfigCARegistryIdentitiesItem_Type_User = "user"
)


// NewConfigCARegistryIdentitiesItem : Instantiate ConfigCARegistryIdentitiesItem (Generic Model Constructor)
func (*BlockchainV2) NewConfigCARegistryIdentitiesItem(name string, pass string, typeVar string) (model *ConfigCARegistryIdentitiesItem, err error) {
	model = &ConfigCARegistryIdentitiesItem{
		Name: core.StringPtr(name),
		Pass: core.StringPtr(pass),
		Type: core.StringPtr(typeVar),
	}
	err = core.ValidateStruct(model, "required parameters")
	return
}

// UnmarshalConfigCARegistryIdentitiesItem unmarshals an instance of ConfigCARegistryIdentitiesItem from the specified map of raw messages.
func UnmarshalConfigCARegistryIdentitiesItem(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ConfigCARegistryIdentitiesItem)
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "pass", &obj.Pass)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "maxenrollments", &obj.Maxenrollments)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "affiliation", &obj.Affiliation)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "attrs", &obj.Attrs, UnmarshalIdentityAttrs)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ConfigCATlsClientauth : ConfigCATlsClientauth struct
type ConfigCATlsClientauth struct {
	Type *string `json:"type" validate:"required"`

	Certfiles []string `json:"certfiles" validate:"required"`
}


// NewConfigCATlsClientauth : Instantiate ConfigCATlsClientauth (Generic Model Constructor)
func (*BlockchainV2) NewConfigCATlsClientauth(typeVar string, certfiles []string) (model *ConfigCATlsClientauth, err error) {
	model = &ConfigCATlsClientauth{
		Type: core.StringPtr(typeVar),
		Certfiles: certfiles,
	}
	err = core.ValidateStruct(model, "required parameters")
	return
}

// UnmarshalConfigCATlsClientauth unmarshals an instance of ConfigCATlsClientauth from the specified map of raw messages.
func UnmarshalConfigCATlsClientauth(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ConfigCATlsClientauth)
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "certfiles", &obj.Certfiles)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ConfigCAUpdate : ConfigCAUpdate struct
type ConfigCAUpdate struct {
	Cors *ConfigCACors `json:"cors,omitempty"`

	// Enable debug to debug the CA.
	Debug *bool `json:"debug,omitempty"`

	// Max size of an acceptable CRL in bytes.
	Crlsizelimit *float64 `json:"crlsizelimit,omitempty"`

	Tls *ConfigCATls `json:"tls,omitempty"`

	Ca *ConfigCACa `json:"ca,omitempty"`

	Crl *ConfigCACrl `json:"crl,omitempty"`

	Registry *ConfigCARegistry `json:"registry,omitempty"`

	Db *ConfigCADb `json:"db,omitempty"`

	// Set the keys to the desired affiliation parent names. The keys 'org1' and 'org2' are examples.
	Affiliations *ConfigCAAffiliations `json:"affiliations,omitempty"`

	Csr *ConfigCACsr `json:"csr,omitempty"`

	Idemix *ConfigCAIdemix `json:"idemix,omitempty"`

	// Configures the Blockchain Crypto Service Providers (bccsp).
	BCCSP *Bccsp `json:"BCCSP,omitempty"`

	Intermediate *ConfigCAIntermediate `json:"intermediate,omitempty"`

	Cfg *ConfigCACfg `json:"cfg,omitempty"`

	Metrics *Metrics `json:"metrics,omitempty"`
}


// UnmarshalConfigCAUpdate unmarshals an instance of ConfigCAUpdate from the specified map of raw messages.
func UnmarshalConfigCAUpdate(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ConfigCAUpdate)
	err = core.UnmarshalModel(m, "cors", &obj.Cors, UnmarshalConfigCACors)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "debug", &obj.Debug)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "crlsizelimit", &obj.Crlsizelimit)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "tls", &obj.Tls, UnmarshalConfigCATls)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "ca", &obj.Ca, UnmarshalConfigCACa)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "crl", &obj.Crl, UnmarshalConfigCACrl)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "registry", &obj.Registry, UnmarshalConfigCARegistry)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "db", &obj.Db, UnmarshalConfigCADb)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "affiliations", &obj.Affiliations, UnmarshalConfigCAAffiliations)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "csr", &obj.Csr, UnmarshalConfigCACsr)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "idemix", &obj.Idemix, UnmarshalConfigCAIdemix)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "BCCSP", &obj.BCCSP, UnmarshalBccsp)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "intermediate", &obj.Intermediate, UnmarshalConfigCAIntermediate)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "cfg", &obj.Cfg, UnmarshalConfigCACfg)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "metrics", &obj.Metrics, UnmarshalMetrics)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ConfigCAAffiliations : Set the keys to the desired affiliation parent names. The keys 'org1' and 'org2' are examples.
type ConfigCAAffiliations struct {
	Org1 []string `json:"org1,omitempty"`

	Org2 []string `json:"org2,omitempty"`

	// Allows users to set arbitrary properties
	additionalProperties map[string]interface{}
}


// SetProperty allows the user to set an arbitrary property on an instance of ConfigCAAffiliations
func (o *ConfigCAAffiliations) SetProperty(key string, value interface{}) {
	if o.additionalProperties == nil {
		o.additionalProperties = make(map[string]interface{})
	}
	o.additionalProperties[key] = value
}

// GetProperty allows the user to retrieve an arbitrary property from an instance of ConfigCAAffiliations
func (o *ConfigCAAffiliations) GetProperty(key string) interface{} {
	return o.additionalProperties[key]
}

// GetProperties allows the user to retrieve the map of arbitrary properties from an instance of ConfigCAAffiliations
func (o *ConfigCAAffiliations) GetProperties() map[string]interface{} {
	return o.additionalProperties
}

// MarshalJSON performs custom serialization for instances of ConfigCAAffiliations
func (o *ConfigCAAffiliations) MarshalJSON() (buffer []byte, err error) {
	m := make(map[string]interface{})
	if len(o.additionalProperties) > 0 {
		for k, v := range o.additionalProperties {
			m[k] = v
		}
	}
	if o.Org1 != nil {
		m["org1"] = o.Org1
	}
	if o.Org2 != nil {
		m["org2"] = o.Org2
	}
	buffer, err = json.Marshal(m)
	return
}

// UnmarshalConfigCAAffiliations unmarshals an instance of ConfigCAAffiliations from the specified map of raw messages.
func UnmarshalConfigCAAffiliations(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ConfigCAAffiliations)
	err = core.UnmarshalPrimitive(m, "org1", &obj.Org1)
	if err != nil {
		return
	}
	delete(m, "org1")
	err = core.UnmarshalPrimitive(m, "org2", &obj.Org2)
	if err != nil {
		return
	}
	delete(m, "org2")
	for k := range m {
		var v interface{}
		e := core.UnmarshalPrimitive(m, k, &v)
		if e != nil {
			err = e
			return
		}
		obj.SetProperty(k, v)
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ConfigCACa : ConfigCACa struct
type ConfigCACa struct {
	// The CA's private key as base 64 encoded PEM.
	Keyfile *string `json:"keyfile,omitempty"`

	// The CA's certificate as base 64 encoded PEM.
	Certfile *string `json:"certfile,omitempty"`

	// The CA's certificate chain as base 64 encoded PEM.
	Chainfile *string `json:"chainfile,omitempty"`
}


// UnmarshalConfigCACa unmarshals an instance of ConfigCACa from the specified map of raw messages.
func UnmarshalConfigCACa(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ConfigCACa)
	err = core.UnmarshalPrimitive(m, "keyfile", &obj.Keyfile)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "certfile", &obj.Certfile)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "chainfile", &obj.Chainfile)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ConfigCACfg : ConfigCACfg struct
type ConfigCACfg struct {
	Identities *ConfigCACfgIdentities `json:"identities" validate:"required"`
}


// NewConfigCACfg : Instantiate ConfigCACfg (Generic Model Constructor)
func (*BlockchainV2) NewConfigCACfg(identities *ConfigCACfgIdentities) (model *ConfigCACfg, err error) {
	model = &ConfigCACfg{
		Identities: identities,
	}
	err = core.ValidateStruct(model, "required parameters")
	return
}

// UnmarshalConfigCACfg unmarshals an instance of ConfigCACfg from the specified map of raw messages.
func UnmarshalConfigCACfg(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ConfigCACfg)
	err = core.UnmarshalModel(m, "identities", &obj.Identities, UnmarshalConfigCACfgIdentities)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ConfigCACors : ConfigCACors struct
type ConfigCACors struct {
	Enabled *bool `json:"enabled" validate:"required"`

	Origins []string `json:"origins" validate:"required"`
}


// NewConfigCACors : Instantiate ConfigCACors (Generic Model Constructor)
func (*BlockchainV2) NewConfigCACors(enabled bool, origins []string) (model *ConfigCACors, err error) {
	model = &ConfigCACors{
		Enabled: core.BoolPtr(enabled),
		Origins: origins,
	}
	err = core.ValidateStruct(model, "required parameters")
	return
}

// UnmarshalConfigCACors unmarshals an instance of ConfigCACors from the specified map of raw messages.
func UnmarshalConfigCACors(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ConfigCACors)
	err = core.UnmarshalPrimitive(m, "enabled", &obj.Enabled)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "origins", &obj.Origins)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ConfigCACrl : ConfigCACrl struct
type ConfigCACrl struct {
	// Expiration of the CRL (Certificate Revocation List) generated by the 'gencrl' requests.
	Expiry *string `json:"expiry" validate:"required"`
}


// NewConfigCACrl : Instantiate ConfigCACrl (Generic Model Constructor)
func (*BlockchainV2) NewConfigCACrl(expiry string) (model *ConfigCACrl, err error) {
	model = &ConfigCACrl{
		Expiry: core.StringPtr(expiry),
	}
	err = core.ValidateStruct(model, "required parameters")
	return
}

// UnmarshalConfigCACrl unmarshals an instance of ConfigCACrl from the specified map of raw messages.
func UnmarshalConfigCACrl(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ConfigCACrl)
	err = core.UnmarshalPrimitive(m, "expiry", &obj.Expiry)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ConfigCACsr : ConfigCACsr struct
type ConfigCACsr struct {
	// The Common Name for the CSRs.
	Cn *string `json:"cn" validate:"required"`

	Keyrequest *ConfigCACsrKeyrequest `json:"keyrequest,omitempty"`

	Names []ConfigCACsrNamesItem `json:"names" validate:"required"`

	Hosts []string `json:"hosts,omitempty"`

	Ca *ConfigCACsrCa `json:"ca" validate:"required"`
}


// NewConfigCACsr : Instantiate ConfigCACsr (Generic Model Constructor)
func (*BlockchainV2) NewConfigCACsr(cn string, names []ConfigCACsrNamesItem, ca *ConfigCACsrCa) (model *ConfigCACsr, err error) {
	model = &ConfigCACsr{
		Cn: core.StringPtr(cn),
		Names: names,
		Ca: ca,
	}
	err = core.ValidateStruct(model, "required parameters")
	return
}

// UnmarshalConfigCACsr unmarshals an instance of ConfigCACsr from the specified map of raw messages.
func UnmarshalConfigCACsr(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ConfigCACsr)
	err = core.UnmarshalPrimitive(m, "cn", &obj.Cn)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "keyrequest", &obj.Keyrequest, UnmarshalConfigCACsrKeyrequest)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "names", &obj.Names, UnmarshalConfigCACsrNamesItem)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "hosts", &obj.Hosts)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "ca", &obj.Ca, UnmarshalConfigCACsrCa)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ConfigCADb : ConfigCADb struct
type ConfigCADb struct {
	// The type of database. Either 'sqlite3', 'postgres', 'mysql'. Defaults 'sqlite3'.
	Type *string `json:"type" validate:"required"`

	// Build this string - "host=\<hostname> port=\<port> user=\<username> password=\<password> dbname=ibmclouddb
	// sslmode=verify-full".
	Datasource *string `json:"datasource" validate:"required"`

	Tls *ConfigCADbTls `json:"tls,omitempty"`
}

// Constants associated with the ConfigCADb.Type property.
// The type of database. Either 'sqlite3', 'postgres', 'mysql'. Defaults 'sqlite3'.
const (
	ConfigCADb_Type_Mysql = "mysql"
	ConfigCADb_Type_Postgres = "postgres"
	ConfigCADb_Type_Sqlite3 = "sqlite3"
)


// NewConfigCADb : Instantiate ConfigCADb (Generic Model Constructor)
func (*BlockchainV2) NewConfigCADb(typeVar string, datasource string) (model *ConfigCADb, err error) {
	model = &ConfigCADb{
		Type: core.StringPtr(typeVar),
		Datasource: core.StringPtr(datasource),
	}
	err = core.ValidateStruct(model, "required parameters")
	return
}

// UnmarshalConfigCADb unmarshals an instance of ConfigCADb from the specified map of raw messages.
func UnmarshalConfigCADb(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ConfigCADb)
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "datasource", &obj.Datasource)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "tls", &obj.Tls, UnmarshalConfigCADbTls)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ConfigCAIdemix : ConfigCAIdemix struct
type ConfigCAIdemix struct {
	// Specifies the revocation pool size.
	Rhpoolsize *float64 `json:"rhpoolsize" validate:"required"`

	// Specifies the expiration for the nonces.
	Nonceexpiration *string `json:"nonceexpiration" validate:"required"`

	// Specifies frequency at which expired nonces are removed from data store.
	Noncesweepinterval *string `json:"noncesweepinterval" validate:"required"`
}


// NewConfigCAIdemix : Instantiate ConfigCAIdemix (Generic Model Constructor)
func (*BlockchainV2) NewConfigCAIdemix(rhpoolsize float64, nonceexpiration string, noncesweepinterval string) (model *ConfigCAIdemix, err error) {
	model = &ConfigCAIdemix{
		Rhpoolsize: core.Float64Ptr(rhpoolsize),
		Nonceexpiration: core.StringPtr(nonceexpiration),
		Noncesweepinterval: core.StringPtr(noncesweepinterval),
	}
	err = core.ValidateStruct(model, "required parameters")
	return
}

// UnmarshalConfigCAIdemix unmarshals an instance of ConfigCAIdemix from the specified map of raw messages.
func UnmarshalConfigCAIdemix(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ConfigCAIdemix)
	err = core.UnmarshalPrimitive(m, "rhpoolsize", &obj.Rhpoolsize)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "nonceexpiration", &obj.Nonceexpiration)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "noncesweepinterval", &obj.Noncesweepinterval)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ConfigCAIntermediate : ConfigCAIntermediate struct
type ConfigCAIntermediate struct {
	Parentserver *ConfigCAIntermediateParentserver `json:"parentserver" validate:"required"`

	Enrollment *ConfigCAIntermediateEnrollment `json:"enrollment,omitempty"`

	Tls *ConfigCAIntermediateTls `json:"tls,omitempty"`
}


// NewConfigCAIntermediate : Instantiate ConfigCAIntermediate (Generic Model Constructor)
func (*BlockchainV2) NewConfigCAIntermediate(parentserver *ConfigCAIntermediateParentserver) (model *ConfigCAIntermediate, err error) {
	model = &ConfigCAIntermediate{
		Parentserver: parentserver,
	}
	err = core.ValidateStruct(model, "required parameters")
	return
}

// UnmarshalConfigCAIntermediate unmarshals an instance of ConfigCAIntermediate from the specified map of raw messages.
func UnmarshalConfigCAIntermediate(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ConfigCAIntermediate)
	err = core.UnmarshalModel(m, "parentserver", &obj.Parentserver, UnmarshalConfigCAIntermediateParentserver)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "enrollment", &obj.Enrollment, UnmarshalConfigCAIntermediateEnrollment)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "tls", &obj.Tls, UnmarshalConfigCAIntermediateTls)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ConfigCARegistry : ConfigCARegistry struct
type ConfigCARegistry struct {
	// Default maximum number of enrollments per id. Set -1 for infinite.
	Maxenrollments *float64 `json:"maxenrollments" validate:"required"`

	Identities []ConfigCARegistryIdentitiesItem `json:"identities" validate:"required"`
}


// NewConfigCARegistry : Instantiate ConfigCARegistry (Generic Model Constructor)
func (*BlockchainV2) NewConfigCARegistry(maxenrollments float64, identities []ConfigCARegistryIdentitiesItem) (model *ConfigCARegistry, err error) {
	model = &ConfigCARegistry{
		Maxenrollments: core.Float64Ptr(maxenrollments),
		Identities: identities,
	}
	err = core.ValidateStruct(model, "required parameters")
	return
}

// UnmarshalConfigCARegistry unmarshals an instance of ConfigCARegistry from the specified map of raw messages.
func UnmarshalConfigCARegistry(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ConfigCARegistry)
	err = core.UnmarshalPrimitive(m, "maxenrollments", &obj.Maxenrollments)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "identities", &obj.Identities, UnmarshalConfigCARegistryIdentitiesItem)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ConfigCATls : ConfigCATls struct
type ConfigCATls struct {
	// The CA's private key as base 64 encoded PEM.
	Keyfile *string `json:"keyfile" validate:"required"`

	// The CA's certificate as base 64 encoded PEM.
	Certfile *string `json:"certfile" validate:"required"`

	Clientauth *ConfigCATlsClientauth `json:"clientauth,omitempty"`
}


// NewConfigCATls : Instantiate ConfigCATls (Generic Model Constructor)
func (*BlockchainV2) NewConfigCATls(keyfile string, certfile string) (model *ConfigCATls, err error) {
	model = &ConfigCATls{
		Keyfile: core.StringPtr(keyfile),
		Certfile: core.StringPtr(certfile),
	}
	err = core.ValidateStruct(model, "required parameters")
	return
}

// UnmarshalConfigCATls unmarshals an instance of ConfigCATls from the specified map of raw messages.
func UnmarshalConfigCATls(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ConfigCATls)
	err = core.UnmarshalPrimitive(m, "keyfile", &obj.Keyfile)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "certfile", &obj.Certfile)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "clientauth", &obj.Clientauth, UnmarshalConfigCATlsClientauth)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ConfigObject : ConfigObject struct
type ConfigObject struct {
	// This `enrollment` field contains data that allows a component to enroll an identity for itself. Use `enrollment` or
	// `msp`, not both.
	Enrollment *ConfigObjectEnrollment `json:"enrollment,omitempty"`

	// The `msp` field contains data to allow a component to configure its MSP with an already enrolled identity. Use `msp`
	// or `enrollment`, not both.
	Msp *ConfigObjectMsp `json:"msp,omitempty"`
}


// UnmarshalConfigObject unmarshals an instance of ConfigObject from the specified map of raw messages.
func UnmarshalConfigObject(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ConfigObject)
	err = core.UnmarshalModel(m, "enrollment", &obj.Enrollment, UnmarshalConfigObjectEnrollment)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "msp", &obj.Msp, UnmarshalConfigObjectMsp)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ConfigObjectEnrollment : This `enrollment` field contains data that allows a component to enroll an identity for itself. Use `enrollment` or
// `msp`, not both.
type ConfigObjectEnrollment struct {
	Component *ConfigObjectEnrollmentComponent `json:"component" validate:"required"`

	Tls *ConfigObjectEnrollmentTls `json:"tls" validate:"required"`
}


// NewConfigObjectEnrollment : Instantiate ConfigObjectEnrollment (Generic Model Constructor)
func (*BlockchainV2) NewConfigObjectEnrollment(component *ConfigObjectEnrollmentComponent, tls *ConfigObjectEnrollmentTls) (model *ConfigObjectEnrollment, err error) {
	model = &ConfigObjectEnrollment{
		Component: component,
		Tls: tls,
	}
	err = core.ValidateStruct(model, "required parameters")
	return
}

// UnmarshalConfigObjectEnrollment unmarshals an instance of ConfigObjectEnrollment from the specified map of raw messages.
func UnmarshalConfigObjectEnrollment(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ConfigObjectEnrollment)
	err = core.UnmarshalModel(m, "component", &obj.Component, UnmarshalConfigObjectEnrollmentComponent)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "tls", &obj.Tls, UnmarshalConfigObjectEnrollmentTls)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ConfigObjectEnrollmentComponent : ConfigObjectEnrollmentComponent struct
type ConfigObjectEnrollmentComponent struct {
	// The CA's hostname. Do not include protocol or port.
	Cahost *string `json:"cahost" validate:"required"`

	// The CA's port.
	Caport *float64 `json:"caport" validate:"required"`

	// The CA's "CAName" attribute.
	Caname *string `json:"caname" validate:"required"`

	Catls *ConfigObjectEnrollmentComponentCatls `json:"catls" validate:"required"`

	// The username of the enroll id.
	Enrollid *string `json:"enrollid" validate:"required"`

	// The password of the enroll id.
	Enrollsecret *string `json:"enrollsecret" validate:"required"`

	Admincerts []string `json:"admincerts,omitempty"`
}


// NewConfigObjectEnrollmentComponent : Instantiate ConfigObjectEnrollmentComponent (Generic Model Constructor)
func (*BlockchainV2) NewConfigObjectEnrollmentComponent(cahost string, caport float64, caname string, catls *ConfigObjectEnrollmentComponentCatls, enrollid string, enrollsecret string) (model *ConfigObjectEnrollmentComponent, err error) {
	model = &ConfigObjectEnrollmentComponent{
		Cahost: core.StringPtr(cahost),
		Caport: core.Float64Ptr(caport),
		Caname: core.StringPtr(caname),
		Catls: catls,
		Enrollid: core.StringPtr(enrollid),
		Enrollsecret: core.StringPtr(enrollsecret),
	}
	err = core.ValidateStruct(model, "required parameters")
	return
}

// UnmarshalConfigObjectEnrollmentComponent unmarshals an instance of ConfigObjectEnrollmentComponent from the specified map of raw messages.
func UnmarshalConfigObjectEnrollmentComponent(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ConfigObjectEnrollmentComponent)
	err = core.UnmarshalPrimitive(m, "cahost", &obj.Cahost)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "caport", &obj.Caport)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "caname", &obj.Caname)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "catls", &obj.Catls, UnmarshalConfigObjectEnrollmentComponentCatls)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "enrollid", &obj.Enrollid)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "enrollsecret", &obj.Enrollsecret)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "admincerts", &obj.Admincerts)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ConfigObjectEnrollmentComponentCatls : ConfigObjectEnrollmentComponentCatls struct
type ConfigObjectEnrollmentComponentCatls struct {
	// The TLS certificate as base 64 encoded PEM. Certificate is used to secure/validate a TLS connection.
	Cacert *string `json:"cacert" validate:"required"`
}


// NewConfigObjectEnrollmentComponentCatls : Instantiate ConfigObjectEnrollmentComponentCatls (Generic Model Constructor)
func (*BlockchainV2) NewConfigObjectEnrollmentComponentCatls(cacert string) (model *ConfigObjectEnrollmentComponentCatls, err error) {
	model = &ConfigObjectEnrollmentComponentCatls{
		Cacert: core.StringPtr(cacert),
	}
	err = core.ValidateStruct(model, "required parameters")
	return
}

// UnmarshalConfigObjectEnrollmentComponentCatls unmarshals an instance of ConfigObjectEnrollmentComponentCatls from the specified map of raw messages.
func UnmarshalConfigObjectEnrollmentComponentCatls(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ConfigObjectEnrollmentComponentCatls)
	err = core.UnmarshalPrimitive(m, "cacert", &obj.Cacert)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ConfigObjectEnrollmentTls : ConfigObjectEnrollmentTls struct
type ConfigObjectEnrollmentTls struct {
	// The CA's hostname. Do not include protocol or port.
	Cahost *string `json:"cahost" validate:"required"`

	// The CA's port.
	Caport *float64 `json:"caport" validate:"required"`

	// The TLS CA's "CAName" attribute.
	Caname *string `json:"caname" validate:"required"`

	Catls *ConfigObjectEnrollmentTlsCatls `json:"catls" validate:"required"`

	// The username of the enroll id.
	Enrollid *string `json:"enrollid,omitempty"`

	// The password of the enroll id.
	Enrollsecret *string `json:"enrollsecret,omitempty"`

	Admincerts []string `json:"admincerts,omitempty"`

	Csr *ConfigObjectEnrollmentTlsCsr `json:"csr,omitempty"`
}


// NewConfigObjectEnrollmentTls : Instantiate ConfigObjectEnrollmentTls (Generic Model Constructor)
func (*BlockchainV2) NewConfigObjectEnrollmentTls(cahost string, caport float64, caname string, catls *ConfigObjectEnrollmentTlsCatls) (model *ConfigObjectEnrollmentTls, err error) {
	model = &ConfigObjectEnrollmentTls{
		Cahost: core.StringPtr(cahost),
		Caport: core.Float64Ptr(caport),
		Caname: core.StringPtr(caname),
		Catls: catls,
	}
	err = core.ValidateStruct(model, "required parameters")
	return
}

// UnmarshalConfigObjectEnrollmentTls unmarshals an instance of ConfigObjectEnrollmentTls from the specified map of raw messages.
func UnmarshalConfigObjectEnrollmentTls(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ConfigObjectEnrollmentTls)
	err = core.UnmarshalPrimitive(m, "cahost", &obj.Cahost)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "caport", &obj.Caport)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "caname", &obj.Caname)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "catls", &obj.Catls, UnmarshalConfigObjectEnrollmentTlsCatls)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "enrollid", &obj.Enrollid)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "enrollsecret", &obj.Enrollsecret)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "admincerts", &obj.Admincerts)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "csr", &obj.Csr, UnmarshalConfigObjectEnrollmentTlsCsr)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ConfigObjectEnrollmentTlsCatls : ConfigObjectEnrollmentTlsCatls struct
type ConfigObjectEnrollmentTlsCatls struct {
	// The TLS certificate as base 64 encoded PEM. Certificate is used to secure/validate a TLS connection.
	Cacert *string `json:"cacert" validate:"required"`
}


// NewConfigObjectEnrollmentTlsCatls : Instantiate ConfigObjectEnrollmentTlsCatls (Generic Model Constructor)
func (*BlockchainV2) NewConfigObjectEnrollmentTlsCatls(cacert string) (model *ConfigObjectEnrollmentTlsCatls, err error) {
	model = &ConfigObjectEnrollmentTlsCatls{
		Cacert: core.StringPtr(cacert),
	}
	err = core.ValidateStruct(model, "required parameters")
	return
}

// UnmarshalConfigObjectEnrollmentTlsCatls unmarshals an instance of ConfigObjectEnrollmentTlsCatls from the specified map of raw messages.
func UnmarshalConfigObjectEnrollmentTlsCatls(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ConfigObjectEnrollmentTlsCatls)
	err = core.UnmarshalPrimitive(m, "cacert", &obj.Cacert)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ConfigObjectEnrollmentTlsCsr : ConfigObjectEnrollmentTlsCsr struct
type ConfigObjectEnrollmentTlsCsr struct {
	Hosts []string `json:"hosts" validate:"required"`
}


// NewConfigObjectEnrollmentTlsCsr : Instantiate ConfigObjectEnrollmentTlsCsr (Generic Model Constructor)
func (*BlockchainV2) NewConfigObjectEnrollmentTlsCsr(hosts []string) (model *ConfigObjectEnrollmentTlsCsr, err error) {
	model = &ConfigObjectEnrollmentTlsCsr{
		Hosts: hosts,
	}
	err = core.ValidateStruct(model, "required parameters")
	return
}

// UnmarshalConfigObjectEnrollmentTlsCsr unmarshals an instance of ConfigObjectEnrollmentTlsCsr from the specified map of raw messages.
func UnmarshalConfigObjectEnrollmentTlsCsr(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ConfigObjectEnrollmentTlsCsr)
	err = core.UnmarshalPrimitive(m, "hosts", &obj.Hosts)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ConfigObjectMsp : The `msp` field contains data to allow a component to configure its MSP with an already enrolled identity. Use `msp`
// or `enrollment`, not both.
type ConfigObjectMsp struct {
	Component *MspConfigData `json:"component" validate:"required"`

	Tls *MspConfigData `json:"tls" validate:"required"`

	Clientauth *MspConfigData `json:"clientauth,omitempty"`
}


// NewConfigObjectMsp : Instantiate ConfigObjectMsp (Generic Model Constructor)
func (*BlockchainV2) NewConfigObjectMsp(component *MspConfigData, tls *MspConfigData) (model *ConfigObjectMsp, err error) {
	model = &ConfigObjectMsp{
		Component: component,
		Tls: tls,
	}
	err = core.ValidateStruct(model, "required parameters")
	return
}

// UnmarshalConfigObjectMsp unmarshals an instance of ConfigObjectMsp from the specified map of raw messages.
func UnmarshalConfigObjectMsp(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ConfigObjectMsp)
	err = core.UnmarshalModel(m, "component", &obj.Component, UnmarshalMspConfigData)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "tls", &obj.Tls, UnmarshalMspConfigData)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "clientauth", &obj.Clientauth, UnmarshalMspConfigData)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ConfigOrdererCreate : Override the [Fabric Orderer configuration
// file](https://github.com/hyperledger/fabric/blob/release-1.4/sampleconfig/orderer.yaml) if you want use custom
// attributes to configure the Orderer. Omit if not.
//
// *The field **names** below are not case-sensitive.*.
type ConfigOrdererCreate struct {
	General *ConfigOrdererGeneral `json:"General,omitempty"`

	// Controls the debugging options for the orderer.
	Debug *ConfigOrdererDebug `json:"Debug,omitempty"`

	Metrics *ConfigOrdererMetrics `json:"Metrics,omitempty"`
}


// UnmarshalConfigOrdererCreate unmarshals an instance of ConfigOrdererCreate from the specified map of raw messages.
func UnmarshalConfigOrdererCreate(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ConfigOrdererCreate)
	err = core.UnmarshalModel(m, "General", &obj.General, UnmarshalConfigOrdererGeneral)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "Debug", &obj.Debug, UnmarshalConfigOrdererDebug)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "Metrics", &obj.Metrics, UnmarshalConfigOrdererMetrics)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ConfigOrdererMetricsStatsd : The statsd configuration.
type ConfigOrdererMetricsStatsd struct {
	// Network protocol to use.
	Network *string `json:"Network,omitempty"`

	// The address of the statsd server. Include hostname/ip and port.
	Address *string `json:"Address,omitempty"`

	// The frequency at which locally cached counters and gauges are pushed to statsd.
	WriteInterval *string `json:"WriteInterval,omitempty"`

	// The string that is prepended to all emitted statsd metrics.
	Prefix *string `json:"Prefix,omitempty"`
}

// Constants associated with the ConfigOrdererMetricsStatsd.Network property.
// Network protocol to use.
const (
	ConfigOrdererMetricsStatsd_Network_Tcp = "tcp"
	ConfigOrdererMetricsStatsd_Network_Udp = "udp"
)


// UnmarshalConfigOrdererMetricsStatsd unmarshals an instance of ConfigOrdererMetricsStatsd from the specified map of raw messages.
func UnmarshalConfigOrdererMetricsStatsd(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ConfigOrdererMetricsStatsd)
	err = core.UnmarshalPrimitive(m, "Network", &obj.Network)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "Address", &obj.Address)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "WriteInterval", &obj.WriteInterval)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "Prefix", &obj.Prefix)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ConfigOrdererUpdate : Update the [Fabric Orderer configuration
// file](https://github.com/hyperledger/fabric/blob/release-1.4/sampleconfig/orderer.yaml) if you want use custom
// attributes to configure the Orderer. Omit if not.
//
// *The field **names** below are not case-sensitive.*.
type ConfigOrdererUpdate struct {
	General *ConfigOrdererGeneralUpdate `json:"General,omitempty"`

	// Controls the debugging options for the orderer.
	Debug *ConfigOrdererDebug `json:"Debug,omitempty"`

	Metrics *ConfigOrdererMetrics `json:"Metrics,omitempty"`
}


// UnmarshalConfigOrdererUpdate unmarshals an instance of ConfigOrdererUpdate from the specified map of raw messages.
func UnmarshalConfigOrdererUpdate(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ConfigOrdererUpdate)
	err = core.UnmarshalModel(m, "General", &obj.General, UnmarshalConfigOrdererGeneralUpdate)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "Debug", &obj.Debug, UnmarshalConfigOrdererDebug)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "Metrics", &obj.Metrics, UnmarshalConfigOrdererMetrics)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ConfigOrdererAuthentication : Contains configuration parameters that are related to authenticating client messages.
type ConfigOrdererAuthentication struct {
	// The maximum acceptable difference between the current server time and the client's time.
	TimeWindow *string `json:"TimeWindow,omitempty"`

	// Indicates if the orderer should ignore expired identities. Should only be used temporarily to recover from an
	// extreme event such as the expiration of administrators. Defaults `false`.
	NoExpirationChecks *bool `json:"NoExpirationChecks,omitempty"`
}


// UnmarshalConfigOrdererAuthentication unmarshals an instance of ConfigOrdererAuthentication from the specified map of raw messages.
func UnmarshalConfigOrdererAuthentication(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ConfigOrdererAuthentication)
	err = core.UnmarshalPrimitive(m, "TimeWindow", &obj.TimeWindow)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "NoExpirationChecks", &obj.NoExpirationChecks)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ConfigOrdererDebug : Controls the debugging options for the orderer.
type ConfigOrdererDebug struct {
	// Path to directory. If set will cause each request to the Broadcast service to be written to a file in this
	// directory.
	BroadcastTraceDir *string `json:"BroadcastTraceDir,omitempty"`

	// Path to directory. If set will cause each request to the Deliver service to be written to a file in this directory.
	DeliverTraceDir *string `json:"DeliverTraceDir,omitempty"`
}


// UnmarshalConfigOrdererDebug unmarshals an instance of ConfigOrdererDebug from the specified map of raw messages.
func UnmarshalConfigOrdererDebug(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ConfigOrdererDebug)
	err = core.UnmarshalPrimitive(m, "BroadcastTraceDir", &obj.BroadcastTraceDir)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "DeliverTraceDir", &obj.DeliverTraceDir)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ConfigOrdererGeneral : ConfigOrdererGeneral struct
type ConfigOrdererGeneral struct {
	// Keep alive settings for the GRPC server.
	Keepalive *ConfigOrdererKeepalive `json:"Keepalive,omitempty"`

	// Configures the Blockchain Crypto Service Providers (bccsp).
	BCCSP *Bccsp `json:"BCCSP,omitempty"`

	// Contains configuration parameters that are related to authenticating client messages.
	Authentication *ConfigOrdererAuthentication `json:"Authentication,omitempty"`
}


// UnmarshalConfigOrdererGeneral unmarshals an instance of ConfigOrdererGeneral from the specified map of raw messages.
func UnmarshalConfigOrdererGeneral(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ConfigOrdererGeneral)
	err = core.UnmarshalModel(m, "Keepalive", &obj.Keepalive, UnmarshalConfigOrdererKeepalive)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "BCCSP", &obj.BCCSP, UnmarshalBccsp)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "Authentication", &obj.Authentication, UnmarshalConfigOrdererAuthentication)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ConfigOrdererGeneralUpdate : ConfigOrdererGeneralUpdate struct
type ConfigOrdererGeneralUpdate struct {
	// Keep alive settings for the GRPC server.
	Keepalive *ConfigOrdererKeepalive `json:"Keepalive,omitempty"`

	// Contains configuration parameters that are related to authenticating client messages.
	Authentication *ConfigOrdererAuthentication `json:"Authentication,omitempty"`
}


// UnmarshalConfigOrdererGeneralUpdate unmarshals an instance of ConfigOrdererGeneralUpdate from the specified map of raw messages.
func UnmarshalConfigOrdererGeneralUpdate(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ConfigOrdererGeneralUpdate)
	err = core.UnmarshalModel(m, "Keepalive", &obj.Keepalive, UnmarshalConfigOrdererKeepalive)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "Authentication", &obj.Authentication, UnmarshalConfigOrdererAuthentication)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ConfigOrdererKeepalive : Keep alive settings for the GRPC server.
type ConfigOrdererKeepalive struct {
	// The minimum time between client pings. If a client sends pings more frequently the server will disconnect from the
	// client.
	ServerMinInterval *string `json:"ServerMinInterval,omitempty"`

	// The time between pings to clients.
	ServerInterval *string `json:"ServerInterval,omitempty"`

	// The duration the server will wait for a response from a client before closing the connection.
	ServerTimeout *string `json:"ServerTimeout,omitempty"`
}


// UnmarshalConfigOrdererKeepalive unmarshals an instance of ConfigOrdererKeepalive from the specified map of raw messages.
func UnmarshalConfigOrdererKeepalive(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ConfigOrdererKeepalive)
	err = core.UnmarshalPrimitive(m, "ServerMinInterval", &obj.ServerMinInterval)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "ServerInterval", &obj.ServerInterval)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "ServerTimeout", &obj.ServerTimeout)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ConfigOrdererMetrics : ConfigOrdererMetrics struct
type ConfigOrdererMetrics struct {
	// The metrics provider to use.
	Provider *string `json:"Provider,omitempty"`

	// The statsd configuration.
	Statsd *ConfigOrdererMetricsStatsd `json:"Statsd,omitempty"`
}

// Constants associated with the ConfigOrdererMetrics.Provider property.
// The metrics provider to use.
const (
	ConfigOrdererMetrics_Provider_Disabled = "disabled"
	ConfigOrdererMetrics_Provider_Prometheus = "prometheus"
	ConfigOrdererMetrics_Provider_Statsd = "statsd"
)


// UnmarshalConfigOrdererMetrics unmarshals an instance of ConfigOrdererMetrics from the specified map of raw messages.
func UnmarshalConfigOrdererMetrics(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ConfigOrdererMetrics)
	err = core.UnmarshalPrimitive(m, "Provider", &obj.Provider)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "Statsd", &obj.Statsd, UnmarshalConfigOrdererMetricsStatsd)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ConfigPeerChaincodeExternalBuildersItem : ConfigPeerChaincodeExternalBuildersItem struct
type ConfigPeerChaincodeExternalBuildersItem struct {
	// The path to a build directory.
	Path *string `json:"path,omitempty"`

	// The name of this builder.
	Name *string `json:"name,omitempty"`

	EnvironmentWhitelist []string `json:"environmentWhitelist,omitempty"`
}


// UnmarshalConfigPeerChaincodeExternalBuildersItem unmarshals an instance of ConfigPeerChaincodeExternalBuildersItem from the specified map of raw messages.
func UnmarshalConfigPeerChaincodeExternalBuildersItem(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ConfigPeerChaincodeExternalBuildersItem)
	err = core.UnmarshalPrimitive(m, "path", &obj.Path)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "environmentWhitelist", &obj.EnvironmentWhitelist)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ConfigPeerChaincodeGolang : ConfigPeerChaincodeGolang struct
type ConfigPeerChaincodeGolang struct {
	// Controls if golang chaincode should be built with dynamic linking or static linking. Defaults `false` (static).
	DynamicLink *bool `json:"dynamicLink,omitempty"`
}


// UnmarshalConfigPeerChaincodeGolang unmarshals an instance of ConfigPeerChaincodeGolang from the specified map of raw messages.
func UnmarshalConfigPeerChaincodeGolang(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ConfigPeerChaincodeGolang)
	err = core.UnmarshalPrimitive(m, "dynamicLink", &obj.DynamicLink)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ConfigPeerChaincodeLogging : ConfigPeerChaincodeLogging struct
type ConfigPeerChaincodeLogging struct {
	// Default logging level for loggers within chaincode containers.
	Level *string `json:"level,omitempty"`

	// Override default level for the 'shim' logger.
	Shim *string `json:"shim,omitempty"`

	// Override the default log format for chaincode container logs.
	Format *string `json:"format,omitempty"`
}

// Constants associated with the ConfigPeerChaincodeLogging.Level property.
// Default logging level for loggers within chaincode containers.
const (
	ConfigPeerChaincodeLogging_Level_Debug = "debug"
	ConfigPeerChaincodeLogging_Level_Error = "error"
	ConfigPeerChaincodeLogging_Level_Fatal = "fatal"
	ConfigPeerChaincodeLogging_Level_Info = "info"
	ConfigPeerChaincodeLogging_Level_Panic = "panic"
	ConfigPeerChaincodeLogging_Level_Warning = "warning"
)

// Constants associated with the ConfigPeerChaincodeLogging.Shim property.
// Override default level for the 'shim' logger.
const (
	ConfigPeerChaincodeLogging_Shim_Debug = "debug"
	ConfigPeerChaincodeLogging_Shim_Error = "error"
	ConfigPeerChaincodeLogging_Shim_Fatal = "fatal"
	ConfigPeerChaincodeLogging_Shim_Info = "info"
	ConfigPeerChaincodeLogging_Shim_Panic = "panic"
	ConfigPeerChaincodeLogging_Shim_Warning = "warning"
)


// UnmarshalConfigPeerChaincodeLogging unmarshals an instance of ConfigPeerChaincodeLogging from the specified map of raw messages.
func UnmarshalConfigPeerChaincodeLogging(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ConfigPeerChaincodeLogging)
	err = core.UnmarshalPrimitive(m, "level", &obj.Level)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "shim", &obj.Shim)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "format", &obj.Format)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ConfigPeerChaincodeSystem : The complete whitelist for system chaincodes. To append a new chaincode add the new id to the default list.
type ConfigPeerChaincodeSystem struct {
	// Adds the system chaincode `cscc` to the whitelist.
	Cscc *bool `json:"cscc,omitempty"`

	// Adds the system chaincode `lscc` to the whitelist.
	Lscc *bool `json:"lscc,omitempty"`

	// Adds the system chaincode `escc` to the whitelist.
	Escc *bool `json:"escc,omitempty"`

	// Adds the system chaincode `vscc` to the whitelist.
	Vscc *bool `json:"vscc,omitempty"`

	// Adds the system chaincode `qscc` to the whitelist.
	Qscc *bool `json:"qscc,omitempty"`
}


// UnmarshalConfigPeerChaincodeSystem unmarshals an instance of ConfigPeerChaincodeSystem from the specified map of raw messages.
func UnmarshalConfigPeerChaincodeSystem(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ConfigPeerChaincodeSystem)
	err = core.UnmarshalPrimitive(m, "cscc", &obj.Cscc)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "lscc", &obj.Lscc)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "escc", &obj.Escc)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "vscc", &obj.Vscc)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "qscc", &obj.Qscc)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ConfigPeerCreate : Override the [Fabric Peer configuration
// file](https://github.com/hyperledger/fabric/blob/release-1.4/sampleconfig/core.yaml) if you want use custom
// attributes to configure the Peer. Omit if not.
//
// *The field **names** below are not case-sensitive.*.
type ConfigPeerCreate struct {
	Peer *ConfigPeerCreatePeer `json:"peer,omitempty"`

	Chaincode *ConfigPeerChaincode `json:"chaincode,omitempty"`

	Metrics *Metrics `json:"metrics,omitempty"`
}


// UnmarshalConfigPeerCreate unmarshals an instance of ConfigPeerCreate from the specified map of raw messages.
func UnmarshalConfigPeerCreate(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ConfigPeerCreate)
	err = core.UnmarshalModel(m, "peer", &obj.Peer, UnmarshalConfigPeerCreatePeer)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "chaincode", &obj.Chaincode, UnmarshalConfigPeerChaincode)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "metrics", &obj.Metrics, UnmarshalMetrics)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ConfigPeerCreatePeer : ConfigPeerCreatePeer struct
type ConfigPeerCreatePeer struct {
	// A unique id used to identify this instance.
	ID *string `json:"id,omitempty"`

	// The ID to logically separate one network from another.
	NetworkID *string `json:"networkId,omitempty"`

	// Keep alive settings between the peer server and clients.
	Keepalive *ConfigPeerKeepalive `json:"keepalive,omitempty"`

	Gossip *ConfigPeerGossip `json:"gossip,omitempty"`

	Authentication *ConfigPeerAuthentication `json:"authentication,omitempty"`

	// Configures the Blockchain Crypto Service Providers (bccsp).
	BCCSP *Bccsp `json:"BCCSP,omitempty"`

	Client *ConfigPeerClient `json:"client,omitempty"`

	Deliveryclient *ConfigPeerDeliveryclient `json:"deliveryclient,omitempty"`

	// Used for administrative operations such as control over logger levels. Only peer administrators can use the service.
	AdminService *ConfigPeerAdminService `json:"adminService,omitempty"`

	// Number of go-routines that will execute transaction validation in parallel. By default, the peer chooses the number
	// of CPUs on the machine. It is recommended to use the default values and not set this field.
	ValidatorPoolSize *float64 `json:"validatorPoolSize,omitempty"`

	// The discovery service is used by clients to query information about peers. Such as - which peers have joined a
	// channel, what is the latest channel config, and what possible sets of peers satisfy the endorsement policy (given a
	// smart contract and a channel).
	Discovery *ConfigPeerDiscovery `json:"discovery,omitempty"`

	Limits *ConfigPeerLimits `json:"limits,omitempty"`
}


// UnmarshalConfigPeerCreatePeer unmarshals an instance of ConfigPeerCreatePeer from the specified map of raw messages.
func UnmarshalConfigPeerCreatePeer(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ConfigPeerCreatePeer)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "networkId", &obj.NetworkID)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "keepalive", &obj.Keepalive, UnmarshalConfigPeerKeepalive)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "gossip", &obj.Gossip, UnmarshalConfigPeerGossip)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "authentication", &obj.Authentication, UnmarshalConfigPeerAuthentication)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "BCCSP", &obj.BCCSP, UnmarshalBccsp)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "client", &obj.Client, UnmarshalConfigPeerClient)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "deliveryclient", &obj.Deliveryclient, UnmarshalConfigPeerDeliveryclient)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "adminService", &obj.AdminService, UnmarshalConfigPeerAdminService)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "validatorPoolSize", &obj.ValidatorPoolSize)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "discovery", &obj.Discovery, UnmarshalConfigPeerDiscovery)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "limits", &obj.Limits, UnmarshalConfigPeerLimits)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ConfigPeerDeliveryclientAddressOverridesItem : ConfigPeerDeliveryclientAddressOverridesItem struct
type ConfigPeerDeliveryclientAddressOverridesItem struct {
	// The address in the channel configuration that will be overridden.
	From *string `json:"from,omitempty"`

	// The address to use.
	To *string `json:"to,omitempty"`

	// The path to the CA's cert file.
	CaCertsFile *string `json:"caCertsFile,omitempty"`
}


// UnmarshalConfigPeerDeliveryclientAddressOverridesItem unmarshals an instance of ConfigPeerDeliveryclientAddressOverridesItem from the specified map of raw messages.
func UnmarshalConfigPeerDeliveryclientAddressOverridesItem(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ConfigPeerDeliveryclientAddressOverridesItem)
	err = core.UnmarshalPrimitive(m, "from", &obj.From)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "to", &obj.To)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "caCertsFile", &obj.CaCertsFile)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ConfigPeerGossipElection : Leader election service configuration.
type ConfigPeerGossipElection struct {
	// Longest time the peer will wait for stable membership during leader election startup.
	StartupGracePeriod *string `json:"startupGracePeriod,omitempty"`

	// Frequency that gossip membership samples to check its stability.
	MembershipSampleInterval *string `json:"membershipSampleInterval,omitempty"`

	// Amount of time after the last declaration message for the peer to perform another leader election.
	LeaderAliveThreshold *string `json:"leaderAliveThreshold,omitempty"`

	// Amount of time between the peer sending a propose message and it declaring itself as a leader.
	LeaderElectionDuration *string `json:"leaderElectionDuration,omitempty"`
}


// UnmarshalConfigPeerGossipElection unmarshals an instance of ConfigPeerGossipElection from the specified map of raw messages.
func UnmarshalConfigPeerGossipElection(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ConfigPeerGossipElection)
	err = core.UnmarshalPrimitive(m, "startupGracePeriod", &obj.StartupGracePeriod)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "membershipSampleInterval", &obj.MembershipSampleInterval)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "leaderAliveThreshold", &obj.LeaderAliveThreshold)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "leaderElectionDuration", &obj.LeaderElectionDuration)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ConfigPeerGossipPvtData : ConfigPeerGossipPvtData struct
type ConfigPeerGossipPvtData struct {
	// Determines the maximum time to attempt to pull private data for a block before that block is committed without the
	// private data.
	PullRetryThreshold *string `json:"pullRetryThreshold,omitempty"`

	// As private data enters the transient store, it is associated with the peer's current ledger's height. This field
	// defines the maximum difference between the current ledger's height on commit, and the private data residing inside
	// the transient store. Private data outside this range is not guaranteed to exist and will be purged periodically.
	TransientstoreMaxBlockRetention *float64 `json:"transientstoreMaxBlockRetention,omitempty"`

	// Maximum time to wait for an acknowledgment from each peer's private data push.
	PushAckTimeout *string `json:"pushAckTimeout,omitempty"`

	// Block to live pulling margin. Used as a buffer to prevent peers from trying to pull private data from others peers
	// that are soon to be purged. "Soon" defined as blocks that will be purged in the next N blocks. This helps a newly
	// joined peer catch up quicker.
	BtlPullMargin *float64 `json:"btlPullMargin,omitempty"`

	// Determines the maximum batch size of missing private data that will be reconciled in a single iteration. The process
	// of reconciliation is done in an endless loop. The "reconciler" in each iteration tries to pull from the other peers
	// with the most recent missing blocks and this maximum batch size limitation.
	ReconcileBatchSize *float64 `json:"reconcileBatchSize,omitempty"`

	// Determines the time "reconciler" sleeps from the end of an iteration until the beginning of the next iteration.
	ReconcileSleepInterval *string `json:"reconcileSleepInterval,omitempty"`

	// Determines whether private data reconciliation is enabled or not.
	ReconciliationEnabled *bool `json:"reconciliationEnabled,omitempty"`

	// Controls whether pulling invalid transaction's private data from other peers need to be skipped during the commit
	// time. If `true` it will be pulled through "reconciler".
	SkipPullingInvalidTransactionsDuringCommit *bool `json:"skipPullingInvalidTransactionsDuringCommit,omitempty"`

	ImplicitCollectionDisseminationPolicy *ConfigPeerGossipPvtDataImplicitCollectionDisseminationPolicy `json:"implicitCollectionDisseminationPolicy,omitempty"`
}


// UnmarshalConfigPeerGossipPvtData unmarshals an instance of ConfigPeerGossipPvtData from the specified map of raw messages.
func UnmarshalConfigPeerGossipPvtData(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ConfigPeerGossipPvtData)
	err = core.UnmarshalPrimitive(m, "pullRetryThreshold", &obj.PullRetryThreshold)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "transientstoreMaxBlockRetention", &obj.TransientstoreMaxBlockRetention)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "pushAckTimeout", &obj.PushAckTimeout)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "btlPullMargin", &obj.BtlPullMargin)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "reconcileBatchSize", &obj.ReconcileBatchSize)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "reconcileSleepInterval", &obj.ReconcileSleepInterval)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "reconciliationEnabled", &obj.ReconciliationEnabled)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "skipPullingInvalidTransactionsDuringCommit", &obj.SkipPullingInvalidTransactionsDuringCommit)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "implicitCollectionDisseminationPolicy", &obj.ImplicitCollectionDisseminationPolicy, UnmarshalConfigPeerGossipPvtDataImplicitCollectionDisseminationPolicy)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ConfigPeerGossipPvtDataImplicitCollectionDisseminationPolicy : ConfigPeerGossipPvtDataImplicitCollectionDisseminationPolicy struct
type ConfigPeerGossipPvtDataImplicitCollectionDisseminationPolicy struct {
	// Defines the minimum number of peers to successfully disseminate private data during endorsement.
	RequiredPeerCount *float64 `json:"requiredPeerCount,omitempty"`

	// Defines the maximum number of peers to attempt to disseminate private data during endorsement.
	MaxPeerCount *float64 `json:"maxPeerCount,omitempty"`
}


// UnmarshalConfigPeerGossipPvtDataImplicitCollectionDisseminationPolicy unmarshals an instance of ConfigPeerGossipPvtDataImplicitCollectionDisseminationPolicy from the specified map of raw messages.
func UnmarshalConfigPeerGossipPvtDataImplicitCollectionDisseminationPolicy(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ConfigPeerGossipPvtDataImplicitCollectionDisseminationPolicy)
	err = core.UnmarshalPrimitive(m, "requiredPeerCount", &obj.RequiredPeerCount)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "maxPeerCount", &obj.MaxPeerCount)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ConfigPeerGossipState : Gossip state transfer related configuration.
type ConfigPeerGossipState struct {
	// Controls if the state transfer is enabled or not. If state transfer is active, it syncs up missing blocks and allows
	// lagging peers to catch up with the rest of the network.
	Enabled *bool `json:"enabled,omitempty"`

	// The frequency to check whether a peer is lagging behind enough to request blocks by using state transfer from
	// another peer.
	CheckInterval *string `json:"checkInterval,omitempty"`

	// Amount of time to wait for state transfer responses from other peers.
	ResponseTimeout *string `json:"responseTimeout,omitempty"`

	// Number of blocks to request by using state transfer from another peer.
	BatchSize *float64 `json:"batchSize,omitempty"`

	// Maximum difference between the lowest and highest block sequence number. In order to ensure that there are no holes
	// the actual buffer size is twice this distance.
	BlockBufferSize *float64 `json:"blockBufferSize,omitempty"`

	// Maximum number of retries of a single state transfer request.
	MaxRetries *float64 `json:"maxRetries,omitempty"`
}


// UnmarshalConfigPeerGossipState unmarshals an instance of ConfigPeerGossipState from the specified map of raw messages.
func UnmarshalConfigPeerGossipState(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ConfigPeerGossipState)
	err = core.UnmarshalPrimitive(m, "enabled", &obj.Enabled)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "checkInterval", &obj.CheckInterval)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "responseTimeout", &obj.ResponseTimeout)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "batchSize", &obj.BatchSize)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "blockBufferSize", &obj.BlockBufferSize)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "maxRetries", &obj.MaxRetries)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ConfigPeerKeepaliveClient : ConfigPeerKeepaliveClient struct
type ConfigPeerKeepaliveClient struct {
	// The time between pings to other peer nodes. Must greater than or equal to the minInterval.
	Interval *string `json:"interval,omitempty"`

	// The duration a client waits for a peer's response before it closes the connection.
	Timeout *string `json:"timeout,omitempty"`
}


// UnmarshalConfigPeerKeepaliveClient unmarshals an instance of ConfigPeerKeepaliveClient from the specified map of raw messages.
func UnmarshalConfigPeerKeepaliveClient(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ConfigPeerKeepaliveClient)
	err = core.UnmarshalPrimitive(m, "interval", &obj.Interval)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "timeout", &obj.Timeout)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ConfigPeerKeepaliveDeliveryClient : ConfigPeerKeepaliveDeliveryClient struct
type ConfigPeerKeepaliveDeliveryClient struct {
	// The time between pings to ordering nodes. Must greater than or equal to the minInterval.
	Interval *string `json:"interval,omitempty"`

	// The duration a client waits for an orderer's response before it closes the connection.
	Timeout *string `json:"timeout,omitempty"`
}


// UnmarshalConfigPeerKeepaliveDeliveryClient unmarshals an instance of ConfigPeerKeepaliveDeliveryClient from the specified map of raw messages.
func UnmarshalConfigPeerKeepaliveDeliveryClient(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ConfigPeerKeepaliveDeliveryClient)
	err = core.UnmarshalPrimitive(m, "interval", &obj.Interval)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "timeout", &obj.Timeout)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ConfigPeerLimitsConcurrency : ConfigPeerLimitsConcurrency struct
type ConfigPeerLimitsConcurrency struct {
	// Limits the number of concurrent requests to the endorser service. The endorser service handles application and
	// system chaincode deployment and invocations (including queries).
	EndorserService interface{} `json:"endorserService,omitempty"`

	// Limits the number of concurrent requests to the deliver service. The deliver service handles block and transaction
	// events.
	DeliverService interface{} `json:"deliverService,omitempty"`
}


// UnmarshalConfigPeerLimitsConcurrency unmarshals an instance of ConfigPeerLimitsConcurrency from the specified map of raw messages.
func UnmarshalConfigPeerLimitsConcurrency(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ConfigPeerLimitsConcurrency)
	err = core.UnmarshalPrimitive(m, "endorserService", &obj.EndorserService)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "deliverService", &obj.DeliverService)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ConfigPeerUpdate : Update the [Fabric Peer configuration
// file](https://github.com/hyperledger/fabric/blob/release-1.4/sampleconfig/core.yaml) if you want use custom
// attributes to configure the Peer. Omit if not.
//
// *The field **names** below are not case-sensitive.*.
type ConfigPeerUpdate struct {
	Peer *ConfigPeerUpdatePeer `json:"peer,omitempty"`

	Chaincode *ConfigPeerChaincode `json:"chaincode,omitempty"`

	Metrics *Metrics `json:"metrics,omitempty"`
}


// UnmarshalConfigPeerUpdate unmarshals an instance of ConfigPeerUpdate from the specified map of raw messages.
func UnmarshalConfigPeerUpdate(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ConfigPeerUpdate)
	err = core.UnmarshalModel(m, "peer", &obj.Peer, UnmarshalConfigPeerUpdatePeer)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "chaincode", &obj.Chaincode, UnmarshalConfigPeerChaincode)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "metrics", &obj.Metrics, UnmarshalMetrics)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ConfigPeerUpdatePeer : ConfigPeerUpdatePeer struct
type ConfigPeerUpdatePeer struct {
	// A unique id used to identify this instance.
	ID *string `json:"id,omitempty"`

	// The ID to logically separate one network from another.
	NetworkID *string `json:"networkId,omitempty"`

	// Keep alive settings between the peer server and clients.
	Keepalive *ConfigPeerKeepalive `json:"keepalive,omitempty"`

	Gossip *ConfigPeerGossip `json:"gossip,omitempty"`

	Authentication *ConfigPeerAuthentication `json:"authentication,omitempty"`

	Client *ConfigPeerClient `json:"client,omitempty"`

	Deliveryclient *ConfigPeerDeliveryclient `json:"deliveryclient,omitempty"`

	// Used for administrative operations such as control over logger levels. Only peer administrators can use the service.
	AdminService *ConfigPeerAdminService `json:"adminService,omitempty"`

	// Number of go-routines that will execute transaction validation in parallel. By default, the peer chooses the number
	// of CPUs on the machine. It is recommended to use the default values and not set this field.
	ValidatorPoolSize *float64 `json:"validatorPoolSize,omitempty"`

	// The discovery service is used by clients to query information about peers. Such as - which peers have joined a
	// channel, what is the latest channel config, and what possible sets of peers satisfy the endorsement policy (given a
	// smart contract and a channel).
	Discovery *ConfigPeerDiscovery `json:"discovery,omitempty"`

	Limits *ConfigPeerLimits `json:"limits,omitempty"`
}


// UnmarshalConfigPeerUpdatePeer unmarshals an instance of ConfigPeerUpdatePeer from the specified map of raw messages.
func UnmarshalConfigPeerUpdatePeer(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ConfigPeerUpdatePeer)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "networkId", &obj.NetworkID)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "keepalive", &obj.Keepalive, UnmarshalConfigPeerKeepalive)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "gossip", &obj.Gossip, UnmarshalConfigPeerGossip)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "authentication", &obj.Authentication, UnmarshalConfigPeerAuthentication)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "client", &obj.Client, UnmarshalConfigPeerClient)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "deliveryclient", &obj.Deliveryclient, UnmarshalConfigPeerDeliveryclient)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "adminService", &obj.AdminService, UnmarshalConfigPeerAdminService)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "validatorPoolSize", &obj.ValidatorPoolSize)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "discovery", &obj.Discovery, UnmarshalConfigPeerDiscovery)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "limits", &obj.Limits, UnmarshalConfigPeerLimits)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ConfigPeerAdminService : Used for administrative operations such as control over logger levels. Only peer administrators can use the service.
type ConfigPeerAdminService struct {
	// The interface and port on which the admin server will listen on. Defaults to the same address as the peer's listen
	// address and port 7051.
	ListenAddress *string `json:"listenAddress" validate:"required"`
}


// NewConfigPeerAdminService : Instantiate ConfigPeerAdminService (Generic Model Constructor)
func (*BlockchainV2) NewConfigPeerAdminService(listenAddress string) (model *ConfigPeerAdminService, err error) {
	model = &ConfigPeerAdminService{
		ListenAddress: core.StringPtr(listenAddress),
	}
	err = core.ValidateStruct(model, "required parameters")
	return
}

// UnmarshalConfigPeerAdminService unmarshals an instance of ConfigPeerAdminService from the specified map of raw messages.
func UnmarshalConfigPeerAdminService(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ConfigPeerAdminService)
	err = core.UnmarshalPrimitive(m, "listenAddress", &obj.ListenAddress)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ConfigPeerAuthentication : ConfigPeerAuthentication struct
type ConfigPeerAuthentication struct {
	// The maximum acceptable difference between the current server time and the client's time.
	Timewindow *string `json:"timewindow" validate:"required"`
}


// NewConfigPeerAuthentication : Instantiate ConfigPeerAuthentication (Generic Model Constructor)
func (*BlockchainV2) NewConfigPeerAuthentication(timewindow string) (model *ConfigPeerAuthentication, err error) {
	model = &ConfigPeerAuthentication{
		Timewindow: core.StringPtr(timewindow),
	}
	err = core.ValidateStruct(model, "required parameters")
	return
}

// UnmarshalConfigPeerAuthentication unmarshals an instance of ConfigPeerAuthentication from the specified map of raw messages.
func UnmarshalConfigPeerAuthentication(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ConfigPeerAuthentication)
	err = core.UnmarshalPrimitive(m, "timewindow", &obj.Timewindow)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ConfigPeerChaincode : ConfigPeerChaincode struct
type ConfigPeerChaincode struct {
	Golang *ConfigPeerChaincodeGolang `json:"golang,omitempty"`

	// List of directories to treat as external builders/launches of chaincode.
	ExternalBuilders []ConfigPeerChaincodeExternalBuildersItem `json:"externalBuilders,omitempty"`

	// Maximum duration to wait for the chaincode build and install process to complete.
	InstallTimeout *string `json:"installTimeout,omitempty"`

	// Time for starting up a container and waiting for Register to come through.
	Startuptimeout *string `json:"startuptimeout,omitempty"`

	// Time for Invoke and Init calls to return. This timeout is used by all chaincodes in all the channels, including
	// system chaincodes. Note that if the image is not available the peer needs to build the image, which will take
	// additional time.
	Executetimeout *string `json:"executetimeout,omitempty"`

	// The complete whitelist for system chaincodes. To append a new chaincode add the new id to the default list.
	System *ConfigPeerChaincodeSystem `json:"system,omitempty"`

	Logging *ConfigPeerChaincodeLogging `json:"logging,omitempty"`
}


// UnmarshalConfigPeerChaincode unmarshals an instance of ConfigPeerChaincode from the specified map of raw messages.
func UnmarshalConfigPeerChaincode(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ConfigPeerChaincode)
	err = core.UnmarshalModel(m, "golang", &obj.Golang, UnmarshalConfigPeerChaincodeGolang)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "externalBuilders", &obj.ExternalBuilders, UnmarshalConfigPeerChaincodeExternalBuildersItem)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "installTimeout", &obj.InstallTimeout)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "startuptimeout", &obj.Startuptimeout)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "executetimeout", &obj.Executetimeout)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "system", &obj.System, UnmarshalConfigPeerChaincodeSystem)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "logging", &obj.Logging, UnmarshalConfigPeerChaincodeLogging)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ConfigPeerClient : ConfigPeerClient struct
type ConfigPeerClient struct {
	// The timeout for a network connection.
	ConnTimeout *string `json:"connTimeout" validate:"required"`
}


// NewConfigPeerClient : Instantiate ConfigPeerClient (Generic Model Constructor)
func (*BlockchainV2) NewConfigPeerClient(connTimeout string) (model *ConfigPeerClient, err error) {
	model = &ConfigPeerClient{
		ConnTimeout: core.StringPtr(connTimeout),
	}
	err = core.ValidateStruct(model, "required parameters")
	return
}

// UnmarshalConfigPeerClient unmarshals an instance of ConfigPeerClient from the specified map of raw messages.
func UnmarshalConfigPeerClient(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ConfigPeerClient)
	err = core.UnmarshalPrimitive(m, "connTimeout", &obj.ConnTimeout)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ConfigPeerDeliveryclient : ConfigPeerDeliveryclient struct
type ConfigPeerDeliveryclient struct {
	// Total time to spend retrying connections to ordering nodes before giving up and returning an error.
	ReconnectTotalTimeThreshold *string `json:"reconnectTotalTimeThreshold,omitempty"`

	// The timeout for a network connection.
	ConnTimeout *string `json:"connTimeout,omitempty"`

	// Maximum delay between consecutive connection retry attempts to ordering nodes.
	ReConnectBackoffThreshold *string `json:"reConnectBackoffThreshold,omitempty"`

	// A list of orderer endpoint addresses in channel configurations that should be overridden. Typically used when the
	// original orderer addresses no longer exist.
	AddressOverrides []ConfigPeerDeliveryclientAddressOverridesItem `json:"addressOverrides,omitempty"`
}


// UnmarshalConfigPeerDeliveryclient unmarshals an instance of ConfigPeerDeliveryclient from the specified map of raw messages.
func UnmarshalConfigPeerDeliveryclient(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ConfigPeerDeliveryclient)
	err = core.UnmarshalPrimitive(m, "reconnectTotalTimeThreshold", &obj.ReconnectTotalTimeThreshold)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "connTimeout", &obj.ConnTimeout)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "reConnectBackoffThreshold", &obj.ReConnectBackoffThreshold)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "addressOverrides", &obj.AddressOverrides, UnmarshalConfigPeerDeliveryclientAddressOverridesItem)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ConfigPeerDiscovery : The discovery service is used by clients to query information about peers. Such as - which peers have joined a
// channel, what is the latest channel config, and what possible sets of peers satisfy the endorsement policy (given a
// smart contract and a channel).
type ConfigPeerDiscovery struct {
	// Determines whether the discover service is available or not.
	Enabled *bool `json:"enabled,omitempty"`

	// Determines whether the authentication cache is enabled or not.
	AuthCacheEnabled *bool `json:"authCacheEnabled,omitempty"`

	// Maximum size of the cache. If exceeded a purge takes place.
	AuthCacheMaxSize *float64 `json:"authCacheMaxSize,omitempty"`

	// The proportion (0 - 1) of entries that remain in the cache after the cache is purged due to overpopulation.
	AuthCachePurgeRetentionRatio *float64 `json:"authCachePurgeRetentionRatio,omitempty"`

	// Whether to allow non-admins to perform non-channel scoped queries. When `false`, it means that only peer admins can
	// perform non-channel scoped queries.
	OrgMembersAllowedAccess *bool `json:"orgMembersAllowedAccess,omitempty"`
}


// UnmarshalConfigPeerDiscovery unmarshals an instance of ConfigPeerDiscovery from the specified map of raw messages.
func UnmarshalConfigPeerDiscovery(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ConfigPeerDiscovery)
	err = core.UnmarshalPrimitive(m, "enabled", &obj.Enabled)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "authCacheEnabled", &obj.AuthCacheEnabled)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "authCacheMaxSize", &obj.AuthCacheMaxSize)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "authCachePurgeRetentionRatio", &obj.AuthCachePurgeRetentionRatio)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "orgMembersAllowedAccess", &obj.OrgMembersAllowedAccess)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ConfigPeerGossip : ConfigPeerGossip struct
type ConfigPeerGossip struct {
	// Decides whether a peer will use a dynamic algorithm for "leader" selection (instead of a static leader). The leader
	// is the peer that establishes a connection with the ordering service (OS). The leader pulls ledger blocks from the
	// OS. It is recommended to use leader election for large networks of peers.
	UseLeaderElection *bool `json:"useLeaderElection,omitempty"`

	// Decides whether this peer should be an organization "leader". It maintains a connection with the ordering service
	// and disseminate blocks to peers in its own organization.
	OrgLeader *bool `json:"orgLeader,omitempty"`

	// The frequency to poll on membershipTracker.
	MembershipTrackerInterval *string `json:"membershipTrackerInterval,omitempty"`

	// Maximum number of blocks that can be stored in memory.
	MaxBlockCountToStore *float64 `json:"maxBlockCountToStore,omitempty"`

	// Maximum time between consecutive message pushes.
	MaxPropagationBurstLatency *string `json:"maxPropagationBurstLatency,omitempty"`

	// Maximum number of messages that are stored until a push to remote peers is triggered.
	MaxPropagationBurstSize *float64 `json:"maxPropagationBurstSize,omitempty"`

	// Number of times a message is pushed to remote peers.
	PropagateIterations *float64 `json:"propagateIterations,omitempty"`

	// Determines the frequency of pull phases.
	PullInterval *string `json:"pullInterval,omitempty"`

	// Number of peers to pull from.
	PullPeerNum *float64 `json:"pullPeerNum,omitempty"`

	// Determines the frequency of pulling stateInfo messages from peers.
	RequestStateInfoInterval *string `json:"requestStateInfoInterval,omitempty"`

	// Determines the frequency of pushing stateInfo messages to peers.
	PublishStateInfoInterval *string `json:"publishStateInfoInterval,omitempty"`

	// Maximum time a stateInfo message is kept.
	StateInfoRetentionInterval *string `json:"stateInfoRetentionInterval,omitempty"`

	// Time after startup to start including certificates in Alive messages.
	PublishCertPeriod *string `json:"publishCertPeriod,omitempty"`

	// Decides whether the peer should skip the verification of block messages.
	SkipBlockVerification *bool `json:"skipBlockVerification,omitempty"`

	// The timeout for dialing a network request.
	DialTimeout *string `json:"dialTimeout,omitempty"`

	// The timeout for a network connection.
	ConnTimeout *string `json:"connTimeout,omitempty"`

	// Number of received messages to hold in buffer.
	RecvBuffSize *float64 `json:"recvBuffSize,omitempty"`

	// Number of sent messages to hold in buffer.
	SendBuffSize *float64 `json:"sendBuffSize,omitempty"`

	// Time to wait before the pull-engine processes incoming digests. Should be slightly smaller than requestWaitTime.
	DigestWaitTime *string `json:"digestWaitTime,omitempty"`

	// Time to wait before pull-engine removes the incoming nonce. Should be slightly bigger than digestWaitTime.
	RequestWaitTime *string `json:"requestWaitTime,omitempty"`

	// Time to wait before the pull-engine ends.
	ResponseWaitTime *string `json:"responseWaitTime,omitempty"`

	// Alive check frequency.
	AliveTimeInterval *string `json:"aliveTimeInterval,omitempty"`

	// Alive expiration timeout.
	AliveExpirationTimeout *string `json:"aliveExpirationTimeout,omitempty"`

	// Reconnect frequency.
	ReconnectInterval *string `json:"reconnectInterval,omitempty"`

	// Leader election service configuration.
	Election *ConfigPeerGossipElection `json:"election,omitempty"`

	PvtData *ConfigPeerGossipPvtData `json:"pvtData,omitempty"`

	// Gossip state transfer related configuration.
	State *ConfigPeerGossipState `json:"state,omitempty"`
}


// UnmarshalConfigPeerGossip unmarshals an instance of ConfigPeerGossip from the specified map of raw messages.
func UnmarshalConfigPeerGossip(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ConfigPeerGossip)
	err = core.UnmarshalPrimitive(m, "useLeaderElection", &obj.UseLeaderElection)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "orgLeader", &obj.OrgLeader)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "membershipTrackerInterval", &obj.MembershipTrackerInterval)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "maxBlockCountToStore", &obj.MaxBlockCountToStore)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "maxPropagationBurstLatency", &obj.MaxPropagationBurstLatency)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "maxPropagationBurstSize", &obj.MaxPropagationBurstSize)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "propagateIterations", &obj.PropagateIterations)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "pullInterval", &obj.PullInterval)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "pullPeerNum", &obj.PullPeerNum)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "requestStateInfoInterval", &obj.RequestStateInfoInterval)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "publishStateInfoInterval", &obj.PublishStateInfoInterval)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "stateInfoRetentionInterval", &obj.StateInfoRetentionInterval)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "publishCertPeriod", &obj.PublishCertPeriod)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "skipBlockVerification", &obj.SkipBlockVerification)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "dialTimeout", &obj.DialTimeout)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "connTimeout", &obj.ConnTimeout)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "recvBuffSize", &obj.RecvBuffSize)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "sendBuffSize", &obj.SendBuffSize)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "digestWaitTime", &obj.DigestWaitTime)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "requestWaitTime", &obj.RequestWaitTime)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "responseWaitTime", &obj.ResponseWaitTime)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "aliveTimeInterval", &obj.AliveTimeInterval)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "aliveExpirationTimeout", &obj.AliveExpirationTimeout)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "reconnectInterval", &obj.ReconnectInterval)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "election", &obj.Election, UnmarshalConfigPeerGossipElection)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "pvtData", &obj.PvtData, UnmarshalConfigPeerGossipPvtData)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "state", &obj.State, UnmarshalConfigPeerGossipState)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ConfigPeerKeepalive : Keep alive settings between the peer server and clients.
type ConfigPeerKeepalive struct {
	// The minimum time between client pings. If a client sends pings more frequently the server disconnects from the
	// client.
	MinInterval *string `json:"minInterval,omitempty"`

	Client *ConfigPeerKeepaliveClient `json:"client,omitempty"`

	DeliveryClient *ConfigPeerKeepaliveDeliveryClient `json:"deliveryClient,omitempty"`
}


// UnmarshalConfigPeerKeepalive unmarshals an instance of ConfigPeerKeepalive from the specified map of raw messages.
func UnmarshalConfigPeerKeepalive(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ConfigPeerKeepalive)
	err = core.UnmarshalPrimitive(m, "minInterval", &obj.MinInterval)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "client", &obj.Client, UnmarshalConfigPeerKeepaliveClient)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "deliveryClient", &obj.DeliveryClient, UnmarshalConfigPeerKeepaliveDeliveryClient)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ConfigPeerLimits : ConfigPeerLimits struct
type ConfigPeerLimits struct {
	Concurrency *ConfigPeerLimitsConcurrency `json:"concurrency,omitempty"`
}


// UnmarshalConfigPeerLimits unmarshals an instance of ConfigPeerLimits from the specified map of raw messages.
func UnmarshalConfigPeerLimits(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ConfigPeerLimits)
	err = core.UnmarshalModel(m, "concurrency", &obj.Concurrency, UnmarshalConfigPeerLimitsConcurrency)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// CpuHealthStats : CpuHealthStats struct
type CpuHealthStats struct {
	// Model of CPU core.
	Model *string `json:"model,omitempty"`

	// Speed of core in MHz.
	Speed *string `json:"speed,omitempty"`

	Times *CpuHealthStatsTimes `json:"times,omitempty"`
}


// UnmarshalCpuHealthStats unmarshals an instance of CpuHealthStats from the specified map of raw messages.
func UnmarshalCpuHealthStats(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(CpuHealthStats)
	err = core.UnmarshalPrimitive(m, "model", &obj.Model)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "speed", &obj.Speed)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "times", &obj.Times, UnmarshalCpuHealthStatsTimes)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// CpuHealthStatsTimes : CpuHealthStatsTimes struct
type CpuHealthStatsTimes struct {
	// ms CPU is in idle.
	Idle *float64 `json:"idle,omitempty"`

	// ms CPU is in irq.
	Irq *float64 `json:"irq,omitempty"`

	// ms CPU is in nice.
	Nice *float64 `json:"nice,omitempty"`

	// ms CPU is in sys.
	Sys *float64 `json:"sys,omitempty"`

	// ms CPU is in user.
	User *float64 `json:"user,omitempty"`
}


// UnmarshalCpuHealthStatsTimes unmarshals an instance of CpuHealthStatsTimes from the specified map of raw messages.
func UnmarshalCpuHealthStatsTimes(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(CpuHealthStatsTimes)
	err = core.UnmarshalPrimitive(m, "idle", &obj.Idle)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "irq", &obj.Irq)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "nice", &obj.Nice)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "sys", &obj.Sys)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "user", &obj.User)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// CreateCaBodyConfigOverride : Set `config_override` to create the root/initial enroll id and enroll secret as well as enabling custom CA
// configurations (such as using postgres). See the [Fabric CA configuration
// file](https://hyperledger-fabric-ca.readthedocs.io/en/release-1.4/serverconfig.html) for available options.
//
// The field `tlsca` is optional. The IBP console will copy the value of `config_override.ca` into
// `config_override.tlsca` if `config_override.tlsca` is omitted (which is recommended).
//
// *The field **names** below are not case-sensitive.*.
type CreateCaBodyConfigOverride struct {
	Ca *ConfigCACreate `json:"ca" validate:"required"`

	Tlsca *ConfigCACreate `json:"tlsca,omitempty"`
}


// NewCreateCaBodyConfigOverride : Instantiate CreateCaBodyConfigOverride (Generic Model Constructor)
func (*BlockchainV2) NewCreateCaBodyConfigOverride(ca *ConfigCACreate) (model *CreateCaBodyConfigOverride, err error) {
	model = &CreateCaBodyConfigOverride{
		Ca: ca,
	}
	err = core.ValidateStruct(model, "required parameters")
	return
}

// UnmarshalCreateCaBodyConfigOverride unmarshals an instance of CreateCaBodyConfigOverride from the specified map of raw messages.
func UnmarshalCreateCaBodyConfigOverride(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(CreateCaBodyConfigOverride)
	err = core.UnmarshalModel(m, "ca", &obj.Ca, UnmarshalConfigCACreate)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "tlsca", &obj.Tlsca, UnmarshalConfigCACreate)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// CreateCaBodyResources : CPU and memory properties. This feature is not available if using a free Kubernetes cluster.
type CreateCaBodyResources struct {
	// This field requires the use of Fabric v1.4.* and higher.
	Ca *ResourceObject `json:"ca" validate:"required"`
}


// NewCreateCaBodyResources : Instantiate CreateCaBodyResources (Generic Model Constructor)
func (*BlockchainV2) NewCreateCaBodyResources(ca *ResourceObject) (model *CreateCaBodyResources, err error) {
	model = &CreateCaBodyResources{
		Ca: ca,
	}
	err = core.ValidateStruct(model, "required parameters")
	return
}

// UnmarshalCreateCaBodyResources unmarshals an instance of CreateCaBodyResources from the specified map of raw messages.
func UnmarshalCreateCaBodyResources(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(CreateCaBodyResources)
	err = core.UnmarshalModel(m, "ca", &obj.Ca, UnmarshalResourceObject)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// CreateCaBodyStorage : Disk space properties. This feature is not available if using a free Kubernetes cluster.
type CreateCaBodyStorage struct {
	Ca *StorageObject `json:"ca" validate:"required"`
}


// NewCreateCaBodyStorage : Instantiate CreateCaBodyStorage (Generic Model Constructor)
func (*BlockchainV2) NewCreateCaBodyStorage(ca *StorageObject) (model *CreateCaBodyStorage, err error) {
	model = &CreateCaBodyStorage{
		Ca: ca,
	}
	err = core.ValidateStruct(model, "required parameters")
	return
}

// UnmarshalCreateCaBodyStorage unmarshals an instance of CreateCaBodyStorage from the specified map of raw messages.
func UnmarshalCreateCaBodyStorage(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(CreateCaBodyStorage)
	err = core.UnmarshalModel(m, "ca", &obj.Ca, UnmarshalStorageObject)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// CreateCaOptions : The CreateCa options.
type CreateCaOptions struct {
	// A descriptive name for this CA. The IBP console tile displays this name.
	DisplayName *string `json:"display_name" validate:"required"`

	// Set `config_override` to create the root/initial enroll id and enroll secret as well as enabling custom CA
	// configurations (such as using postgres). See the [Fabric CA configuration
	// file](https://hyperledger-fabric-ca.readthedocs.io/en/release-1.4/serverconfig.html) for available options.
	//
	// The field `tlsca` is optional. The IBP console will copy the value of `config_override.ca` into
	// `config_override.tlsca` if `config_override.tlsca` is omitted (which is recommended).
	//
	// *The field **names** below are not case-sensitive.*.
	ConfigOverride *CreateCaBodyConfigOverride `json:"config_override" validate:"required"`

	// CPU and memory properties. This feature is not available if using a free Kubernetes cluster.
	Resources *CreateCaBodyResources `json:"resources,omitempty"`

	// Disk space properties. This feature is not available if using a free Kubernetes cluster.
	Storage *CreateCaBodyStorage `json:"storage,omitempty"`

	// Specify the Kubernetes zone for the deployment. The deployment will use a k8s node in this zone. Find the list of
	// possible zones by retrieving your Kubernetes node labels: `kubectl get nodes --show-labels`. [More
	// information](https://kubernetes.io/docs/setup/best-practices/multiple-zones).
	Zone *string `json:"zone,omitempty"`

	// The number of replica pods running at any given time.
	Replicas *float64 `json:"replicas,omitempty"`

	Tags []string `json:"tags,omitempty"`

	// The connection details of the HSM (Hardware Security Module).
	Hsm *Hsm `json:"hsm,omitempty"`

	// Specify the Kubernetes region for the deployment. The deployment will use a k8s node in this region. Find the list
	// of possible regions by retrieving your Kubernetes node labels: `kubectl get nodes --show-labels`. [More
	// info](https://kubernetes.io/docs/setup/best-practices/multiple-zones).
	Region *string `json:"region,omitempty"`

	// The Hyperledger Fabric release version to use.
	Version *string `json:"version,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewCreateCaOptions : Instantiate CreateCaOptions
func (*BlockchainV2) NewCreateCaOptions(displayName string, configOverride *CreateCaBodyConfigOverride) *CreateCaOptions {
	return &CreateCaOptions{
		DisplayName: core.StringPtr(displayName),
		ConfigOverride: configOverride,
	}
}

// SetDisplayName : Allow user to set DisplayName
func (options *CreateCaOptions) SetDisplayName(displayName string) *CreateCaOptions {
	options.DisplayName = core.StringPtr(displayName)
	return options
}

// SetConfigOverride : Allow user to set ConfigOverride
func (options *CreateCaOptions) SetConfigOverride(configOverride *CreateCaBodyConfigOverride) *CreateCaOptions {
	options.ConfigOverride = configOverride
	return options
}

// SetResources : Allow user to set Resources
func (options *CreateCaOptions) SetResources(resources *CreateCaBodyResources) *CreateCaOptions {
	options.Resources = resources
	return options
}

// SetStorage : Allow user to set Storage
func (options *CreateCaOptions) SetStorage(storage *CreateCaBodyStorage) *CreateCaOptions {
	options.Storage = storage
	return options
}

// SetZone : Allow user to set Zone
func (options *CreateCaOptions) SetZone(zone string) *CreateCaOptions {
	options.Zone = core.StringPtr(zone)
	return options
}

// SetReplicas : Allow user to set Replicas
func (options *CreateCaOptions) SetReplicas(replicas float64) *CreateCaOptions {
	options.Replicas = core.Float64Ptr(replicas)
	return options
}

// SetTags : Allow user to set Tags
func (options *CreateCaOptions) SetTags(tags []string) *CreateCaOptions {
	options.Tags = tags
	return options
}

// SetHsm : Allow user to set Hsm
func (options *CreateCaOptions) SetHsm(hsm *Hsm) *CreateCaOptions {
	options.Hsm = hsm
	return options
}

// SetRegion : Allow user to set Region
func (options *CreateCaOptions) SetRegion(region string) *CreateCaOptions {
	options.Region = core.StringPtr(region)
	return options
}

// SetVersion : Allow user to set Version
func (options *CreateCaOptions) SetVersion(version string) *CreateCaOptions {
	options.Version = core.StringPtr(version)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *CreateCaOptions) SetHeaders(param map[string]string) *CreateCaOptions {
	options.Headers = param
	return options
}

// CreateOrdererOptions : The CreateOrderer options.
type CreateOrdererOptions struct {
	// The type of Fabric orderer. Currently, only the type `"raft"` is supported.
	// [etcd/raft](/docs/blockchain?topic=blockchain-ibp-console-build-network#ibp-console-build-network-ordering-console).
	OrdererType *string `json:"orderer_type" validate:"required"`

	// The MSP id that is related to this component.
	MspID *string `json:"msp_id" validate:"required"`

	// A descriptive base name for each ordering node. One or more child IBP console tiles display this name.
	DisplayName *string `json:"display_name" validate:"required"`

	// An array of config objects. Add one per desired raft node. 1 or 5 nodes are recommended.
	//
	// See this [topic](/docs/blockchain?topic=blockchain-ibp-v2-apis#ibp-v2-apis-config) for instructions on how to build
	// a config object.
	Config []ConfigObject `json:"config" validate:"required"`

	// A descriptive name for an ordering service. The parent IBP console tile displays this name.
	//
	// This field should only be set if you are creating a new OS cluster or when appending to an unknown (external) OS
	// cluster. An unknown/external cluster is one that this IBP console has not imported or created.
	ClusterName *string `json:"cluster_name,omitempty"`

	// This field should only be set if you are appending a new raft node to an **existing** raft cluster. When appending
	// to a known (internal) OS cluster set `cluster_id` to the same value used by the OS cluster. When appending to an
	// unknown (external) OS cluster set `cluster_id` to a unique string.
	//
	// Setting this field means the `config` array should be of length 1, since it is not possible to add multiple raft
	// nodes at the same time in Fabric.
	//
	// If this field is set the orderer will be "pre-created" and start without a genesis block. It is effectively dead
	// until it is configured. This is the first step to **append** a node to a raft cluster. The next step is to add this
	// node as a consenter to the system-channel by using Fabric-APIs. Then, init this node by sending the updated
	// system-channel config-block with the [Submit config block to orderer](#submit-config-block-to-orderer) API. The node
	// will not be usable until these steps are completed.
	ClusterID *string `json:"cluster_id,omitempty"`

	// Set to `true` only if you are appending to an unknown (external) OS cluster. Else set it to `false` or omit the
	// field. An unknown/external cluster is one that this IBP console has not imported or created.
	ExternalAppend *string `json:"external_append,omitempty"`

	// An array of configuration override objects. 1 object per component. Must be the same size as the `config` array.
	ConfigOverride []ConfigOrdererCreate `json:"config_override,omitempty"`

	// CPU and memory properties. This feature is not available if using a free Kubernetes cluster.
	Resources *CreateOrdererRaftBodyResources `json:"resources,omitempty"`

	// Disk space properties. This feature is not available if using a free Kubernetes cluster.
	Storage *CreateOrdererRaftBodyStorage `json:"storage,omitempty"`

	// The name of the system channel. Defaults to `testchainid`.
	SystemChannelID *string `json:"system_channel_id,omitempty"`

	// An array of Kubernetes zones for the deployment. 1 zone per component. Must be the same size as the `config` array.
	Zone []string `json:"zone,omitempty"`

	Tags []string `json:"tags,omitempty"`

	// An array of Kubernetes regions for the deployment. One region per component. Must be the same size as the `config`
	// array.
	Region []string `json:"region,omitempty"`

	// The connection details of the HSM (Hardware Security Module).
	Hsm *Hsm `json:"hsm,omitempty"`

	// The Hyperledger Fabric release version to use.
	Version *string `json:"version,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// Constants associated with the CreateOrdererOptions.OrdererType property.
// The type of Fabric orderer. Currently, only the type `"raft"` is supported.
// [etcd/raft](/docs/blockchain?topic=blockchain-ibp-console-build-network#ibp-console-build-network-ordering-console).
const (
	CreateOrdererOptions_OrdererType_Raft = "raft"
)

// NewCreateOrdererOptions : Instantiate CreateOrdererOptions
func (*BlockchainV2) NewCreateOrdererOptions(ordererType string, mspID string, displayName string, config []ConfigObject) *CreateOrdererOptions {
	return &CreateOrdererOptions{
		OrdererType: core.StringPtr(ordererType),
		MspID: core.StringPtr(mspID),
		DisplayName: core.StringPtr(displayName),
		Config: config,
	}
}

// SetOrdererType : Allow user to set OrdererType
func (options *CreateOrdererOptions) SetOrdererType(ordererType string) *CreateOrdererOptions {
	options.OrdererType = core.StringPtr(ordererType)
	return options
}

// SetMspID : Allow user to set MspID
func (options *CreateOrdererOptions) SetMspID(mspID string) *CreateOrdererOptions {
	options.MspID = core.StringPtr(mspID)
	return options
}

// SetDisplayName : Allow user to set DisplayName
func (options *CreateOrdererOptions) SetDisplayName(displayName string) *CreateOrdererOptions {
	options.DisplayName = core.StringPtr(displayName)
	return options
}

// SetConfig : Allow user to set Config
func (options *CreateOrdererOptions) SetConfig(config []ConfigObject) *CreateOrdererOptions {
	options.Config = config
	return options
}

// SetClusterName : Allow user to set ClusterName
func (options *CreateOrdererOptions) SetClusterName(clusterName string) *CreateOrdererOptions {
	options.ClusterName = core.StringPtr(clusterName)
	return options
}

// SetClusterID : Allow user to set ClusterID
func (options *CreateOrdererOptions) SetClusterID(clusterID string) *CreateOrdererOptions {
	options.ClusterID = core.StringPtr(clusterID)
	return options
}

// SetExternalAppend : Allow user to set ExternalAppend
func (options *CreateOrdererOptions) SetExternalAppend(externalAppend string) *CreateOrdererOptions {
	options.ExternalAppend = core.StringPtr(externalAppend)
	return options
}

// SetConfigOverride : Allow user to set ConfigOverride
func (options *CreateOrdererOptions) SetConfigOverride(configOverride []ConfigOrdererCreate) *CreateOrdererOptions {
	options.ConfigOverride = configOverride
	return options
}

// SetResources : Allow user to set Resources
func (options *CreateOrdererOptions) SetResources(resources *CreateOrdererRaftBodyResources) *CreateOrdererOptions {
	options.Resources = resources
	return options
}

// SetStorage : Allow user to set Storage
func (options *CreateOrdererOptions) SetStorage(storage *CreateOrdererRaftBodyStorage) *CreateOrdererOptions {
	options.Storage = storage
	return options
}

// SetSystemChannelID : Allow user to set SystemChannelID
func (options *CreateOrdererOptions) SetSystemChannelID(systemChannelID string) *CreateOrdererOptions {
	options.SystemChannelID = core.StringPtr(systemChannelID)
	return options
}

// SetZone : Allow user to set Zone
func (options *CreateOrdererOptions) SetZone(zone []string) *CreateOrdererOptions {
	options.Zone = zone
	return options
}

// SetTags : Allow user to set Tags
func (options *CreateOrdererOptions) SetTags(tags []string) *CreateOrdererOptions {
	options.Tags = tags
	return options
}

// SetRegion : Allow user to set Region
func (options *CreateOrdererOptions) SetRegion(region []string) *CreateOrdererOptions {
	options.Region = region
	return options
}

// SetHsm : Allow user to set Hsm
func (options *CreateOrdererOptions) SetHsm(hsm *Hsm) *CreateOrdererOptions {
	options.Hsm = hsm
	return options
}

// SetVersion : Allow user to set Version
func (options *CreateOrdererOptions) SetVersion(version string) *CreateOrdererOptions {
	options.Version = core.StringPtr(version)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *CreateOrdererOptions) SetHeaders(param map[string]string) *CreateOrdererOptions {
	options.Headers = param
	return options
}

// CreateOrdererRaftBodyResources : CPU and memory properties. This feature is not available if using a free Kubernetes cluster.
type CreateOrdererRaftBodyResources struct {
	// This field requires the use of Fabric v1.4.* and higher.
	Orderer *ResourceObject `json:"orderer" validate:"required"`

	// This field requires the use of Fabric v1.4.* and higher.
	Proxy *ResourceObject `json:"proxy,omitempty"`
}


// NewCreateOrdererRaftBodyResources : Instantiate CreateOrdererRaftBodyResources (Generic Model Constructor)
func (*BlockchainV2) NewCreateOrdererRaftBodyResources(orderer *ResourceObject) (model *CreateOrdererRaftBodyResources, err error) {
	model = &CreateOrdererRaftBodyResources{
		Orderer: orderer,
	}
	err = core.ValidateStruct(model, "required parameters")
	return
}

// UnmarshalCreateOrdererRaftBodyResources unmarshals an instance of CreateOrdererRaftBodyResources from the specified map of raw messages.
func UnmarshalCreateOrdererRaftBodyResources(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(CreateOrdererRaftBodyResources)
	err = core.UnmarshalModel(m, "orderer", &obj.Orderer, UnmarshalResourceObject)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "proxy", &obj.Proxy, UnmarshalResourceObject)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// CreateOrdererRaftBodyStorage : Disk space properties. This feature is not available if using a free Kubernetes cluster.
type CreateOrdererRaftBodyStorage struct {
	Orderer *StorageObject `json:"orderer" validate:"required"`
}


// NewCreateOrdererRaftBodyStorage : Instantiate CreateOrdererRaftBodyStorage (Generic Model Constructor)
func (*BlockchainV2) NewCreateOrdererRaftBodyStorage(orderer *StorageObject) (model *CreateOrdererRaftBodyStorage, err error) {
	model = &CreateOrdererRaftBodyStorage{
		Orderer: orderer,
	}
	err = core.ValidateStruct(model, "required parameters")
	return
}

// UnmarshalCreateOrdererRaftBodyStorage unmarshals an instance of CreateOrdererRaftBodyStorage from the specified map of raw messages.
func UnmarshalCreateOrdererRaftBodyStorage(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(CreateOrdererRaftBodyStorage)
	err = core.UnmarshalModel(m, "orderer", &obj.Orderer, UnmarshalStorageObject)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// CreatePeerBodyStorage : Disk space properties. This feature is not available if using a free Kubernetes cluster.
type CreatePeerBodyStorage struct {
	Peer *StorageObject `json:"peer" validate:"required"`

	Statedb *StorageObject `json:"statedb,omitempty"`
}


// NewCreatePeerBodyStorage : Instantiate CreatePeerBodyStorage (Generic Model Constructor)
func (*BlockchainV2) NewCreatePeerBodyStorage(peer *StorageObject) (model *CreatePeerBodyStorage, err error) {
	model = &CreatePeerBodyStorage{
		Peer: peer,
	}
	err = core.ValidateStruct(model, "required parameters")
	return
}

// UnmarshalCreatePeerBodyStorage unmarshals an instance of CreatePeerBodyStorage from the specified map of raw messages.
func UnmarshalCreatePeerBodyStorage(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(CreatePeerBodyStorage)
	err = core.UnmarshalModel(m, "peer", &obj.Peer, UnmarshalStorageObject)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "statedb", &obj.Statedb, UnmarshalStorageObject)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// CreatePeerOptions : The CreatePeer options.
type CreatePeerOptions struct {
	// The MSP id that is related to this component.
	MspID *string `json:"msp_id" validate:"required"`

	// A descriptive name for this peer. The IBP console tile displays this name.
	DisplayName *string `json:"display_name" validate:"required"`

	Config *ConfigObject `json:"config" validate:"required"`

	// Override the [Fabric Peer configuration
	// file](https://github.com/hyperledger/fabric/blob/release-1.4/sampleconfig/core.yaml) if you want use custom
	// attributes to configure the Peer. Omit if not.
	//
	// *The field **names** below are not case-sensitive.*.
	ConfigOverride *ConfigPeerCreate `json:"config_override,omitempty"`

	// CPU and memory properties. This feature is not available if using a free Kubernetes cluster.
	Resources *PeerResources `json:"resources,omitempty"`

	// Disk space properties. This feature is not available if using a free Kubernetes cluster.
	Storage *CreatePeerBodyStorage `json:"storage,omitempty"`

	// Specify the Kubernetes zone for the deployment. The deployment will use a k8s node in this zone. Find the list of
	// possible zones by retrieving your Kubernetes node labels: `kubectl get nodes --show-labels`. [More
	// information](https://kubernetes.io/docs/setup/best-practices/multiple-zones).
	Zone *string `json:"zone,omitempty"`

	// Select the state database for the peer. Can be either "couchdb" or "leveldb". The default is "couchdb".
	StateDb *string `json:"state_db,omitempty"`

	Tags []string `json:"tags,omitempty"`

	// The connection details of the HSM (Hardware Security Module).
	Hsm *Hsm `json:"hsm,omitempty"`

	// Specify the Kubernetes region for the deployment. The deployment will use a k8s node in this region. Find the list
	// of possible regions by retrieving your Kubernetes node labels: `kubectl get nodes --show-labels`. [More
	// info](https://kubernetes.io/docs/setup/best-practices/multiple-zones).
	Region *string `json:"region,omitempty"`

	// The Hyperledger Fabric release version to use.
	Version *string `json:"version,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// Constants associated with the CreatePeerOptions.StateDb property.
// Select the state database for the peer. Can be either "couchdb" or "leveldb". The default is "couchdb".
const (
	CreatePeerOptions_StateDb_Couchdb = "couchdb"
	CreatePeerOptions_StateDb_Leveldb = "leveldb"
)

// NewCreatePeerOptions : Instantiate CreatePeerOptions
func (*BlockchainV2) NewCreatePeerOptions(mspID string, displayName string, config *ConfigObject) *CreatePeerOptions {
	return &CreatePeerOptions{
		MspID: core.StringPtr(mspID),
		DisplayName: core.StringPtr(displayName),
		Config: config,
	}
}

// SetMspID : Allow user to set MspID
func (options *CreatePeerOptions) SetMspID(mspID string) *CreatePeerOptions {
	options.MspID = core.StringPtr(mspID)
	return options
}

// SetDisplayName : Allow user to set DisplayName
func (options *CreatePeerOptions) SetDisplayName(displayName string) *CreatePeerOptions {
	options.DisplayName = core.StringPtr(displayName)
	return options
}

// SetConfig : Allow user to set Config
func (options *CreatePeerOptions) SetConfig(config *ConfigObject) *CreatePeerOptions {
	options.Config = config
	return options
}

// SetConfigOverride : Allow user to set ConfigOverride
func (options *CreatePeerOptions) SetConfigOverride(configOverride *ConfigPeerCreate) *CreatePeerOptions {
	options.ConfigOverride = configOverride
	return options
}

// SetResources : Allow user to set Resources
func (options *CreatePeerOptions) SetResources(resources *PeerResources) *CreatePeerOptions {
	options.Resources = resources
	return options
}

// SetStorage : Allow user to set Storage
func (options *CreatePeerOptions) SetStorage(storage *CreatePeerBodyStorage) *CreatePeerOptions {
	options.Storage = storage
	return options
}

// SetZone : Allow user to set Zone
func (options *CreatePeerOptions) SetZone(zone string) *CreatePeerOptions {
	options.Zone = core.StringPtr(zone)
	return options
}

// SetStateDb : Allow user to set StateDb
func (options *CreatePeerOptions) SetStateDb(stateDb string) *CreatePeerOptions {
	options.StateDb = core.StringPtr(stateDb)
	return options
}

// SetTags : Allow user to set Tags
func (options *CreatePeerOptions) SetTags(tags []string) *CreatePeerOptions {
	options.Tags = tags
	return options
}

// SetHsm : Allow user to set Hsm
func (options *CreatePeerOptions) SetHsm(hsm *Hsm) *CreatePeerOptions {
	options.Hsm = hsm
	return options
}

// SetRegion : Allow user to set Region
func (options *CreatePeerOptions) SetRegion(region string) *CreatePeerOptions {
	options.Region = core.StringPtr(region)
	return options
}

// SetVersion : Allow user to set Version
func (options *CreatePeerOptions) SetVersion(version string) *CreatePeerOptions {
	options.Version = core.StringPtr(version)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *CreatePeerOptions) SetHeaders(param map[string]string) *CreatePeerOptions {
	options.Headers = param
	return options
}

// DeleteAllComponentsOptions : The DeleteAllComponents options.
type DeleteAllComponentsOptions struct {

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDeleteAllComponentsOptions : Instantiate DeleteAllComponentsOptions
func (*BlockchainV2) NewDeleteAllComponentsOptions() *DeleteAllComponentsOptions {
	return &DeleteAllComponentsOptions{}
}

// SetHeaders : Allow user to set Headers
func (options *DeleteAllComponentsOptions) SetHeaders(param map[string]string) *DeleteAllComponentsOptions {
	options.Headers = param
	return options
}

// DeleteAllNotificationsOptions : The DeleteAllNotifications options.
type DeleteAllNotificationsOptions struct {

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDeleteAllNotificationsOptions : Instantiate DeleteAllNotificationsOptions
func (*BlockchainV2) NewDeleteAllNotificationsOptions() *DeleteAllNotificationsOptions {
	return &DeleteAllNotificationsOptions{}
}

// SetHeaders : Allow user to set Headers
func (options *DeleteAllNotificationsOptions) SetHeaders(param map[string]string) *DeleteAllNotificationsOptions {
	options.Headers = param
	return options
}

// DeleteAllNotificationsResponse : DeleteAllNotificationsResponse struct
type DeleteAllNotificationsResponse struct {
	// Response message. "ok" indicates the api completed successfully.
	Message *string `json:"message,omitempty"`

	// Text showing what was deleted.
	Details *string `json:"details,omitempty"`
}


// UnmarshalDeleteAllNotificationsResponse unmarshals an instance of DeleteAllNotificationsResponse from the specified map of raw messages.
func UnmarshalDeleteAllNotificationsResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(DeleteAllNotificationsResponse)
	err = core.UnmarshalPrimitive(m, "message", &obj.Message)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "details", &obj.Details)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// DeleteAllSessionsOptions : The DeleteAllSessions options.
type DeleteAllSessionsOptions struct {

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDeleteAllSessionsOptions : Instantiate DeleteAllSessionsOptions
func (*BlockchainV2) NewDeleteAllSessionsOptions() *DeleteAllSessionsOptions {
	return &DeleteAllSessionsOptions{}
}

// SetHeaders : Allow user to set Headers
func (options *DeleteAllSessionsOptions) SetHeaders(param map[string]string) *DeleteAllSessionsOptions {
	options.Headers = param
	return options
}

// DeleteAllSessionsResponse : Describes the outcome of the api.
type DeleteAllSessionsResponse struct {
	// Response message. "ok" indicates the api completed successfully.
	Message *string `json:"message,omitempty"`

	// How many session entries that were deleted.
	Deleted *float64 `json:"deleted,omitempty"`
}


// UnmarshalDeleteAllSessionsResponse unmarshals an instance of DeleteAllSessionsResponse from the specified map of raw messages.
func UnmarshalDeleteAllSessionsResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(DeleteAllSessionsResponse)
	err = core.UnmarshalPrimitive(m, "message", &obj.Message)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "deleted", &obj.Deleted)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// DeleteComponentOptions : The DeleteComponent options.
type DeleteComponentOptions struct {
	// The `id` of the component to delete. Use the [Get all components](#get-all-components) API to determine the id of
	// the component to be deleted.
	ID *string `json:"id" validate:"required"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDeleteComponentOptions : Instantiate DeleteComponentOptions
func (*BlockchainV2) NewDeleteComponentOptions(id string) *DeleteComponentOptions {
	return &DeleteComponentOptions{
		ID: core.StringPtr(id),
	}
}

// SetID : Allow user to set ID
func (options *DeleteComponentOptions) SetID(id string) *DeleteComponentOptions {
	options.ID = core.StringPtr(id)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteComponentOptions) SetHeaders(param map[string]string) *DeleteComponentOptions {
	options.Headers = param
	return options
}

// DeleteComponentResponse : DeleteComponentResponse struct
type DeleteComponentResponse struct {
	Message *string `json:"message,omitempty"`

	// The type of this component. Such as: "fabric-peer", "fabric-ca", "fabric-orderer", etc.
	Type *string `json:"type,omitempty"`

	// The unique identifier of this component.
	ID *string `json:"id,omitempty"`

	// A descriptive name for this peer. The IBP console tile displays this name.
	DisplayName *string `json:"display_name,omitempty"`
}


// UnmarshalDeleteComponentResponse unmarshals an instance of DeleteComponentResponse from the specified map of raw messages.
func UnmarshalDeleteComponentResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(DeleteComponentResponse)
	err = core.UnmarshalPrimitive(m, "message", &obj.Message)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "display_name", &obj.DisplayName)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// DeleteComponentsByTagOptions : The DeleteComponentsByTag options.
type DeleteComponentsByTagOptions struct {
	// The tag to filter components on. Not case-sensitive.
	Tag *string `json:"tag" validate:"required"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDeleteComponentsByTagOptions : Instantiate DeleteComponentsByTagOptions
func (*BlockchainV2) NewDeleteComponentsByTagOptions(tag string) *DeleteComponentsByTagOptions {
	return &DeleteComponentsByTagOptions{
		Tag: core.StringPtr(tag),
	}
}

// SetTag : Allow user to set Tag
func (options *DeleteComponentsByTagOptions) SetTag(tag string) *DeleteComponentsByTagOptions {
	options.Tag = core.StringPtr(tag)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteComponentsByTagOptions) SetHeaders(param map[string]string) *DeleteComponentsByTagOptions {
	options.Headers = param
	return options
}

// DeleteMultiComponentsResponse : DeleteMultiComponentsResponse struct
type DeleteMultiComponentsResponse struct {
	Deleted []DeleteComponentResponse `json:"deleted,omitempty"`
}


// UnmarshalDeleteMultiComponentsResponse unmarshals an instance of DeleteMultiComponentsResponse from the specified map of raw messages.
func UnmarshalDeleteMultiComponentsResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(DeleteMultiComponentsResponse)
	err = core.UnmarshalModel(m, "deleted", &obj.Deleted, UnmarshalDeleteComponentResponse)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// DeleteSigTxOptions : The DeleteSigTx options.
type DeleteSigTxOptions struct {
	// The unique transaction ID of this signature collection.
	ID *string `json:"id" validate:"required"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDeleteSigTxOptions : Instantiate DeleteSigTxOptions
func (*BlockchainV2) NewDeleteSigTxOptions(id string) *DeleteSigTxOptions {
	return &DeleteSigTxOptions{
		ID: core.StringPtr(id),
	}
}

// SetID : Allow user to set ID
func (options *DeleteSigTxOptions) SetID(id string) *DeleteSigTxOptions {
	options.ID = core.StringPtr(id)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteSigTxOptions) SetHeaders(param map[string]string) *DeleteSigTxOptions {
	options.Headers = param
	return options
}

// DeleteSignatureCollectionResponse : DeleteSignatureCollectionResponse struct
type DeleteSignatureCollectionResponse struct {
	// Response message. "ok" indicates the api completed successfully.
	Message *string `json:"message,omitempty"`

	// The unique transaction ID of this signature collection. Must start with a letter.
	TxID *string `json:"tx_id,omitempty"`
}


// UnmarshalDeleteSignatureCollectionResponse unmarshals an instance of DeleteSignatureCollectionResponse from the specified map of raw messages.
func UnmarshalDeleteSignatureCollectionResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(DeleteSignatureCollectionResponse)
	err = core.UnmarshalPrimitive(m, "message", &obj.Message)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "tx_id", &obj.TxID)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// EditAdminCertsOptions : The EditAdminCerts options.
type EditAdminCertsOptions struct {
	// The `id` of the component to edit. Use the [Get all components](#get-all-components) API to determine the id of the
	// component.
	ID *string `json:"id" validate:"required"`

	// The admin certificates to add to the file system.
	AppendAdminCerts []string `json:"append_admin_certs,omitempty"`

	// The admin certificates to remove from the file system. To see the current list run the [Get a
	// component](#get-component-data) API with the `deployment_attrs=included` parameter.
	RemoveAdminCerts []string `json:"remove_admin_certs,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewEditAdminCertsOptions : Instantiate EditAdminCertsOptions
func (*BlockchainV2) NewEditAdminCertsOptions(id string) *EditAdminCertsOptions {
	return &EditAdminCertsOptions{
		ID: core.StringPtr(id),
	}
}

// SetID : Allow user to set ID
func (options *EditAdminCertsOptions) SetID(id string) *EditAdminCertsOptions {
	options.ID = core.StringPtr(id)
	return options
}

// SetAppendAdminCerts : Allow user to set AppendAdminCerts
func (options *EditAdminCertsOptions) SetAppendAdminCerts(appendAdminCerts []string) *EditAdminCertsOptions {
	options.AppendAdminCerts = appendAdminCerts
	return options
}

// SetRemoveAdminCerts : Allow user to set RemoveAdminCerts
func (options *EditAdminCertsOptions) SetRemoveAdminCerts(removeAdminCerts []string) *EditAdminCertsOptions {
	options.RemoveAdminCerts = removeAdminCerts
	return options
}

// SetHeaders : Allow user to set Headers
func (options *EditAdminCertsOptions) SetHeaders(param map[string]string) *EditAdminCertsOptions {
	options.Headers = param
	return options
}

// EditAdminCertsResponse : EditAdminCertsResponse struct
type EditAdminCertsResponse struct {
	// The total number of admin certificate additions and deletions.
	ChangesMade *float64 `json:"changes_made,omitempty"`

	// Array of certs there were set.
	SetAdminCerts []EditAdminCertsResponseSetAdminCertsItem `json:"set_admin_certs,omitempty"`
}


// UnmarshalEditAdminCertsResponse unmarshals an instance of EditAdminCertsResponse from the specified map of raw messages.
func UnmarshalEditAdminCertsResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(EditAdminCertsResponse)
	err = core.UnmarshalPrimitive(m, "changes_made", &obj.ChangesMade)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "set_admin_certs", &obj.SetAdminCerts, UnmarshalEditAdminCertsResponseSetAdminCertsItem)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// EditAdminCertsResponseSetAdminCertsItem : EditAdminCertsResponseSetAdminCertsItem struct
type EditAdminCertsResponseSetAdminCertsItem struct {
	// A certificate as base 64 encoded PEM. Also known as the signing certificate of an organization admin.
	Base64Pem *string `json:"base_64_pem,omitempty"`

	// The issuer string in the certificate.
	Issuer *string `json:"issuer,omitempty"`

	// UTC timestamp of the last ms the certificate is valid.
	NotAfterTs *float64 `json:"not_after_ts,omitempty"`

	// UTC timestamp of the earliest ms the certificate is valid.
	NotBeforeTs *float64 `json:"not_before_ts,omitempty"`

	// The "unique" id of the certificates.
	SerialNumberHex *string `json:"serial_number_hex,omitempty"`

	// The crypto algorithm that signed the public key in the certificate.
	SignatureAlgorithm *string `json:"signature_algorithm,omitempty"`

	// The subject string in the certificate.
	Subject *string `json:"subject,omitempty"`

	// The X.509 version/format.
	X509Version *float64 `json:"X509_version,omitempty"`

	// A friendly (human readable) duration until certificate expiration.
	TimeLeft *string `json:"time_left,omitempty"`
}


// UnmarshalEditAdminCertsResponseSetAdminCertsItem unmarshals an instance of EditAdminCertsResponseSetAdminCertsItem from the specified map of raw messages.
func UnmarshalEditAdminCertsResponseSetAdminCertsItem(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(EditAdminCertsResponseSetAdminCertsItem)
	err = core.UnmarshalPrimitive(m, "base_64_pem", &obj.Base64Pem)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "issuer", &obj.Issuer)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "not_after_ts", &obj.NotAfterTs)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "not_before_ts", &obj.NotBeforeTs)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "serial_number_hex", &obj.SerialNumberHex)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "signature_algorithm", &obj.SignatureAlgorithm)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "subject", &obj.Subject)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "X509_version", &obj.X509Version)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "time_left", &obj.TimeLeft)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// EditCaOptions : The EditCa options.
type EditCaOptions struct {
	// The `id` of the component to modify. Use the [Get all components](#get-all-components) API to determine the
	// component id.
	ID *string `json:"id" validate:"required"`

	// A descriptive name for this CA. The IBP console tile displays this name.
	DisplayName *string `json:"display_name,omitempty"`

	// The URL for the CA. Typically, client applications would send requests to this URL. Include the protocol,
	// hostname/ip and port.
	ApiURL *string `json:"api_url,omitempty"`

	// The operations URL for the CA. Include the protocol, hostname/ip and port.
	OperationsURL *string `json:"operations_url,omitempty"`

	// The CA's "CAName" attribute.
	CaName *string `json:"ca_name,omitempty"`

	// Indicates where the component is running.
	Location *string `json:"location,omitempty"`

	Tags []string `json:"tags,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewEditCaOptions : Instantiate EditCaOptions
func (*BlockchainV2) NewEditCaOptions(id string) *EditCaOptions {
	return &EditCaOptions{
		ID: core.StringPtr(id),
	}
}

// SetID : Allow user to set ID
func (options *EditCaOptions) SetID(id string) *EditCaOptions {
	options.ID = core.StringPtr(id)
	return options
}

// SetDisplayName : Allow user to set DisplayName
func (options *EditCaOptions) SetDisplayName(displayName string) *EditCaOptions {
	options.DisplayName = core.StringPtr(displayName)
	return options
}

// SetApiURL : Allow user to set ApiURL
func (options *EditCaOptions) SetApiURL(apiURL string) *EditCaOptions {
	options.ApiURL = core.StringPtr(apiURL)
	return options
}

// SetOperationsURL : Allow user to set OperationsURL
func (options *EditCaOptions) SetOperationsURL(operationsURL string) *EditCaOptions {
	options.OperationsURL = core.StringPtr(operationsURL)
	return options
}

// SetCaName : Allow user to set CaName
func (options *EditCaOptions) SetCaName(caName string) *EditCaOptions {
	options.CaName = core.StringPtr(caName)
	return options
}

// SetLocation : Allow user to set Location
func (options *EditCaOptions) SetLocation(location string) *EditCaOptions {
	options.Location = core.StringPtr(location)
	return options
}

// SetTags : Allow user to set Tags
func (options *EditCaOptions) SetTags(tags []string) *EditCaOptions {
	options.Tags = tags
	return options
}

// SetHeaders : Allow user to set Headers
func (options *EditCaOptions) SetHeaders(param map[string]string) *EditCaOptions {
	options.Headers = param
	return options
}

// EditLogSettingsBody : File system logging settings. All fields are optional (only send the settings that you want to change). _Changes to
// this field will restart the IBP console server(s)_.
type EditLogSettingsBody struct {
	// The client side (browser) logging settings. _Changes to this field will restart the IBP console server(s)_.
	Client *LoggingSettingsClient `json:"client,omitempty"`

	// The server side logging settings. _Changes to this field will restart the IBP console server(s)_.
	Server *LoggingSettingsServer `json:"server,omitempty"`
}


// UnmarshalEditLogSettingsBody unmarshals an instance of EditLogSettingsBody from the specified map of raw messages.
func UnmarshalEditLogSettingsBody(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(EditLogSettingsBody)
	err = core.UnmarshalModel(m, "client", &obj.Client, UnmarshalLoggingSettingsClient)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "server", &obj.Server, UnmarshalLoggingSettingsServer)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// EditMspOptions : The EditMsp options.
type EditMspOptions struct {
	// The `id` of the component to modify. Use the [Get all components](#get-all-components) API to determine the
	// component id.
	ID *string `json:"id" validate:"required"`

	// The MSP id that is related to this component.
	MspID *string `json:"msp_id,omitempty"`

	// A descriptive name for this MSP. The IBP console tile displays this name.
	DisplayName *string `json:"display_name,omitempty"`

	// An array that contains one or more base 64 encoded PEM root certificates for the MSP.
	RootCerts []string `json:"root_certs,omitempty"`

	// An array that contains base 64 encoded PEM intermediate certificates.
	IntermediateCerts []string `json:"intermediate_certs,omitempty"`

	// An array that contains base 64 encoded PEM identity certificates for administrators. Also known as signing
	// certificates of an organization administrator.
	Admins []string `json:"admins,omitempty"`

	// An array that contains one or more base 64 encoded PEM TLS root certificates.
	TlsRootCerts []string `json:"tls_root_certs,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewEditMspOptions : Instantiate EditMspOptions
func (*BlockchainV2) NewEditMspOptions(id string) *EditMspOptions {
	return &EditMspOptions{
		ID: core.StringPtr(id),
	}
}

// SetID : Allow user to set ID
func (options *EditMspOptions) SetID(id string) *EditMspOptions {
	options.ID = core.StringPtr(id)
	return options
}

// SetMspID : Allow user to set MspID
func (options *EditMspOptions) SetMspID(mspID string) *EditMspOptions {
	options.MspID = core.StringPtr(mspID)
	return options
}

// SetDisplayName : Allow user to set DisplayName
func (options *EditMspOptions) SetDisplayName(displayName string) *EditMspOptions {
	options.DisplayName = core.StringPtr(displayName)
	return options
}

// SetRootCerts : Allow user to set RootCerts
func (options *EditMspOptions) SetRootCerts(rootCerts []string) *EditMspOptions {
	options.RootCerts = rootCerts
	return options
}

// SetIntermediateCerts : Allow user to set IntermediateCerts
func (options *EditMspOptions) SetIntermediateCerts(intermediateCerts []string) *EditMspOptions {
	options.IntermediateCerts = intermediateCerts
	return options
}

// SetAdmins : Allow user to set Admins
func (options *EditMspOptions) SetAdmins(admins []string) *EditMspOptions {
	options.Admins = admins
	return options
}

// SetTlsRootCerts : Allow user to set TlsRootCerts
func (options *EditMspOptions) SetTlsRootCerts(tlsRootCerts []string) *EditMspOptions {
	options.TlsRootCerts = tlsRootCerts
	return options
}

// SetHeaders : Allow user to set Headers
func (options *EditMspOptions) SetHeaders(param map[string]string) *EditMspOptions {
	options.Headers = param
	return options
}

// EditOrdererOptions : The EditOrderer options.
type EditOrdererOptions struct {
	// The `id` of the component to modify. Use the [Get all components](#get-all-components) API to determine the
	// component id.
	ID *string `json:"id" validate:"required"`

	// A descriptive name for an ordering service. The parent IBP console tile displays this name.
	ClusterName *string `json:"cluster_name,omitempty"`

	// A descriptive base name for each ordering node. One or more child IBP console tiles display this name.
	DisplayName *string `json:"display_name,omitempty"`

	// The gRPC URL for the orderer. Typically, client applications would send requests to this URL. Include the protocol,
	// hostname/ip and port.
	ApiURL *string `json:"api_url,omitempty"`

	// Used by Fabric health checker to monitor health status of the node. For more information, see [Fabric
	// documentation](https://hyperledger-fabric.readthedocs.io/en/release-1.4/operations_service.html). Include the
	// protocol, hostname/ip and port.
	OperationsURL *string `json:"operations_url,omitempty"`

	// The gRPC web proxy URL in front of the orderer. Include the protocol, hostname/ip and port.
	GrpcwpURL *string `json:"grpcwp_url,omitempty"`

	// The MSP id that is related to this component.
	MspID *string `json:"msp_id,omitempty"`

	// The state of a pre-created orderer node. A value of `true` means that the orderer node was added as a system channel
	// consenter. This is a manual field. Set it yourself after finishing the raft append flow to indicate that this node
	// is ready for use. See the [Submit config block to orderer](#submit-config-block-to-orderer) API description for more
	// details about appending raft nodes.
	ConsenterProposalFin *bool `json:"consenter_proposal_fin,omitempty"`

	// Indicates where the component is running.
	Location *string `json:"location,omitempty"`

	// The name of the system channel. Defaults to `testchainid`.
	SystemChannelID *string `json:"system_channel_id,omitempty"`

	Tags []string `json:"tags,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewEditOrdererOptions : Instantiate EditOrdererOptions
func (*BlockchainV2) NewEditOrdererOptions(id string) *EditOrdererOptions {
	return &EditOrdererOptions{
		ID: core.StringPtr(id),
	}
}

// SetID : Allow user to set ID
func (options *EditOrdererOptions) SetID(id string) *EditOrdererOptions {
	options.ID = core.StringPtr(id)
	return options
}

// SetClusterName : Allow user to set ClusterName
func (options *EditOrdererOptions) SetClusterName(clusterName string) *EditOrdererOptions {
	options.ClusterName = core.StringPtr(clusterName)
	return options
}

// SetDisplayName : Allow user to set DisplayName
func (options *EditOrdererOptions) SetDisplayName(displayName string) *EditOrdererOptions {
	options.DisplayName = core.StringPtr(displayName)
	return options
}

// SetApiURL : Allow user to set ApiURL
func (options *EditOrdererOptions) SetApiURL(apiURL string) *EditOrdererOptions {
	options.ApiURL = core.StringPtr(apiURL)
	return options
}

// SetOperationsURL : Allow user to set OperationsURL
func (options *EditOrdererOptions) SetOperationsURL(operationsURL string) *EditOrdererOptions {
	options.OperationsURL = core.StringPtr(operationsURL)
	return options
}

// SetGrpcwpURL : Allow user to set GrpcwpURL
func (options *EditOrdererOptions) SetGrpcwpURL(grpcwpURL string) *EditOrdererOptions {
	options.GrpcwpURL = core.StringPtr(grpcwpURL)
	return options
}

// SetMspID : Allow user to set MspID
func (options *EditOrdererOptions) SetMspID(mspID string) *EditOrdererOptions {
	options.MspID = core.StringPtr(mspID)
	return options
}

// SetConsenterProposalFin : Allow user to set ConsenterProposalFin
func (options *EditOrdererOptions) SetConsenterProposalFin(consenterProposalFin bool) *EditOrdererOptions {
	options.ConsenterProposalFin = core.BoolPtr(consenterProposalFin)
	return options
}

// SetLocation : Allow user to set Location
func (options *EditOrdererOptions) SetLocation(location string) *EditOrdererOptions {
	options.Location = core.StringPtr(location)
	return options
}

// SetSystemChannelID : Allow user to set SystemChannelID
func (options *EditOrdererOptions) SetSystemChannelID(systemChannelID string) *EditOrdererOptions {
	options.SystemChannelID = core.StringPtr(systemChannelID)
	return options
}

// SetTags : Allow user to set Tags
func (options *EditOrdererOptions) SetTags(tags []string) *EditOrdererOptions {
	options.Tags = tags
	return options
}

// SetHeaders : Allow user to set Headers
func (options *EditOrdererOptions) SetHeaders(param map[string]string) *EditOrdererOptions {
	options.Headers = param
	return options
}

// EditPeerOptions : The EditPeer options.
type EditPeerOptions struct {
	// The `id` of the component to modify. Use the [Get all components](#get-all-components) API to determine the
	// component id.
	ID *string `json:"id" validate:"required"`

	// A descriptive name for this peer. The IBP console tile displays this name.
	DisplayName *string `json:"display_name,omitempty"`

	// The gRPC URL for the peer. Typically, client applications would send requests to this URL. Include the protocol,
	// hostname/ip and port.
	ApiURL *string `json:"api_url,omitempty"`

	// Used by Fabric health checker to monitor health status of the node. For more information, see [Fabric
	// documentation](https://hyperledger-fabric.readthedocs.io/en/release-1.4/operations_service.html). Include the
	// protocol, hostname/ip and port.
	OperationsURL *string `json:"operations_url,omitempty"`

	// The gRPC web proxy URL in front of the peer. Include the protocol, hostname/ip and port.
	GrpcwpURL *string `json:"grpcwp_url,omitempty"`

	// The MSP id that is related to this component.
	MspID *string `json:"msp_id,omitempty"`

	// Indicates where the component is running.
	Location *string `json:"location,omitempty"`

	Tags []string `json:"tags,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewEditPeerOptions : Instantiate EditPeerOptions
func (*BlockchainV2) NewEditPeerOptions(id string) *EditPeerOptions {
	return &EditPeerOptions{
		ID: core.StringPtr(id),
	}
}

// SetID : Allow user to set ID
func (options *EditPeerOptions) SetID(id string) *EditPeerOptions {
	options.ID = core.StringPtr(id)
	return options
}

// SetDisplayName : Allow user to set DisplayName
func (options *EditPeerOptions) SetDisplayName(displayName string) *EditPeerOptions {
	options.DisplayName = core.StringPtr(displayName)
	return options
}

// SetApiURL : Allow user to set ApiURL
func (options *EditPeerOptions) SetApiURL(apiURL string) *EditPeerOptions {
	options.ApiURL = core.StringPtr(apiURL)
	return options
}

// SetOperationsURL : Allow user to set OperationsURL
func (options *EditPeerOptions) SetOperationsURL(operationsURL string) *EditPeerOptions {
	options.OperationsURL = core.StringPtr(operationsURL)
	return options
}

// SetGrpcwpURL : Allow user to set GrpcwpURL
func (options *EditPeerOptions) SetGrpcwpURL(grpcwpURL string) *EditPeerOptions {
	options.GrpcwpURL = core.StringPtr(grpcwpURL)
	return options
}

// SetMspID : Allow user to set MspID
func (options *EditPeerOptions) SetMspID(mspID string) *EditPeerOptions {
	options.MspID = core.StringPtr(mspID)
	return options
}

// SetLocation : Allow user to set Location
func (options *EditPeerOptions) SetLocation(location string) *EditPeerOptions {
	options.Location = core.StringPtr(location)
	return options
}

// SetTags : Allow user to set Tags
func (options *EditPeerOptions) SetTags(tags []string) *EditPeerOptions {
	options.Tags = tags
	return options
}

// SetHeaders : Allow user to set Headers
func (options *EditPeerOptions) SetHeaders(param map[string]string) *EditPeerOptions {
	options.Headers = param
	return options
}

// EditSettingsBodyInactivityTimeouts : EditSettingsBodyInactivityTimeouts struct
type EditSettingsBodyInactivityTimeouts struct {
	// Indicates if the auto log out logic is enabled or disabled. Defaults `false`. _Refresh browser after changes_.
	Enabled *bool `json:"enabled,omitempty"`

	// Maximum time in milliseconds for a browser client to be idle. Once exceeded the user is logged out. Defaults to
	// `90000` ms (1.5 minutes). _Refresh browser after changes_.
	MaxIdleTime *float64 `json:"max_idle_time,omitempty"`
}


// UnmarshalEditSettingsBodyInactivityTimeouts unmarshals an instance of EditSettingsBodyInactivityTimeouts from the specified map of raw messages.
func UnmarshalEditSettingsBodyInactivityTimeouts(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(EditSettingsBodyInactivityTimeouts)
	err = core.UnmarshalPrimitive(m, "enabled", &obj.Enabled)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "max_idle_time", &obj.MaxIdleTime)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// EditSettingsOptions : The EditSettings options.
type EditSettingsOptions struct {
	InactivityTimeouts *EditSettingsBodyInactivityTimeouts `json:"inactivity_timeouts,omitempty"`

	// File system logging settings. All fields are optional (only send the settings that you want to change). _Changes to
	// this field will restart the IBP console server(s)_.
	FileLogging *EditLogSettingsBody `json:"file_logging,omitempty"`

	// The base limit for the maximum number of `/api/_*` API requests (aka UI requests) in 1 minute. Defaults `25`. [Rate
	// Limits](#rate-limits). _Changes to this field will restart the IBP console server(s)_.
	MaxReqPerMin *float64 `json:"max_req_per_min,omitempty"`

	// The base limit for the maximum number of `/ak/api/_*` API requests (aka external api key requests) in 1 minute.
	// Defaults `25`. [Rate Limits](#rate-limits). _Changes to this field will restart the IBP console server(s)_.
	MaxReqPerMinAk *float64 `json:"max_req_per_min_ak,omitempty"`

	// Maximum time in milliseconds to wait for a get-block transaction. Defaults to `10000` ms (10 seconds). _Refresh
	// browser after changes_.
	FabricGetBlockTimeoutMs *float64 `json:"fabric_get_block_timeout_ms,omitempty"`

	// Maximum time in milliseconds to wait for a instantiate chaincode transaction. Defaults to `300000` ms (5 minutes).
	// _Refresh browser after changes_.
	FabricInstantiateTimeoutMs *float64 `json:"fabric_instantiate_timeout_ms,omitempty"`

	// Maximum time in milliseconds to wait for a join-channel transaction. Defaults to `25000` ms (25 seconds). _Refresh
	// browser after changes_.
	FabricJoinChannelTimeoutMs *float64 `json:"fabric_join_channel_timeout_ms,omitempty"`

	// Maximum time in milliseconds to wait for a install chaincode transaction (Fabric v1.x). Defaults to `300000` ms (5
	// minutes). _Refresh browser after changes_.
	FabricInstallCcTimeoutMs *float64 `json:"fabric_install_cc_timeout_ms,omitempty"`

	// Maximum time in milliseconds to wait for a install chaincode transaction (Fabric v2.x). Defaults to `300000` ms (5
	// minutes). _Refresh browser after changes_.
	FabricLcInstallCcTimeoutMs *float64 `json:"fabric_lc_install_cc_timeout_ms,omitempty"`

	// Maximum time in milliseconds to wait for a get-chaincode transaction (Fabric v2.x). Defaults to `180000` ms (3
	// minutes). _Refresh browser after changes_.
	FabricLcGetCcTimeoutMs *float64 `json:"fabric_lc_get_cc_timeout_ms,omitempty"`

	// Default maximum time in milliseconds to wait for a Fabric transaction. Defaults to `10000` ms (10 seconds). _Refresh
	// browser after changes_.
	FabricGeneralTimeoutMs *float64 `json:"fabric_general_timeout_ms,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewEditSettingsOptions : Instantiate EditSettingsOptions
func (*BlockchainV2) NewEditSettingsOptions() *EditSettingsOptions {
	return &EditSettingsOptions{}
}

// SetInactivityTimeouts : Allow user to set InactivityTimeouts
func (options *EditSettingsOptions) SetInactivityTimeouts(inactivityTimeouts *EditSettingsBodyInactivityTimeouts) *EditSettingsOptions {
	options.InactivityTimeouts = inactivityTimeouts
	return options
}

// SetFileLogging : Allow user to set FileLogging
func (options *EditSettingsOptions) SetFileLogging(fileLogging *EditLogSettingsBody) *EditSettingsOptions {
	options.FileLogging = fileLogging
	return options
}

// SetMaxReqPerMin : Allow user to set MaxReqPerMin
func (options *EditSettingsOptions) SetMaxReqPerMin(maxReqPerMin float64) *EditSettingsOptions {
	options.MaxReqPerMin = core.Float64Ptr(maxReqPerMin)
	return options
}

// SetMaxReqPerMinAk : Allow user to set MaxReqPerMinAk
func (options *EditSettingsOptions) SetMaxReqPerMinAk(maxReqPerMinAk float64) *EditSettingsOptions {
	options.MaxReqPerMinAk = core.Float64Ptr(maxReqPerMinAk)
	return options
}

// SetFabricGetBlockTimeoutMs : Allow user to set FabricGetBlockTimeoutMs
func (options *EditSettingsOptions) SetFabricGetBlockTimeoutMs(fabricGetBlockTimeoutMs float64) *EditSettingsOptions {
	options.FabricGetBlockTimeoutMs = core.Float64Ptr(fabricGetBlockTimeoutMs)
	return options
}

// SetFabricInstantiateTimeoutMs : Allow user to set FabricInstantiateTimeoutMs
func (options *EditSettingsOptions) SetFabricInstantiateTimeoutMs(fabricInstantiateTimeoutMs float64) *EditSettingsOptions {
	options.FabricInstantiateTimeoutMs = core.Float64Ptr(fabricInstantiateTimeoutMs)
	return options
}

// SetFabricJoinChannelTimeoutMs : Allow user to set FabricJoinChannelTimeoutMs
func (options *EditSettingsOptions) SetFabricJoinChannelTimeoutMs(fabricJoinChannelTimeoutMs float64) *EditSettingsOptions {
	options.FabricJoinChannelTimeoutMs = core.Float64Ptr(fabricJoinChannelTimeoutMs)
	return options
}

// SetFabricInstallCcTimeoutMs : Allow user to set FabricInstallCcTimeoutMs
func (options *EditSettingsOptions) SetFabricInstallCcTimeoutMs(fabricInstallCcTimeoutMs float64) *EditSettingsOptions {
	options.FabricInstallCcTimeoutMs = core.Float64Ptr(fabricInstallCcTimeoutMs)
	return options
}

// SetFabricLcInstallCcTimeoutMs : Allow user to set FabricLcInstallCcTimeoutMs
func (options *EditSettingsOptions) SetFabricLcInstallCcTimeoutMs(fabricLcInstallCcTimeoutMs float64) *EditSettingsOptions {
	options.FabricLcInstallCcTimeoutMs = core.Float64Ptr(fabricLcInstallCcTimeoutMs)
	return options
}

// SetFabricLcGetCcTimeoutMs : Allow user to set FabricLcGetCcTimeoutMs
func (options *EditSettingsOptions) SetFabricLcGetCcTimeoutMs(fabricLcGetCcTimeoutMs float64) *EditSettingsOptions {
	options.FabricLcGetCcTimeoutMs = core.Float64Ptr(fabricLcGetCcTimeoutMs)
	return options
}

// SetFabricGeneralTimeoutMs : Allow user to set FabricGeneralTimeoutMs
func (options *EditSettingsOptions) SetFabricGeneralTimeoutMs(fabricGeneralTimeoutMs float64) *EditSettingsOptions {
	options.FabricGeneralTimeoutMs = core.Float64Ptr(fabricGeneralTimeoutMs)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *EditSettingsOptions) SetHeaders(param map[string]string) *EditSettingsOptions {
	options.Headers = param
	return options
}

// FabVersionObject : FabVersionObject struct
type FabVersionObject struct {
	// Indicates if this is the Fabric version that will be used if none is selected.
	Default *bool `json:"default,omitempty"`

	// The Fabric version.
	Version *string `json:"version,omitempty"`

	// Detailed image information for this Fabric release.
	Image interface{} `json:"image,omitempty"`
}


// UnmarshalFabVersionObject unmarshals an instance of FabVersionObject from the specified map of raw messages.
func UnmarshalFabVersionObject(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(FabVersionObject)
	err = core.UnmarshalPrimitive(m, "default", &obj.Default)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "version", &obj.Version)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "image", &obj.Image)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// FabricVersionDictionary : A supported release of Fabric for this component type.
type FabricVersionDictionary struct {
	X1462 *FabVersionObject `json:"1.4.6-2,omitempty"`

	X2100 *FabVersionObject `json:"2.1.0-0,omitempty"`

	// Allows users to set arbitrary properties
	additionalProperties map[string]interface{}
}


// SetProperty allows the user to set an arbitrary property on an instance of FabricVersionDictionary
func (o *FabricVersionDictionary) SetProperty(key string, value interface{}) {
	if o.additionalProperties == nil {
		o.additionalProperties = make(map[string]interface{})
	}
	o.additionalProperties[key] = value
}

// GetProperty allows the user to retrieve an arbitrary property from an instance of FabricVersionDictionary
func (o *FabricVersionDictionary) GetProperty(key string) interface{} {
	return o.additionalProperties[key]
}

// GetProperties allows the user to retrieve the map of arbitrary properties from an instance of FabricVersionDictionary
func (o *FabricVersionDictionary) GetProperties() map[string]interface{} {
	return o.additionalProperties
}

// MarshalJSON performs custom serialization for instances of FabricVersionDictionary
func (o *FabricVersionDictionary) MarshalJSON() (buffer []byte, err error) {
	m := make(map[string]interface{})
	if len(o.additionalProperties) > 0 {
		for k, v := range o.additionalProperties {
			m[k] = v
		}
	}
	if o.X1462 != nil {
		m["1.4.6-2"] = o.X1462
	}
	if o.X2100 != nil {
		m["2.1.0-0"] = o.X2100
	}
	buffer, err = json.Marshal(m)
	return
}

// UnmarshalFabricVersionDictionary unmarshals an instance of FabricVersionDictionary from the specified map of raw messages.
func UnmarshalFabricVersionDictionary(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(FabricVersionDictionary)
	err = core.UnmarshalModel(m, "1.4.6-2", &obj.X1462, UnmarshalFabVersionObject)
	if err != nil {
		return
	}
	delete(m, "1.4.6-2")
	err = core.UnmarshalModel(m, "2.1.0-0", &obj.X2100, UnmarshalFabVersionObject)
	if err != nil {
		return
	}
	delete(m, "2.1.0-0")
	for k := range m {
		var v interface{}
		e := core.UnmarshalPrimitive(m, k, &v)
		if e != nil {
			err = e
			return
		}
		obj.SetProperty(k, v)
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// GenericComponentResponse : Contains the details of a component. Not all components have the same fields, see description of each field for
// details.
type GenericComponentResponse struct {
	// The unique identifier of this component. [Available on all component types].
	ID *string `json:"id,omitempty"`

	// The type of this component [Available on all component types].
	Type *string `json:"type,omitempty"`

	// The displayed name of this component. [Available on all component types].
	DisplayName *string `json:"display_name,omitempty"`

	// The URL for the grpc web proxy for this component. [Available on peer/orderer components].
	GrpcwpURL *string `json:"grpcwp_url,omitempty"`

	// The gRPC URL for the component. Typically, client applications would send requests to this URL. [Available on
	// ca/peer/orderer components].
	ApiURL *string `json:"api_url,omitempty"`

	// Used by Fabric health checker to monitor health status of the node. For more information, see [Fabric
	// documentation](https://hyperledger-fabric.readthedocs.io/en/release-1.4/operations_service.html). [Available on
	// ca/peer/orderer components].
	OperationsURL *string `json:"operations_url,omitempty"`

	// The MSP id that is related to this component. [Available on all components].
	MspID *string `json:"msp_id,omitempty"`

	// Indicates where the component is running.
	Location *string `json:"location,omitempty"`

	// The CA's "CAName" attribute. [Available on ca components].
	CaName *string `json:"ca_name,omitempty"`

	// The Fabric "admin certificates" for the component. [Available on peer/orderer components w/query parameter
	// 'deployment_attrs'].
	AdminCerts []string `json:"admin_certs,omitempty"`

	// The type of ledger database for a peer. [Available on peer components w/query parameter 'deployment_attrs'].
	StateDb *string `json:"state_db,omitempty"`

	// UNIX timestamp of component creation, UTC, ms. [Available on all components].
	Timestamp *float64 `json:"timestamp,omitempty"`

	// The **cached** Kubernetes resource attributes for this component. [Available on ca/peer/orderer components w/query
	// parameter 'deployment_attrs'].
	Resources *GenericComponentResponseResources `json:"resources,omitempty"`

	// The versioning of the IBP console format of this JSON.
	SchemeVersion *string `json:"scheme_version,omitempty"`

	// The **cached** Kubernetes storage attributes for this component. [Available on ca/peer/orderer components w/query
	// parameter 'deployment_attrs'].
	Storage *GenericComponentResponseStorage `json:"storage,omitempty"`

	Tags []string `json:"tags,omitempty"`

	// The TLS certificate as base 64 encoded PEM. Certificate is used to secure/validate a TLS connection.
	TlsCert *string `json:"tls_cert,omitempty"`

	// The cached Hyperledger Fabric version for this component. [Available on ca/peer/orderer components w/query parameter
	// 'deployment_attrs'].
	Version *string `json:"version,omitempty"`

	// The Kubernetes zone of this component's deployment. [Available on ca/peer/orderer components w/query parameter
	// 'deployment_attrs'].
	Zone *string `json:"zone,omitempty"`
}

// Constants associated with the GenericComponentResponse.Type property.
// The type of this component [Available on all component types].
const (
	GenericComponentResponse_Type_FabricCa = "fabric-ca"
	GenericComponentResponse_Type_FabricOrderer = "fabric-orderer"
	GenericComponentResponse_Type_FabricPeer = "fabric-peer"
)


// UnmarshalGenericComponentResponse unmarshals an instance of GenericComponentResponse from the specified map of raw messages.
func UnmarshalGenericComponentResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(GenericComponentResponse)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "display_name", &obj.DisplayName)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "grpcwp_url", &obj.GrpcwpURL)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "api_url", &obj.ApiURL)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "operations_url", &obj.OperationsURL)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "msp_id", &obj.MspID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "location", &obj.Location)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "ca_name", &obj.CaName)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "admin_certs", &obj.AdminCerts)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "state_db", &obj.StateDb)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "timestamp", &obj.Timestamp)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "resources", &obj.Resources, UnmarshalGenericComponentResponseResources)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "scheme_version", &obj.SchemeVersion)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "storage", &obj.Storage, UnmarshalGenericComponentResponseStorage)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "tags", &obj.Tags)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "tls_cert", &obj.TlsCert)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "version", &obj.Version)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "zone", &obj.Zone)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// GenericComponentResponseResources : The **cached** Kubernetes resource attributes for this component. [Available on ca/peer/orderer components w/query
// parameter 'deployment_attrs'].
type GenericComponentResponseResources struct {
	Ca *GenericResources `json:"ca,omitempty"`

	Peer *GenericResources `json:"peer,omitempty"`

	Orderer *GenericResources `json:"orderer,omitempty"`

	Proxy *GenericResources `json:"proxy,omitempty"`

	Couchdb *GenericResources `json:"couchdb,omitempty"`
}


// UnmarshalGenericComponentResponseResources unmarshals an instance of GenericComponentResponseResources from the specified map of raw messages.
func UnmarshalGenericComponentResponseResources(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(GenericComponentResponseResources)
	err = core.UnmarshalModel(m, "ca", &obj.Ca, UnmarshalGenericResources)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "peer", &obj.Peer, UnmarshalGenericResources)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "orderer", &obj.Orderer, UnmarshalGenericResources)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "proxy", &obj.Proxy, UnmarshalGenericResources)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "couchdb", &obj.Couchdb, UnmarshalGenericResources)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// GenericComponentResponseStorage : The **cached** Kubernetes storage attributes for this component. [Available on ca/peer/orderer components w/query
// parameter 'deployment_attrs'].
type GenericComponentResponseStorage struct {
	Ca *StorageObject `json:"ca,omitempty"`

	Peer *StorageObject `json:"peer,omitempty"`

	Orderer *StorageObject `json:"orderer,omitempty"`

	Couchdb *StorageObject `json:"couchdb,omitempty"`
}


// UnmarshalGenericComponentResponseStorage unmarshals an instance of GenericComponentResponseStorage from the specified map of raw messages.
func UnmarshalGenericComponentResponseStorage(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(GenericComponentResponseStorage)
	err = core.UnmarshalModel(m, "ca", &obj.Ca, UnmarshalStorageObject)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "peer", &obj.Peer, UnmarshalStorageObject)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "orderer", &obj.Orderer, UnmarshalStorageObject)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "couchdb", &obj.Couchdb, UnmarshalStorageObject)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// GenericResourceLimits : GenericResourceLimits struct
type GenericResourceLimits struct {
	Cpu *string `json:"cpu,omitempty"`

	Memory *string `json:"memory,omitempty"`
}


// UnmarshalGenericResourceLimits unmarshals an instance of GenericResourceLimits from the specified map of raw messages.
func UnmarshalGenericResourceLimits(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(GenericResourceLimits)
	err = core.UnmarshalPrimitive(m, "cpu", &obj.Cpu)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "memory", &obj.Memory)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// GenericResources : GenericResources struct
type GenericResources struct {
	Requests *GenericResourcesRequests `json:"requests,omitempty"`

	Limits *GenericResourceLimits `json:"limits,omitempty"`
}


// UnmarshalGenericResources unmarshals an instance of GenericResources from the specified map of raw messages.
func UnmarshalGenericResources(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(GenericResources)
	err = core.UnmarshalModel(m, "requests", &obj.Requests, UnmarshalGenericResourcesRequests)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "limits", &obj.Limits, UnmarshalGenericResourceLimits)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// GenericResourcesRequests : GenericResourcesRequests struct
type GenericResourcesRequests struct {
	Cpu *string `json:"cpu,omitempty"`

	Memory *string `json:"memory,omitempty"`
}


// UnmarshalGenericResourcesRequests unmarshals an instance of GenericResourcesRequests from the specified map of raw messages.
func UnmarshalGenericResourcesRequests(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(GenericResourcesRequests)
	err = core.UnmarshalPrimitive(m, "cpu", &obj.Cpu)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "memory", &obj.Memory)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// GetAthenaHealthStatsResponse : Contains various health statistics like up time and cache sizes.
type GetAthenaHealthStatsResponse struct {
	OPTOOLS *GetAthenaHealthStatsResponseOPTOOLS `json:"OPTOOLS,omitempty"`

	OS *GetAthenaHealthStatsResponseOS `json:"OS,omitempty"`
}


// UnmarshalGetAthenaHealthStatsResponse unmarshals an instance of GetAthenaHealthStatsResponse from the specified map of raw messages.
func UnmarshalGetAthenaHealthStatsResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(GetAthenaHealthStatsResponse)
	err = core.UnmarshalModel(m, "OPTOOLS", &obj.OPTOOLS, UnmarshalGetAthenaHealthStatsResponseOPTOOLS)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "OS", &obj.OS, UnmarshalGetAthenaHealthStatsResponseOS)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// GetAthenaHealthStatsResponseOPTOOLS : GetAthenaHealthStatsResponseOPTOOLS struct
type GetAthenaHealthStatsResponseOPTOOLS struct {
	// Random/unique id for a process running IBP console.
	InstanceID *string `json:"instance_id,omitempty"`

	// UTC UNIX timestamp of the current time according to the server. In milliseconds.
	Now *float64 `json:"now,omitempty"`

	// UTC UNIX timestamp of when the server started. In milliseconds.
	Born *float64 `json:"born,omitempty"`

	// Total time the IBP console server has been running.
	UpTime *string `json:"up_time,omitempty"`

	MemoryUsage *GetAthenaHealthStatsResponseOPTOOLSMemoryUsage `json:"memory_usage,omitempty"`

	SessionCacheStats *CacheData `json:"session_cache_stats,omitempty"`

	CouchCacheStats *CacheData `json:"couch_cache_stats,omitempty"`

	IamCacheStats *CacheData `json:"iam_cache_stats,omitempty"`

	ProxyCache *CacheData `json:"proxy_cache,omitempty"`
}


// UnmarshalGetAthenaHealthStatsResponseOPTOOLS unmarshals an instance of GetAthenaHealthStatsResponseOPTOOLS from the specified map of raw messages.
func UnmarshalGetAthenaHealthStatsResponseOPTOOLS(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(GetAthenaHealthStatsResponseOPTOOLS)
	err = core.UnmarshalPrimitive(m, "instance_id", &obj.InstanceID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "now", &obj.Now)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "born", &obj.Born)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "up_time", &obj.UpTime)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "memory_usage", &obj.MemoryUsage, UnmarshalGetAthenaHealthStatsResponseOPTOOLSMemoryUsage)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "session_cache_stats", &obj.SessionCacheStats, UnmarshalCacheData)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "couch_cache_stats", &obj.CouchCacheStats, UnmarshalCacheData)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "iam_cache_stats", &obj.IamCacheStats, UnmarshalCacheData)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "proxy_cache", &obj.ProxyCache, UnmarshalCacheData)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// GetAthenaHealthStatsResponseOPTOOLSMemoryUsage : GetAthenaHealthStatsResponseOPTOOLSMemoryUsage struct
type GetAthenaHealthStatsResponseOPTOOLSMemoryUsage struct {
	// Resident set size - total memory allocated for the process.
	Rss *string `json:"rss,omitempty"`

	// Memory allocated for the heap of V8.
	HeapTotal *string `json:"heapTotal,omitempty"`

	// Current heap used by V8.
	HeapUsed *string `json:"heapUsed,omitempty"`

	// Memory used by bound C++ objects.
	External *string `json:"external,omitempty"`
}


// UnmarshalGetAthenaHealthStatsResponseOPTOOLSMemoryUsage unmarshals an instance of GetAthenaHealthStatsResponseOPTOOLSMemoryUsage from the specified map of raw messages.
func UnmarshalGetAthenaHealthStatsResponseOPTOOLSMemoryUsage(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(GetAthenaHealthStatsResponseOPTOOLSMemoryUsage)
	err = core.UnmarshalPrimitive(m, "rss", &obj.Rss)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "heapTotal", &obj.HeapTotal)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "heapUsed", &obj.HeapUsed)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "external", &obj.External)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// GetAthenaHealthStatsResponseOS : GetAthenaHealthStatsResponseOS struct
type GetAthenaHealthStatsResponseOS struct {
	// CPU architecture.
	Arch *string `json:"arch,omitempty"`

	// Operating system name.
	Type *string `json:"type,omitempty"`

	// Endianness of the CPU. LE = Little Endian, BE = Big Endian.
	Endian *string `json:"endian,omitempty"`

	// CPU load in 1, 5, & 15 minute averages. n/a on windows.
	Loadavg *string `json:"loadavg,omitempty"`

	Cpus []CpuHealthStats `json:"cpus,omitempty"`

	// Total memory known to the operating system.
	TotalMemory *string `json:"total_memory,omitempty"`

	// Free memory on the operating system.
	FreeMemory *string `json:"free_memory,omitempty"`

	// Time operating system has been running.
	UpTime *string `json:"up_time,omitempty"`
}


// UnmarshalGetAthenaHealthStatsResponseOS unmarshals an instance of GetAthenaHealthStatsResponseOS from the specified map of raw messages.
func UnmarshalGetAthenaHealthStatsResponseOS(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(GetAthenaHealthStatsResponseOS)
	err = core.UnmarshalPrimitive(m, "arch", &obj.Arch)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "endian", &obj.Endian)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "loadavg", &obj.Loadavg)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "cpus", &obj.Cpus, UnmarshalCpuHealthStats)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "total_memory", &obj.TotalMemory)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "free_memory", &obj.FreeMemory)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "up_time", &obj.UpTime)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// GetComponentByTagOptions : The GetComponentByTag options.
type GetComponentByTagOptions struct {
	// The tag to filter components on. Not case-sensitive.
	Tag *string `json:"tag" validate:"required"`

	// Set to 'included' if the response should include Kubernetes deployment attributes such as 'resources', 'storage',
	// 'zone', 'region', 'admin_certs', etc. Default responses will not include these fields.
	//
	// **This parameter will not work on *imported* components.**
	//
	// It's recommended to use `cache=skip` as well if up-to-date deployment data is needed.
	DeploymentAttrs *string `json:"deployment_attrs,omitempty"`

	// Set to 'included' if the response should include parsed PEM data along with base 64 encoded PEM string. Parsed
	// certificate data will include fields such as the serial number, issuer, expiration, subject, subject alt names, etc.
	// Default responses will not include these fields.
	ParsedCerts *string `json:"parsed_certs,omitempty"`

	// Set to 'skip' if the response should skip local data and fetch live data wherever possible. Expect longer response
	// times if the cache is skipped. Default responses will use the cache.
	Cache *string `json:"cache,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// Constants associated with the GetComponentByTagOptions.DeploymentAttrs property.
// Set to 'included' if the response should include Kubernetes deployment attributes such as 'resources', 'storage',
// 'zone', 'region', 'admin_certs', etc. Default responses will not include these fields.
//
// **This parameter will not work on *imported* components.**
//
// It's recommended to use `cache=skip` as well if up-to-date deployment data is needed.
const (
	GetComponentByTagOptions_DeploymentAttrs_Included = "included"
	GetComponentByTagOptions_DeploymentAttrs_Omitted = "omitted"
)

// Constants associated with the GetComponentByTagOptions.ParsedCerts property.
// Set to 'included' if the response should include parsed PEM data along with base 64 encoded PEM string. Parsed
// certificate data will include fields such as the serial number, issuer, expiration, subject, subject alt names, etc.
// Default responses will not include these fields.
const (
	GetComponentByTagOptions_ParsedCerts_Included = "included"
	GetComponentByTagOptions_ParsedCerts_Omitted = "omitted"
)

// Constants associated with the GetComponentByTagOptions.Cache property.
// Set to 'skip' if the response should skip local data and fetch live data wherever possible. Expect longer response
// times if the cache is skipped. Default responses will use the cache.
const (
	GetComponentByTagOptions_Cache_Skip = "skip"
	GetComponentByTagOptions_Cache_Use = "use"
)

// NewGetComponentByTagOptions : Instantiate GetComponentByTagOptions
func (*BlockchainV2) NewGetComponentByTagOptions(tag string) *GetComponentByTagOptions {
	return &GetComponentByTagOptions{
		Tag: core.StringPtr(tag),
	}
}

// SetTag : Allow user to set Tag
func (options *GetComponentByTagOptions) SetTag(tag string) *GetComponentByTagOptions {
	options.Tag = core.StringPtr(tag)
	return options
}

// SetDeploymentAttrs : Allow user to set DeploymentAttrs
func (options *GetComponentByTagOptions) SetDeploymentAttrs(deploymentAttrs string) *GetComponentByTagOptions {
	options.DeploymentAttrs = core.StringPtr(deploymentAttrs)
	return options
}

// SetParsedCerts : Allow user to set ParsedCerts
func (options *GetComponentByTagOptions) SetParsedCerts(parsedCerts string) *GetComponentByTagOptions {
	options.ParsedCerts = core.StringPtr(parsedCerts)
	return options
}

// SetCache : Allow user to set Cache
func (options *GetComponentByTagOptions) SetCache(cache string) *GetComponentByTagOptions {
	options.Cache = core.StringPtr(cache)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *GetComponentByTagOptions) SetHeaders(param map[string]string) *GetComponentByTagOptions {
	options.Headers = param
	return options
}

// GetComponentOptions : The GetComponent options.
type GetComponentOptions struct {
	// The `id` of the component to retrieve. Use the [Get all components](#get-all-components) API to determine the
	// component id.
	ID *string `json:"id" validate:"required"`

	// Set to 'included' if the response should include Kubernetes deployment attributes such as 'resources', 'storage',
	// 'zone', 'region', 'admin_certs', etc. Default responses will not include these fields.
	//
	// **This parameter will not work on *imported* components.**
	//
	// It's recommended to use `cache=skip` as well if up-to-date deployment data is needed.
	DeploymentAttrs *string `json:"deployment_attrs,omitempty"`

	// Set to 'included' if the response should include parsed PEM data along with base 64 encoded PEM string. Parsed
	// certificate data will include fields such as the serial number, issuer, expiration, subject, subject alt names, etc.
	// Default responses will not include these fields.
	ParsedCerts *string `json:"parsed_certs,omitempty"`

	// Set to 'skip' if the response should skip local data and fetch live data wherever possible. Expect longer response
	// times if the cache is skipped. Default responses will use the cache.
	Cache *string `json:"cache,omitempty"`

	// Set to 'included' if the response should fetch CA attributes, inspect certificates, and append extra fields to CA
	// and MSP component responses.
	// - CA components will have fields appended/updated with data fetched from the `/cainfo?ca=ca` endpoint of a CA, such
	// as: `ca_name`, `root_cert`, `fabric_version`, `issuer_public_key` and `issued_known_msps`. The field
	// `issued_known_msps` indicates imported IBP MSPs that this CA has issued. Meaning the MSP's root cert contains a
	// signature that is derived from this CA's root cert. Only imported MSPs are checked. Default responses will not
	// include these fields.
	// - MSP components will have the field `issued_by_ca_id` appended. This field indicates the id of an IBP console CA
	// that issued this MSP. Meaning the MSP's root cert contains a signature that is derived from this CA's root cert.
	// Only imported/created CAs are checked. Default responses will not include these fields.
	CaAttrs *string `json:"ca_attrs,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// Constants associated with the GetComponentOptions.DeploymentAttrs property.
// Set to 'included' if the response should include Kubernetes deployment attributes such as 'resources', 'storage',
// 'zone', 'region', 'admin_certs', etc. Default responses will not include these fields.
//
// **This parameter will not work on *imported* components.**
//
// It's recommended to use `cache=skip` as well if up-to-date deployment data is needed.
const (
	GetComponentOptions_DeploymentAttrs_Included = "included"
	GetComponentOptions_DeploymentAttrs_Omitted = "omitted"
)

// Constants associated with the GetComponentOptions.ParsedCerts property.
// Set to 'included' if the response should include parsed PEM data along with base 64 encoded PEM string. Parsed
// certificate data will include fields such as the serial number, issuer, expiration, subject, subject alt names, etc.
// Default responses will not include these fields.
const (
	GetComponentOptions_ParsedCerts_Included = "included"
	GetComponentOptions_ParsedCerts_Omitted = "omitted"
)

// Constants associated with the GetComponentOptions.Cache property.
// Set to 'skip' if the response should skip local data and fetch live data wherever possible. Expect longer response
// times if the cache is skipped. Default responses will use the cache.
const (
	GetComponentOptions_Cache_Skip = "skip"
	GetComponentOptions_Cache_Use = "use"
)

// Constants associated with the GetComponentOptions.CaAttrs property.
// Set to 'included' if the response should fetch CA attributes, inspect certificates, and append extra fields to CA and
// MSP component responses.
// - CA components will have fields appended/updated with data fetched from the `/cainfo?ca=ca` endpoint of a CA, such
// as: `ca_name`, `root_cert`, `fabric_version`, `issuer_public_key` and `issued_known_msps`. The field
// `issued_known_msps` indicates imported IBP MSPs that this CA has issued. Meaning the MSP's root cert contains a
// signature that is derived from this CA's root cert. Only imported MSPs are checked. Default responses will not
// include these fields.
// - MSP components will have the field `issued_by_ca_id` appended. This field indicates the id of an IBP console CA
// that issued this MSP. Meaning the MSP's root cert contains a signature that is derived from this CA's root cert. Only
// imported/created CAs are checked. Default responses will not include these fields.
const (
	GetComponentOptions_CaAttrs_Included = "included"
	GetComponentOptions_CaAttrs_Omitted = "omitted"
)

// NewGetComponentOptions : Instantiate GetComponentOptions
func (*BlockchainV2) NewGetComponentOptions(id string) *GetComponentOptions {
	return &GetComponentOptions{
		ID: core.StringPtr(id),
	}
}

// SetID : Allow user to set ID
func (options *GetComponentOptions) SetID(id string) *GetComponentOptions {
	options.ID = core.StringPtr(id)
	return options
}

// SetDeploymentAttrs : Allow user to set DeploymentAttrs
func (options *GetComponentOptions) SetDeploymentAttrs(deploymentAttrs string) *GetComponentOptions {
	options.DeploymentAttrs = core.StringPtr(deploymentAttrs)
	return options
}

// SetParsedCerts : Allow user to set ParsedCerts
func (options *GetComponentOptions) SetParsedCerts(parsedCerts string) *GetComponentOptions {
	options.ParsedCerts = core.StringPtr(parsedCerts)
	return options
}

// SetCache : Allow user to set Cache
func (options *GetComponentOptions) SetCache(cache string) *GetComponentOptions {
	options.Cache = core.StringPtr(cache)
	return options
}

// SetCaAttrs : Allow user to set CaAttrs
func (options *GetComponentOptions) SetCaAttrs(caAttrs string) *GetComponentOptions {
	options.CaAttrs = core.StringPtr(caAttrs)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *GetComponentOptions) SetHeaders(param map[string]string) *GetComponentOptions {
	options.Headers = param
	return options
}

// GetComponentsByTypeOptions : The GetComponentsByType options.
type GetComponentsByTypeOptions struct {
	// The type to filter components on.
	ComponentType *string `json:"component-type" validate:"required"`

	// Set to 'included' if the response should include Kubernetes deployment attributes such as 'resources', 'storage',
	// 'zone', 'region', 'admin_certs', etc. Default responses will not include these fields.
	//
	// **This parameter will not work on *imported* components.**
	//
	// It's recommended to use `cache=skip` as well if up-to-date deployment data is needed.
	DeploymentAttrs *string `json:"deployment_attrs,omitempty"`

	// Set to 'included' if the response should include parsed PEM data along with base 64 encoded PEM string. Parsed
	// certificate data will include fields such as the serial number, issuer, expiration, subject, subject alt names, etc.
	// Default responses will not include these fields.
	ParsedCerts *string `json:"parsed_certs,omitempty"`

	// Set to 'skip' if the response should skip local data and fetch live data wherever possible. Expect longer response
	// times if the cache is skipped. Default responses will use the cache.
	Cache *string `json:"cache,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// Constants associated with the GetComponentsByTypeOptions.ComponentType property.
// The type to filter components on.
const (
	GetComponentsByTypeOptions_ComponentType_FabricCa = "fabric-ca"
	GetComponentsByTypeOptions_ComponentType_FabricOrderer = "fabric-orderer"
	GetComponentsByTypeOptions_ComponentType_FabricPeer = "fabric-peer"
	GetComponentsByTypeOptions_ComponentType_Msp = "msp"
)

// Constants associated with the GetComponentsByTypeOptions.DeploymentAttrs property.
// Set to 'included' if the response should include Kubernetes deployment attributes such as 'resources', 'storage',
// 'zone', 'region', 'admin_certs', etc. Default responses will not include these fields.
//
// **This parameter will not work on *imported* components.**
//
// It's recommended to use `cache=skip` as well if up-to-date deployment data is needed.
const (
	GetComponentsByTypeOptions_DeploymentAttrs_Included = "included"
	GetComponentsByTypeOptions_DeploymentAttrs_Omitted = "omitted"
)

// Constants associated with the GetComponentsByTypeOptions.ParsedCerts property.
// Set to 'included' if the response should include parsed PEM data along with base 64 encoded PEM string. Parsed
// certificate data will include fields such as the serial number, issuer, expiration, subject, subject alt names, etc.
// Default responses will not include these fields.
const (
	GetComponentsByTypeOptions_ParsedCerts_Included = "included"
	GetComponentsByTypeOptions_ParsedCerts_Omitted = "omitted"
)

// Constants associated with the GetComponentsByTypeOptions.Cache property.
// Set to 'skip' if the response should skip local data and fetch live data wherever possible. Expect longer response
// times if the cache is skipped. Default responses will use the cache.
const (
	GetComponentsByTypeOptions_Cache_Skip = "skip"
	GetComponentsByTypeOptions_Cache_Use = "use"
)

// NewGetComponentsByTypeOptions : Instantiate GetComponentsByTypeOptions
func (*BlockchainV2) NewGetComponentsByTypeOptions(componentType string) *GetComponentsByTypeOptions {
	return &GetComponentsByTypeOptions{
		ComponentType: core.StringPtr(componentType),
	}
}

// SetComponentType : Allow user to set ComponentType
func (options *GetComponentsByTypeOptions) SetComponentType(componentType string) *GetComponentsByTypeOptions {
	options.ComponentType = core.StringPtr(componentType)
	return options
}

// SetDeploymentAttrs : Allow user to set DeploymentAttrs
func (options *GetComponentsByTypeOptions) SetDeploymentAttrs(deploymentAttrs string) *GetComponentsByTypeOptions {
	options.DeploymentAttrs = core.StringPtr(deploymentAttrs)
	return options
}

// SetParsedCerts : Allow user to set ParsedCerts
func (options *GetComponentsByTypeOptions) SetParsedCerts(parsedCerts string) *GetComponentsByTypeOptions {
	options.ParsedCerts = core.StringPtr(parsedCerts)
	return options
}

// SetCache : Allow user to set Cache
func (options *GetComponentsByTypeOptions) SetCache(cache string) *GetComponentsByTypeOptions {
	options.Cache = core.StringPtr(cache)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *GetComponentsByTypeOptions) SetHeaders(param map[string]string) *GetComponentsByTypeOptions {
	options.Headers = param
	return options
}

// GetFabVersionsOptions : The GetFabVersions options.
type GetFabVersionsOptions struct {
	// Set to 'skip' if the response should skip local data and fetch live data wherever possible. Expect longer response
	// times if the cache is skipped. Default responses will use the cache.
	Cache *string `json:"cache,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// Constants associated with the GetFabVersionsOptions.Cache property.
// Set to 'skip' if the response should skip local data and fetch live data wherever possible. Expect longer response
// times if the cache is skipped. Default responses will use the cache.
const (
	GetFabVersionsOptions_Cache_Skip = "skip"
	GetFabVersionsOptions_Cache_Use = "use"
)

// NewGetFabVersionsOptions : Instantiate GetFabVersionsOptions
func (*BlockchainV2) NewGetFabVersionsOptions() *GetFabVersionsOptions {
	return &GetFabVersionsOptions{}
}

// SetCache : Allow user to set Cache
func (options *GetFabVersionsOptions) SetCache(cache string) *GetFabVersionsOptions {
	options.Cache = core.StringPtr(cache)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *GetFabVersionsOptions) SetHeaders(param map[string]string) *GetFabVersionsOptions {
	options.Headers = param
	return options
}

// GetFabricVersionsResponse : GetFabricVersionsResponse struct
type GetFabricVersionsResponse struct {
	Versions *GetFabricVersionsResponseVersions `json:"versions,omitempty"`
}


// UnmarshalGetFabricVersionsResponse unmarshals an instance of GetFabricVersionsResponse from the specified map of raw messages.
func UnmarshalGetFabricVersionsResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(GetFabricVersionsResponse)
	err = core.UnmarshalModel(m, "versions", &obj.Versions, UnmarshalGetFabricVersionsResponseVersions)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// GetFabricVersionsResponseVersions : GetFabricVersionsResponseVersions struct
type GetFabricVersionsResponseVersions struct {
	// A supported release of Fabric for this component type.
	Ca *FabricVersionDictionary `json:"ca,omitempty"`

	// A supported release of Fabric for this component type.
	Peer *FabricVersionDictionary `json:"peer,omitempty"`

	// A supported release of Fabric for this component type.
	Orderer *FabricVersionDictionary `json:"orderer,omitempty"`
}


// UnmarshalGetFabricVersionsResponseVersions unmarshals an instance of GetFabricVersionsResponseVersions from the specified map of raw messages.
func UnmarshalGetFabricVersionsResponseVersions(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(GetFabricVersionsResponseVersions)
	err = core.UnmarshalModel(m, "ca", &obj.Ca, UnmarshalFabricVersionDictionary)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "peer", &obj.Peer, UnmarshalFabricVersionDictionary)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "orderer", &obj.Orderer, UnmarshalFabricVersionDictionary)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// GetHealthOptions : The GetHealth options.
type GetHealthOptions struct {

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetHealthOptions : Instantiate GetHealthOptions
func (*BlockchainV2) NewGetHealthOptions() *GetHealthOptions {
	return &GetHealthOptions{}
}

// SetHeaders : Allow user to set Headers
func (options *GetHealthOptions) SetHeaders(param map[string]string) *GetHealthOptions {
	options.Headers = param
	return options
}

// GetMSPCertificateResponse : GetMSPCertificateResponse struct
type GetMSPCertificateResponse struct {
	Msps []MspPublicData `json:"msps,omitempty"`
}


// UnmarshalGetMSPCertificateResponse unmarshals an instance of GetMSPCertificateResponse from the specified map of raw messages.
func UnmarshalGetMSPCertificateResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(GetMSPCertificateResponse)
	err = core.UnmarshalModel(m, "msps", &obj.Msps, UnmarshalMspPublicData)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// GetMspCertificateOptions : The GetMspCertificate options.
type GetMspCertificateOptions struct {
	// The `msp_id` to fetch.
	MspID *string `json:"msp_id" validate:"required"`

	// Set to 'skip' if the response should skip local data and fetch live data wherever possible. Expect longer response
	// times if the cache is skipped. Default responses will use the cache.
	Cache *string `json:"cache,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// Constants associated with the GetMspCertificateOptions.Cache property.
// Set to 'skip' if the response should skip local data and fetch live data wherever possible. Expect longer response
// times if the cache is skipped. Default responses will use the cache.
const (
	GetMspCertificateOptions_Cache_Skip = "skip"
	GetMspCertificateOptions_Cache_Use = "use"
)

// NewGetMspCertificateOptions : Instantiate GetMspCertificateOptions
func (*BlockchainV2) NewGetMspCertificateOptions(mspID string) *GetMspCertificateOptions {
	return &GetMspCertificateOptions{
		MspID: core.StringPtr(mspID),
	}
}

// SetMspID : Allow user to set MspID
func (options *GetMspCertificateOptions) SetMspID(mspID string) *GetMspCertificateOptions {
	options.MspID = core.StringPtr(mspID)
	return options
}

// SetCache : Allow user to set Cache
func (options *GetMspCertificateOptions) SetCache(cache string) *GetMspCertificateOptions {
	options.Cache = core.StringPtr(cache)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *GetMspCertificateOptions) SetHeaders(param map[string]string) *GetMspCertificateOptions {
	options.Headers = param
	return options
}

// GetMultiComponentsResponse : Contains the details of multiple components the UI has onboarded.
type GetMultiComponentsResponse struct {
	// Array of components the UI has onboarded.
	Components []GenericComponentResponse `json:"components,omitempty"`
}


// UnmarshalGetMultiComponentsResponse unmarshals an instance of GetMultiComponentsResponse from the specified map of raw messages.
func UnmarshalGetMultiComponentsResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(GetMultiComponentsResponse)
	err = core.UnmarshalModel(m, "components", &obj.Components, UnmarshalGenericComponentResponse)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// GetNotificationsResponse : GetNotificationsResponse struct
type GetNotificationsResponse struct {
	// Number of notifications in database.
	Total *float64 `json:"total,omitempty"`

	// Number of notifications returned.
	Returning *float64 `json:"returning,omitempty"`

	// This array is ordered by creation date.
	Notifications []NotificationData `json:"notifications,omitempty"`
}


// UnmarshalGetNotificationsResponse unmarshals an instance of GetNotificationsResponse from the specified map of raw messages.
func UnmarshalGetNotificationsResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(GetNotificationsResponse)
	err = core.UnmarshalPrimitive(m, "total", &obj.Total)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "returning", &obj.Returning)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "notifications", &obj.Notifications, UnmarshalNotificationData)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// GetPostmanOptions : The GetPostman options.
type GetPostmanOptions struct {
	// - **bearer** - IAM Bearer Auth - *[Available on IBM Cloud]* - The same bearer token used to authenticate this
	// request will be copied into the Postman collection examples. The query parameter `token` must also be set with your
	// IAM bearer/access token value.
	// - **api_key** - IAM Api Key Auth - *[Available on IBM Cloud]* - The IAM api key will be copied into the Postman
	// collection examples. The query parameter `api_key` must also be set with your IAM API Key value.
	// - **basic** - Basic Auth - *[Available on OpenShift & IBM Cloud Private]* - A basic auth username and password will
	// be copied into the Postman collection examples. The query parameters `username` & `password` must also be set with
	// your IBP api key credentials. The IBP api key is the username and the api secret is the password.
	AuthType *string `json:"auth_type" validate:"required"`

	// The IAM access/bearer token to use for auth in the collection.
	Token *string `json:"token,omitempty"`

	// The IAM api key to use for auth in the collection.
	ApiKey *string `json:"api_key,omitempty"`

	// The basic auth username to use for auth in the collection.
	Username *string `json:"username,omitempty"`

	// The basic auth password to use for auth in the collection.
	Password *string `json:"password,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// Constants associated with the GetPostmanOptions.AuthType property.
// - **bearer** - IAM Bearer Auth - *[Available on IBM Cloud]* - The same bearer token used to authenticate this request
// will be copied into the Postman collection examples. The query parameter `token` must also be set with your IAM
// bearer/access token value.
// - **api_key** - IAM Api Key Auth - *[Available on IBM Cloud]* - The IAM api key will be copied into the Postman
// collection examples. The query parameter `api_key` must also be set with your IAM API Key value.
// - **basic** - Basic Auth - *[Available on OpenShift & IBM Cloud Private]* - A basic auth username and password will
// be copied into the Postman collection examples. The query parameters `username` & `password` must also be set with
// your IBP api key credentials. The IBP api key is the username and the api secret is the password.
const (
	GetPostmanOptions_AuthType_ApiKey = "api_key"
	GetPostmanOptions_AuthType_Basic = "basic"
	GetPostmanOptions_AuthType_Bearer = "bearer"
)

// NewGetPostmanOptions : Instantiate GetPostmanOptions
func (*BlockchainV2) NewGetPostmanOptions(authType string) *GetPostmanOptions {
	return &GetPostmanOptions{
		AuthType: core.StringPtr(authType),
	}
}

// SetAuthType : Allow user to set AuthType
func (options *GetPostmanOptions) SetAuthType(authType string) *GetPostmanOptions {
	options.AuthType = core.StringPtr(authType)
	return options
}

// SetToken : Allow user to set Token
func (options *GetPostmanOptions) SetToken(token string) *GetPostmanOptions {
	options.Token = core.StringPtr(token)
	return options
}

// SetApiKey : Allow user to set ApiKey
func (options *GetPostmanOptions) SetApiKey(apiKey string) *GetPostmanOptions {
	options.ApiKey = core.StringPtr(apiKey)
	return options
}

// SetUsername : Allow user to set Username
func (options *GetPostmanOptions) SetUsername(username string) *GetPostmanOptions {
	options.Username = core.StringPtr(username)
	return options
}

// SetPassword : Allow user to set Password
func (options *GetPostmanOptions) SetPassword(password string) *GetPostmanOptions {
	options.Password = core.StringPtr(password)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *GetPostmanOptions) SetHeaders(param map[string]string) *GetPostmanOptions {
	options.Headers = param
	return options
}

// GetPublicSettingsResponse : Contains the details of all public settings for the UI.
type GetPublicSettingsResponse struct {
	// The path to the activity tracker file. This file holds details of all activity. Defaults to '?' (disabled).
	ACTIVITYTRACKERPATH *string `json:"ACTIVITY_TRACKER_PATH,omitempty"`

	// Random/unique id for the process running the IBP console server.
	ATHENAID *string `json:"ATHENA_ID,omitempty"`

	// The type of auth protecting the UI.
	AUTHSCHEME *string `json:"AUTH_SCHEME,omitempty"`

	// Route used for an SSO callback uri. Only used if AUTH_SCHEME is "iam".
	CALLBACKURI *string `json:"CALLBACK_URI,omitempty"`

	CLUSTERDATA *GetPublicSettingsResponseCLUSTERDATA `json:"CLUSTER_DATA,omitempty"`

	// URL used for a configtxlator rest server.
	CONFIGTXLATORURL *string `json:"CONFIGTXLATOR_URL,omitempty"`

	// metadata about the IBM Cloud service instance. [Only populated if using IBM Cloud].
	CRN *GetPublicSettingsResponseCRN `json:"CRN,omitempty"`

	CRNSTRING *string `json:"CRN_STRING,omitempty"`

	// array of strings that define the Content Security Policy headers for the IBP console.
	CSPHEADERVALUES []string `json:"CSP_HEADER_VALUES,omitempty"`

	// The id of the database for internal documents.
	DBSYSTEM *string `json:"DB_SYSTEM,omitempty"`

	// URL of the companion application for the IBP console.
	DEPLOYERURL *string `json:"DEPLOYER_URL,omitempty"`

	// Browser cookies will use this value for their domain property. Thus it should match the URL's domain in the browser.
	// `null` is valid if serving over http.
	DOMAIN *string `json:"DOMAIN,omitempty"`

	// Either "dev" "staging" or "prod". Controls different security settings and minor things such as the amount of time
	// to cache content.
	ENVIRONMENT *string `json:"ENVIRONMENT,omitempty"`

	// Contains the Hyperledger Fabric capabilities flags that should be used.
	FABRICCAPABILITIES *GetPublicSettingsResponseFABRICCAPABILITIES `json:"FABRIC_CAPABILITIES,omitempty"`

	// Configures th IBP console to enable/disable features.
	FEATUREFLAGS interface{} `json:"FEATURE_FLAGS,omitempty"`

	// File logging settings.
	FILELOGGING *GetPublicSettingsResponseFILELOGGING `json:"FILE_LOGGING,omitempty"`

	// The external URL to reach the IBP console.
	HOSTURL *string `json:"HOST_URL,omitempty"`

	// If true an in memory cache will be used to interface with the IBM IAM (an authorization) service. [Only applies if
	// IBP is running in IBM Cloud].
	IAMCACHEENABLED *bool `json:"IAM_CACHE_ENABLED,omitempty"`

	// The URL to reach the IBM IAM service. [Only applies if IBP is running in IBM Cloud].
	IAMURL *string `json:"IAM_URL,omitempty"`

	// The URL to use during SSO login with the IBM IAM service. [Only applies if IBP is running in IBM Cloud].
	IBMIDCALLBACKURL *string `json:"IBM_ID_CALLBACK_URL,omitempty"`

	// If true the config file will not be loaded during startup. Thus settings in the config file will not take effect.
	IGNORECONFIGFILE *bool `json:"IGNORE_CONFIG_FILE,omitempty"`

	INACTIVITYTIMEOUTS *GetPublicSettingsResponseINACTIVITYTIMEOUTS `json:"INACTIVITY_TIMEOUTS,omitempty"`

	// What type of infrastructure is being used to run the IBP console. "ibmcloud", "azure", "other".
	INFRASTRUCTURE *string `json:"INFRASTRUCTURE,omitempty"`

	LANDINGURL *string `json:"LANDING_URL,omitempty"`

	// path for user login.
	LOGINURI *string `json:"LOGIN_URI,omitempty"`

	// path for user logout.
	LOGOUTURI *string `json:"LOGOUT_URI,omitempty"`

	// The number of `/api/_*` requests per minute to allow. Exceeding this limit results in 429 error responses.
	MAXREQPERMIN *float64 `json:"MAX_REQ_PER_MIN,omitempty"`

	// The number of `/ak/api/_*` requests per minute to allow. Exceeding this limit results in 429 error responses.
	MAXREQPERMINAK *float64 `json:"MAX_REQ_PER_MIN_AK,omitempty"`

	// If true an in memory cache will be used against couchdb requests.
	MEMORYCACHEENABLED *bool `json:"MEMORY_CACHE_ENABLED,omitempty"`

	// Internal port that IBP console is running on.
	PORT *string `json:"PORT,omitempty"`

	// If true an in memory cache will be used for internal proxy requests.
	PROXYCACHEENABLED *bool `json:"PROXY_CACHE_ENABLED,omitempty"`

	// If `"always"` requests to Fabric components will go through the IBP console server. If `true` requests to Fabric
	// components with IP based URLs will go through the IBP console server, while Fabric components with hostname based
	// URLs will go directly from the browser to the component. If `false` all requests to Fabric components will go
	// directly from the browser to the component.
	PROXYTLSFABRICREQS *string `json:"PROXY_TLS_FABRIC_REQS,omitempty"`

	// The URL to use to proxy an http request to a Fabric component.
	PROXYTLSHTTPURL *string `json:"PROXY_TLS_HTTP_URL,omitempty"`

	// The URL to use to proxy WebSocket request to a Fabric component.
	PROXYTLSWSURL interface{} `json:"PROXY_TLS_WS_URL,omitempty"`

	// If it's "local", things like https are disabled.
	REGION *string `json:"REGION,omitempty"`

	// If true an in memory cache will be used for browser session data.
	SESSIONCACHEENABLED *bool `json:"SESSION_CACHE_ENABLED,omitempty"`

	// Various timeouts for different Fabric operations.
	TIMEOUTS interface{} `json:"TIMEOUTS,omitempty"`

	TIMESTAMPS *SettingsTimestampData `json:"TIMESTAMPS,omitempty"`

	// Controls if Fabric transaction details are visible on the UI.
	TRANSACTIONVISIBILITY interface{} `json:"TRANSACTION_VISIBILITY,omitempty"`

	// Controls if proxy headers such as `X-Forwarded-*` should be parsed to gather data such as the client's IP.
	TRUSTPROXY *string `json:"TRUST_PROXY,omitempty"`

	// Controls if signatures in a signature collection APIs should skip verification or not.
	TRUSTUNKNOWNCERTS *bool `json:"TRUST_UNKNOWN_CERTS,omitempty"`

	// The various commit hashes of components powering this IBP console.
	VERSIONS *GetPublicSettingsResponseVERSIONS `json:"VERSIONS,omitempty"`
}


// UnmarshalGetPublicSettingsResponse unmarshals an instance of GetPublicSettingsResponse from the specified map of raw messages.
func UnmarshalGetPublicSettingsResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(GetPublicSettingsResponse)
	err = core.UnmarshalPrimitive(m, "ACTIVITY_TRACKER_PATH", &obj.ACTIVITYTRACKERPATH)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "ATHENA_ID", &obj.ATHENAID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "AUTH_SCHEME", &obj.AUTHSCHEME)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "CALLBACK_URI", &obj.CALLBACKURI)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "CLUSTER_DATA", &obj.CLUSTERDATA, UnmarshalGetPublicSettingsResponseCLUSTERDATA)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "CONFIGTXLATOR_URL", &obj.CONFIGTXLATORURL)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "CRN", &obj.CRN, UnmarshalGetPublicSettingsResponseCRN)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "CRN_STRING", &obj.CRNSTRING)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "CSP_HEADER_VALUES", &obj.CSPHEADERVALUES)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "DB_SYSTEM", &obj.DBSYSTEM)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "DEPLOYER_URL", &obj.DEPLOYERURL)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "DOMAIN", &obj.DOMAIN)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "ENVIRONMENT", &obj.ENVIRONMENT)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "FABRIC_CAPABILITIES", &obj.FABRICCAPABILITIES, UnmarshalGetPublicSettingsResponseFABRICCAPABILITIES)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "FEATURE_FLAGS", &obj.FEATUREFLAGS)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "FILE_LOGGING", &obj.FILELOGGING, UnmarshalGetPublicSettingsResponseFILELOGGING)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "HOST_URL", &obj.HOSTURL)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "IAM_CACHE_ENABLED", &obj.IAMCACHEENABLED)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "IAM_URL", &obj.IAMURL)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "IBM_ID_CALLBACK_URL", &obj.IBMIDCALLBACKURL)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "IGNORE_CONFIG_FILE", &obj.IGNORECONFIGFILE)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "INACTIVITY_TIMEOUTS", &obj.INACTIVITYTIMEOUTS, UnmarshalGetPublicSettingsResponseINACTIVITYTIMEOUTS)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "INFRASTRUCTURE", &obj.INFRASTRUCTURE)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "LANDING_URL", &obj.LANDINGURL)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "LOGIN_URI", &obj.LOGINURI)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "LOGOUT_URI", &obj.LOGOUTURI)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "MAX_REQ_PER_MIN", &obj.MAXREQPERMIN)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "MAX_REQ_PER_MIN_AK", &obj.MAXREQPERMINAK)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "MEMORY_CACHE_ENABLED", &obj.MEMORYCACHEENABLED)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "PORT", &obj.PORT)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "PROXY_CACHE_ENABLED", &obj.PROXYCACHEENABLED)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "PROXY_TLS_FABRIC_REQS", &obj.PROXYTLSFABRICREQS)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "PROXY_TLS_HTTP_URL", &obj.PROXYTLSHTTPURL)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "PROXY_TLS_WS_URL", &obj.PROXYTLSWSURL)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "REGION", &obj.REGION)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "SESSION_CACHE_ENABLED", &obj.SESSIONCACHEENABLED)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "TIMEOUTS", &obj.TIMEOUTS)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "TIMESTAMPS", &obj.TIMESTAMPS, UnmarshalSettingsTimestampData)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "TRANSACTION_VISIBILITY", &obj.TRANSACTIONVISIBILITY)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "TRUST_PROXY", &obj.TRUSTPROXY)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "TRUST_UNKNOWN_CERTS", &obj.TRUSTUNKNOWNCERTS)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "VERSIONS", &obj.VERSIONS, UnmarshalGetPublicSettingsResponseVERSIONS)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// GetPublicSettingsResponseCLUSTERDATA : GetPublicSettingsResponseCLUSTERDATA struct
type GetPublicSettingsResponseCLUSTERDATA struct {
	// Indicates whether this is a paid or free IBP console.
	Type *string `json:"type,omitempty"`
}


// UnmarshalGetPublicSettingsResponseCLUSTERDATA unmarshals an instance of GetPublicSettingsResponseCLUSTERDATA from the specified map of raw messages.
func UnmarshalGetPublicSettingsResponseCLUSTERDATA(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(GetPublicSettingsResponseCLUSTERDATA)
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// GetPublicSettingsResponseCRN : metadata about the IBM Cloud service instance. [Only populated if using IBM Cloud].
type GetPublicSettingsResponseCRN struct {
	AccountID *string `json:"account_id,omitempty"`

	CName *string `json:"c_name,omitempty"`

	CType *string `json:"c_type,omitempty"`

	InstanceID *string `json:"instance_id,omitempty"`

	Location *string `json:"location,omitempty"`

	ResourceID *string `json:"resource_id,omitempty"`

	ResourceType *string `json:"resource_type,omitempty"`

	ServiceName *string `json:"service_name,omitempty"`

	Version *string `json:"version,omitempty"`
}


// UnmarshalGetPublicSettingsResponseCRN unmarshals an instance of GetPublicSettingsResponseCRN from the specified map of raw messages.
func UnmarshalGetPublicSettingsResponseCRN(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(GetPublicSettingsResponseCRN)
	err = core.UnmarshalPrimitive(m, "account_id", &obj.AccountID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "c_name", &obj.CName)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "c_type", &obj.CType)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "instance_id", &obj.InstanceID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "location", &obj.Location)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "resource_id", &obj.ResourceID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "resource_type", &obj.ResourceType)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "service_name", &obj.ServiceName)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "version", &obj.Version)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// GetPublicSettingsResponseFABRICCAPABILITIES : Contains the Hyperledger Fabric capabilities flags that should be used.
type GetPublicSettingsResponseFABRICCAPABILITIES struct {
	Application []string `json:"application,omitempty"`

	Channel []string `json:"channel,omitempty"`

	Orderer []string `json:"orderer,omitempty"`
}


// UnmarshalGetPublicSettingsResponseFABRICCAPABILITIES unmarshals an instance of GetPublicSettingsResponseFABRICCAPABILITIES from the specified map of raw messages.
func UnmarshalGetPublicSettingsResponseFABRICCAPABILITIES(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(GetPublicSettingsResponseFABRICCAPABILITIES)
	err = core.UnmarshalPrimitive(m, "application", &obj.Application)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "channel", &obj.Channel)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "orderer", &obj.Orderer)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// GetPublicSettingsResponseFILELOGGING : File logging settings.
type GetPublicSettingsResponseFILELOGGING struct {
	// The logging settings for the client and server.
	Server *LogSettingsResponse `json:"server,omitempty"`

	// The logging settings for the client and server.
	Client *LogSettingsResponse `json:"client,omitempty"`
}


// UnmarshalGetPublicSettingsResponseFILELOGGING unmarshals an instance of GetPublicSettingsResponseFILELOGGING from the specified map of raw messages.
func UnmarshalGetPublicSettingsResponseFILELOGGING(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(GetPublicSettingsResponseFILELOGGING)
	err = core.UnmarshalModel(m, "server", &obj.Server, UnmarshalLogSettingsResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "client", &obj.Client, UnmarshalLogSettingsResponse)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// GetPublicSettingsResponseINACTIVITYTIMEOUTS : GetPublicSettingsResponseINACTIVITYTIMEOUTS struct
type GetPublicSettingsResponseINACTIVITYTIMEOUTS struct {
	Enabled *bool `json:"enabled,omitempty"`

	// How long to wait before auto-logging out a user. In milliseconds.
	MaxIdleTime *float64 `json:"max_idle_time,omitempty"`
}


// UnmarshalGetPublicSettingsResponseINACTIVITYTIMEOUTS unmarshals an instance of GetPublicSettingsResponseINACTIVITYTIMEOUTS from the specified map of raw messages.
func UnmarshalGetPublicSettingsResponseINACTIVITYTIMEOUTS(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(GetPublicSettingsResponseINACTIVITYTIMEOUTS)
	err = core.UnmarshalPrimitive(m, "enabled", &obj.Enabled)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "max_idle_time", &obj.MaxIdleTime)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// GetPublicSettingsResponseVERSIONS : The various commit hashes of components powering this IBP console.
type GetPublicSettingsResponseVERSIONS struct {
	// The commit hash of Apollo code (front-end).
	Apollo *string `json:"apollo,omitempty"`

	// The commit hash of Athena code (back-end).
	Athena *string `json:"athena,omitempty"`

	// The commit hash of Stitch code (fabric-sdk).
	Stitch *string `json:"stitch,omitempty"`

	// The tag of the build powering this IBP console.
	Tag *string `json:"tag,omitempty"`
}


// UnmarshalGetPublicSettingsResponseVERSIONS unmarshals an instance of GetPublicSettingsResponseVERSIONS from the specified map of raw messages.
func UnmarshalGetPublicSettingsResponseVERSIONS(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(GetPublicSettingsResponseVERSIONS)
	err = core.UnmarshalPrimitive(m, "apollo", &obj.Apollo)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "athena", &obj.Athena)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "stitch", &obj.Stitch)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "tag", &obj.Tag)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// GetSettingsOptions : The GetSettings options.
type GetSettingsOptions struct {

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetSettingsOptions : Instantiate GetSettingsOptions
func (*BlockchainV2) NewGetSettingsOptions() *GetSettingsOptions {
	return &GetSettingsOptions{}
}

// SetHeaders : Allow user to set Headers
func (options *GetSettingsOptions) SetHeaders(param map[string]string) *GetSettingsOptions {
	options.Headers = param
	return options
}

// GetSwaggerOptions : The GetSwagger options.
type GetSwaggerOptions struct {

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetSwaggerOptions : Instantiate GetSwaggerOptions
func (*BlockchainV2) NewGetSwaggerOptions() *GetSwaggerOptions {
	return &GetSwaggerOptions{}
}

// SetHeaders : Allow user to set Headers
func (options *GetSwaggerOptions) SetHeaders(param map[string]string) *GetSwaggerOptions {
	options.Headers = param
	return options
}

// ImportCaOptions : The ImportCa options.
type ImportCaOptions struct {
	// A descriptive name for this component.
	DisplayName *string `json:"display_name" validate:"required"`

	// The URL for the CA. Typically, client applications would send requests to this URL. Include the protocol,
	// hostname/ip and port.
	ApiURL *string `json:"api_url" validate:"required"`

	CaName *string `json:"ca_name" validate:"required"`

	TlscaName *string `json:"tlsca_name" validate:"required"`

	// The TLS certificate as base 64 encoded PEM. Certificate is used to secure/validate a TLS connection.
	TlsCert *string `json:"tls_cert" validate:"required"`

	// Indicates where the component is running.
	Location *string `json:"location,omitempty"`

	// The operations URL for the CA. Include the protocol, hostname/ip and port.
	OperationsURL *string `json:"operations_url,omitempty"`

	Tags []string `json:"tags,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewImportCaOptions : Instantiate ImportCaOptions
func (*BlockchainV2) NewImportCaOptions(displayName string, apiURL string, caName string, tlscaName string, tlsCert string) *ImportCaOptions {
	return &ImportCaOptions{
		DisplayName: core.StringPtr(displayName),
		ApiURL: core.StringPtr(apiURL),
		CaName: core.StringPtr(caName),
		TlscaName: core.StringPtr(tlscaName),
		TlsCert: core.StringPtr(tlsCert),
	}
}

// SetDisplayName : Allow user to set DisplayName
func (options *ImportCaOptions) SetDisplayName(displayName string) *ImportCaOptions {
	options.DisplayName = core.StringPtr(displayName)
	return options
}

// SetApiURL : Allow user to set ApiURL
func (options *ImportCaOptions) SetApiURL(apiURL string) *ImportCaOptions {
	options.ApiURL = core.StringPtr(apiURL)
	return options
}

// SetCaName : Allow user to set CaName
func (options *ImportCaOptions) SetCaName(caName string) *ImportCaOptions {
	options.CaName = core.StringPtr(caName)
	return options
}

// SetTlscaName : Allow user to set TlscaName
func (options *ImportCaOptions) SetTlscaName(tlscaName string) *ImportCaOptions {
	options.TlscaName = core.StringPtr(tlscaName)
	return options
}

// SetTlsCert : Allow user to set TlsCert
func (options *ImportCaOptions) SetTlsCert(tlsCert string) *ImportCaOptions {
	options.TlsCert = core.StringPtr(tlsCert)
	return options
}

// SetLocation : Allow user to set Location
func (options *ImportCaOptions) SetLocation(location string) *ImportCaOptions {
	options.Location = core.StringPtr(location)
	return options
}

// SetOperationsURL : Allow user to set OperationsURL
func (options *ImportCaOptions) SetOperationsURL(operationsURL string) *ImportCaOptions {
	options.OperationsURL = core.StringPtr(operationsURL)
	return options
}

// SetTags : Allow user to set Tags
func (options *ImportCaOptions) SetTags(tags []string) *ImportCaOptions {
	options.Tags = tags
	return options
}

// SetHeaders : Allow user to set Headers
func (options *ImportCaOptions) SetHeaders(param map[string]string) *ImportCaOptions {
	options.Headers = param
	return options
}

// ImportMspOptions : The ImportMsp options.
type ImportMspOptions struct {
	// The MSP id that is related to this component.
	MspID *string `json:"msp_id" validate:"required"`

	// A descriptive name for this MSP. The IBP console tile displays this name.
	DisplayName *string `json:"display_name" validate:"required"`

	// An array that contains one or more base 64 encoded PEM root certificates for the MSP.
	RootCerts []string `json:"root_certs" validate:"required"`

	// An array that contains base 64 encoded PEM intermediate certificates.
	IntermediateCerts []string `json:"intermediate_certs,omitempty"`

	// An array that contains base 64 encoded PEM identity certificates for administrators. Also known as signing
	// certificates of an organization administrator.
	Admins []string `json:"admins,omitempty"`

	// An array that contains one or more base 64 encoded PEM TLS root certificates.
	TlsRootCerts []string `json:"tls_root_certs,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewImportMspOptions : Instantiate ImportMspOptions
func (*BlockchainV2) NewImportMspOptions(mspID string, displayName string, rootCerts []string) *ImportMspOptions {
	return &ImportMspOptions{
		MspID: core.StringPtr(mspID),
		DisplayName: core.StringPtr(displayName),
		RootCerts: rootCerts,
	}
}

// SetMspID : Allow user to set MspID
func (options *ImportMspOptions) SetMspID(mspID string) *ImportMspOptions {
	options.MspID = core.StringPtr(mspID)
	return options
}

// SetDisplayName : Allow user to set DisplayName
func (options *ImportMspOptions) SetDisplayName(displayName string) *ImportMspOptions {
	options.DisplayName = core.StringPtr(displayName)
	return options
}

// SetRootCerts : Allow user to set RootCerts
func (options *ImportMspOptions) SetRootCerts(rootCerts []string) *ImportMspOptions {
	options.RootCerts = rootCerts
	return options
}

// SetIntermediateCerts : Allow user to set IntermediateCerts
func (options *ImportMspOptions) SetIntermediateCerts(intermediateCerts []string) *ImportMspOptions {
	options.IntermediateCerts = intermediateCerts
	return options
}

// SetAdmins : Allow user to set Admins
func (options *ImportMspOptions) SetAdmins(admins []string) *ImportMspOptions {
	options.Admins = admins
	return options
}

// SetTlsRootCerts : Allow user to set TlsRootCerts
func (options *ImportMspOptions) SetTlsRootCerts(tlsRootCerts []string) *ImportMspOptions {
	options.TlsRootCerts = tlsRootCerts
	return options
}

// SetHeaders : Allow user to set Headers
func (options *ImportMspOptions) SetHeaders(param map[string]string) *ImportMspOptions {
	options.Headers = param
	return options
}

// ImportOrdererOptions : The ImportOrderer options.
type ImportOrdererOptions struct {
	// A descriptive name for an ordering service. The parent IBP console tile displays this name.
	ClusterName *string `json:"cluster_name" validate:"required"`

	// A descriptive base name for each ordering node. One or more child IBP console tiles display this name.
	DisplayName *string `json:"display_name" validate:"required"`

	// The MSP id that is related to this component.
	MspID *string `json:"msp_id" validate:"required"`

	// The gRPC web proxy URL in front of the orderer. Include the protocol, hostname/ip and port.
	GrpcwpURL *string `json:"grpcwp_url" validate:"required"`

	// The TLS CA's root certificate as base 64 encoded PEM.
	TlsCaRootCert *string `json:"tls_ca_root_cert" validate:"required"`

	// Indicates where the component is running.
	Location *string `json:"location,omitempty"`

	// The gRPC URL for the orderer. Typically, client applications would send requests to this URL. Include the protocol,
	// hostname/ip and port.
	ApiURL *string `json:"api_url,omitempty"`

	// Used by Fabric health checker to monitor health status of the node. For more information, see [Fabric
	// documentation](https://hyperledger-fabric.readthedocs.io/en/release-1.4/operations_service.html). Include the
	// protocol, hostname/ip and port.
	OperationsURL *string `json:"operations_url,omitempty"`

	// The name of the system channel. Defaults to `testchainid`.
	SystemChannelID *string `json:"system_channel_id,omitempty"`

	Tags []string `json:"tags,omitempty"`

	// This field is used by IBP console when adding the node as a consenter to the system channel. The field indicates
	// what certificate the orderer will use when establishing mutual tls connections between other orderers when acting as
	// the **server or client**. The certificate should be base 64 encoded PEM.
	//
	// When setting `tls_cert` the fields `client_tls_cert` and `server_tls_cert` are not allowed.
	TlsCert *string `json:"tls_cert,omitempty"`

	// This field is used by IBP console when adding the node as a consenter to the system channel. The field indicates
	// what certificate the orderer will use when establishing mutual tls connections between other orderers when acting as
	// the **server**. The certificate should be base 64 encoded PEM.
	//
	// When setting `server_tls_cert` the field `client_tls_cert` is required and `tls_cert` is not allowed.
	ServerTlsCert *string `json:"server_tls_cert,omitempty"`

	// This field is used by IBP console when adding the node as a consenter to the system channel. The field indicates
	// what certificate the orderer will use when establishing mutual tls connections between other orderers when acting as
	// the **client**. The certificate should be base 64 encoded PEM.
	//
	// When setting `client_tls_cert` the field `server_tls_cert` is required and `tls_cert` is not allowed.
	ClientTlsCert *string `json:"client_tls_cert,omitempty"`

	// A unique id to identify this rafter cluster. Generated if not provided.
	ClusterID *string `json:"cluster_id,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewImportOrdererOptions : Instantiate ImportOrdererOptions
func (*BlockchainV2) NewImportOrdererOptions(clusterName string, displayName string, mspID string, grpcwpURL string, tlsCaRootCert string) *ImportOrdererOptions {
	return &ImportOrdererOptions{
		ClusterName: core.StringPtr(clusterName),
		DisplayName: core.StringPtr(displayName),
		MspID: core.StringPtr(mspID),
		GrpcwpURL: core.StringPtr(grpcwpURL),
		TlsCaRootCert: core.StringPtr(tlsCaRootCert),
	}
}

// SetClusterName : Allow user to set ClusterName
func (options *ImportOrdererOptions) SetClusterName(clusterName string) *ImportOrdererOptions {
	options.ClusterName = core.StringPtr(clusterName)
	return options
}

// SetDisplayName : Allow user to set DisplayName
func (options *ImportOrdererOptions) SetDisplayName(displayName string) *ImportOrdererOptions {
	options.DisplayName = core.StringPtr(displayName)
	return options
}

// SetMspID : Allow user to set MspID
func (options *ImportOrdererOptions) SetMspID(mspID string) *ImportOrdererOptions {
	options.MspID = core.StringPtr(mspID)
	return options
}

// SetGrpcwpURL : Allow user to set GrpcwpURL
func (options *ImportOrdererOptions) SetGrpcwpURL(grpcwpURL string) *ImportOrdererOptions {
	options.GrpcwpURL = core.StringPtr(grpcwpURL)
	return options
}

// SetTlsCaRootCert : Allow user to set TlsCaRootCert
func (options *ImportOrdererOptions) SetTlsCaRootCert(tlsCaRootCert string) *ImportOrdererOptions {
	options.TlsCaRootCert = core.StringPtr(tlsCaRootCert)
	return options
}

// SetLocation : Allow user to set Location
func (options *ImportOrdererOptions) SetLocation(location string) *ImportOrdererOptions {
	options.Location = core.StringPtr(location)
	return options
}

// SetApiURL : Allow user to set ApiURL
func (options *ImportOrdererOptions) SetApiURL(apiURL string) *ImportOrdererOptions {
	options.ApiURL = core.StringPtr(apiURL)
	return options
}

// SetOperationsURL : Allow user to set OperationsURL
func (options *ImportOrdererOptions) SetOperationsURL(operationsURL string) *ImportOrdererOptions {
	options.OperationsURL = core.StringPtr(operationsURL)
	return options
}

// SetSystemChannelID : Allow user to set SystemChannelID
func (options *ImportOrdererOptions) SetSystemChannelID(systemChannelID string) *ImportOrdererOptions {
	options.SystemChannelID = core.StringPtr(systemChannelID)
	return options
}

// SetTags : Allow user to set Tags
func (options *ImportOrdererOptions) SetTags(tags []string) *ImportOrdererOptions {
	options.Tags = tags
	return options
}

// SetTlsCert : Allow user to set TlsCert
func (options *ImportOrdererOptions) SetTlsCert(tlsCert string) *ImportOrdererOptions {
	options.TlsCert = core.StringPtr(tlsCert)
	return options
}

// SetServerTlsCert : Allow user to set ServerTlsCert
func (options *ImportOrdererOptions) SetServerTlsCert(serverTlsCert string) *ImportOrdererOptions {
	options.ServerTlsCert = core.StringPtr(serverTlsCert)
	return options
}

// SetClientTlsCert : Allow user to set ClientTlsCert
func (options *ImportOrdererOptions) SetClientTlsCert(clientTlsCert string) *ImportOrdererOptions {
	options.ClientTlsCert = core.StringPtr(clientTlsCert)
	return options
}

// SetClusterID : Allow user to set ClusterID
func (options *ImportOrdererOptions) SetClusterID(clusterID string) *ImportOrdererOptions {
	options.ClusterID = core.StringPtr(clusterID)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *ImportOrdererOptions) SetHeaders(param map[string]string) *ImportOrdererOptions {
	options.Headers = param
	return options
}

// ImportPeerOptions : The ImportPeer options.
type ImportPeerOptions struct {
	// A descriptive name for this peer. The IBP console tile displays this name.
	DisplayName *string `json:"display_name" validate:"required"`

	// The MSP id that is related to this component.
	MspID *string `json:"msp_id" validate:"required"`

	// The gRPC web proxy URL in front of the peer. Include the protocol, hostname/ip and port.
	GrpcwpURL *string `json:"grpcwp_url" validate:"required"`

	// The TLS CA's root certificate as base 64 encoded PEM.
	TlsCaRootCert *string `json:"tls_ca_root_cert" validate:"required"`

	// Indicates where the component is running.
	Location *string `json:"location,omitempty"`

	// The gRPC URL for the peer. Typically, client applications would send requests to this URL. Include the protocol,
	// hostname/ip and port.
	ApiURL *string `json:"api_url,omitempty"`

	// Used by Fabric health checker to monitor health status of the node. For more information, see [Fabric
	// documentation](https://hyperledger-fabric.readthedocs.io/en/release-1.4/operations_service.html). Include the
	// protocol, hostname/ip and port.
	OperationsURL *string `json:"operations_url,omitempty"`

	Tags []string `json:"tags,omitempty"`

	// The TLS certificate as base 64 encoded PEM. Certificate is used to secure/validate a TLS connection.
	TlsCert *string `json:"tls_cert,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewImportPeerOptions : Instantiate ImportPeerOptions
func (*BlockchainV2) NewImportPeerOptions(displayName string, mspID string, grpcwpURL string, tlsCaRootCert string) *ImportPeerOptions {
	return &ImportPeerOptions{
		DisplayName: core.StringPtr(displayName),
		MspID: core.StringPtr(mspID),
		GrpcwpURL: core.StringPtr(grpcwpURL),
		TlsCaRootCert: core.StringPtr(tlsCaRootCert),
	}
}

// SetDisplayName : Allow user to set DisplayName
func (options *ImportPeerOptions) SetDisplayName(displayName string) *ImportPeerOptions {
	options.DisplayName = core.StringPtr(displayName)
	return options
}

// SetMspID : Allow user to set MspID
func (options *ImportPeerOptions) SetMspID(mspID string) *ImportPeerOptions {
	options.MspID = core.StringPtr(mspID)
	return options
}

// SetGrpcwpURL : Allow user to set GrpcwpURL
func (options *ImportPeerOptions) SetGrpcwpURL(grpcwpURL string) *ImportPeerOptions {
	options.GrpcwpURL = core.StringPtr(grpcwpURL)
	return options
}

// SetTlsCaRootCert : Allow user to set TlsCaRootCert
func (options *ImportPeerOptions) SetTlsCaRootCert(tlsCaRootCert string) *ImportPeerOptions {
	options.TlsCaRootCert = core.StringPtr(tlsCaRootCert)
	return options
}

// SetLocation : Allow user to set Location
func (options *ImportPeerOptions) SetLocation(location string) *ImportPeerOptions {
	options.Location = core.StringPtr(location)
	return options
}

// SetApiURL : Allow user to set ApiURL
func (options *ImportPeerOptions) SetApiURL(apiURL string) *ImportPeerOptions {
	options.ApiURL = core.StringPtr(apiURL)
	return options
}

// SetOperationsURL : Allow user to set OperationsURL
func (options *ImportPeerOptions) SetOperationsURL(operationsURL string) *ImportPeerOptions {
	options.OperationsURL = core.StringPtr(operationsURL)
	return options
}

// SetTags : Allow user to set Tags
func (options *ImportPeerOptions) SetTags(tags []string) *ImportPeerOptions {
	options.Tags = tags
	return options
}

// SetTlsCert : Allow user to set TlsCert
func (options *ImportPeerOptions) SetTlsCert(tlsCert string) *ImportPeerOptions {
	options.TlsCert = core.StringPtr(tlsCert)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *ImportPeerOptions) SetHeaders(param map[string]string) *ImportPeerOptions {
	options.Headers = param
	return options
}

// ListComponentsOptions : The ListComponents options.
type ListComponentsOptions struct {
	// Set to 'included' if the response should include Kubernetes deployment attributes such as 'resources', 'storage',
	// 'zone', 'region', 'admin_certs', etc. Default responses will not include these fields.
	//
	// **This parameter will not work on *imported* components.**
	//
	// It's recommended to use `cache=skip` as well if up-to-date deployment data is needed.
	DeploymentAttrs *string `json:"deployment_attrs,omitempty"`

	// Set to 'included' if the response should include parsed PEM data along with base 64 encoded PEM string. Parsed
	// certificate data will include fields such as the serial number, issuer, expiration, subject, subject alt names, etc.
	// Default responses will not include these fields.
	ParsedCerts *string `json:"parsed_certs,omitempty"`

	// Set to 'skip' if the response should skip local data and fetch live data wherever possible. Expect longer response
	// times if the cache is skipped. Default responses will use the cache.
	Cache *string `json:"cache,omitempty"`

	// Set to 'included' if the response should fetch CA attributes, inspect certificates, and append extra fields to CA
	// and MSP component responses.
	// - CA components will have fields appended/updated with data fetched from the `/cainfo?ca=ca` endpoint of a CA, such
	// as: `ca_name`, `root_cert`, `fabric_version`, `issuer_public_key` and `issued_known_msps`. The field
	// `issued_known_msps` indicates imported IBP MSPs that this CA has issued. Meaning the MSP's root cert contains a
	// signature that is derived from this CA's root cert. Only imported MSPs are checked. Default responses will not
	// include these fields.
	// - MSP components will have the field `issued_by_ca_id` appended. This field indicates the id of an IBP console CA
	// that issued this MSP. Meaning the MSP's root cert contains a signature that is derived from this CA's root cert.
	// Only imported/created CAs are checked. Default responses will not include these fields.
	CaAttrs *string `json:"ca_attrs,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// Constants associated with the ListComponentsOptions.DeploymentAttrs property.
// Set to 'included' if the response should include Kubernetes deployment attributes such as 'resources', 'storage',
// 'zone', 'region', 'admin_certs', etc. Default responses will not include these fields.
//
// **This parameter will not work on *imported* components.**
//
// It's recommended to use `cache=skip` as well if up-to-date deployment data is needed.
const (
	ListComponentsOptions_DeploymentAttrs_Included = "included"
	ListComponentsOptions_DeploymentAttrs_Omitted = "omitted"
)

// Constants associated with the ListComponentsOptions.ParsedCerts property.
// Set to 'included' if the response should include parsed PEM data along with base 64 encoded PEM string. Parsed
// certificate data will include fields such as the serial number, issuer, expiration, subject, subject alt names, etc.
// Default responses will not include these fields.
const (
	ListComponentsOptions_ParsedCerts_Included = "included"
	ListComponentsOptions_ParsedCerts_Omitted = "omitted"
)

// Constants associated with the ListComponentsOptions.Cache property.
// Set to 'skip' if the response should skip local data and fetch live data wherever possible. Expect longer response
// times if the cache is skipped. Default responses will use the cache.
const (
	ListComponentsOptions_Cache_Skip = "skip"
	ListComponentsOptions_Cache_Use = "use"
)

// Constants associated with the ListComponentsOptions.CaAttrs property.
// Set to 'included' if the response should fetch CA attributes, inspect certificates, and append extra fields to CA and
// MSP component responses.
// - CA components will have fields appended/updated with data fetched from the `/cainfo?ca=ca` endpoint of a CA, such
// as: `ca_name`, `root_cert`, `fabric_version`, `issuer_public_key` and `issued_known_msps`. The field
// `issued_known_msps` indicates imported IBP MSPs that this CA has issued. Meaning the MSP's root cert contains a
// signature that is derived from this CA's root cert. Only imported MSPs are checked. Default responses will not
// include these fields.
// - MSP components will have the field `issued_by_ca_id` appended. This field indicates the id of an IBP console CA
// that issued this MSP. Meaning the MSP's root cert contains a signature that is derived from this CA's root cert. Only
// imported/created CAs are checked. Default responses will not include these fields.
const (
	ListComponentsOptions_CaAttrs_Included = "included"
	ListComponentsOptions_CaAttrs_Omitted = "omitted"
)

// NewListComponentsOptions : Instantiate ListComponentsOptions
func (*BlockchainV2) NewListComponentsOptions() *ListComponentsOptions {
	return &ListComponentsOptions{}
}

// SetDeploymentAttrs : Allow user to set DeploymentAttrs
func (options *ListComponentsOptions) SetDeploymentAttrs(deploymentAttrs string) *ListComponentsOptions {
	options.DeploymentAttrs = core.StringPtr(deploymentAttrs)
	return options
}

// SetParsedCerts : Allow user to set ParsedCerts
func (options *ListComponentsOptions) SetParsedCerts(parsedCerts string) *ListComponentsOptions {
	options.ParsedCerts = core.StringPtr(parsedCerts)
	return options
}

// SetCache : Allow user to set Cache
func (options *ListComponentsOptions) SetCache(cache string) *ListComponentsOptions {
	options.Cache = core.StringPtr(cache)
	return options
}

// SetCaAttrs : Allow user to set CaAttrs
func (options *ListComponentsOptions) SetCaAttrs(caAttrs string) *ListComponentsOptions {
	options.CaAttrs = core.StringPtr(caAttrs)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *ListComponentsOptions) SetHeaders(param map[string]string) *ListComponentsOptions {
	options.Headers = param
	return options
}

// ListNotificationsOptions : The ListNotifications options.
type ListNotificationsOptions struct {
	// The number of notifications to return. The default value is 100.
	Limit *float64 `json:"limit,omitempty"`

	// `skip` is used to paginate through a long list of sorted entries. For example, if there are 100 notifications, you
	// can issue the API with limit=10 and skip=0 to get the first 1-10. To get the next 10, you can set limit=10 and
	// skip=10 so that the values of entries 11-20 are returned.
	Skip *float64 `json:"skip,omitempty"`

	// Filter response to only contain notifications for a particular component id. The default response will include all
	// notifications.
	ComponentID *string `json:"component_id,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewListNotificationsOptions : Instantiate ListNotificationsOptions
func (*BlockchainV2) NewListNotificationsOptions() *ListNotificationsOptions {
	return &ListNotificationsOptions{}
}

// SetLimit : Allow user to set Limit
func (options *ListNotificationsOptions) SetLimit(limit float64) *ListNotificationsOptions {
	options.Limit = core.Float64Ptr(limit)
	return options
}

// SetSkip : Allow user to set Skip
func (options *ListNotificationsOptions) SetSkip(skip float64) *ListNotificationsOptions {
	options.Skip = core.Float64Ptr(skip)
	return options
}

// SetComponentID : Allow user to set ComponentID
func (options *ListNotificationsOptions) SetComponentID(componentID string) *ListNotificationsOptions {
	options.ComponentID = core.StringPtr(componentID)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *ListNotificationsOptions) SetHeaders(param map[string]string) *ListNotificationsOptions {
	options.Headers = param
	return options
}

// LogSettingsResponse : The logging settings for the client and server.
type LogSettingsResponse struct {
	// The client side (browser) logging settings. _Changes to this field will restart the IBP console server(s)_.
	Client *LoggingSettingsClient `json:"client,omitempty"`

	// The server side logging settings. _Changes to this field will restart the IBP console server(s)_.
	Server *LoggingSettingsServer `json:"server,omitempty"`
}


// UnmarshalLogSettingsResponse unmarshals an instance of LogSettingsResponse from the specified map of raw messages.
func UnmarshalLogSettingsResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(LogSettingsResponse)
	err = core.UnmarshalModel(m, "client", &obj.Client, UnmarshalLoggingSettingsClient)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "server", &obj.Server, UnmarshalLoggingSettingsServer)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// LoggingSettingsClient : The client side (browser) logging settings. _Changes to this field will restart the IBP console server(s)_.
type LoggingSettingsClient struct {
	// If `true` logging will be stored to a file on the file system.
	Enabled *bool `json:"enabled,omitempty"`

	// Valid log levels: "error", "warn", "info", "verbose", "debug", or "silly".
	Level *string `json:"level,omitempty"`

	// If `true` log file names will have a random suffix.
	UniqueName *bool `json:"unique_name,omitempty"`
}

// Constants associated with the LoggingSettingsClient.Level property.
// Valid log levels: "error", "warn", "info", "verbose", "debug", or "silly".
const (
	LoggingSettingsClient_Level_Debug = "debug"
	LoggingSettingsClient_Level_Error = "error"
	LoggingSettingsClient_Level_Info = "info"
	LoggingSettingsClient_Level_Silly = "silly"
	LoggingSettingsClient_Level_Verbose = "verbose"
	LoggingSettingsClient_Level_Warn = "warn"
)


// UnmarshalLoggingSettingsClient unmarshals an instance of LoggingSettingsClient from the specified map of raw messages.
func UnmarshalLoggingSettingsClient(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(LoggingSettingsClient)
	err = core.UnmarshalPrimitive(m, "enabled", &obj.Enabled)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "level", &obj.Level)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "unique_name", &obj.UniqueName)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// LoggingSettingsServer : The server side logging settings. _Changes to this field will restart the IBP console server(s)_.
type LoggingSettingsServer struct {
	// If `true` logging will be stored to a file on the file system.
	Enabled *bool `json:"enabled,omitempty"`

	// Valid log levels: "error", "warn", "info", "verbose", "debug", or "silly".
	Level *string `json:"level,omitempty"`

	// If `true` log file names will have a random suffix.
	UniqueName *bool `json:"unique_name,omitempty"`
}

// Constants associated with the LoggingSettingsServer.Level property.
// Valid log levels: "error", "warn", "info", "verbose", "debug", or "silly".
const (
	LoggingSettingsServer_Level_Debug = "debug"
	LoggingSettingsServer_Level_Error = "error"
	LoggingSettingsServer_Level_Info = "info"
	LoggingSettingsServer_Level_Silly = "silly"
	LoggingSettingsServer_Level_Verbose = "verbose"
	LoggingSettingsServer_Level_Warn = "warn"
)


// UnmarshalLoggingSettingsServer unmarshals an instance of LoggingSettingsServer from the specified map of raw messages.
func UnmarshalLoggingSettingsServer(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(LoggingSettingsServer)
	err = core.UnmarshalPrimitive(m, "enabled", &obj.Enabled)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "level", &obj.Level)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "unique_name", &obj.UniqueName)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Metrics : Metrics struct
type Metrics struct {
	// Metrics provider to use. Can be either 'statsd', 'prometheus', or 'disabled'.
	Provider *string `json:"provider" validate:"required"`

	Statsd *MetricsStatsd `json:"statsd,omitempty"`
}

// Constants associated with the Metrics.Provider property.
// Metrics provider to use. Can be either 'statsd', 'prometheus', or 'disabled'.
const (
	Metrics_Provider_Disabled = "disabled"
	Metrics_Provider_Prometheus = "prometheus"
	Metrics_Provider_Statsd = "statsd"
)


// NewMetrics : Instantiate Metrics (Generic Model Constructor)
func (*BlockchainV2) NewMetrics(provider string) (model *Metrics, err error) {
	model = &Metrics{
		Provider: core.StringPtr(provider),
	}
	err = core.ValidateStruct(model, "required parameters")
	return
}

// UnmarshalMetrics unmarshals an instance of Metrics from the specified map of raw messages.
func UnmarshalMetrics(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Metrics)
	err = core.UnmarshalPrimitive(m, "provider", &obj.Provider)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "statsd", &obj.Statsd, UnmarshalMetricsStatsd)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// MetricsStatsd : MetricsStatsd struct
type MetricsStatsd struct {
	// Either UDP or TCP.
	Network *string `json:"network" validate:"required"`

	// The address of the statsd server. Include hostname/ip and port.
	Address *string `json:"address" validate:"required"`

	// The frequency at which locally cached counters and gauges are pushed to statsd.
	WriteInterval *string `json:"writeInterval" validate:"required"`

	// The string that is prepended to all emitted statsd metrics.
	Prefix *string `json:"prefix" validate:"required"`
}

// Constants associated with the MetricsStatsd.Network property.
// Either UDP or TCP.
const (
	MetricsStatsd_Network_Tcp = "tcp"
	MetricsStatsd_Network_Udp = "udp"
)


// NewMetricsStatsd : Instantiate MetricsStatsd (Generic Model Constructor)
func (*BlockchainV2) NewMetricsStatsd(network string, address string, writeInterval string, prefix string) (model *MetricsStatsd, err error) {
	model = &MetricsStatsd{
		Network: core.StringPtr(network),
		Address: core.StringPtr(address),
		WriteInterval: core.StringPtr(writeInterval),
		Prefix: core.StringPtr(prefix),
	}
	err = core.ValidateStruct(model, "required parameters")
	return
}

// UnmarshalMetricsStatsd unmarshals an instance of MetricsStatsd from the specified map of raw messages.
func UnmarshalMetricsStatsd(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(MetricsStatsd)
	err = core.UnmarshalPrimitive(m, "network", &obj.Network)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "address", &obj.Address)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "writeInterval", &obj.WriteInterval)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "prefix", &obj.Prefix)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// MspConfigData : MspConfigData struct
type MspConfigData struct {
	// An identity private key (base 64 encoded PEM) for this component.
	Keystore *string `json:"keystore" validate:"required"`

	// An identity certificate (base 64 encoded PEM) for this component that was signed by the CA.
	Signcerts *string `json:"signcerts" validate:"required"`

	// An array that contains one or more base 64 encoded PEM CA root certificates.
	Cacerts []string `json:"cacerts" validate:"required"`

	// An array that contains base 64 encoded PEM intermediate CA certificates.
	Intermediatecerts []string `json:"intermediatecerts,omitempty"`

	// An array that contains base 64 encoded PEM identity certificates for administrators. Also known as signing
	// certificates of an organization administrator.
	Admincerts []string `json:"admincerts,omitempty"`
}


// NewMspConfigData : Instantiate MspConfigData (Generic Model Constructor)
func (*BlockchainV2) NewMspConfigData(keystore string, signcerts string, cacerts []string) (model *MspConfigData, err error) {
	model = &MspConfigData{
		Keystore: core.StringPtr(keystore),
		Signcerts: core.StringPtr(signcerts),
		Cacerts: cacerts,
	}
	err = core.ValidateStruct(model, "required parameters")
	return
}

// UnmarshalMspConfigData unmarshals an instance of MspConfigData from the specified map of raw messages.
func UnmarshalMspConfigData(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(MspConfigData)
	err = core.UnmarshalPrimitive(m, "keystore", &obj.Keystore)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "signcerts", &obj.Signcerts)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "cacerts", &obj.Cacerts)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "intermediatecerts", &obj.Intermediatecerts)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "admincerts", &obj.Admincerts)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// MspPublicData : MspPublicData struct
type MspPublicData struct {
	// The MSP id that is related to this component.
	MspID *string `json:"msp_id,omitempty"`

	// An array that contains one or more base 64 encoded PEM root certificates for the MSP.
	RootCerts []string `json:"root_certs,omitempty"`

	// An array that contains base 64 encoded PEM identity certificates for administrators. Also known as signing
	// certificates of an organization administrator.
	Admins []string `json:"admins,omitempty"`

	// An array that contains one or more base 64 encoded PEM TLS root certificates.
	TlsRootCerts []string `json:"tls_root_certs,omitempty"`
}


// UnmarshalMspPublicData unmarshals an instance of MspPublicData from the specified map of raw messages.
func UnmarshalMspPublicData(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(MspPublicData)
	err = core.UnmarshalPrimitive(m, "msp_id", &obj.MspID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "root_certs", &obj.RootCerts)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "admins", &obj.Admins)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "tls_root_certs", &obj.TlsRootCerts)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// MspResponse : Contains the details of an MSP (Membership Service Provider).
type MspResponse struct {
	// The unique identifier of this component.
	ID *string `json:"id,omitempty"`

	// The type of this component. Such as: "fabric-peer", "fabric-ca", "fabric-orderer", etc.
	Type *string `json:"type,omitempty"`

	// A descriptive name for this MSP. The IBP console tile displays this name.
	DisplayName *string `json:"display_name,omitempty"`

	// The MSP id that is related to this component.
	MspID *string `json:"msp_id,omitempty"`

	// UTC UNIX timestamp of component onboarding to the UI. In milliseconds.
	Timestamp *float64 `json:"timestamp,omitempty"`

	Tags []string `json:"tags,omitempty"`

	// An array that contains one or more base 64 encoded PEM root certificates for the MSP.
	RootCerts []string `json:"root_certs,omitempty"`

	// An array that contains base 64 encoded PEM intermediate certificates.
	IntermediateCerts []string `json:"intermediate_certs,omitempty"`

	// An array that contains base 64 encoded PEM identity certificates for administrators. Also known as signing
	// certificates of an organization administrator.
	Admins []string `json:"admins,omitempty"`

	// The versioning of the IBP console format of this JSON.
	SchemeVersion *string `json:"scheme_version,omitempty"`

	// An array that contains one or more base 64 encoded PEM TLS root certificates.
	TlsRootCerts []string `json:"tls_root_certs,omitempty"`
}


// UnmarshalMspResponse unmarshals an instance of MspResponse from the specified map of raw messages.
func UnmarshalMspResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(MspResponse)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "display_name", &obj.DisplayName)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "msp_id", &obj.MspID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "timestamp", &obj.Timestamp)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "tags", &obj.Tags)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "root_certs", &obj.RootCerts)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "intermediate_certs", &obj.IntermediateCerts)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "admins", &obj.Admins)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "scheme_version", &obj.SchemeVersion)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "tls_root_certs", &obj.TlsRootCerts)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// NotificationData : NotificationData struct
type NotificationData struct {
	// Unique id for the notification.
	ID *string `json:"id,omitempty"`

	// Values can be "notification", "webhook_tx" or "other".
	Type *string `json:"type,omitempty"`

	// Values can be "pending", "error", or "success".
	Status *string `json:"status,omitempty"`

	// The end user who initiated the action for the notification.
	By *string `json:"by,omitempty"`

	// Text describing the outcome of the transaction.
	Message *string `json:"message,omitempty"`

	// UTC UNIX timestamp of the notification's creation. In milliseconds.
	TsDisplay *float64 `json:"ts_display,omitempty"`
}


// UnmarshalNotificationData unmarshals an instance of NotificationData from the specified map of raw messages.
func UnmarshalNotificationData(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(NotificationData)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "status", &obj.Status)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "by", &obj.By)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "message", &obj.Message)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "ts_display", &obj.TsDisplay)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// OrdererResponse : Contains the details of an ordering node.
type OrdererResponse struct {
	// The unique identifier of this component.
	ID *string `json:"id,omitempty"`

	// The unique id for the component in Kubernetes. Not available if component was imported.
	DepComponentID *string `json:"dep_component_id,omitempty"`

	// The type of this component. Such as: "fabric-peer", "fabric-ca", "fabric-orderer", etc.
	Type *string `json:"type,omitempty"`

	// A descriptive base name for each ordering node. One or more child IBP console tiles display this name.
	DisplayName *string `json:"display_name,omitempty"`

	// The gRPC web proxy URL in front of the orderer. Include the protocol, hostname/ip and port.
	GrpcwpURL *string `json:"grpcwp_url,omitempty"`

	// The gRPC URL for the orderer. Typically, client applications would send requests to this URL. Include the protocol,
	// hostname/ip and port.
	ApiURL *string `json:"api_url,omitempty"`

	// Used by Fabric health checker to monitor health status of the node. For more information, see [Fabric
	// documentation](https://hyperledger-fabric.readthedocs.io/en/release-1.4/operations_service.html). Include the
	// protocol, hostname/ip and port.
	OperationsURL *string `json:"operations_url,omitempty"`

	// The MSP id that is related to this component.
	MspID *string `json:"msp_id,omitempty"`

	// The **cached** configuration override that was set for the Kubernetes deployment. Field does not exist if an
	// override was not set of if the component was imported.
	ConfigOverride interface{} `json:"config_override,omitempty"`

	// The state of a pre-created orderer node. A value of `true` means that the orderer node was added as a system channel
	// consenter. This is a manual field. Set it yourself after finishing the raft append flow to indicate that this node
	// is ready for use. See the [Submit config block to orderer](#submit-config-block-to-orderer) API description for more
	// details about appending raft nodes.
	ConsenterProposalFin *bool `json:"consenter_proposal_fin,omitempty"`

	// Indicates where the component is running.
	Location *string `json:"location,omitempty"`

	// UTC UNIX timestamp of component onboarding to the UI. In milliseconds.
	Timestamp *float64 `json:"timestamp,omitempty"`

	// The **cached** Kubernetes resource attributes for this component. Not available if orderer was imported.
	Resources *OrdererResponseResources `json:"resources,omitempty"`

	// The versioning of the IBP console format of this JSON.
	SchemeVersion *string `json:"scheme_version,omitempty"`

	// The **cached** Kubernetes storage attributes for this component. Not available if orderer was imported.
	Storage *OrdererResponseStorage `json:"storage,omitempty"`

	// The name of the system channel. Defaults to `testchainid`.
	SystemChannelID *string `json:"system_channel_id,omitempty"`

	Tags []string `json:"tags,omitempty"`

	// This field is used by IBP console when adding the node as a consenter to the system channel. The field indicates
	// what certificate the orderer will use when establishing mutual tls connections between other orderers when acting as
	// the **server or client**. The certificate should be base 64 encoded PEM.
	//
	// When setting `tls_cert` the fields `client_tls_cert` and `server_tls_cert` are not allowed.
	TlsCert *string `json:"tls_cert,omitempty"`

	// This field is used by IBP console when adding the node as a consenter to the system channel. The field indicates
	// what certificate the orderer will use when establishing mutual tls connections between other orderers when acting as
	// the **server**. The certificate should be base 64 encoded PEM.
	//
	// When setting `server_tls_cert` the field `client_tls_cert` is required and `tls_cert` is not allowed.
	ServerTlsCert *string `json:"server_tls_cert,omitempty"`

	// This field is used by IBP console when adding the node as a consenter to the system channel. The field indicates
	// what certificate the orderer will use when establishing mutual tls connections between other orderers when acting as
	// the **client**. The certificate should be base 64 encoded PEM.
	//
	// When setting `client_tls_cert` the field `server_tls_cert` is required and `tls_cert` is not allowed.
	ClientTlsCert *string `json:"client_tls_cert,omitempty"`

	// The type of Fabric orderer. Currently, only the type `"raft"` is supported.
	// [etcd/raft](/docs/blockchain?topic=blockchain-ibp-console-build-network#ibp-console-build-network-ordering-console).
	OrdererType *string `json:"orderer_type,omitempty"`

	// The cached Hyperledger Fabric release version.
	Version *string `json:"version,omitempty"`

	// Specify the Kubernetes zone for the deployment. The deployment will use a k8s node in this zone. Find the list of
	// possible zones by retrieving your Kubernetes node labels: `kubectl get nodes --show-labels`. [More
	// information](https://kubernetes.io/docs/setup/best-practices/multiple-zones).
	Zone *string `json:"zone,omitempty"`
}

// Constants associated with the OrdererResponse.OrdererType property.
// The type of Fabric orderer. Currently, only the type `"raft"` is supported.
// [etcd/raft](/docs/blockchain?topic=blockchain-ibp-console-build-network#ibp-console-build-network-ordering-console).
const (
	OrdererResponse_OrdererType_Raft = "raft"
)


// UnmarshalOrdererResponse unmarshals an instance of OrdererResponse from the specified map of raw messages.
func UnmarshalOrdererResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(OrdererResponse)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "dep_component_id", &obj.DepComponentID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "display_name", &obj.DisplayName)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "grpcwp_url", &obj.GrpcwpURL)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "api_url", &obj.ApiURL)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "operations_url", &obj.OperationsURL)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "msp_id", &obj.MspID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "config_override", &obj.ConfigOverride)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "consenter_proposal_fin", &obj.ConsenterProposalFin)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "location", &obj.Location)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "timestamp", &obj.Timestamp)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "resources", &obj.Resources, UnmarshalOrdererResponseResources)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "scheme_version", &obj.SchemeVersion)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "storage", &obj.Storage, UnmarshalOrdererResponseStorage)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "system_channel_id", &obj.SystemChannelID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "tags", &obj.Tags)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "tls_cert", &obj.TlsCert)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "server_tls_cert", &obj.ServerTlsCert)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "client_tls_cert", &obj.ClientTlsCert)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "orderer_type", &obj.OrdererType)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "version", &obj.Version)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "zone", &obj.Zone)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// OrdererResponseResources : The **cached** Kubernetes resource attributes for this component. Not available if orderer was imported.
type OrdererResponseResources struct {
	Orderer *GenericResources `json:"orderer,omitempty"`

	Proxy *GenericResources `json:"proxy,omitempty"`
}


// UnmarshalOrdererResponseResources unmarshals an instance of OrdererResponseResources from the specified map of raw messages.
func UnmarshalOrdererResponseResources(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(OrdererResponseResources)
	err = core.UnmarshalModel(m, "orderer", &obj.Orderer, UnmarshalGenericResources)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "proxy", &obj.Proxy, UnmarshalGenericResources)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// OrdererResponseStorage : The **cached** Kubernetes storage attributes for this component. Not available if orderer was imported.
type OrdererResponseStorage struct {
	Orderer *StorageObject `json:"orderer,omitempty"`
}


// UnmarshalOrdererResponseStorage unmarshals an instance of OrdererResponseStorage from the specified map of raw messages.
func UnmarshalOrdererResponseStorage(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(OrdererResponseStorage)
	err = core.UnmarshalModel(m, "orderer", &obj.Orderer, UnmarshalStorageObject)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// PeerResources : CPU and memory properties. This feature is not available if using a free Kubernetes cluster.
type PeerResources struct {
	// This field requires the use of Fabric v2.1.* and higher.
	Chaincodelauncher *ResourceObjectFabV2 `json:"chaincodelauncher,omitempty"`

	// This field requires the use of Fabric v1.4.* and higher.
	Couchdb *ResourceObject `json:"couchdb,omitempty"`

	// This field requires the use of Fabric v1.4.* and **lower**.
	Dind *ResourceObjectFabV1 `json:"dind,omitempty"`

	// This field requires the use of Fabric v1.4.* and **lower**.
	Fluentd *ResourceObjectFabV1 `json:"fluentd,omitempty"`

	// This field requires the use of Fabric v1.4.* and higher.
	Peer *ResourceObject `json:"peer,omitempty"`

	// This field requires the use of Fabric v1.4.* and higher.
	Proxy *ResourceObject `json:"proxy,omitempty"`
}


// UnmarshalPeerResources unmarshals an instance of PeerResources from the specified map of raw messages.
func UnmarshalPeerResources(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(PeerResources)
	err = core.UnmarshalModel(m, "chaincodelauncher", &obj.Chaincodelauncher, UnmarshalResourceObjectFabV2)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "couchdb", &obj.Couchdb, UnmarshalResourceObject)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "dind", &obj.Dind, UnmarshalResourceObjectFabV1)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "fluentd", &obj.Fluentd, UnmarshalResourceObjectFabV1)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "peer", &obj.Peer, UnmarshalResourceObject)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "proxy", &obj.Proxy, UnmarshalResourceObject)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// PeerResponse : Contains the details of a peer.
type PeerResponse struct {
	// The unique identifier of this component.
	ID *string `json:"id,omitempty"`

	// The unique id for the component in Kubernetes. Not available if component was imported.
	DepComponentID *string `json:"dep_component_id,omitempty"`

	// The type of this component. Such as: "fabric-peer", "fabric-ca", "fabric-orderer", etc.
	Type *string `json:"type,omitempty"`

	// A descriptive name for this peer. The IBP console tile displays this name.
	DisplayName *string `json:"display_name,omitempty"`

	// The gRPC web proxy URL in front of the peer. Include the protocol, hostname/ip and port.
	GrpcwpURL *string `json:"grpcwp_url,omitempty"`

	// The gRPC URL for the peer. Typically, client applications would send requests to this URL. Include the protocol,
	// hostname/ip and port.
	ApiURL *string `json:"api_url,omitempty"`

	// Used by Fabric health checker to monitor health status of the node. For more information, see [Fabric
	// documentation](https://hyperledger-fabric.readthedocs.io/en/release-1.4/operations_service.html). Include the
	// protocol, hostname/ip and port.
	OperationsURL *string `json:"operations_url,omitempty"`

	// The MSP id that is related to this component.
	MspID *string `json:"msp_id,omitempty"`

	// The **cached** configuration override that was set for the Kubernetes deployment. Field does not exist if an
	// override was not set of if the component was imported.
	ConfigOverride interface{} `json:"config_override,omitempty"`

	// Indicates where the component is running.
	Location *string `json:"location,omitempty"`

	// Select the state database for the peer. Can be either "couchdb" or "leveldb". The default is "couchdb".
	StateDb *string `json:"state_db,omitempty"`

	// UTC UNIX timestamp of component onboarding to the UI. In milliseconds.
	Timestamp *float64 `json:"timestamp,omitempty"`

	// The **cached** Kubernetes resource attributes for this component. Not available if peer was imported.
	Resources *PeerResponseResources `json:"resources,omitempty"`

	// The versioning of the IBP console format of this JSON.
	SchemeVersion *string `json:"scheme_version,omitempty"`

	// The **cached** Kubernetes storage attributes for this component. Not available if peer was imported.
	Storage *PeerResponseStorage `json:"storage,omitempty"`

	Tags []string `json:"tags,omitempty"`

	// The TLS certificate as base 64 encoded PEM. Certificate is used to secure/validate a TLS connection.
	TlsCert *string `json:"tls_cert,omitempty"`

	// The cached Hyperledger Fabric release version.
	Version *string `json:"version,omitempty"`

	// Specify the Kubernetes zone for the deployment. The deployment will use a k8s node in this zone. Find the list of
	// possible zones by retrieving your Kubernetes node labels: `kubectl get nodes --show-labels`. [More
	// information](https://kubernetes.io/docs/setup/best-practices/multiple-zones).
	Zone *string `json:"zone,omitempty"`
}

// Constants associated with the PeerResponse.StateDb property.
// Select the state database for the peer. Can be either "couchdb" or "leveldb". The default is "couchdb".
const (
	PeerResponse_StateDb_Couchdb = "couchdb"
	PeerResponse_StateDb_Leveldb = "leveldb"
)


// UnmarshalPeerResponse unmarshals an instance of PeerResponse from the specified map of raw messages.
func UnmarshalPeerResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(PeerResponse)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "dep_component_id", &obj.DepComponentID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "display_name", &obj.DisplayName)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "grpcwp_url", &obj.GrpcwpURL)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "api_url", &obj.ApiURL)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "operations_url", &obj.OperationsURL)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "msp_id", &obj.MspID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "config_override", &obj.ConfigOverride)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "location", &obj.Location)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "state_db", &obj.StateDb)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "timestamp", &obj.Timestamp)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "resources", &obj.Resources, UnmarshalPeerResponseResources)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "scheme_version", &obj.SchemeVersion)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "storage", &obj.Storage, UnmarshalPeerResponseStorage)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "tags", &obj.Tags)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "tls_cert", &obj.TlsCert)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "version", &obj.Version)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "zone", &obj.Zone)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// PeerResponseResources : The **cached** Kubernetes resource attributes for this component. Not available if peer was imported.
type PeerResponseResources struct {
	Peer *GenericResources `json:"peer,omitempty"`

	Proxy *GenericResources `json:"proxy,omitempty"`

	Couchdb *GenericResources `json:"couchdb,omitempty"`
}


// UnmarshalPeerResponseResources unmarshals an instance of PeerResponseResources from the specified map of raw messages.
func UnmarshalPeerResponseResources(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(PeerResponseResources)
	err = core.UnmarshalModel(m, "peer", &obj.Peer, UnmarshalGenericResources)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "proxy", &obj.Proxy, UnmarshalGenericResources)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "couchdb", &obj.Couchdb, UnmarshalGenericResources)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// PeerResponseStorage : The **cached** Kubernetes storage attributes for this component. Not available if peer was imported.
type PeerResponseStorage struct {
	Peer *StorageObject `json:"peer,omitempty"`

	Couchdb *StorageObject `json:"couchdb,omitempty"`
}


// UnmarshalPeerResponseStorage unmarshals an instance of PeerResponseStorage from the specified map of raw messages.
func UnmarshalPeerResponseStorage(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(PeerResponseStorage)
	err = core.UnmarshalModel(m, "peer", &obj.Peer, UnmarshalStorageObject)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "couchdb", &obj.Couchdb, UnmarshalStorageObject)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// RemoveComponentOptions : The RemoveComponent options.
type RemoveComponentOptions struct {
	// The `id` of the imported component to remove. Use the [Get all components](#get-all-components) API to determine the
	// component id.
	ID *string `json:"id" validate:"required"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewRemoveComponentOptions : Instantiate RemoveComponentOptions
func (*BlockchainV2) NewRemoveComponentOptions(id string) *RemoveComponentOptions {
	return &RemoveComponentOptions{
		ID: core.StringPtr(id),
	}
}

// SetID : Allow user to set ID
func (options *RemoveComponentOptions) SetID(id string) *RemoveComponentOptions {
	options.ID = core.StringPtr(id)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *RemoveComponentOptions) SetHeaders(param map[string]string) *RemoveComponentOptions {
	options.Headers = param
	return options
}

// RemoveComponentsByTagOptions : The RemoveComponentsByTag options.
type RemoveComponentsByTagOptions struct {
	// The tag to filter components on. Not case-sensitive.
	Tag *string `json:"tag" validate:"required"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewRemoveComponentsByTagOptions : Instantiate RemoveComponentsByTagOptions
func (*BlockchainV2) NewRemoveComponentsByTagOptions(tag string) *RemoveComponentsByTagOptions {
	return &RemoveComponentsByTagOptions{
		Tag: core.StringPtr(tag),
	}
}

// SetTag : Allow user to set Tag
func (options *RemoveComponentsByTagOptions) SetTag(tag string) *RemoveComponentsByTagOptions {
	options.Tag = core.StringPtr(tag)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *RemoveComponentsByTagOptions) SetHeaders(param map[string]string) *RemoveComponentsByTagOptions {
	options.Headers = param
	return options
}

// RemoveMultiComponentsResponse : RemoveMultiComponentsResponse struct
type RemoveMultiComponentsResponse struct {
	Removed []DeleteComponentResponse `json:"removed,omitempty"`
}


// UnmarshalRemoveMultiComponentsResponse unmarshals an instance of RemoveMultiComponentsResponse from the specified map of raw messages.
func UnmarshalRemoveMultiComponentsResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(RemoveMultiComponentsResponse)
	err = core.UnmarshalModel(m, "removed", &obj.Removed, UnmarshalDeleteComponentResponse)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ResourceLimits : ResourceLimits struct
type ResourceLimits struct {
	// Maximum CPU for subcomponent. Must be >= "requests.cpu". Defaults to the same value in "requests.cpu". [Resource
	// details](/docs/blockchain?topic=blockchain-ibp-console-govern-components#ibp-console-govern-components-allocate-resources).
	Cpu *string `json:"cpu,omitempty"`

	// Maximum memory for subcomponent. Must be >= "requests.memory". Defaults to the same value in "requests.memory".
	// [Resource
	// details](/docs/blockchain?topic=blockchain-ibp-console-govern-components#ibp-console-govern-components-allocate-resources).
	Memory *string `json:"memory,omitempty"`
}


// UnmarshalResourceLimits unmarshals an instance of ResourceLimits from the specified map of raw messages.
func UnmarshalResourceLimits(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ResourceLimits)
	err = core.UnmarshalPrimitive(m, "cpu", &obj.Cpu)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "memory", &obj.Memory)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ResourceObject : This field requires the use of Fabric v1.4.* and higher.
type ResourceObject struct {
	Requests *ResourceRequests `json:"requests" validate:"required"`

	Limits *ResourceLimits `json:"limits,omitempty"`
}


// NewResourceObject : Instantiate ResourceObject (Generic Model Constructor)
func (*BlockchainV2) NewResourceObject(requests *ResourceRequests) (model *ResourceObject, err error) {
	model = &ResourceObject{
		Requests: requests,
	}
	err = core.ValidateStruct(model, "required parameters")
	return
}

// UnmarshalResourceObject unmarshals an instance of ResourceObject from the specified map of raw messages.
func UnmarshalResourceObject(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ResourceObject)
	err = core.UnmarshalModel(m, "requests", &obj.Requests, UnmarshalResourceRequests)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "limits", &obj.Limits, UnmarshalResourceLimits)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ResourceObjectFabV1 : This field requires the use of Fabric v1.4.* and **lower**.
type ResourceObjectFabV1 struct {
	Requests *ResourceRequests `json:"requests" validate:"required"`

	Limits *ResourceLimits `json:"limits,omitempty"`
}


// NewResourceObjectFabV1 : Instantiate ResourceObjectFabV1 (Generic Model Constructor)
func (*BlockchainV2) NewResourceObjectFabV1(requests *ResourceRequests) (model *ResourceObjectFabV1, err error) {
	model = &ResourceObjectFabV1{
		Requests: requests,
	}
	err = core.ValidateStruct(model, "required parameters")
	return
}

// UnmarshalResourceObjectFabV1 unmarshals an instance of ResourceObjectFabV1 from the specified map of raw messages.
func UnmarshalResourceObjectFabV1(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ResourceObjectFabV1)
	err = core.UnmarshalModel(m, "requests", &obj.Requests, UnmarshalResourceRequests)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "limits", &obj.Limits, UnmarshalResourceLimits)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ResourceObjectFabV2 : This field requires the use of Fabric v2.1.* and higher.
type ResourceObjectFabV2 struct {
	Requests *ResourceRequests `json:"requests" validate:"required"`

	Limits *ResourceLimits `json:"limits,omitempty"`
}


// NewResourceObjectFabV2 : Instantiate ResourceObjectFabV2 (Generic Model Constructor)
func (*BlockchainV2) NewResourceObjectFabV2(requests *ResourceRequests) (model *ResourceObjectFabV2, err error) {
	model = &ResourceObjectFabV2{
		Requests: requests,
	}
	err = core.ValidateStruct(model, "required parameters")
	return
}

// UnmarshalResourceObjectFabV2 unmarshals an instance of ResourceObjectFabV2 from the specified map of raw messages.
func UnmarshalResourceObjectFabV2(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ResourceObjectFabV2)
	err = core.UnmarshalModel(m, "requests", &obj.Requests, UnmarshalResourceRequests)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "limits", &obj.Limits, UnmarshalResourceLimits)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ResourceRequests : ResourceRequests struct
type ResourceRequests struct {
	// Desired CPU for subcomponent. [Resource
	// details](/docs/blockchain?topic=blockchain-ibp-console-govern-components#ibp-console-govern-components-allocate-resources).
	Cpu *string `json:"cpu,omitempty"`

	// Desired memory for subcomponent. [Resource
	// details](/docs/blockchain?topic=blockchain-ibp-console-govern-components#ibp-console-govern-components-allocate-resources).
	Memory *string `json:"memory,omitempty"`
}


// UnmarshalResourceRequests unmarshals an instance of ResourceRequests from the specified map of raw messages.
func UnmarshalResourceRequests(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ResourceRequests)
	err = core.UnmarshalPrimitive(m, "cpu", &obj.Cpu)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "memory", &obj.Memory)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// RestartAthenaResponse : Describes the outcome of the api.
type RestartAthenaResponse struct {
	// Text describing the outcome of the api.
	Message *string `json:"message,omitempty"`
}


// UnmarshalRestartAthenaResponse unmarshals an instance of RestartAthenaResponse from the specified map of raw messages.
func UnmarshalRestartAthenaResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(RestartAthenaResponse)
	err = core.UnmarshalPrimitive(m, "message", &obj.Message)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// RestartOptions : The Restart options.
type RestartOptions struct {

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewRestartOptions : Instantiate RestartOptions
func (*BlockchainV2) NewRestartOptions() *RestartOptions {
	return &RestartOptions{}
}

// SetHeaders : Allow user to set Headers
func (options *RestartOptions) SetHeaders(param map[string]string) *RestartOptions {
	options.Headers = param
	return options
}

// SettingsTimestampData : SettingsTimestampData struct
type SettingsTimestampData struct {
	// UTC UNIX timestamp of the current time according to the server. In milliseconds.
	Now *float64 `json:"now,omitempty"`

	// UTC UNIX timestamp of when the server started. In milliseconds.
	Born *float64 `json:"born,omitempty"`

	// Time remaining until the server performs a hard-refresh of its settings.
	NextSettingsUpdate *string `json:"next_settings_update,omitempty"`

	// Total time the IBP console server has been running.
	UpTime *string `json:"up_time,omitempty"`
}


// UnmarshalSettingsTimestampData unmarshals an instance of SettingsTimestampData from the specified map of raw messages.
func UnmarshalSettingsTimestampData(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(SettingsTimestampData)
	err = core.UnmarshalPrimitive(m, "now", &obj.Now)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "born", &obj.Born)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "next_settings_update", &obj.NextSettingsUpdate)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "up_time", &obj.UpTime)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// StorageObject : StorageObject struct
type StorageObject struct {
	// Maximum disk space for subcomponent. [Resource
	// details](/docs/blockchain?topic=blockchain-ibp-console-govern-components#ibp-console-govern-components-allocate-resources).
	Size *string `json:"size,omitempty"`

	// Kubernetes storage class for subcomponent's disk space.
	Class *string `json:"class,omitempty"`
}


// UnmarshalStorageObject unmarshals an instance of StorageObject from the specified map of raw messages.
func UnmarshalStorageObject(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(StorageObject)
	err = core.UnmarshalPrimitive(m, "size", &obj.Size)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "class", &obj.Class)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// SubmitBlockOptions : The SubmitBlock options.
type SubmitBlockOptions struct {
	// The `id` of the component to modify. Use the [Get all components](#get-all-components) API to determine the
	// component id.
	ID *string `json:"id" validate:"required"`

	// The latest config block of the system channel. Base 64 encoded. To obtain this block, you must use a **Fabric API**.
	// This config block should list this ordering node as a valid consenter on the system-channel.
	B64Block *string `json:"b64_block,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewSubmitBlockOptions : Instantiate SubmitBlockOptions
func (*BlockchainV2) NewSubmitBlockOptions(id string) *SubmitBlockOptions {
	return &SubmitBlockOptions{
		ID: core.StringPtr(id),
	}
}

// SetID : Allow user to set ID
func (options *SubmitBlockOptions) SetID(id string) *SubmitBlockOptions {
	options.ID = core.StringPtr(id)
	return options
}

// SetB64Block : Allow user to set B64Block
func (options *SubmitBlockOptions) SetB64Block(b64Block string) *SubmitBlockOptions {
	options.B64Block = core.StringPtr(b64Block)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *SubmitBlockOptions) SetHeaders(param map[string]string) *SubmitBlockOptions {
	options.Headers = param
	return options
}

// UpdateCaBodyConfigOverride : Update the [Fabric CA configuration
// file](https://hyperledger-fabric-ca.readthedocs.io/en/release-1.4/serverconfig.html) if you want use custom
// attributes to configure advanced CA features. Omit if not.
//
// *The field **names** below are not case-sensitive.*.
type UpdateCaBodyConfigOverride struct {
	Ca *ConfigCAUpdate `json:"ca" validate:"required"`
}


// NewUpdateCaBodyConfigOverride : Instantiate UpdateCaBodyConfigOverride (Generic Model Constructor)
func (*BlockchainV2) NewUpdateCaBodyConfigOverride(ca *ConfigCAUpdate) (model *UpdateCaBodyConfigOverride, err error) {
	model = &UpdateCaBodyConfigOverride{
		Ca: ca,
	}
	err = core.ValidateStruct(model, "required parameters")
	return
}

// UnmarshalUpdateCaBodyConfigOverride unmarshals an instance of UpdateCaBodyConfigOverride from the specified map of raw messages.
func UnmarshalUpdateCaBodyConfigOverride(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(UpdateCaBodyConfigOverride)
	err = core.UnmarshalModel(m, "ca", &obj.Ca, UnmarshalConfigCAUpdate)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// UpdateCaBodyResources : CPU and memory properties. This feature is not available if using a free Kubernetes cluster.
type UpdateCaBodyResources struct {
	// This field requires the use of Fabric v1.4.* and higher.
	Ca *ResourceObject `json:"ca" validate:"required"`
}


// NewUpdateCaBodyResources : Instantiate UpdateCaBodyResources (Generic Model Constructor)
func (*BlockchainV2) NewUpdateCaBodyResources(ca *ResourceObject) (model *UpdateCaBodyResources, err error) {
	model = &UpdateCaBodyResources{
		Ca: ca,
	}
	err = core.ValidateStruct(model, "required parameters")
	return
}

// UnmarshalUpdateCaBodyResources unmarshals an instance of UpdateCaBodyResources from the specified map of raw messages.
func UnmarshalUpdateCaBodyResources(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(UpdateCaBodyResources)
	err = core.UnmarshalModel(m, "ca", &obj.Ca, UnmarshalResourceObject)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// UpdateCaOptions : The UpdateCa options.
type UpdateCaOptions struct {
	// The `id` of the component to modify. Use the [Get all components](#get-all-components) API to determine the
	// component id.
	ID *string `json:"id" validate:"required"`

	// CPU and memory properties. This feature is not available if using a free Kubernetes cluster.
	Resources *UpdateCaBodyResources `json:"resources,omitempty"`

	// Specify the Kubernetes zone for the deployment. The deployment will use a k8s node in this zone. Find the list of
	// possible zones by retrieving your Kubernetes node labels: `kubectl get nodes --show-labels`. [More
	// information](https://kubernetes.io/docs/setup/best-practices/multiple-zones).
	Zone *string `json:"zone,omitempty"`

	// Update the [Fabric CA configuration
	// file](https://hyperledger-fabric-ca.readthedocs.io/en/release-1.4/serverconfig.html) if you want use custom
	// attributes to configure advanced CA features. Omit if not.
	//
	// *The field **names** below are not case-sensitive.*.
	ConfigOverride *UpdateCaBodyConfigOverride `json:"config_override,omitempty"`

	// The number of replica pods running at any given time.
	Replicas *float64 `json:"replicas,omitempty"`

	// The Hyperledger Fabric release version to use.
	Version *string `json:"version,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewUpdateCaOptions : Instantiate UpdateCaOptions
func (*BlockchainV2) NewUpdateCaOptions(id string) *UpdateCaOptions {
	return &UpdateCaOptions{
		ID: core.StringPtr(id),
	}
}

// SetID : Allow user to set ID
func (options *UpdateCaOptions) SetID(id string) *UpdateCaOptions {
	options.ID = core.StringPtr(id)
	return options
}

// SetResources : Allow user to set Resources
func (options *UpdateCaOptions) SetResources(resources *UpdateCaBodyResources) *UpdateCaOptions {
	options.Resources = resources
	return options
}

// SetZone : Allow user to set Zone
func (options *UpdateCaOptions) SetZone(zone string) *UpdateCaOptions {
	options.Zone = core.StringPtr(zone)
	return options
}

// SetConfigOverride : Allow user to set ConfigOverride
func (options *UpdateCaOptions) SetConfigOverride(configOverride *UpdateCaBodyConfigOverride) *UpdateCaOptions {
	options.ConfigOverride = configOverride
	return options
}

// SetReplicas : Allow user to set Replicas
func (options *UpdateCaOptions) SetReplicas(replicas float64) *UpdateCaOptions {
	options.Replicas = core.Float64Ptr(replicas)
	return options
}

// SetVersion : Allow user to set Version
func (options *UpdateCaOptions) SetVersion(version string) *UpdateCaOptions {
	options.Version = core.StringPtr(version)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *UpdateCaOptions) SetHeaders(param map[string]string) *UpdateCaOptions {
	options.Headers = param
	return options
}

// UpdateOrdererBodyResources : CPU and memory properties. This feature is not available if using a free Kubernetes cluster.
type UpdateOrdererBodyResources struct {
	// This field requires the use of Fabric v1.4.* and higher.
	Orderer *ResourceObject `json:"orderer,omitempty"`

	// This field requires the use of Fabric v1.4.* and higher.
	Proxy *ResourceObject `json:"proxy,omitempty"`
}


// UnmarshalUpdateOrdererBodyResources unmarshals an instance of UpdateOrdererBodyResources from the specified map of raw messages.
func UnmarshalUpdateOrdererBodyResources(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(UpdateOrdererBodyResources)
	err = core.UnmarshalModel(m, "orderer", &obj.Orderer, UnmarshalResourceObject)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "proxy", &obj.Proxy, UnmarshalResourceObject)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// UpdateOrdererOptions : The UpdateOrderer options.
type UpdateOrdererOptions struct {
	// The `id` of the component to modify. Use the [Get all components](#get-all-components) API to determine the
	// component id.
	ID *string `json:"id" validate:"required"`

	// Update the [Fabric Orderer configuration
	// file](https://github.com/hyperledger/fabric/blob/release-1.4/sampleconfig/orderer.yaml) if you want use custom
	// attributes to configure the Orderer. Omit if not.
	//
	// *The field **names** below are not case-sensitive.*.
	ConfigOverride *ConfigOrdererUpdate `json:"config_override,omitempty"`

	// CPU and memory properties. This feature is not available if using a free Kubernetes cluster.
	Resources *UpdateOrdererBodyResources `json:"resources,omitempty"`

	// Specify the Kubernetes zone for the deployment. The deployment will use a k8s node in this zone. Find the list of
	// possible zones by retrieving your Kubernetes node labels: `kubectl get nodes --show-labels`. [More
	// information](https://kubernetes.io/docs/setup/best-practices/multiple-zones).
	Zone *string `json:"zone,omitempty"`

	// The Hyperledger Fabric release version to use.
	Version *string `json:"version,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewUpdateOrdererOptions : Instantiate UpdateOrdererOptions
func (*BlockchainV2) NewUpdateOrdererOptions(id string) *UpdateOrdererOptions {
	return &UpdateOrdererOptions{
		ID: core.StringPtr(id),
	}
}

// SetID : Allow user to set ID
func (options *UpdateOrdererOptions) SetID(id string) *UpdateOrdererOptions {
	options.ID = core.StringPtr(id)
	return options
}

// SetConfigOverride : Allow user to set ConfigOverride
func (options *UpdateOrdererOptions) SetConfigOverride(configOverride *ConfigOrdererUpdate) *UpdateOrdererOptions {
	options.ConfigOverride = configOverride
	return options
}

// SetResources : Allow user to set Resources
func (options *UpdateOrdererOptions) SetResources(resources *UpdateOrdererBodyResources) *UpdateOrdererOptions {
	options.Resources = resources
	return options
}

// SetZone : Allow user to set Zone
func (options *UpdateOrdererOptions) SetZone(zone string) *UpdateOrdererOptions {
	options.Zone = core.StringPtr(zone)
	return options
}

// SetVersion : Allow user to set Version
func (options *UpdateOrdererOptions) SetVersion(version string) *UpdateOrdererOptions {
	options.Version = core.StringPtr(version)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *UpdateOrdererOptions) SetHeaders(param map[string]string) *UpdateOrdererOptions {
	options.Headers = param
	return options
}

// UpdatePeerOptions : The UpdatePeer options.
type UpdatePeerOptions struct {
	// The `id` of the component to modify. Use the [Get all components](#get-all-components) API to determine the
	// component id.
	ID *string `json:"id" validate:"required"`

	// Update the [Fabric Peer configuration
	// file](https://github.com/hyperledger/fabric/blob/release-1.4/sampleconfig/core.yaml) if you want use custom
	// attributes to configure the Peer. Omit if not.
	//
	// *The field **names** below are not case-sensitive.*.
	ConfigOverride *ConfigPeerUpdate `json:"config_override,omitempty"`

	// CPU and memory properties. This feature is not available if using a free Kubernetes cluster.
	Resources *PeerResources `json:"resources,omitempty"`

	// Specify the Kubernetes zone for the deployment. The deployment will use a k8s node in this zone. Find the list of
	// possible zones by retrieving your Kubernetes node labels: `kubectl get nodes --show-labels`. [More
	// information](https://kubernetes.io/docs/setup/best-practices/multiple-zones).
	Zone *string `json:"zone,omitempty"`

	// The Hyperledger Fabric release version to use.
	Version *string `json:"version,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewUpdatePeerOptions : Instantiate UpdatePeerOptions
func (*BlockchainV2) NewUpdatePeerOptions(id string) *UpdatePeerOptions {
	return &UpdatePeerOptions{
		ID: core.StringPtr(id),
	}
}

// SetID : Allow user to set ID
func (options *UpdatePeerOptions) SetID(id string) *UpdatePeerOptions {
	options.ID = core.StringPtr(id)
	return options
}

// SetConfigOverride : Allow user to set ConfigOverride
func (options *UpdatePeerOptions) SetConfigOverride(configOverride *ConfigPeerUpdate) *UpdatePeerOptions {
	options.ConfigOverride = configOverride
	return options
}

// SetResources : Allow user to set Resources
func (options *UpdatePeerOptions) SetResources(resources *PeerResources) *UpdatePeerOptions {
	options.Resources = resources
	return options
}

// SetZone : Allow user to set Zone
func (options *UpdatePeerOptions) SetZone(zone string) *UpdatePeerOptions {
	options.Zone = core.StringPtr(zone)
	return options
}

// SetVersion : Allow user to set Version
func (options *UpdatePeerOptions) SetVersion(version string) *UpdatePeerOptions {
	options.Version = core.StringPtr(version)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *UpdatePeerOptions) SetHeaders(param map[string]string) *UpdatePeerOptions {
	options.Headers = param
	return options
}

// Hsm : The connection details of the HSM (Hardware Security Module).
type Hsm struct {
	// The url to the HSM. Include the protocol, hostname, and port.
	Pkcs11endpoint *string `json:"pkcs11endpoint" validate:"required"`
}


// NewHsm : Instantiate Hsm (Generic Model Constructor)
func (*BlockchainV2) NewHsm(pkcs11endpoint string) (model *Hsm, err error) {
	model = &Hsm{
		Pkcs11endpoint: core.StringPtr(pkcs11endpoint),
	}
	err = core.ValidateStruct(model, "required parameters")
	return
}

// UnmarshalHsm unmarshals an instance of Hsm from the specified map of raw messages.
func UnmarshalHsm(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Hsm)
	err = core.UnmarshalPrimitive(m, "pkcs11endpoint", &obj.Pkcs11endpoint)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// IdentityAttrs : IdentityAttrs struct
type IdentityAttrs struct {
	HfRegistrarRoles *string `json:"hf.Registrar.Roles,omitempty"`

	HfRegistrarDelegateRoles *string `json:"hf.Registrar.DelegateRoles,omitempty"`

	HfRevoker *bool `json:"hf.Revoker,omitempty"`

	HfIntermediateCA *bool `json:"hf.IntermediateCA,omitempty"`

	HfGenCRL *bool `json:"hf.GenCRL,omitempty"`

	HfRegistrarAttributes *string `json:"hf.Registrar.Attributes,omitempty"`

	HfAffiliationMgr *bool `json:"hf.AffiliationMgr,omitempty"`
}


// UnmarshalIdentityAttrs unmarshals an instance of IdentityAttrs from the specified map of raw messages.
func UnmarshalIdentityAttrs(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(IdentityAttrs)
	err = core.UnmarshalPrimitive(m, "hf.Registrar.Roles", &obj.HfRegistrarRoles)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "hf.Registrar.DelegateRoles", &obj.HfRegistrarDelegateRoles)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "hf.Revoker", &obj.HfRevoker)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "hf.IntermediateCA", &obj.HfIntermediateCA)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "hf.GenCRL", &obj.HfGenCRL)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "hf.Registrar.Attributes", &obj.HfRegistrarAttributes)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "hf.AffiliationMgr", &obj.HfAffiliationMgr)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}
