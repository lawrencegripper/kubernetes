package backups

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

// RestorableDroppedDatabasesClient is the provides read functionality for
// Azure SQL Database Backups
type RestorableDroppedDatabasesClient struct {
	ManagementClient
}

// NewRestorableDroppedDatabasesClient creates an instance of the
// RestorableDroppedDatabasesClient client.
func NewRestorableDroppedDatabasesClient(subscriptionID string) RestorableDroppedDatabasesClient {
	return NewRestorableDroppedDatabasesClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewRestorableDroppedDatabasesClientWithBaseURI creates an instance of the
// RestorableDroppedDatabasesClient client.
func NewRestorableDroppedDatabasesClientWithBaseURI(baseURI string, subscriptionID string) RestorableDroppedDatabasesClient {
	return RestorableDroppedDatabasesClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Get gets a deleted database that can be restored
//
// resourceGroupName is the name of the resource group that contains the
// resource. You can obtain this value from the Azure Resource Manager API or
// the portal. serverName is the name of the server.
// restorableDroppededDatabaseID is the id of the deleted database in the form
// of databaseName,deletionTimeInFileTimeFormat
func (client RestorableDroppedDatabasesClient) Get(resourceGroupName string, serverName string, restorableDroppededDatabaseID string) (result RestorableDroppedDatabase, err error) {
	req, err := client.GetPreparer(resourceGroupName, serverName, restorableDroppededDatabaseID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "backups.RestorableDroppedDatabasesClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "backups.RestorableDroppedDatabasesClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "backups.RestorableDroppedDatabasesClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client RestorableDroppedDatabasesClient) GetPreparer(resourceGroupName string, serverName string, restorableDroppededDatabaseID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName":             autorest.Encode("path", resourceGroupName),
		"restorableDroppededDatabaseId": autorest.Encode("path", restorableDroppededDatabaseID),
		"serverName":                    autorest.Encode("path", serverName),
		"subscriptionId":                autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2014-04-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/restorableDroppedDatabases/{restorableDroppededDatabaseId}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client RestorableDroppedDatabasesClient) GetSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client RestorableDroppedDatabasesClient) GetResponder(resp *http.Response) (result RestorableDroppedDatabase, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByServer gets a list of deleted databases that can be restored
//
// resourceGroupName is the name of the resource group that contains the
// resource. You can obtain this value from the Azure Resource Manager API or
// the portal. serverName is the name of the server.
func (client RestorableDroppedDatabasesClient) ListByServer(resourceGroupName string, serverName string) (result RestorableDroppedDatabaseListResult, err error) {
	req, err := client.ListByServerPreparer(resourceGroupName, serverName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "backups.RestorableDroppedDatabasesClient", "ListByServer", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByServerSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "backups.RestorableDroppedDatabasesClient", "ListByServer", resp, "Failure sending request")
		return
	}

	result, err = client.ListByServerResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "backups.RestorableDroppedDatabasesClient", "ListByServer", resp, "Failure responding to request")
	}

	return
}

// ListByServerPreparer prepares the ListByServer request.
func (client RestorableDroppedDatabasesClient) ListByServerPreparer(resourceGroupName string, serverName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serverName":        autorest.Encode("path", serverName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2014-04-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/restorableDroppedDatabases", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// ListByServerSender sends the ListByServer request. The method will close the
// http.Response Body if it receives an error.
func (client RestorableDroppedDatabasesClient) ListByServerSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// ListByServerResponder handles the response to the ListByServer request. The method always
// closes the http.Response Body.
func (client RestorableDroppedDatabasesClient) ListByServerResponder(resp *http.Response) (result RestorableDroppedDatabaseListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
