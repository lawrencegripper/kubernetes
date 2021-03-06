package servicefabric

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

// ClusterVersionsClient is the client for the ClusterVersions methods of the
// Servicefabric service.
type ClusterVersionsClient struct {
	ManagementClient
}

// NewClusterVersionsClient creates an instance of the ClusterVersionsClient
// client.
func NewClusterVersionsClient(subscriptionID string) ClusterVersionsClient {
	return NewClusterVersionsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewClusterVersionsClientWithBaseURI creates an instance of the
// ClusterVersionsClient client.
func NewClusterVersionsClientWithBaseURI(baseURI string, subscriptionID string) ClusterVersionsClient {
	return ClusterVersionsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Get get cluster code versions by environment and version
//
// location is the location for the cluster code versions, this is different
// from cluster location environment is cluster operating system, the default
// means all clusterVersion is the cluster code version
func (client ClusterVersionsClient) Get(location string, environment string, clusterVersion string) (result ClusterCodeVersionsResult, err error) {
	req, err := client.GetPreparer(location, environment, clusterVersion)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicefabric.ClusterVersionsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "servicefabric.ClusterVersionsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicefabric.ClusterVersionsClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client ClusterVersionsClient) GetPreparer(location string, environment string, clusterVersion string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"clusterVersion": autorest.Encode("path", clusterVersion),
		"environment":    autorest.Encode("path", environment),
		"location":       autorest.Encode("path", location),
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-09-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.ServiceFabric/locations/{location}/environments/{environment}/clusterVersions/{clusterVersion}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client ClusterVersionsClient) GetSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client ClusterVersionsClient) GetResponder(resp *http.Response) (result ClusterCodeVersionsResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List list cluster code versions by location
//
// location is the location for the cluster code versions, this is different
// from cluster location
func (client ClusterVersionsClient) List(location string) (result ClusterCodeVersionsListResult, err error) {
	req, err := client.ListPreparer(location)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicefabric.ClusterVersionsClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "servicefabric.ClusterVersionsClient", "List", resp, "Failure sending request")
		return
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicefabric.ClusterVersionsClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client ClusterVersionsClient) ListPreparer(location string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"location":       autorest.Encode("path", location),
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-09-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.ServiceFabric/locations/{location}/clusterVersions", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client ClusterVersionsClient) ListSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client ClusterVersionsClient) ListResponder(resp *http.Response) (result ClusterCodeVersionsListResult, err error) {
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
func (client ClusterVersionsClient) ListNextResults(lastResults ClusterCodeVersionsListResult) (result ClusterCodeVersionsListResult, err error) {
	req, err := lastResults.ClusterCodeVersionsListResultPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "servicefabric.ClusterVersionsClient", "List", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "servicefabric.ClusterVersionsClient", "List", resp, "Failure sending next results request")
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicefabric.ClusterVersionsClient", "List", resp, "Failure responding to next results request")
	}

	return
}

// ListComplete gets all elements from the list without paging.
func (client ClusterVersionsClient) ListComplete(location string, cancel <-chan struct{}) (<-chan ClusterCodeVersionsResult, <-chan error) {
	resultChan := make(chan ClusterCodeVersionsResult)
	errChan := make(chan error, 1)
	go func() {
		defer func() {
			close(resultChan)
			close(errChan)
		}()
		list, err := client.List(location)
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

// ListByEnvironment list cluster code versions by environment
//
// location is the location for the cluster code versions, this is different
// from cluster location environment is cluster operating system, the default
// means all
func (client ClusterVersionsClient) ListByEnvironment(location string, environment string) (result ClusterCodeVersionsListResult, err error) {
	req, err := client.ListByEnvironmentPreparer(location, environment)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicefabric.ClusterVersionsClient", "ListByEnvironment", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByEnvironmentSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "servicefabric.ClusterVersionsClient", "ListByEnvironment", resp, "Failure sending request")
		return
	}

	result, err = client.ListByEnvironmentResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicefabric.ClusterVersionsClient", "ListByEnvironment", resp, "Failure responding to request")
	}

	return
}

// ListByEnvironmentPreparer prepares the ListByEnvironment request.
func (client ClusterVersionsClient) ListByEnvironmentPreparer(location string, environment string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"environment":    autorest.Encode("path", environment),
		"location":       autorest.Encode("path", location),
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-09-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.ServiceFabric/locations/{location}/environments/{environment}/clusterVersions", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// ListByEnvironmentSender sends the ListByEnvironment request. The method will close the
// http.Response Body if it receives an error.
func (client ClusterVersionsClient) ListByEnvironmentSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// ListByEnvironmentResponder handles the response to the ListByEnvironment request. The method always
// closes the http.Response Body.
func (client ClusterVersionsClient) ListByEnvironmentResponder(resp *http.Response) (result ClusterCodeVersionsListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByEnvironmentNextResults retrieves the next set of results, if any.
func (client ClusterVersionsClient) ListByEnvironmentNextResults(lastResults ClusterCodeVersionsListResult) (result ClusterCodeVersionsListResult, err error) {
	req, err := lastResults.ClusterCodeVersionsListResultPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "servicefabric.ClusterVersionsClient", "ListByEnvironment", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}

	resp, err := client.ListByEnvironmentSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "servicefabric.ClusterVersionsClient", "ListByEnvironment", resp, "Failure sending next results request")
	}

	result, err = client.ListByEnvironmentResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicefabric.ClusterVersionsClient", "ListByEnvironment", resp, "Failure responding to next results request")
	}

	return
}

// ListByEnvironmentComplete gets all elements from the list without paging.
func (client ClusterVersionsClient) ListByEnvironmentComplete(location string, environment string, cancel <-chan struct{}) (<-chan ClusterCodeVersionsResult, <-chan error) {
	resultChan := make(chan ClusterCodeVersionsResult)
	errChan := make(chan error, 1)
	go func() {
		defer func() {
			close(resultChan)
			close(errChan)
		}()
		list, err := client.ListByEnvironment(location, environment)
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
			list, err = client.ListByEnvironmentNextResults(list)
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

// ListByVersion list cluster code versions by version
//
// location is the location for the cluster code versions, this is different
// from cluster location clusterVersion is the cluster code version
func (client ClusterVersionsClient) ListByVersion(location string, clusterVersion string) (result ClusterCodeVersionsListResult, err error) {
	req, err := client.ListByVersionPreparer(location, clusterVersion)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicefabric.ClusterVersionsClient", "ListByVersion", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByVersionSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "servicefabric.ClusterVersionsClient", "ListByVersion", resp, "Failure sending request")
		return
	}

	result, err = client.ListByVersionResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicefabric.ClusterVersionsClient", "ListByVersion", resp, "Failure responding to request")
	}

	return
}

// ListByVersionPreparer prepares the ListByVersion request.
func (client ClusterVersionsClient) ListByVersionPreparer(location string, clusterVersion string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"clusterVersion": autorest.Encode("path", clusterVersion),
		"location":       autorest.Encode("path", location),
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-09-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.ServiceFabric/locations/{location}/clusterVersions/{clusterVersion}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// ListByVersionSender sends the ListByVersion request. The method will close the
// http.Response Body if it receives an error.
func (client ClusterVersionsClient) ListByVersionSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// ListByVersionResponder handles the response to the ListByVersion request. The method always
// closes the http.Response Body.
func (client ClusterVersionsClient) ListByVersionResponder(resp *http.Response) (result ClusterCodeVersionsListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByVersionNextResults retrieves the next set of results, if any.
func (client ClusterVersionsClient) ListByVersionNextResults(lastResults ClusterCodeVersionsListResult) (result ClusterCodeVersionsListResult, err error) {
	req, err := lastResults.ClusterCodeVersionsListResultPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "servicefabric.ClusterVersionsClient", "ListByVersion", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}

	resp, err := client.ListByVersionSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "servicefabric.ClusterVersionsClient", "ListByVersion", resp, "Failure sending next results request")
	}

	result, err = client.ListByVersionResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicefabric.ClusterVersionsClient", "ListByVersion", resp, "Failure responding to next results request")
	}

	return
}

// ListByVersionComplete gets all elements from the list without paging.
func (client ClusterVersionsClient) ListByVersionComplete(location string, clusterVersion string, cancel <-chan struct{}) (<-chan ClusterCodeVersionsResult, <-chan error) {
	resultChan := make(chan ClusterCodeVersionsResult)
	errChan := make(chan error, 1)
	go func() {
		defer func() {
			close(resultChan)
			close(errChan)
		}()
		list, err := client.ListByVersion(location, clusterVersion)
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
			list, err = client.ListByVersionNextResults(list)
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
