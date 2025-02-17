//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armreservations_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/reservations/armreservations"
)

// x-ms-original-file: specification/reservations/resource-manager/Microsoft.Capacity/stable/2021-07-01/examples/CalculateReservationOrder.json
func ExampleReservationOrderClient_Calculate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armreservations.NewReservationOrderClient(cred, nil)
	_, err = client.Calculate(ctx,
		armreservations.PurchaseRequest{
			Location: to.StringPtr("<location>"),
			Properties: &armreservations.PurchaseRequestProperties{
				AppliedScopeType: armreservations.AppliedScopeTypeShared.ToPtr(),
				AppliedScopes:    []*string{},
				BillingPlan:      armreservations.ReservationBillingPlanMonthly.ToPtr(),
				BillingScopeID:   to.StringPtr("<billing-scope-id>"),
				DisplayName:      to.StringPtr("<display-name>"),
				Quantity:         to.Int32Ptr(1),
				ReservedResourceProperties: &armreservations.PurchaseRequestPropertiesReservedResourceProperties{
					InstanceFlexibility: armreservations.InstanceFlexibilityOn.ToPtr(),
				},
				ReservedResourceType: armreservations.ReservedResourceTypeVirtualMachines.ToPtr(),
				Term:                 armreservations.ReservationTermP1Y.ToPtr(),
			},
			SKU: &armreservations.SKUName{
				Name: to.StringPtr("<name>"),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
}

// x-ms-original-file: specification/reservations/resource-manager/Microsoft.Capacity/stable/2021-07-01/examples/GetReservationOrders.json
func ExampleReservationOrderClient_List() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armreservations.NewReservationOrderClient(cred, nil)
	pager := client.List(nil)
	for pager.NextPage(ctx) {
		if err := pager.Err(); err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range pager.PageResponse().Value {
			log.Printf("ReservationOrderResponse.ID: %s\n", *v.ID)
		}
	}
}

// x-ms-original-file: specification/reservations/resource-manager/Microsoft.Capacity/stable/2021-07-01/examples/PurchaseReservationOrder.json
func ExampleReservationOrderClient_BeginPurchase() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armreservations.NewReservationOrderClient(cred, nil)
	poller, err := client.BeginPurchase(ctx,
		"<reservation-order-id>",
		armreservations.PurchaseRequest{
			Location: to.StringPtr("<location>"),
			Properties: &armreservations.PurchaseRequestProperties{
				AppliedScopeType: armreservations.AppliedScopeTypeShared.ToPtr(),
				AppliedScopes:    []*string{},
				BillingPlan:      armreservations.ReservationBillingPlanMonthly.ToPtr(),
				BillingScopeID:   to.StringPtr("<billing-scope-id>"),
				DisplayName:      to.StringPtr("<display-name>"),
				Quantity:         to.Int32Ptr(1),
				Renew:            to.BoolPtr(false),
				ReservedResourceProperties: &armreservations.PurchaseRequestPropertiesReservedResourceProperties{
					InstanceFlexibility: armreservations.InstanceFlexibilityOn.ToPtr(),
				},
				ReservedResourceType: armreservations.ReservedResourceTypeVirtualMachines.ToPtr(),
				Term:                 armreservations.ReservationTermP1Y.ToPtr(),
			},
			SKU: &armreservations.SKUName{
				Name: to.StringPtr("<name>"),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("ReservationOrderResponse.ID: %s\n", *res.ID)
}

// x-ms-original-file: specification/reservations/resource-manager/Microsoft.Capacity/stable/2021-07-01/examples/GetReservationOrderDetails.json
func ExampleReservationOrderClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armreservations.NewReservationOrderClient(cred, nil)
	res, err := client.Get(ctx,
		"<reservation-order-id>",
		&armreservations.ReservationOrderGetOptions{Expand: nil})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("ReservationOrderResponse.ID: %s\n", *res.ID)
}

// x-ms-original-file: specification/reservations/resource-manager/Microsoft.Capacity/stable/2021-07-01/examples/ChangeDirectoryReservationOrder.json
func ExampleReservationOrderClient_ChangeDirectory() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armreservations.NewReservationOrderClient(cred, nil)
	_, err = client.ChangeDirectory(ctx,
		"<reservation-order-id>",
		armreservations.ChangeDirectoryRequest{
			DestinationTenantID: to.StringPtr("<destination-tenant-id>"),
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
}
