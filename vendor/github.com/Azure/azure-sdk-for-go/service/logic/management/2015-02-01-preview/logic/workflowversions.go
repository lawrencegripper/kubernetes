package logic

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

// WorkflowVersionsClient is the rEST API for Azure Logic Apps.
type WorkflowVersionsClient struct {
	ManagementClient
}

// NewWorkflowVersionsClient creates an instance of the WorkflowVersionsClient
// client.
func NewWorkflowVersionsClient(subscriptionID string) WorkflowVersionsClient {
	return NewWorkflowVersionsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewWorkflowVersionsClientWithBaseURI creates an instance of the
// WorkflowVersionsClient client.
func NewWorkflowVersionsClientWithBaseURI(baseURI string, subscriptionID string) WorkflowVersionsClient {
	return WorkflowVersionsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Get gets a workflow version.
//
// resourceGroupName is the resource group name. workflowName is the workflow
// name. versionID is the workflow versionId.
func (client WorkflowVersionsClient) Get(resourceGroupName string, workflowName string, versionID string) (result WorkflowVersion, err error) {
	req, err := client.GetPreparer(resourceGroupName, workflowName, versionID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "logic.WorkflowVersionsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "logic.WorkflowVersionsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "logic.WorkflowVersionsClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client WorkflowVersionsClient) GetPreparer(resourceGroupName string, workflowName string, versionID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"versionId":         autorest.Encode("path", versionID),
		"workflowName":      autorest.Encode("path", workflowName),
	}

	const APIVersion = "2015-02-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/workflows/{workflowName}/versions/{versionId}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client WorkflowVersionsClient) GetSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client WorkflowVersionsClient) GetResponder(resp *http.Response) (result WorkflowVersion, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
