//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armresourcehealth

import "net/http"

// AvailabilityStatusesGetByResourceResponse contains the response from method AvailabilityStatuses.GetByResource.
type AvailabilityStatusesGetByResourceResponse struct {
	AvailabilityStatusesGetByResourceResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// AvailabilityStatusesGetByResourceResult contains the result from method AvailabilityStatuses.GetByResource.
type AvailabilityStatusesGetByResourceResult struct {
	AvailabilityStatus
}

// AvailabilityStatusesListByResourceGroupResponse contains the response from method AvailabilityStatuses.ListByResourceGroup.
type AvailabilityStatusesListByResourceGroupResponse struct {
	AvailabilityStatusesListByResourceGroupResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// AvailabilityStatusesListByResourceGroupResult contains the result from method AvailabilityStatuses.ListByResourceGroup.
type AvailabilityStatusesListByResourceGroupResult struct {
	AvailabilityStatusListResult
}

// AvailabilityStatusesListBySubscriptionIDResponse contains the response from method AvailabilityStatuses.ListBySubscriptionID.
type AvailabilityStatusesListBySubscriptionIDResponse struct {
	AvailabilityStatusesListBySubscriptionIDResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// AvailabilityStatusesListBySubscriptionIDResult contains the result from method AvailabilityStatuses.ListBySubscriptionID.
type AvailabilityStatusesListBySubscriptionIDResult struct {
	AvailabilityStatusListResult
}

// AvailabilityStatusesListResponse contains the response from method AvailabilityStatuses.List.
type AvailabilityStatusesListResponse struct {
	AvailabilityStatusesListResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// AvailabilityStatusesListResult contains the result from method AvailabilityStatuses.List.
type AvailabilityStatusesListResult struct {
	AvailabilityStatusListResult
}

// ChildAvailabilityStatusesGetByResourceResponse contains the response from method ChildAvailabilityStatuses.GetByResource.
type ChildAvailabilityStatusesGetByResourceResponse struct {
	ChildAvailabilityStatusesGetByResourceResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ChildAvailabilityStatusesGetByResourceResult contains the result from method ChildAvailabilityStatuses.GetByResource.
type ChildAvailabilityStatusesGetByResourceResult struct {
	AvailabilityStatus
}

// ChildAvailabilityStatusesListResponse contains the response from method ChildAvailabilityStatuses.List.
type ChildAvailabilityStatusesListResponse struct {
	ChildAvailabilityStatusesListResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ChildAvailabilityStatusesListResult contains the result from method ChildAvailabilityStatuses.List.
type ChildAvailabilityStatusesListResult struct {
	AvailabilityStatusListResult
}

// ChildResourcesListResponse contains the response from method ChildResources.List.
type ChildResourcesListResponse struct {
	ChildResourcesListResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ChildResourcesListResult contains the result from method ChildResources.List.
type ChildResourcesListResult struct {
	AvailabilityStatusListResult
}

// EmergingIssuesGetResponse contains the response from method EmergingIssues.Get.
type EmergingIssuesGetResponse struct {
	EmergingIssuesGetResultEnvelope
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// EmergingIssuesGetResultEnvelope contains the result from method EmergingIssues.Get.
type EmergingIssuesGetResultEnvelope struct {
	EmergingIssuesGetResult
}

// EmergingIssuesListResponse contains the response from method EmergingIssues.List.
type EmergingIssuesListResponse struct {
	EmergingIssuesListResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// EmergingIssuesListResult contains the result from method EmergingIssues.List.
type EmergingIssuesListResult struct {
	EmergingIssueListResult
}

// OperationsListResponse contains the response from method Operations.List.
type OperationsListResponse struct {
	OperationsListResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// OperationsListResult contains the result from method Operations.List.
type OperationsListResult struct {
	OperationListResult
}
