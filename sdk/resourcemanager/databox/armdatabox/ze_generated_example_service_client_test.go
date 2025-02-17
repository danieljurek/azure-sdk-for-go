//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdatabox_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/databox/armdatabox"
)

// x-ms-original-file: specification/databox/resource-manager/Microsoft.DataBox/preview/2021-08-01-preview/examples/AvailableSkusPost.json
func ExampleServiceClient_ListAvailableSKUsByResourceGroup() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armdatabox.NewServiceClient("<subscription-id>", cred, nil)
	pager := client.ListAvailableSKUsByResourceGroup("<resource-group-name>",
		"<location>",
		armdatabox.AvailableSKURequest{
			Country:      to.StringPtr("<country>"),
			Location:     to.StringPtr("<location>"),
			TransferType: armdatabox.TransferTypeImportToAzure.ToPtr(),
		},
		nil)
	for pager.NextPage(ctx) {
		if err := pager.Err(); err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
	}
}

// x-ms-original-file: specification/databox/resource-manager/Microsoft.DataBox/preview/2021-08-01-preview/examples/ValidateAddressPost.json
func ExampleServiceClient_ValidateAddress() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armdatabox.NewServiceClient("<subscription-id>", cred, nil)
	_, err = client.ValidateAddress(ctx,
		"<location>",
		armdatabox.ValidateAddress{
			ValidationInputRequest: armdatabox.ValidationInputRequest{
				ValidationType: armdatabox.ValidationInputDiscriminatorValidateAddress.ToPtr(),
			},
			DeviceType: armdatabox.SKUNameDataBox.ToPtr(),
			ShippingAddress: &armdatabox.ShippingAddress{
				AddressType:     armdatabox.AddressTypeCommercial.ToPtr(),
				City:            to.StringPtr("<city>"),
				CompanyName:     to.StringPtr("<company-name>"),
				Country:         to.StringPtr("<country>"),
				PostalCode:      to.StringPtr("<postal-code>"),
				StateOrProvince: to.StringPtr("<state-or-province>"),
				StreetAddress1:  to.StringPtr("<street-address1>"),
				StreetAddress2:  to.StringPtr("<street-address2>"),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
}

// x-ms-original-file: specification/databox/resource-manager/Microsoft.DataBox/preview/2021-08-01-preview/examples/ValidateInputsByResourceGroup.json
func ExampleServiceClient_ValidateInputsByResourceGroup() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armdatabox.NewServiceClient("<subscription-id>", cred, nil)
	_, err = client.ValidateInputsByResourceGroup(ctx,
		"<resource-group-name>",
		"<location>",
		&armdatabox.CreateJobValidations{
			ValidationRequest: armdatabox.ValidationRequest{
				IndividualRequestDetails: []armdatabox.ValidationInputRequestClassification{
					&armdatabox.DataTransferDetailsValidationRequest{
						ValidationInputRequest: armdatabox.ValidationInputRequest{
							ValidationType: armdatabox.ValidationInputDiscriminatorValidateDataTransferDetails.ToPtr(),
						},
						DataImportDetails: []*armdatabox.DataImportDetails{
							{
								AccountDetails: &armdatabox.StorageAccountDetails{
									DataAccountDetails: armdatabox.DataAccountDetails{
										DataAccountType: armdatabox.DataAccountTypeStorageAccount.ToPtr(),
									},
									StorageAccountID: to.StringPtr("<storage-account-id>"),
								},
							}},
						DeviceType:   armdatabox.SKUNameDataBox.ToPtr(),
						TransferType: armdatabox.TransferTypeImportToAzure.ToPtr(),
					},
					&armdatabox.ValidateAddress{
						ValidationInputRequest: armdatabox.ValidationInputRequest{
							ValidationType: armdatabox.ValidationInputDiscriminatorValidateAddress.ToPtr(),
						},
						DeviceType: armdatabox.SKUNameDataBox.ToPtr(),
						ShippingAddress: &armdatabox.ShippingAddress{
							AddressType:     armdatabox.AddressTypeCommercial.ToPtr(),
							City:            to.StringPtr("<city>"),
							CompanyName:     to.StringPtr("<company-name>"),
							Country:         to.StringPtr("<country>"),
							PostalCode:      to.StringPtr("<postal-code>"),
							StateOrProvince: to.StringPtr("<state-or-province>"),
							StreetAddress1:  to.StringPtr("<street-address1>"),
							StreetAddress2:  to.StringPtr("<street-address2>"),
						},
						TransportPreferences: &armdatabox.TransportPreferences{
							PreferredShipmentType: armdatabox.TransportShipmentTypesMicrosoftManaged.ToPtr(),
						},
					},
					&armdatabox.SubscriptionIsAllowedToCreateJobValidationRequest{
						ValidationInputRequest: armdatabox.ValidationInputRequest{
							ValidationType: armdatabox.ValidationInputDiscriminatorValidateSubscriptionIsAllowedToCreateJob.ToPtr(),
						},
					},
					&armdatabox.SKUAvailabilityValidationRequest{
						ValidationInputRequest: armdatabox.ValidationInputRequest{
							ValidationType: armdatabox.ValidationInputDiscriminatorValidateSKUAvailability.ToPtr(),
						},
						Country:      to.StringPtr("<country>"),
						DeviceType:   armdatabox.SKUNameDataBox.ToPtr(),
						Location:     to.StringPtr("<location>"),
						TransferType: armdatabox.TransferTypeImportToAzure.ToPtr(),
					},
					&armdatabox.CreateOrderLimitForSubscriptionValidationRequest{
						ValidationInputRequest: armdatabox.ValidationInputRequest{
							ValidationType: armdatabox.ValidationInputDiscriminatorValidateCreateOrderLimit.ToPtr(),
						},
						DeviceType: armdatabox.SKUNameDataBox.ToPtr(),
					},
					&armdatabox.PreferencesValidationRequest{
						ValidationInputRequest: armdatabox.ValidationInputRequest{
							ValidationType: armdatabox.ValidationInputDiscriminatorValidatePreferences.ToPtr(),
						},
						DeviceType: armdatabox.SKUNameDataBox.ToPtr(),
						Preference: &armdatabox.Preferences{
							TransportPreferences: &armdatabox.TransportPreferences{
								PreferredShipmentType: armdatabox.TransportShipmentTypesMicrosoftManaged.ToPtr(),
							},
						},
					}},
				ValidationCategory: to.StringPtr("<validation-category>"),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
}

// x-ms-original-file: specification/databox/resource-manager/Microsoft.DataBox/preview/2021-08-01-preview/examples/ValidateInputs.json
func ExampleServiceClient_ValidateInputs() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armdatabox.NewServiceClient("<subscription-id>", cred, nil)
	_, err = client.ValidateInputs(ctx,
		"<location>",
		&armdatabox.CreateJobValidations{
			ValidationRequest: armdatabox.ValidationRequest{
				IndividualRequestDetails: []armdatabox.ValidationInputRequestClassification{
					&armdatabox.DataTransferDetailsValidationRequest{
						ValidationInputRequest: armdatabox.ValidationInputRequest{
							ValidationType: armdatabox.ValidationInputDiscriminatorValidateDataTransferDetails.ToPtr(),
						},
						DataImportDetails: []*armdatabox.DataImportDetails{
							{
								AccountDetails: &armdatabox.StorageAccountDetails{
									DataAccountDetails: armdatabox.DataAccountDetails{
										DataAccountType: armdatabox.DataAccountTypeStorageAccount.ToPtr(),
									},
									StorageAccountID: to.StringPtr("<storage-account-id>"),
								},
							}},
						DeviceType:   armdatabox.SKUNameDataBox.ToPtr(),
						TransferType: armdatabox.TransferTypeImportToAzure.ToPtr(),
					},
					&armdatabox.ValidateAddress{
						ValidationInputRequest: armdatabox.ValidationInputRequest{
							ValidationType: armdatabox.ValidationInputDiscriminatorValidateAddress.ToPtr(),
						},
						DeviceType: armdatabox.SKUNameDataBox.ToPtr(),
						ShippingAddress: &armdatabox.ShippingAddress{
							AddressType:     armdatabox.AddressTypeCommercial.ToPtr(),
							City:            to.StringPtr("<city>"),
							CompanyName:     to.StringPtr("<company-name>"),
							Country:         to.StringPtr("<country>"),
							PostalCode:      to.StringPtr("<postal-code>"),
							StateOrProvince: to.StringPtr("<state-or-province>"),
							StreetAddress1:  to.StringPtr("<street-address1>"),
							StreetAddress2:  to.StringPtr("<street-address2>"),
						},
						TransportPreferences: &armdatabox.TransportPreferences{
							PreferredShipmentType: armdatabox.TransportShipmentTypesMicrosoftManaged.ToPtr(),
						},
					},
					&armdatabox.SubscriptionIsAllowedToCreateJobValidationRequest{
						ValidationInputRequest: armdatabox.ValidationInputRequest{
							ValidationType: armdatabox.ValidationInputDiscriminatorValidateSubscriptionIsAllowedToCreateJob.ToPtr(),
						},
					},
					&armdatabox.SKUAvailabilityValidationRequest{
						ValidationInputRequest: armdatabox.ValidationInputRequest{
							ValidationType: armdatabox.ValidationInputDiscriminatorValidateSKUAvailability.ToPtr(),
						},
						Country:      to.StringPtr("<country>"),
						DeviceType:   armdatabox.SKUNameDataBox.ToPtr(),
						Location:     to.StringPtr("<location>"),
						TransferType: armdatabox.TransferTypeImportToAzure.ToPtr(),
					},
					&armdatabox.CreateOrderLimitForSubscriptionValidationRequest{
						ValidationInputRequest: armdatabox.ValidationInputRequest{
							ValidationType: armdatabox.ValidationInputDiscriminatorValidateCreateOrderLimit.ToPtr(),
						},
						DeviceType: armdatabox.SKUNameDataBox.ToPtr(),
					},
					&armdatabox.PreferencesValidationRequest{
						ValidationInputRequest: armdatabox.ValidationInputRequest{
							ValidationType: armdatabox.ValidationInputDiscriminatorValidatePreferences.ToPtr(),
						},
						DeviceType: armdatabox.SKUNameDataBox.ToPtr(),
						Preference: &armdatabox.Preferences{
							TransportPreferences: &armdatabox.TransportPreferences{
								PreferredShipmentType: armdatabox.TransportShipmentTypesMicrosoftManaged.ToPtr(),
							},
						},
					}},
				ValidationCategory: to.StringPtr("<validation-category>"),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
}

// x-ms-original-file: specification/databox/resource-manager/Microsoft.DataBox/preview/2021-08-01-preview/examples/RegionConfiguration.json
func ExampleServiceClient_RegionConfiguration() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armdatabox.NewServiceClient("<subscription-id>", cred, nil)
	_, err = client.RegionConfiguration(ctx,
		"<location>",
		armdatabox.RegionConfigurationRequest{
			ScheduleAvailabilityRequest: &armdatabox.DataBoxScheduleAvailabilityRequest{
				ScheduleAvailabilityRequest: armdatabox.ScheduleAvailabilityRequest{
					SKUName:         armdatabox.SKUNameDataBox.ToPtr(),
					StorageLocation: to.StringPtr("<storage-location>"),
				},
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
}

// x-ms-original-file: specification/databox/resource-manager/Microsoft.DataBox/preview/2021-08-01-preview/examples/RegionConfigurationByResourceGroup.json
func ExampleServiceClient_RegionConfigurationByResourceGroup() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armdatabox.NewServiceClient("<subscription-id>", cred, nil)
	_, err = client.RegionConfigurationByResourceGroup(ctx,
		"<resource-group-name>",
		"<location>",
		armdatabox.RegionConfigurationRequest{
			ScheduleAvailabilityRequest: &armdatabox.DataBoxScheduleAvailabilityRequest{
				ScheduleAvailabilityRequest: armdatabox.ScheduleAvailabilityRequest{
					SKUName:         armdatabox.SKUNameDataBox.ToPtr(),
					StorageLocation: to.StringPtr("<storage-location>"),
				},
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
}
