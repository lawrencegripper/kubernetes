package metricdefinitionsapi

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

// MetricDefinitionsClient is the client for the MetricDefinitions methods of
// the Metricdefinitionsapi service.
type MetricDefinitionsClient struct {
	ManagementClient
}

// NewMetricDefinitionsClient creates an instance of the
// MetricDefinitionsClient client.
func NewMetricDefinitionsClient() MetricDefinitionsClient {
	return NewMetricDefinitionsClientWithBaseURI(DefaultBaseURI)
}

// NewMetricDefinitionsClientWithBaseURI creates an instance of the
// MetricDefinitionsClient client.
func NewMetricDefinitionsClientWithBaseURI(baseURI string) MetricDefinitionsClient {
	return MetricDefinitionsClient{NewWithBaseURI(baseURI)}
}

// List lists the metric definitions for the resource.
//
// resourceURI is the identifier of the resource. filter is reduces the set of
// data collected by retrieving particular metric definitions from all the
// definitions available for the resource.<br>For example, to get just the
// definition for the 'CPU percentage' counter: $filter=name.value eq
// '\Processor(_Total)\% Processor Time'.<br>Multiple metrics can be retrieved
// by joining together *'name eq <value>'* clauses separated by *or* logical
// operators.<br>**NOTE**: No other syntax is allowed.
func (client MetricDefinitionsClient) List(resourceURI string, filter string) (result MetricDefinitionCollection, err error) {
	req, err := client.ListPreparer(resourceURI, filter)
	if err != nil {
		err = autorest.NewErrorWithError(err, "metricdefinitionsapi.MetricDefinitionsClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "metricdefinitionsapi.MetricDefinitionsClient", "List", resp, "Failure sending request")
		return
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "metricdefinitionsapi.MetricDefinitionsClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client MetricDefinitionsClient) ListPreparer(resourceURI string, filter string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceUri": resourceURI,
	}

	const APIVersion = "2016-03-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/{resourceUri}/providers/microsoft.insights/metricDefinitions", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client MetricDefinitionsClient) ListSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client MetricDefinitionsClient) ListResponder(resp *http.Response) (result MetricDefinitionCollection, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
