package service

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator 1.0.1.0
// Changes may cause incorrect behavior and will be lost if the code is
// regenerated.

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"net/http"
)

// ReplicationFabricsClient is the client for the ReplicationFabrics methods of
// the Service service.
type ReplicationFabricsClient struct {
	ManagementClient
}

// NewReplicationFabricsClient creates an instance of the
// ReplicationFabricsClient client.
func NewReplicationFabricsClient(subscriptionID string, resourceGroupName string, resourceName string) ReplicationFabricsClient {
	return NewReplicationFabricsClientWithBaseURI(DefaultBaseURI, subscriptionID, resourceGroupName, resourceName)
}

// NewReplicationFabricsClientWithBaseURI creates an instance of the
// ReplicationFabricsClient client.
func NewReplicationFabricsClientWithBaseURI(baseURI string, subscriptionID string, resourceGroupName string, resourceName string) ReplicationFabricsClient {
	return ReplicationFabricsClient{NewWithBaseURI(baseURI, subscriptionID, resourceGroupName, resourceName)}
}

// CheckConsistency the operation to perform a consistency check on the fabric.
// This method may poll for completion. Polling can be canceled by passing the
// cancel channel argument. The channel will be used to cancel polling and any
// outstanding HTTP requests.
//
// fabricName is fabric name.
func (client ReplicationFabricsClient) CheckConsistency(fabricName string, cancel <-chan struct{}) (<-chan Fabric, <-chan error) {
	resultChan := make(chan Fabric, 1)
	errChan := make(chan error, 1)
	go func() {
		var err error
		var result Fabric
		defer func() {
			if err != nil {
				errChan <- err
			}
			resultChan <- result
			close(resultChan)
			close(errChan)
		}()
		req, err := client.CheckConsistencyPreparer(fabricName, cancel)
		if err != nil {
			err = autorest.NewErrorWithError(err, "service.ReplicationFabricsClient", "CheckConsistency", nil, "Failure preparing request")
			return
		}

		resp, err := client.CheckConsistencySender(req)
		if err != nil {
			result.Response = autorest.Response{Response: resp}
			err = autorest.NewErrorWithError(err, "service.ReplicationFabricsClient", "CheckConsistency", resp, "Failure sending request")
			return
		}

		result, err = client.CheckConsistencyResponder(resp)
		if err != nil {
			err = autorest.NewErrorWithError(err, "service.ReplicationFabricsClient", "CheckConsistency", resp, "Failure responding to request")
		}
	}()
	return resultChan, errChan
}

// CheckConsistencyPreparer prepares the CheckConsistency request.
func (client ReplicationFabricsClient) CheckConsistencyPreparer(fabricName string, cancel <-chan struct{}) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"fabricName":        autorest.Encode("path", fabricName),
		"resourceGroupName": autorest.Encode("path", client.ResourceGroupName),
		"resourceName":      autorest.Encode("path", client.ResourceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-08-10"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/checkConsistency", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{Cancel: cancel})
}

// CheckConsistencySender sends the CheckConsistency request. The method will close the
// http.Response Body if it receives an error.
func (client ReplicationFabricsClient) CheckConsistencySender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		azure.DoPollForAsynchronous(client.PollingDelay))
}

// CheckConsistencyResponder handles the response to the CheckConsistency request. The method always
// closes the http.Response Body.
func (client ReplicationFabricsClient) CheckConsistencyResponder(resp *http.Response) (result Fabric, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Create the operation to create an Azure Site Recovery fabric (for e.g.
// Hyper-V site) This method may poll for completion. Polling can be canceled
// by passing the cancel channel argument. The channel will be used to cancel
// polling and any outstanding HTTP requests.
//
// fabricName is name of the ASR fabric. input is fabric creation input.
func (client ReplicationFabricsClient) Create(fabricName string, input FabricCreationInput, cancel <-chan struct{}) (<-chan Fabric, <-chan error) {
	resultChan := make(chan Fabric, 1)
	errChan := make(chan error, 1)
	go func() {
		var err error
		var result Fabric
		defer func() {
			if err != nil {
				errChan <- err
			}
			resultChan <- result
			close(resultChan)
			close(errChan)
		}()
		req, err := client.CreatePreparer(fabricName, input, cancel)
		if err != nil {
			err = autorest.NewErrorWithError(err, "service.ReplicationFabricsClient", "Create", nil, "Failure preparing request")
			return
		}

		resp, err := client.CreateSender(req)
		if err != nil {
			result.Response = autorest.Response{Response: resp}
			err = autorest.NewErrorWithError(err, "service.ReplicationFabricsClient", "Create", resp, "Failure sending request")
			return
		}

		result, err = client.CreateResponder(resp)
		if err != nil {
			err = autorest.NewErrorWithError(err, "service.ReplicationFabricsClient", "Create", resp, "Failure responding to request")
		}
	}()
	return resultChan, errChan
}

// CreatePreparer prepares the Create request.
func (client ReplicationFabricsClient) CreatePreparer(fabricName string, input FabricCreationInput, cancel <-chan struct{}) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"fabricName":        autorest.Encode("path", fabricName),
		"resourceGroupName": autorest.Encode("path", client.ResourceGroupName),
		"resourceName":      autorest.Encode("path", client.ResourceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-08-10"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}", pathParameters),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{Cancel: cancel})
}

// CreateSender sends the Create request. The method will close the
// http.Response Body if it receives an error.
func (client ReplicationFabricsClient) CreateSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		azure.DoPollForAsynchronous(client.PollingDelay))
}

// CreateResponder handles the response to the Create request. The method always
// closes the http.Response Body.
func (client ReplicationFabricsClient) CreateResponder(resp *http.Response) (result Fabric, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete the operation to delete or remove an Azure Site Recovery fabric. This
// method may poll for completion. Polling can be canceled by passing the
// cancel channel argument. The channel will be used to cancel polling and any
// outstanding HTTP requests.
//
// fabricName is aSR fabric to delete
func (client ReplicationFabricsClient) Delete(fabricName string, cancel <-chan struct{}) (<-chan autorest.Response, <-chan error) {
	resultChan := make(chan autorest.Response, 1)
	errChan := make(chan error, 1)
	go func() {
		var err error
		var result autorest.Response
		defer func() {
			if err != nil {
				errChan <- err
			}
			resultChan <- result
			close(resultChan)
			close(errChan)
		}()
		req, err := client.DeletePreparer(fabricName, cancel)
		if err != nil {
			err = autorest.NewErrorWithError(err, "service.ReplicationFabricsClient", "Delete", nil, "Failure preparing request")
			return
		}

		resp, err := client.DeleteSender(req)
		if err != nil {
			result.Response = resp
			err = autorest.NewErrorWithError(err, "service.ReplicationFabricsClient", "Delete", resp, "Failure sending request")
			return
		}

		result, err = client.DeleteResponder(resp)
		if err != nil {
			err = autorest.NewErrorWithError(err, "service.ReplicationFabricsClient", "Delete", resp, "Failure responding to request")
		}
	}()
	return resultChan, errChan
}

// DeletePreparer prepares the Delete request.
func (client ReplicationFabricsClient) DeletePreparer(fabricName string, cancel <-chan struct{}) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"fabricName":        autorest.Encode("path", fabricName),
		"resourceGroupName": autorest.Encode("path", client.ResourceGroupName),
		"resourceName":      autorest.Encode("path", client.ResourceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-08-10"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/remove", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{Cancel: cancel})
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client ReplicationFabricsClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		azure.DoPollForAsynchronous(client.PollingDelay))
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client ReplicationFabricsClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get gets the details of an Azure Site Recovery fabric.
//
// fabricName is fabric name.
func (client ReplicationFabricsClient) Get(fabricName string) (result Fabric, err error) {
	req, err := client.GetPreparer(fabricName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "service.ReplicationFabricsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "service.ReplicationFabricsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "service.ReplicationFabricsClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client ReplicationFabricsClient) GetPreparer(fabricName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"fabricName":        autorest.Encode("path", fabricName),
		"resourceGroupName": autorest.Encode("path", client.ResourceGroupName),
		"resourceName":      autorest.Encode("path", client.ResourceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-08-10"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client ReplicationFabricsClient) GetSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client ReplicationFabricsClient) GetResponder(resp *http.Response) (result Fabric, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List gets a list of the Azure Site Recovery fabrics in the vault.
func (client ReplicationFabricsClient) List() (result FabricCollection, err error) {
	req, err := client.ListPreparer()
	if err != nil {
		err = autorest.NewErrorWithError(err, "service.ReplicationFabricsClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "service.ReplicationFabricsClient", "List", resp, "Failure sending request")
		return
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "service.ReplicationFabricsClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client ReplicationFabricsClient) ListPreparer() (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", client.ResourceGroupName),
		"resourceName":      autorest.Encode("path", client.ResourceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-08-10"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client ReplicationFabricsClient) ListSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client ReplicationFabricsClient) ListResponder(resp *http.Response) (result FabricCollection, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListNextResults retrieves the next set of results, if any.
func (client ReplicationFabricsClient) ListNextResults(lastResults FabricCollection) (result FabricCollection, err error) {
	req, err := lastResults.FabricCollectionPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "service.ReplicationFabricsClient", "List", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "service.ReplicationFabricsClient", "List", resp, "Failure sending next results request")
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "service.ReplicationFabricsClient", "List", resp, "Failure responding to next results request")
	}

	return
}

// ListComplete gets all elements from the list without paging.
func (client ReplicationFabricsClient) ListComplete(cancel <-chan struct{}) (<-chan Fabric, <-chan error) {
	resultChan := make(chan Fabric)
	errChan := make(chan error, 1)
	go func() {
		defer func() {
			close(resultChan)
			close(errChan)
		}()
		list, err := client.List()
		if err != nil {
			errChan <- err
			return
		}
		if list.Value != nil {
			for _, item := range *list.Value {
				select {
				case <-cancel:
					return
				case resultChan <- item:
					// Intentionally left blank
				}
			}
		}
		for list.NextLink != nil {
			list, err = client.ListNextResults(list)
			if err != nil {
				errChan <- err
				return
			}
			if list.Value != nil {
				for _, item := range *list.Value {
					select {
					case <-cancel:
						return
					case resultChan <- item:
						// Intentionally left blank
					}
				}
			}
		}
	}()
	return resultChan, errChan
}

// MigrateToAad the operation to migrate an Azure Site Recovery fabric to AAD.
// This method may poll for completion. Polling can be canceled by passing the
// cancel channel argument. The channel will be used to cancel polling and any
// outstanding HTTP requests.
//
// fabricName is aSR fabric to migrate.
func (client ReplicationFabricsClient) MigrateToAad(fabricName string, cancel <-chan struct{}) (<-chan autorest.Response, <-chan error) {
	resultChan := make(chan autorest.Response, 1)
	errChan := make(chan error, 1)
	go func() {
		var err error
		var result autorest.Response
		defer func() {
			if err != nil {
				errChan <- err
			}
			resultChan <- result
			close(resultChan)
			close(errChan)
		}()
		req, err := client.MigrateToAadPreparer(fabricName, cancel)
		if err != nil {
			err = autorest.NewErrorWithError(err, "service.ReplicationFabricsClient", "MigrateToAad", nil, "Failure preparing request")
			return
		}

		resp, err := client.MigrateToAadSender(req)
		if err != nil {
			result.Response = resp
			err = autorest.NewErrorWithError(err, "service.ReplicationFabricsClient", "MigrateToAad", resp, "Failure sending request")
			return
		}

		result, err = client.MigrateToAadResponder(resp)
		if err != nil {
			err = autorest.NewErrorWithError(err, "service.ReplicationFabricsClient", "MigrateToAad", resp, "Failure responding to request")
		}
	}()
	return resultChan, errChan
}

// MigrateToAadPreparer prepares the MigrateToAad request.
func (client ReplicationFabricsClient) MigrateToAadPreparer(fabricName string, cancel <-chan struct{}) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"fabricName":        autorest.Encode("path", fabricName),
		"resourceGroupName": autorest.Encode("path", client.ResourceGroupName),
		"resourceName":      autorest.Encode("path", client.ResourceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-08-10"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/migratetoaad", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{Cancel: cancel})
}

// MigrateToAadSender sends the MigrateToAad request. The method will close the
// http.Response Body if it receives an error.
func (client ReplicationFabricsClient) MigrateToAadSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		azure.DoPollForAsynchronous(client.PollingDelay))
}

// MigrateToAadResponder handles the response to the MigrateToAad request. The method always
// closes the http.Response Body.
func (client ReplicationFabricsClient) MigrateToAadResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Purge the operation to purge(force delete) an Azure Site Recovery fabric.
// This method may poll for completion. Polling can be canceled by passing the
// cancel channel argument. The channel will be used to cancel polling and any
// outstanding HTTP requests.
//
// fabricName is aSR fabric to purge.
func (client ReplicationFabricsClient) Purge(fabricName string, cancel <-chan struct{}) (<-chan autorest.Response, <-chan error) {
	resultChan := make(chan autorest.Response, 1)
	errChan := make(chan error, 1)
	go func() {
		var err error
		var result autorest.Response
		defer func() {
			if err != nil {
				errChan <- err
			}
			resultChan <- result
			close(resultChan)
			close(errChan)
		}()
		req, err := client.PurgePreparer(fabricName, cancel)
		if err != nil {
			err = autorest.NewErrorWithError(err, "service.ReplicationFabricsClient", "Purge", nil, "Failure preparing request")
			return
		}

		resp, err := client.PurgeSender(req)
		if err != nil {
			result.Response = resp
			err = autorest.NewErrorWithError(err, "service.ReplicationFabricsClient", "Purge", resp, "Failure sending request")
			return
		}

		result, err = client.PurgeResponder(resp)
		if err != nil {
			err = autorest.NewErrorWithError(err, "service.ReplicationFabricsClient", "Purge", resp, "Failure responding to request")
		}
	}()
	return resultChan, errChan
}

// PurgePreparer prepares the Purge request.
func (client ReplicationFabricsClient) PurgePreparer(fabricName string, cancel <-chan struct{}) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"fabricName":        autorest.Encode("path", fabricName),
		"resourceGroupName": autorest.Encode("path", client.ResourceGroupName),
		"resourceName":      autorest.Encode("path", client.ResourceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-08-10"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{Cancel: cancel})
}

// PurgeSender sends the Purge request. The method will close the
// http.Response Body if it receives an error.
func (client ReplicationFabricsClient) PurgeSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		azure.DoPollForAsynchronous(client.PollingDelay))
}

// PurgeResponder handles the response to the Purge request. The method always
// closes the http.Response Body.
func (client ReplicationFabricsClient) PurgeResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// ReassociateGateway the operation to move replications from a process server
// to another process server. This method may poll for completion. Polling can
// be canceled by passing the cancel channel argument. The channel will be used
// to cancel polling and any outstanding HTTP requests.
//
// fabricName is the name of the fabric containing the process server.
// failoverProcessServerRequest is the input to the failover process server
// operation.
func (client ReplicationFabricsClient) ReassociateGateway(fabricName string, failoverProcessServerRequest FailoverProcessServerRequest, cancel <-chan struct{}) (<-chan Fabric, <-chan error) {
	resultChan := make(chan Fabric, 1)
	errChan := make(chan error, 1)
	go func() {
		var err error
		var result Fabric
		defer func() {
			if err != nil {
				errChan <- err
			}
			resultChan <- result
			close(resultChan)
			close(errChan)
		}()
		req, err := client.ReassociateGatewayPreparer(fabricName, failoverProcessServerRequest, cancel)
		if err != nil {
			err = autorest.NewErrorWithError(err, "service.ReplicationFabricsClient", "ReassociateGateway", nil, "Failure preparing request")
			return
		}

		resp, err := client.ReassociateGatewaySender(req)
		if err != nil {
			result.Response = autorest.Response{Response: resp}
			err = autorest.NewErrorWithError(err, "service.ReplicationFabricsClient", "ReassociateGateway", resp, "Failure sending request")
			return
		}

		result, err = client.ReassociateGatewayResponder(resp)
		if err != nil {
			err = autorest.NewErrorWithError(err, "service.ReplicationFabricsClient", "ReassociateGateway", resp, "Failure responding to request")
		}
	}()
	return resultChan, errChan
}

// ReassociateGatewayPreparer prepares the ReassociateGateway request.
func (client ReplicationFabricsClient) ReassociateGatewayPreparer(fabricName string, failoverProcessServerRequest FailoverProcessServerRequest, cancel <-chan struct{}) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"fabricName":        autorest.Encode("path", fabricName),
		"resourceGroupName": autorest.Encode("path", client.ResourceGroupName),
		"resourceName":      autorest.Encode("path", client.ResourceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-08-10"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/reassociateGateway", pathParameters),
		autorest.WithJSON(failoverProcessServerRequest),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{Cancel: cancel})
}

// ReassociateGatewaySender sends the ReassociateGateway request. The method will close the
// http.Response Body if it receives an error.
func (client ReplicationFabricsClient) ReassociateGatewaySender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		azure.DoPollForAsynchronous(client.PollingDelay))
}

// ReassociateGatewayResponder handles the response to the ReassociateGateway request. The method always
// closes the http.Response Body.
func (client ReplicationFabricsClient) ReassociateGatewayResponder(resp *http.Response) (result Fabric, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// RenewCertificate renews the connection certificate for the ASR replication
// fabric. This method may poll for completion. Polling can be canceled by
// passing the cancel channel argument. The channel will be used to cancel
// polling and any outstanding HTTP requests.
//
// fabricName is fabric name to renew certs for. renewCertificate is renew
// certificate input.
func (client ReplicationFabricsClient) RenewCertificate(fabricName string, renewCertificate RenewCertificateInput, cancel <-chan struct{}) (<-chan Fabric, <-chan error) {
	resultChan := make(chan Fabric, 1)
	errChan := make(chan error, 1)
	go func() {
		var err error
		var result Fabric
		defer func() {
			if err != nil {
				errChan <- err
			}
			resultChan <- result
			close(resultChan)
			close(errChan)
		}()
		req, err := client.RenewCertificatePreparer(fabricName, renewCertificate, cancel)
		if err != nil {
			err = autorest.NewErrorWithError(err, "service.ReplicationFabricsClient", "RenewCertificate", nil, "Failure preparing request")
			return
		}

		resp, err := client.RenewCertificateSender(req)
		if err != nil {
			result.Response = autorest.Response{Response: resp}
			err = autorest.NewErrorWithError(err, "service.ReplicationFabricsClient", "RenewCertificate", resp, "Failure sending request")
			return
		}

		result, err = client.RenewCertificateResponder(resp)
		if err != nil {
			err = autorest.NewErrorWithError(err, "service.ReplicationFabricsClient", "RenewCertificate", resp, "Failure responding to request")
		}
	}()
	return resultChan, errChan
}

// RenewCertificatePreparer prepares the RenewCertificate request.
func (client ReplicationFabricsClient) RenewCertificatePreparer(fabricName string, renewCertificate RenewCertificateInput, cancel <-chan struct{}) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"fabricName":        autorest.Encode("path", fabricName),
		"resourceGroupName": autorest.Encode("path", client.ResourceGroupName),
		"resourceName":      autorest.Encode("path", client.ResourceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-08-10"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/renewCertificate", pathParameters),
		autorest.WithJSON(renewCertificate),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{Cancel: cancel})
}

// RenewCertificateSender sends the RenewCertificate request. The method will close the
// http.Response Body if it receives an error.
func (client ReplicationFabricsClient) RenewCertificateSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		azure.DoPollForAsynchronous(client.PollingDelay))
}

// RenewCertificateResponder handles the response to the RenewCertificate request. The method always
// closes the http.Response Body.
func (client ReplicationFabricsClient) RenewCertificateResponder(resp *http.Response) (result Fabric, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
