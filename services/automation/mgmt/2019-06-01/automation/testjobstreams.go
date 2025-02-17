package automation

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/validation"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// TestJobStreamsClient is the automation Client
type TestJobStreamsClient struct {
	BaseClient
}

// NewTestJobStreamsClient creates an instance of the TestJobStreamsClient client.
func NewTestJobStreamsClient(subscriptionID string) TestJobStreamsClient {
	return NewTestJobStreamsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewTestJobStreamsClientWithBaseURI creates an instance of the TestJobStreamsClient client using a custom endpoint.
// Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewTestJobStreamsClientWithBaseURI(baseURI string, subscriptionID string) TestJobStreamsClient {
	return TestJobStreamsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Get retrieve a test job stream of the test job identified by runbook name and stream id.
// Parameters:
// resourceGroupName - name of an Azure Resource group.
// automationAccountName - the name of the automation account.
// runbookName - the runbook name.
// jobStreamID - the job stream id.
func (client TestJobStreamsClient) Get(ctx context.Context, resourceGroupName string, automationAccountName string, runbookName string, jobStreamID string) (result JobStream, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/TestJobStreamsClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("automation.TestJobStreamsClient", "Get", err.Error())
	}

	req, err := client.GetPreparer(ctx, resourceGroupName, automationAccountName, runbookName, jobStreamID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "automation.TestJobStreamsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "automation.TestJobStreamsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "automation.TestJobStreamsClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client TestJobStreamsClient) GetPreparer(ctx context.Context, resourceGroupName string, automationAccountName string, runbookName string, jobStreamID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"automationAccountName": autorest.Encode("path", automationAccountName),
		"jobStreamId":           autorest.Encode("path", jobStreamID),
		"resourceGroupName":     autorest.Encode("path", resourceGroupName),
		"runbookName":           autorest.Encode("path", runbookName),
		"subscriptionId":        autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-06-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/runbooks/{runbookName}/draft/testJob/streams/{jobStreamId}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client TestJobStreamsClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client TestJobStreamsClient) GetResponder(resp *http.Response) (result JobStream, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByTestJob retrieve a list of test job streams identified by runbook name.
// Parameters:
// resourceGroupName - name of an Azure Resource group.
// automationAccountName - the name of the automation account.
// runbookName - the runbook name.
// filter - the filter to apply on the operation.
func (client TestJobStreamsClient) ListByTestJob(ctx context.Context, resourceGroupName string, automationAccountName string, runbookName string, filter string) (result JobStreamListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/TestJobStreamsClient.ListByTestJob")
		defer func() {
			sc := -1
			if result.jslr.Response.Response != nil {
				sc = result.jslr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("automation.TestJobStreamsClient", "ListByTestJob", err.Error())
	}

	result.fn = client.listByTestJobNextResults
	req, err := client.ListByTestJobPreparer(ctx, resourceGroupName, automationAccountName, runbookName, filter)
	if err != nil {
		err = autorest.NewErrorWithError(err, "automation.TestJobStreamsClient", "ListByTestJob", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByTestJobSender(req)
	if err != nil {
		result.jslr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "automation.TestJobStreamsClient", "ListByTestJob", resp, "Failure sending request")
		return
	}

	result.jslr, err = client.ListByTestJobResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "automation.TestJobStreamsClient", "ListByTestJob", resp, "Failure responding to request")
		return
	}
	if result.jslr.hasNextLink() && result.jslr.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListByTestJobPreparer prepares the ListByTestJob request.
func (client TestJobStreamsClient) ListByTestJobPreparer(ctx context.Context, resourceGroupName string, automationAccountName string, runbookName string, filter string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"automationAccountName": autorest.Encode("path", automationAccountName),
		"resourceGroupName":     autorest.Encode("path", resourceGroupName),
		"runbookName":           autorest.Encode("path", runbookName),
		"subscriptionId":        autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-06-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/runbooks/{runbookName}/draft/testJob/streams", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByTestJobSender sends the ListByTestJob request. The method will close the
// http.Response Body if it receives an error.
func (client TestJobStreamsClient) ListByTestJobSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListByTestJobResponder handles the response to the ListByTestJob request. The method always
// closes the http.Response Body.
func (client TestJobStreamsClient) ListByTestJobResponder(resp *http.Response) (result JobStreamListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByTestJobNextResults retrieves the next set of results, if any.
func (client TestJobStreamsClient) listByTestJobNextResults(ctx context.Context, lastResults JobStreamListResult) (result JobStreamListResult, err error) {
	req, err := lastResults.jobStreamListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "automation.TestJobStreamsClient", "listByTestJobNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByTestJobSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "automation.TestJobStreamsClient", "listByTestJobNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByTestJobResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "automation.TestJobStreamsClient", "listByTestJobNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByTestJobComplete enumerates all values, automatically crossing page boundaries as required.
func (client TestJobStreamsClient) ListByTestJobComplete(ctx context.Context, resourceGroupName string, automationAccountName string, runbookName string, filter string) (result JobStreamListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/TestJobStreamsClient.ListByTestJob")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByTestJob(ctx, resourceGroupName, automationAccountName, runbookName, filter)
	return
}
