//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armservicefabric

import (
	"context"
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// ApplicationTypeVersionsClient contains the methods for the ApplicationTypeVersions group.
// Don't use this type directly, use NewApplicationTypeVersionsClient() instead.
type ApplicationTypeVersionsClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewApplicationTypeVersionsClient creates a new instance of ApplicationTypeVersionsClient with the specified values.
func NewApplicationTypeVersionsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *ApplicationTypeVersionsClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Host) == 0 {
		cp.Host = arm.AzurePublicCloud
	}
	return &ApplicationTypeVersionsClient{subscriptionID: subscriptionID, ep: string(cp.Host), pl: armruntime.NewPipeline(module, version, credential, &cp)}
}

// BeginCreateOrUpdate - Create or update a Service Fabric application type version resource with the specified name.
// If the operation fails it returns the *ErrorModel error type.
func (client *ApplicationTypeVersionsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, clusterName string, applicationTypeName string, version string, parameters ApplicationTypeVersionResource, options *ApplicationTypeVersionsBeginCreateOrUpdateOptions) (ApplicationTypeVersionsCreateOrUpdatePollerResponse, error) {
	resp, err := client.createOrUpdate(ctx, resourceGroupName, clusterName, applicationTypeName, version, parameters, options)
	if err != nil {
		return ApplicationTypeVersionsCreateOrUpdatePollerResponse{}, err
	}
	result := ApplicationTypeVersionsCreateOrUpdatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("ApplicationTypeVersionsClient.CreateOrUpdate", "", resp, client.pl, client.createOrUpdateHandleError)
	if err != nil {
		return ApplicationTypeVersionsCreateOrUpdatePollerResponse{}, err
	}
	result.Poller = &ApplicationTypeVersionsCreateOrUpdatePoller{
		pt: pt,
	}
	return result, nil
}

// CreateOrUpdate - Create or update a Service Fabric application type version resource with the specified name.
// If the operation fails it returns the *ErrorModel error type.
func (client *ApplicationTypeVersionsClient) createOrUpdate(ctx context.Context, resourceGroupName string, clusterName string, applicationTypeName string, version string, parameters ApplicationTypeVersionResource, options *ApplicationTypeVersionsBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, clusterName, applicationTypeName, version, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusAccepted) {
		return nil, client.createOrUpdateHandleError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *ApplicationTypeVersionsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, clusterName string, applicationTypeName string, version string, parameters ApplicationTypeVersionResource, options *ApplicationTypeVersionsBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceFabric/clusters/{clusterName}/applicationTypes/{applicationTypeName}/versions/{version}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if clusterName == "" {
		return nil, errors.New("parameter clusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterName}", url.PathEscape(clusterName))
	if applicationTypeName == "" {
		return nil, errors.New("parameter applicationTypeName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{applicationTypeName}", url.PathEscape(applicationTypeName))
	if version == "" {
		return nil, errors.New("parameter version cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{version}", url.PathEscape(version))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *ApplicationTypeVersionsClient) createOrUpdateHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorModel{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// BeginDelete - Delete a Service Fabric application type version resource with the specified name.
// If the operation fails it returns the *ErrorModel error type.
func (client *ApplicationTypeVersionsClient) BeginDelete(ctx context.Context, resourceGroupName string, clusterName string, applicationTypeName string, version string, options *ApplicationTypeVersionsBeginDeleteOptions) (ApplicationTypeVersionsDeletePollerResponse, error) {
	resp, err := client.deleteOperation(ctx, resourceGroupName, clusterName, applicationTypeName, version, options)
	if err != nil {
		return ApplicationTypeVersionsDeletePollerResponse{}, err
	}
	result := ApplicationTypeVersionsDeletePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("ApplicationTypeVersionsClient.Delete", "", resp, client.pl, client.deleteHandleError)
	if err != nil {
		return ApplicationTypeVersionsDeletePollerResponse{}, err
	}
	result.Poller = &ApplicationTypeVersionsDeletePoller{
		pt: pt,
	}
	return result, nil
}

// Delete - Delete a Service Fabric application type version resource with the specified name.
// If the operation fails it returns the *ErrorModel error type.
func (client *ApplicationTypeVersionsClient) deleteOperation(ctx context.Context, resourceGroupName string, clusterName string, applicationTypeName string, version string, options *ApplicationTypeVersionsBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, clusterName, applicationTypeName, version, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusAccepted, http.StatusNoContent) {
		return nil, client.deleteHandleError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *ApplicationTypeVersionsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, clusterName string, applicationTypeName string, version string, options *ApplicationTypeVersionsBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceFabric/clusters/{clusterName}/applicationTypes/{applicationTypeName}/versions/{version}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if clusterName == "" {
		return nil, errors.New("parameter clusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterName}", url.PathEscape(clusterName))
	if applicationTypeName == "" {
		return nil, errors.New("parameter applicationTypeName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{applicationTypeName}", url.PathEscape(applicationTypeName))
	if version == "" {
		return nil, errors.New("parameter version cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{version}", url.PathEscape(version))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// deleteHandleError handles the Delete error response.
func (client *ApplicationTypeVersionsClient) deleteHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorModel{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// Get - Get a Service Fabric application type version resource created or in the process of being created in the Service Fabric application type name resource.
// If the operation fails it returns the *ErrorModel error type.
func (client *ApplicationTypeVersionsClient) Get(ctx context.Context, resourceGroupName string, clusterName string, applicationTypeName string, version string, options *ApplicationTypeVersionsGetOptions) (ApplicationTypeVersionsGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, clusterName, applicationTypeName, version, options)
	if err != nil {
		return ApplicationTypeVersionsGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ApplicationTypeVersionsGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ApplicationTypeVersionsGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ApplicationTypeVersionsClient) getCreateRequest(ctx context.Context, resourceGroupName string, clusterName string, applicationTypeName string, version string, options *ApplicationTypeVersionsGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceFabric/clusters/{clusterName}/applicationTypes/{applicationTypeName}/versions/{version}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if clusterName == "" {
		return nil, errors.New("parameter clusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterName}", url.PathEscape(clusterName))
	if applicationTypeName == "" {
		return nil, errors.New("parameter applicationTypeName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{applicationTypeName}", url.PathEscape(applicationTypeName))
	if version == "" {
		return nil, errors.New("parameter version cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{version}", url.PathEscape(version))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ApplicationTypeVersionsClient) getHandleResponse(resp *http.Response) (ApplicationTypeVersionsGetResponse, error) {
	result := ApplicationTypeVersionsGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ApplicationTypeVersionResource); err != nil {
		return ApplicationTypeVersionsGetResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *ApplicationTypeVersionsClient) getHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorModel{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// List - Gets all application type version resources created or in the process of being created in the Service Fabric application type name resource.
// If the operation fails it returns the *ErrorModel error type.
func (client *ApplicationTypeVersionsClient) List(ctx context.Context, resourceGroupName string, clusterName string, applicationTypeName string, options *ApplicationTypeVersionsListOptions) (ApplicationTypeVersionsListResponse, error) {
	req, err := client.listCreateRequest(ctx, resourceGroupName, clusterName, applicationTypeName, options)
	if err != nil {
		return ApplicationTypeVersionsListResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ApplicationTypeVersionsListResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ApplicationTypeVersionsListResponse{}, client.listHandleError(resp)
	}
	return client.listHandleResponse(resp)
}

// listCreateRequest creates the List request.
func (client *ApplicationTypeVersionsClient) listCreateRequest(ctx context.Context, resourceGroupName string, clusterName string, applicationTypeName string, options *ApplicationTypeVersionsListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceFabric/clusters/{clusterName}/applicationTypes/{applicationTypeName}/versions"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if clusterName == "" {
		return nil, errors.New("parameter clusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterName}", url.PathEscape(clusterName))
	if applicationTypeName == "" {
		return nil, errors.New("parameter applicationTypeName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{applicationTypeName}", url.PathEscape(applicationTypeName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *ApplicationTypeVersionsClient) listHandleResponse(resp *http.Response) (ApplicationTypeVersionsListResponse, error) {
	result := ApplicationTypeVersionsListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ApplicationTypeVersionResourceList); err != nil {
		return ApplicationTypeVersionsListResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// listHandleError handles the List error response.
func (client *ApplicationTypeVersionsClient) listHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorModel{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}
