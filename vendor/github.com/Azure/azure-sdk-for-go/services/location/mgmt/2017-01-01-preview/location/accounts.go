package location

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

// AccountsClient is the resource Provider
type AccountsClient struct {
	BaseClient
}

// NewAccountsClient creates an instance of the AccountsClient client.
func NewAccountsClient(subscriptionID string) AccountsClient {
	return NewAccountsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewAccountsClientWithBaseURI creates an instance of the AccountsClient client.
func NewAccountsClientWithBaseURI(baseURI string, subscriptionID string) AccountsClient {
	return AccountsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate create or update a Location Based Services Account. A Location Based Services Account holds the keys
// which allow access to the Location Based Services REST APIs.
//
// resourceGroupName is the name of the Azure Resource Group. accountName is the name of the Location Based
// Services Account. locationBasedServicesAccountCreateParameters is the new or updated parameters for the Location
// Based Services Account.
func (client AccountsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, accountName string, locationBasedServicesAccountCreateParameters BasedServicesAccountCreateParameters) (result BasedServicesAccount, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: locationBasedServicesAccountCreateParameters,
			Constraints: []validation.Constraint{{Target: "locationBasedServicesAccountCreateParameters.LocationProperty", Name: validation.Null, Rule: true, Chain: nil},
				{Target: "locationBasedServicesAccountCreateParameters.Sku", Name: validation.Null, Rule: true,
					Chain: []validation.Constraint{{Target: "locationBasedServicesAccountCreateParameters.Sku.Name", Name: validation.Null, Rule: true, Chain: nil}}}}}}); err != nil {
		return result, validation.NewError("location.AccountsClient", "CreateOrUpdate", err.Error())
	}

	req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, accountName, locationBasedServicesAccountCreateParameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "location.AccountsClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateOrUpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "location.AccountsClient", "CreateOrUpdate", resp, "Failure sending request")
		return
	}

	result, err = client.CreateOrUpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "location.AccountsClient", "CreateOrUpdate", resp, "Failure responding to request")
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client AccountsClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, accountName string, locationBasedServicesAccountCreateParameters BasedServicesAccountCreateParameters) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"accountName":       autorest.Encode("path", accountName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-01-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.LocationBasedServices/accounts/{accountName}", pathParameters),
		autorest.WithJSON(locationBasedServicesAccountCreateParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client AccountsClient) CreateOrUpdateSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client AccountsClient) CreateOrUpdateResponder(resp *http.Response) (result BasedServicesAccount, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete delete a Location Based Services Account
//
// resourceGroupName is the name of the Azure Resource Group. accountName is the name of the Location Based
// Services Account.
func (client AccountsClient) Delete(ctx context.Context, resourceGroupName string, accountName string) (result autorest.Response, err error) {
	req, err := client.DeletePreparer(ctx, resourceGroupName, accountName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "location.AccountsClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "location.AccountsClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "location.AccountsClient", "Delete", resp, "Failure responding to request")
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client AccountsClient) DeletePreparer(ctx context.Context, resourceGroupName string, accountName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"accountName":       autorest.Encode("path", accountName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-01-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.LocationBasedServices/accounts/{accountName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client AccountsClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client AccountsClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get get a Location Based Services Account
//
// resourceGroupName is the name of the Azure Resource Group. accountName is the name of the Location Based
// Services Account.
func (client AccountsClient) Get(ctx context.Context, resourceGroupName string, accountName string) (result BasedServicesAccount, err error) {
	req, err := client.GetPreparer(ctx, resourceGroupName, accountName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "location.AccountsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "location.AccountsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "location.AccountsClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client AccountsClient) GetPreparer(ctx context.Context, resourceGroupName string, accountName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"accountName":       autorest.Encode("path", accountName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-01-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.LocationBasedServices/accounts/{accountName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client AccountsClient) GetSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client AccountsClient) GetResponder(resp *http.Response) (result BasedServicesAccount, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNotFound),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByResourceGroup get all Location Based Services Accounts in a Resource Group
//
// resourceGroupName is the name of the Azure Resource Group.
func (client AccountsClient) ListByResourceGroup(ctx context.Context, resourceGroupName string) (result BasedServicesAccounts, err error) {
	req, err := client.ListByResourceGroupPreparer(ctx, resourceGroupName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "location.AccountsClient", "ListByResourceGroup", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByResourceGroupSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "location.AccountsClient", "ListByResourceGroup", resp, "Failure sending request")
		return
	}

	result, err = client.ListByResourceGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "location.AccountsClient", "ListByResourceGroup", resp, "Failure responding to request")
	}

	return
}

// ListByResourceGroupPreparer prepares the ListByResourceGroup request.
func (client AccountsClient) ListByResourceGroupPreparer(ctx context.Context, resourceGroupName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-01-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.LocationBasedServices/accounts", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByResourceGroupSender sends the ListByResourceGroup request. The method will close the
// http.Response Body if it receives an error.
func (client AccountsClient) ListByResourceGroupSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// ListByResourceGroupResponder handles the response to the ListByResourceGroup request. The method always
// closes the http.Response Body.
func (client AccountsClient) ListByResourceGroupResponder(resp *http.Response) (result BasedServicesAccounts, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListBySubscription get all Location Based Services Accounts in a Subscription
func (client AccountsClient) ListBySubscription(ctx context.Context) (result BasedServicesAccounts, err error) {
	req, err := client.ListBySubscriptionPreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "location.AccountsClient", "ListBySubscription", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListBySubscriptionSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "location.AccountsClient", "ListBySubscription", resp, "Failure sending request")
		return
	}

	result, err = client.ListBySubscriptionResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "location.AccountsClient", "ListBySubscription", resp, "Failure responding to request")
	}

	return
}

// ListBySubscriptionPreparer prepares the ListBySubscription request.
func (client AccountsClient) ListBySubscriptionPreparer(ctx context.Context) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-01-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.LocationBasedServices/accounts", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListBySubscriptionSender sends the ListBySubscription request. The method will close the
// http.Response Body if it receives an error.
func (client AccountsClient) ListBySubscriptionSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// ListBySubscriptionResponder handles the response to the ListBySubscription request. The method always
// closes the http.Response Body.
func (client AccountsClient) ListBySubscriptionResponder(resp *http.Response) (result BasedServicesAccounts, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListKeys get the keys to use with the Location Based Services APIs. A key is used to authenticate and authorize
// access to the Location Based Services REST APIs. Only one key is needed at a time; two are given to provide seamless
// key regeneration.
//
// resourceGroupName is the name of the Azure Resource Group. accountName is the name of the Location Based
// Services Account.
func (client AccountsClient) ListKeys(ctx context.Context, resourceGroupName string, accountName string) (result BasedServicesAccountKeys, err error) {
	req, err := client.ListKeysPreparer(ctx, resourceGroupName, accountName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "location.AccountsClient", "ListKeys", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListKeysSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "location.AccountsClient", "ListKeys", resp, "Failure sending request")
		return
	}

	result, err = client.ListKeysResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "location.AccountsClient", "ListKeys", resp, "Failure responding to request")
	}

	return
}

// ListKeysPreparer prepares the ListKeys request.
func (client AccountsClient) ListKeysPreparer(ctx context.Context, resourceGroupName string, accountName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"accountName":       autorest.Encode("path", accountName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-01-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.LocationBasedServices/accounts/{accountName}/listKeys", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListKeysSender sends the ListKeys request. The method will close the
// http.Response Body if it receives an error.
func (client AccountsClient) ListKeysSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// ListKeysResponder handles the response to the ListKeys request. The method always
// closes the http.Response Body.
func (client AccountsClient) ListKeysResponder(resp *http.Response) (result BasedServicesAccountKeys, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNotFound),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListOperations list operations available for the Location Based Services Resource Provider
func (client AccountsClient) ListOperations(ctx context.Context) (result BasedServicesOperations, err error) {
	req, err := client.ListOperationsPreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "location.AccountsClient", "ListOperations", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListOperationsSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "location.AccountsClient", "ListOperations", resp, "Failure sending request")
		return
	}

	result, err = client.ListOperationsResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "location.AccountsClient", "ListOperations", resp, "Failure responding to request")
	}

	return
}

// ListOperationsPreparer prepares the ListOperations request.
func (client AccountsClient) ListOperationsPreparer(ctx context.Context) (*http.Request, error) {
	const APIVersion = "2017-01-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/providers/Microsoft.LocationBasedServices/operations"),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListOperationsSender sends the ListOperations request. The method will close the
// http.Response Body if it receives an error.
func (client AccountsClient) ListOperationsSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// ListOperationsResponder handles the response to the ListOperations request. The method always
// closes the http.Response Body.
func (client AccountsClient) ListOperationsResponder(resp *http.Response) (result BasedServicesOperations, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Move moves Location Based Services Accounts from one ResourceGroup (or Subscription) to another
//
// resourceGroupName is the name of the resource group that contains Location Based Services Account to move.
// moveRequest is the details of the Location Based Services Account move.
func (client AccountsClient) Move(ctx context.Context, resourceGroupName string, moveRequest BasedServicesAccountsMoveRequest) (result autorest.Response, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: moveRequest,
			Constraints: []validation.Constraint{{Target: "moveRequest.TargetResourceGroup", Name: validation.Null, Rule: true, Chain: nil},
				{Target: "moveRequest.ResourceIds", Name: validation.Null, Rule: true, Chain: nil}}}}); err != nil {
		return result, validation.NewError("location.AccountsClient", "Move", err.Error())
	}

	req, err := client.MovePreparer(ctx, resourceGroupName, moveRequest)
	if err != nil {
		err = autorest.NewErrorWithError(err, "location.AccountsClient", "Move", nil, "Failure preparing request")
		return
	}

	resp, err := client.MoveSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "location.AccountsClient", "Move", resp, "Failure sending request")
		return
	}

	result, err = client.MoveResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "location.AccountsClient", "Move", resp, "Failure responding to request")
	}

	return
}

// MovePreparer prepares the Move request.
func (client AccountsClient) MovePreparer(ctx context.Context, resourceGroupName string, moveRequest BasedServicesAccountsMoveRequest) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-01-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/moveResources", pathParameters),
		autorest.WithJSON(moveRequest),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// MoveSender sends the Move request. The method will close the
// http.Response Body if it receives an error.
func (client AccountsClient) MoveSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// MoveResponder handles the response to the Move request. The method always
// closes the http.Response Body.
func (client AccountsClient) MoveResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}

// RegenerateKeys regenerate either the primary or secondary key for use with the Location Based Services APIs. The old
// key will stop working immediately.
//
// resourceGroupName is the name of the Azure Resource Group. accountName is the name of the Location Based
// Services Account. keySpecification is which key to regenerate:  primary or secondary.
func (client AccountsClient) RegenerateKeys(ctx context.Context, resourceGroupName string, accountName string, keySpecification BasedServicesKeySpecification) (result BasedServicesAccountKeys, err error) {
	req, err := client.RegenerateKeysPreparer(ctx, resourceGroupName, accountName, keySpecification)
	if err != nil {
		err = autorest.NewErrorWithError(err, "location.AccountsClient", "RegenerateKeys", nil, "Failure preparing request")
		return
	}

	resp, err := client.RegenerateKeysSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "location.AccountsClient", "RegenerateKeys", resp, "Failure sending request")
		return
	}

	result, err = client.RegenerateKeysResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "location.AccountsClient", "RegenerateKeys", resp, "Failure responding to request")
	}

	return
}

// RegenerateKeysPreparer prepares the RegenerateKeys request.
func (client AccountsClient) RegenerateKeysPreparer(ctx context.Context, resourceGroupName string, accountName string, keySpecification BasedServicesKeySpecification) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"accountName":       autorest.Encode("path", accountName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-01-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.LocationBasedServices/accounts/{accountName}/regenerateKey", pathParameters),
		autorest.WithJSON(keySpecification),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// RegenerateKeysSender sends the RegenerateKeys request. The method will close the
// http.Response Body if it receives an error.
func (client AccountsClient) RegenerateKeysSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// RegenerateKeysResponder handles the response to the RegenerateKeys request. The method always
// closes the http.Response Body.
func (client AccountsClient) RegenerateKeysResponder(resp *http.Response) (result BasedServicesAccountKeys, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNotFound),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Update updates a Location Based Services Account. Only a subset of the parameters may be updated after creation,
// such as Sku and Tags.
//
// resourceGroupName is the name of the Azure Resource Group. accountName is the name of the Location Based
// Services Account. locationBasedServicesAccountUpdateParameters is the updated parameters for the Location Based
// Services Account.
func (client AccountsClient) Update(ctx context.Context, resourceGroupName string, accountName string, locationBasedServicesAccountUpdateParameters BasedServicesAccountUpdateParameters) (result BasedServicesAccount, err error) {
	req, err := client.UpdatePreparer(ctx, resourceGroupName, accountName, locationBasedServicesAccountUpdateParameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "location.AccountsClient", "Update", nil, "Failure preparing request")
		return
	}

	resp, err := client.UpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "location.AccountsClient", "Update", resp, "Failure sending request")
		return
	}

	result, err = client.UpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "location.AccountsClient", "Update", resp, "Failure responding to request")
	}

	return
}

// UpdatePreparer prepares the Update request.
func (client AccountsClient) UpdatePreparer(ctx context.Context, resourceGroupName string, accountName string, locationBasedServicesAccountUpdateParameters BasedServicesAccountUpdateParameters) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"accountName":       autorest.Encode("path", accountName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-01-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.LocationBasedServices/accounts/{accountName}", pathParameters),
		autorest.WithJSON(locationBasedServicesAccountUpdateParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UpdateSender sends the Update request. The method will close the
// http.Response Body if it receives an error.
func (client AccountsClient) UpdateSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// UpdateResponder handles the response to the Update request. The method always
// closes the http.Response Body.
func (client AccountsClient) UpdateResponder(resp *http.Response) (result BasedServicesAccount, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNotFound),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
