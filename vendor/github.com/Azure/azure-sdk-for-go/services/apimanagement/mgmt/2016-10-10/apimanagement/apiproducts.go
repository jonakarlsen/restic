package apimanagement

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
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/validation"
	"net/http"
)

// APIProductsClient is the apiManagement Client
type APIProductsClient struct {
	BaseClient
}

// NewAPIProductsClient creates an instance of the APIProductsClient client.
func NewAPIProductsClient(subscriptionID string) APIProductsClient {
	return NewAPIProductsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewAPIProductsClientWithBaseURI creates an instance of the APIProductsClient client.
func NewAPIProductsClientWithBaseURI(baseURI string, subscriptionID string) APIProductsClient {
	return APIProductsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// ListByApis lists all API associated products.
//
// resourceGroupName is the name of the resource group. serviceName is the name of the API Management service.
// apiid is API identifier. Must be unique in the current API Management service instance. filter is | Field |
// Supported operators    | Supported functions                         |
// |-------|------------------------|---------------------------------------------|
// | name  | ge, le, eq, ne, gt, lt | substringof, contains, startswith, endswith | top is number of records to
// return. skip is number of records to skip.
func (client APIProductsClient) ListByApis(ctx context.Context, resourceGroupName string, serviceName string, apiid string, filter string, top *int32, skip *int32) (result ProductCollectionPage, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: serviceName,
			Constraints: []validation.Constraint{{Target: "serviceName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "serviceName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "serviceName", Name: validation.Pattern, Rule: `^[a-zA-Z](?:[a-zA-Z0-9-]*[a-zA-Z0-9])?$`, Chain: nil}}},
		{TargetValue: apiid,
			Constraints: []validation.Constraint{{Target: "apiid", Name: validation.MaxLength, Rule: 256, Chain: nil},
				{Target: "apiid", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "apiid", Name: validation.Pattern, Rule: `^[^*#&+:<>?]+$`, Chain: nil}}},
		{TargetValue: top,
			Constraints: []validation.Constraint{{Target: "top", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "top", Name: validation.InclusiveMinimum, Rule: 1, Chain: nil}}}}},
		{TargetValue: skip,
			Constraints: []validation.Constraint{{Target: "skip", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "skip", Name: validation.InclusiveMinimum, Rule: 0, Chain: nil}}}}}}); err != nil {
		return result, validation.NewError("apimanagement.APIProductsClient", "ListByApis", err.Error())
	}

	result.fn = client.listByApisNextResults
	req, err := client.ListByApisPreparer(ctx, resourceGroupName, serviceName, apiid, filter, top, skip)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.APIProductsClient", "ListByApis", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByApisSender(req)
	if err != nil {
		result.pc.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "apimanagement.APIProductsClient", "ListByApis", resp, "Failure sending request")
		return
	}

	result.pc, err = client.ListByApisResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.APIProductsClient", "ListByApis", resp, "Failure responding to request")
	}

	return
}

// ListByApisPreparer prepares the ListByApis request.
func (client APIProductsClient) ListByApisPreparer(ctx context.Context, resourceGroupName string, serviceName string, apiid string, filter string, top *int32, skip *int32) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"apiId":             autorest.Encode("path", apiid),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serviceName":       autorest.Encode("path", serviceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-10-10"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}
	if top != nil {
		queryParameters["$top"] = autorest.Encode("query", *top)
	}
	if skip != nil {
		queryParameters["$skip"] = autorest.Encode("query", *skip)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apis/{apiId}/products", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByApisSender sends the ListByApis request. The method will close the
// http.Response Body if it receives an error.
func (client APIProductsClient) ListByApisSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// ListByApisResponder handles the response to the ListByApis request. The method always
// closes the http.Response Body.
func (client APIProductsClient) ListByApisResponder(resp *http.Response) (result ProductCollection, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByApisNextResults retrieves the next set of results, if any.
func (client APIProductsClient) listByApisNextResults(lastResults ProductCollection) (result ProductCollection, err error) {
	req, err := lastResults.productCollectionPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "apimanagement.APIProductsClient", "listByApisNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByApisSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "apimanagement.APIProductsClient", "listByApisNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByApisResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.APIProductsClient", "listByApisNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByApisComplete enumerates all values, automatically crossing page boundaries as required.
func (client APIProductsClient) ListByApisComplete(ctx context.Context, resourceGroupName string, serviceName string, apiid string, filter string, top *int32, skip *int32) (result ProductCollectionIterator, err error) {
	result.page, err = client.ListByApis(ctx, resourceGroupName, serviceName, apiid, filter, top, skip)
	return
}
