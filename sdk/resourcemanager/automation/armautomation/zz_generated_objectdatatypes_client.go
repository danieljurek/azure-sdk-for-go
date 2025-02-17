//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armautomation

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

// ObjectDataTypesClient contains the methods for the ObjectDataTypes group.
// Don't use this type directly, use NewObjectDataTypesClient() instead.
type ObjectDataTypesClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewObjectDataTypesClient creates a new instance of ObjectDataTypesClient with the specified values.
func NewObjectDataTypesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *ObjectDataTypesClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Host) == 0 {
		cp.Host = arm.AzurePublicCloud
	}
	return &ObjectDataTypesClient{subscriptionID: subscriptionID, ep: string(cp.Host), pl: armruntime.NewPipeline(module, version, credential, &cp)}
}

// ListFieldsByModuleAndType - Retrieve a list of fields of a given type identified by module name.
// If the operation fails it returns the *ErrorResponse error type.
func (client *ObjectDataTypesClient) ListFieldsByModuleAndType(ctx context.Context, resourceGroupName string, automationAccountName string, moduleName string, typeName string, options *ObjectDataTypesListFieldsByModuleAndTypeOptions) (ObjectDataTypesListFieldsByModuleAndTypeResponse, error) {
	req, err := client.listFieldsByModuleAndTypeCreateRequest(ctx, resourceGroupName, automationAccountName, moduleName, typeName, options)
	if err != nil {
		return ObjectDataTypesListFieldsByModuleAndTypeResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ObjectDataTypesListFieldsByModuleAndTypeResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ObjectDataTypesListFieldsByModuleAndTypeResponse{}, client.listFieldsByModuleAndTypeHandleError(resp)
	}
	return client.listFieldsByModuleAndTypeHandleResponse(resp)
}

// listFieldsByModuleAndTypeCreateRequest creates the ListFieldsByModuleAndType request.
func (client *ObjectDataTypesClient) listFieldsByModuleAndTypeCreateRequest(ctx context.Context, resourceGroupName string, automationAccountName string, moduleName string, typeName string, options *ObjectDataTypesListFieldsByModuleAndTypeOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/modules/{moduleName}/objectDataTypes/{typeName}/fields"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if automationAccountName == "" {
		return nil, errors.New("parameter automationAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{automationAccountName}", url.PathEscape(automationAccountName))
	if moduleName == "" {
		return nil, errors.New("parameter moduleName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{moduleName}", url.PathEscape(moduleName))
	if typeName == "" {
		return nil, errors.New("parameter typeName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{typeName}", url.PathEscape(typeName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-01-13-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listFieldsByModuleAndTypeHandleResponse handles the ListFieldsByModuleAndType response.
func (client *ObjectDataTypesClient) listFieldsByModuleAndTypeHandleResponse(resp *http.Response) (ObjectDataTypesListFieldsByModuleAndTypeResponse, error) {
	result := ObjectDataTypesListFieldsByModuleAndTypeResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.TypeFieldListResult); err != nil {
		return ObjectDataTypesListFieldsByModuleAndTypeResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// listFieldsByModuleAndTypeHandleError handles the ListFieldsByModuleAndType error response.
func (client *ObjectDataTypesClient) listFieldsByModuleAndTypeHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// ListFieldsByType - Retrieve a list of fields of a given type across all accessible modules.
// If the operation fails it returns the *ErrorResponse error type.
func (client *ObjectDataTypesClient) ListFieldsByType(ctx context.Context, resourceGroupName string, automationAccountName string, typeName string, options *ObjectDataTypesListFieldsByTypeOptions) (ObjectDataTypesListFieldsByTypeResponse, error) {
	req, err := client.listFieldsByTypeCreateRequest(ctx, resourceGroupName, automationAccountName, typeName, options)
	if err != nil {
		return ObjectDataTypesListFieldsByTypeResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ObjectDataTypesListFieldsByTypeResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ObjectDataTypesListFieldsByTypeResponse{}, client.listFieldsByTypeHandleError(resp)
	}
	return client.listFieldsByTypeHandleResponse(resp)
}

// listFieldsByTypeCreateRequest creates the ListFieldsByType request.
func (client *ObjectDataTypesClient) listFieldsByTypeCreateRequest(ctx context.Context, resourceGroupName string, automationAccountName string, typeName string, options *ObjectDataTypesListFieldsByTypeOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/objectDataTypes/{typeName}/fields"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if automationAccountName == "" {
		return nil, errors.New("parameter automationAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{automationAccountName}", url.PathEscape(automationAccountName))
	if typeName == "" {
		return nil, errors.New("parameter typeName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{typeName}", url.PathEscape(typeName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-01-13-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listFieldsByTypeHandleResponse handles the ListFieldsByType response.
func (client *ObjectDataTypesClient) listFieldsByTypeHandleResponse(resp *http.Response) (ObjectDataTypesListFieldsByTypeResponse, error) {
	result := ObjectDataTypesListFieldsByTypeResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.TypeFieldListResult); err != nil {
		return ObjectDataTypesListFieldsByTypeResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// listFieldsByTypeHandleError handles the ListFieldsByType error response.
func (client *ObjectDataTypesClient) listFieldsByTypeHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}
