//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armhybridkubernetes

import (
	"context"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"net/http"
	"time"
)

// ConnectedClusterCreatePollerResponse contains the response from method ConnectedCluster.Create.
type ConnectedClusterCreatePollerResponse struct {
	// Poller contains an initialized poller.
	Poller *ConnectedClusterCreatePoller

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
// freq: the time to wait between intervals in absence of a Retry-After header. Allowed minimum is one second.
// A good starting value is 30 seconds. Note that some resources might benefit from a different value.
func (l ConnectedClusterCreatePollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (ConnectedClusterCreateResponse, error) {
	respType := ConnectedClusterCreateResponse{}
	resp, err := l.Poller.pt.PollUntilDone(ctx, freq, &respType.ConnectedCluster)
	if err != nil {
		return respType, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// Resume rehydrates a ConnectedClusterCreatePollerResponse from the provided client and resume token.
func (l *ConnectedClusterCreatePollerResponse) Resume(ctx context.Context, client *ConnectedClusterClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("ConnectedClusterClient.Create", token, client.pl, client.createHandleError)
	if err != nil {
		return err
	}
	poller := &ConnectedClusterCreatePoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return err
	}
	l.Poller = poller
	l.RawResponse = resp
	return nil
}

// ConnectedClusterCreateResponse contains the response from method ConnectedCluster.Create.
type ConnectedClusterCreateResponse struct {
	ConnectedClusterCreateResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ConnectedClusterCreateResult contains the result from method ConnectedCluster.Create.
type ConnectedClusterCreateResult struct {
	ConnectedCluster
}

// ConnectedClusterDeletePollerResponse contains the response from method ConnectedCluster.Delete.
type ConnectedClusterDeletePollerResponse struct {
	// Poller contains an initialized poller.
	Poller *ConnectedClusterDeletePoller

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
// freq: the time to wait between intervals in absence of a Retry-After header. Allowed minimum is one second.
// A good starting value is 30 seconds. Note that some resources might benefit from a different value.
func (l ConnectedClusterDeletePollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (ConnectedClusterDeleteResponse, error) {
	respType := ConnectedClusterDeleteResponse{}
	resp, err := l.Poller.pt.PollUntilDone(ctx, freq, nil)
	if err != nil {
		return respType, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// Resume rehydrates a ConnectedClusterDeletePollerResponse from the provided client and resume token.
func (l *ConnectedClusterDeletePollerResponse) Resume(ctx context.Context, client *ConnectedClusterClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("ConnectedClusterClient.Delete", token, client.pl, client.deleteHandleError)
	if err != nil {
		return err
	}
	poller := &ConnectedClusterDeletePoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return err
	}
	l.Poller = poller
	l.RawResponse = resp
	return nil
}

// ConnectedClusterDeleteResponse contains the response from method ConnectedCluster.Delete.
type ConnectedClusterDeleteResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ConnectedClusterGetResponse contains the response from method ConnectedCluster.Get.
type ConnectedClusterGetResponse struct {
	ConnectedClusterGetResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ConnectedClusterGetResult contains the result from method ConnectedCluster.Get.
type ConnectedClusterGetResult struct {
	ConnectedCluster
}

// ConnectedClusterListByResourceGroupResponse contains the response from method ConnectedCluster.ListByResourceGroup.
type ConnectedClusterListByResourceGroupResponse struct {
	ConnectedClusterListByResourceGroupResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ConnectedClusterListByResourceGroupResult contains the result from method ConnectedCluster.ListByResourceGroup.
type ConnectedClusterListByResourceGroupResult struct {
	ConnectedClusterList
}

// ConnectedClusterListBySubscriptionResponse contains the response from method ConnectedCluster.ListBySubscription.
type ConnectedClusterListBySubscriptionResponse struct {
	ConnectedClusterListBySubscriptionResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ConnectedClusterListBySubscriptionResult contains the result from method ConnectedCluster.ListBySubscription.
type ConnectedClusterListBySubscriptionResult struct {
	ConnectedClusterList
}

// ConnectedClusterListClusterUserCredentialResponse contains the response from method ConnectedCluster.ListClusterUserCredential.
type ConnectedClusterListClusterUserCredentialResponse struct {
	ConnectedClusterListClusterUserCredentialResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ConnectedClusterListClusterUserCredentialResult contains the result from method ConnectedCluster.ListClusterUserCredential.
type ConnectedClusterListClusterUserCredentialResult struct {
	CredentialResults
}

// ConnectedClusterUpdateResponse contains the response from method ConnectedCluster.Update.
type ConnectedClusterUpdateResponse struct {
	ConnectedClusterUpdateResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ConnectedClusterUpdateResult contains the result from method ConnectedCluster.Update.
type ConnectedClusterUpdateResult struct {
	ConnectedCluster
}

// OperationsGetResponse contains the response from method Operations.Get.
type OperationsGetResponse struct {
	OperationsGetResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// OperationsGetResult contains the result from method Operations.Get.
type OperationsGetResult struct {
	OperationList
}
