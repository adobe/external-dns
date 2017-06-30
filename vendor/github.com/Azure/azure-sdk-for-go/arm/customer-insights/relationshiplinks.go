package customerinsights

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
	"github.com/Azure/go-autorest/autorest/validation"
	"net/http"
)

// RelationshipLinksClient is the the Azure Customer Insights management API
// provides a RESTful set of web services that interact with Azure Customer
// Insights service to manage your resources. The API has entities that capture
// the relationship between an end user and the Azure Customer Insights
// service.
type RelationshipLinksClient struct {
	ManagementClient
}

// NewRelationshipLinksClient creates an instance of the
// RelationshipLinksClient client.
func NewRelationshipLinksClient(subscriptionID string) RelationshipLinksClient {
	return NewRelationshipLinksClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewRelationshipLinksClientWithBaseURI creates an instance of the
// RelationshipLinksClient client.
func NewRelationshipLinksClientWithBaseURI(baseURI string, subscriptionID string) RelationshipLinksClient {
	return RelationshipLinksClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate creates a relationship link or updates an existing
// relationship link within a hub. This method may poll for completion. Polling
// can be canceled by passing the cancel channel argument. The channel will be
// used to cancel polling and any outstanding HTTP requests.
//
// resourceGroupName is the name of the resource group. hubName is the name of
// the hub. relationshipLinkName is the name of the relationship link.
// parameters is parameters supplied to the CreateOrUpdate relationship link
// operation.
func (client RelationshipLinksClient) CreateOrUpdate(resourceGroupName string, hubName string, relationshipLinkName string, parameters RelationshipLinkResourceFormat, cancel <-chan struct{}) (<-chan RelationshipLinkResourceFormat, <-chan error) {
	resultChan := make(chan RelationshipLinkResourceFormat, 1)
	errChan := make(chan error, 1)
	if err := validation.Validate([]validation.Validation{
		{TargetValue: relationshipLinkName,
			Constraints: []validation.Constraint{{Target: "relationshipLinkName", Name: validation.MaxLength, Rule: 512, Chain: nil},
				{Target: "relationshipLinkName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "relationshipLinkName", Name: validation.Pattern, Rule: `^[a-zA-Z][a-zA-Z0-9_]+$`, Chain: nil}}},
		{TargetValue: parameters,
			Constraints: []validation.Constraint{{Target: "parameters.RelationshipLinkDefinition", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "parameters.RelationshipLinkDefinition.InteractionType", Name: validation.Null, Rule: true, Chain: nil},
					{Target: "parameters.RelationshipLinkDefinition.ProfilePropertyReferences", Name: validation.Null, Rule: true, Chain: nil},
					{Target: "parameters.RelationshipLinkDefinition.RelatedProfilePropertyReferences", Name: validation.Null, Rule: true, Chain: nil},
					{Target: "parameters.RelationshipLinkDefinition.RelationshipName", Name: validation.Null, Rule: true, Chain: nil},
				}}}}}); err != nil {
		errChan <- validation.NewErrorWithValidationError(err, "customerinsights.RelationshipLinksClient", "CreateOrUpdate")
		close(errChan)
		close(resultChan)
		return resultChan, errChan
	}

	go func() {
		var err error
		var result RelationshipLinkResourceFormat
		defer func() {
			resultChan <- result
			errChan <- err
			close(resultChan)
			close(errChan)
		}()
		req, err := client.CreateOrUpdatePreparer(resourceGroupName, hubName, relationshipLinkName, parameters, cancel)
		if err != nil {
			err = autorest.NewErrorWithError(err, "customerinsights.RelationshipLinksClient", "CreateOrUpdate", nil, "Failure preparing request")
			return
		}

		resp, err := client.CreateOrUpdateSender(req)
		if err != nil {
			result.Response = autorest.Response{Response: resp}
			err = autorest.NewErrorWithError(err, "customerinsights.RelationshipLinksClient", "CreateOrUpdate", resp, "Failure sending request")
			return
		}

		result, err = client.CreateOrUpdateResponder(resp)
		if err != nil {
			err = autorest.NewErrorWithError(err, "customerinsights.RelationshipLinksClient", "CreateOrUpdate", resp, "Failure responding to request")
		}
	}()
	return resultChan, errChan
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client RelationshipLinksClient) CreateOrUpdatePreparer(resourceGroupName string, hubName string, relationshipLinkName string, parameters RelationshipLinkResourceFormat, cancel <-chan struct{}) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"hubName":              autorest.Encode("path", hubName),
		"relationshipLinkName": autorest.Encode("path", relationshipLinkName),
		"resourceGroupName":    autorest.Encode("path", resourceGroupName),
		"subscriptionId":       autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-01-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomerInsights/hubs/{hubName}/relationshipLinks/{relationshipLinkName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{Cancel: cancel})
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client RelationshipLinksClient) CreateOrUpdateSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		azure.DoPollForAsynchronous(client.PollingDelay))
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client RelationshipLinksClient) CreateOrUpdateResponder(resp *http.Response) (result RelationshipLinkResourceFormat, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes a relationship link within a hub. This method may poll for
// completion. Polling can be canceled by passing the cancel channel argument.
// The channel will be used to cancel polling and any outstanding HTTP
// requests.
//
// resourceGroupName is the name of the resource group. hubName is the name of
// the hub. relationshipLinkName is the name of the relationship.
func (client RelationshipLinksClient) Delete(resourceGroupName string, hubName string, relationshipLinkName string, cancel <-chan struct{}) (<-chan autorest.Response, <-chan error) {
	resultChan := make(chan autorest.Response, 1)
	errChan := make(chan error, 1)
	go func() {
		var err error
		var result autorest.Response
		defer func() {
			resultChan <- result
			errChan <- err
			close(resultChan)
			close(errChan)
		}()
		req, err := client.DeletePreparer(resourceGroupName, hubName, relationshipLinkName, cancel)
		if err != nil {
			err = autorest.NewErrorWithError(err, "customerinsights.RelationshipLinksClient", "Delete", nil, "Failure preparing request")
			return
		}

		resp, err := client.DeleteSender(req)
		if err != nil {
			result.Response = resp
			err = autorest.NewErrorWithError(err, "customerinsights.RelationshipLinksClient", "Delete", resp, "Failure sending request")
			return
		}

		result, err = client.DeleteResponder(resp)
		if err != nil {
			err = autorest.NewErrorWithError(err, "customerinsights.RelationshipLinksClient", "Delete", resp, "Failure responding to request")
		}
	}()
	return resultChan, errChan
}

// DeletePreparer prepares the Delete request.
func (client RelationshipLinksClient) DeletePreparer(resourceGroupName string, hubName string, relationshipLinkName string, cancel <-chan struct{}) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"hubName":              autorest.Encode("path", hubName),
		"relationshipLinkName": autorest.Encode("path", relationshipLinkName),
		"resourceGroupName":    autorest.Encode("path", resourceGroupName),
		"subscriptionId":       autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-01-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomerInsights/hubs/{hubName}/relationshipLinks/{relationshipLinkName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{Cancel: cancel})
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client RelationshipLinksClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		azure.DoPollForAsynchronous(client.PollingDelay))
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client RelationshipLinksClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusAccepted, http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get gets information about the specified relationship Link.
//
// resourceGroupName is the name of the resource group. hubName is the name of
// the hub. relationshipLinkName is the name of the relationship link.
func (client RelationshipLinksClient) Get(resourceGroupName string, hubName string, relationshipLinkName string) (result RelationshipLinkResourceFormat, err error) {
	req, err := client.GetPreparer(resourceGroupName, hubName, relationshipLinkName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "customerinsights.RelationshipLinksClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "customerinsights.RelationshipLinksClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "customerinsights.RelationshipLinksClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client RelationshipLinksClient) GetPreparer(resourceGroupName string, hubName string, relationshipLinkName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"hubName":              autorest.Encode("path", hubName),
		"relationshipLinkName": autorest.Encode("path", relationshipLinkName),
		"resourceGroupName":    autorest.Encode("path", resourceGroupName),
		"subscriptionId":       autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-01-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomerInsights/hubs/{hubName}/relationshipLinks/{relationshipLinkName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client RelationshipLinksClient) GetSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client RelationshipLinksClient) GetResponder(resp *http.Response) (result RelationshipLinkResourceFormat, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByHub gets all relationship links in the hub.
//
// resourceGroupName is the name of the resource group. hubName is the name of
// the hub.
func (client RelationshipLinksClient) ListByHub(resourceGroupName string, hubName string) (result RelationshipLinkListResult, err error) {
	req, err := client.ListByHubPreparer(resourceGroupName, hubName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "customerinsights.RelationshipLinksClient", "ListByHub", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByHubSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "customerinsights.RelationshipLinksClient", "ListByHub", resp, "Failure sending request")
		return
	}

	result, err = client.ListByHubResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "customerinsights.RelationshipLinksClient", "ListByHub", resp, "Failure responding to request")
	}

	return
}

// ListByHubPreparer prepares the ListByHub request.
func (client RelationshipLinksClient) ListByHubPreparer(resourceGroupName string, hubName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"hubName":           autorest.Encode("path", hubName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-01-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomerInsights/hubs/{hubName}/relationshipLinks", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// ListByHubSender sends the ListByHub request. The method will close the
// http.Response Body if it receives an error.
func (client RelationshipLinksClient) ListByHubSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// ListByHubResponder handles the response to the ListByHub request. The method always
// closes the http.Response Body.
func (client RelationshipLinksClient) ListByHubResponder(resp *http.Response) (result RelationshipLinkListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByHubNextResults retrieves the next set of results, if any.
func (client RelationshipLinksClient) ListByHubNextResults(lastResults RelationshipLinkListResult) (result RelationshipLinkListResult, err error) {
	req, err := lastResults.RelationshipLinkListResultPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "customerinsights.RelationshipLinksClient", "ListByHub", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}

	resp, err := client.ListByHubSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "customerinsights.RelationshipLinksClient", "ListByHub", resp, "Failure sending next results request")
	}

	result, err = client.ListByHubResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "customerinsights.RelationshipLinksClient", "ListByHub", resp, "Failure responding to next results request")
	}

	return
}