// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle Cloud VMware Solution API
//
// Use the Oracle Cloud VMware API to create SDDCs and manage ESXi hosts and software.
// For more information, see Oracle Cloud VMware Solution (https://docs.cloud.oracle.com/iaas/Content/VMware/Concepts/ocvsoverview.htm).
//

package ocvp

import (
	"strings"
)

// OperationTypesEnum Enum with underlying type: string
type OperationTypesEnum string

// Set of constants representing the allowable values for OperationTypesEnum
const (
	OperationTypesCreateSddc              OperationTypesEnum = "CREATE_SDDC"
	OperationTypesDeleteSddc              OperationTypesEnum = "DELETE_SDDC"
	OperationTypesCreateCluster           OperationTypesEnum = "CREATE_CLUSTER"
	OperationTypesDeleteCluster           OperationTypesEnum = "DELETE_CLUSTER"
	OperationTypesCreateEsxiHost          OperationTypesEnum = "CREATE_ESXI_HOST"
	OperationTypesDeleteEsxiHost          OperationTypesEnum = "DELETE_ESXI_HOST"
	OperationTypesUpgradeHcx              OperationTypesEnum = "UPGRADE_HCX"
	OperationTypesDowngradeHcx            OperationTypesEnum = "DOWNGRADE_HCX"
	OperationTypesCancelDowngradeHcx      OperationTypesEnum = "CANCEL_DOWNGRADE_HCX"
	OperationTypesRefreshHcxLicenseStatus OperationTypesEnum = "REFRESH_HCX_LICENSE_STATUS"
	OperationTypesSwapBilling             OperationTypesEnum = "SWAP_BILLING"
	OperationTypesReplaceHost             OperationTypesEnum = "REPLACE_HOST"
	OperationTypesInPlaceUpgrade          OperationTypesEnum = "IN_PLACE_UPGRADE"
)

var mappingOperationTypesEnum = map[string]OperationTypesEnum{
	"CREATE_SDDC":                OperationTypesCreateSddc,
	"DELETE_SDDC":                OperationTypesDeleteSddc,
	"CREATE_CLUSTER":             OperationTypesCreateCluster,
	"DELETE_CLUSTER":             OperationTypesDeleteCluster,
	"CREATE_ESXI_HOST":           OperationTypesCreateEsxiHost,
	"DELETE_ESXI_HOST":           OperationTypesDeleteEsxiHost,
	"UPGRADE_HCX":                OperationTypesUpgradeHcx,
	"DOWNGRADE_HCX":              OperationTypesDowngradeHcx,
	"CANCEL_DOWNGRADE_HCX":       OperationTypesCancelDowngradeHcx,
	"REFRESH_HCX_LICENSE_STATUS": OperationTypesRefreshHcxLicenseStatus,
	"SWAP_BILLING":               OperationTypesSwapBilling,
	"REPLACE_HOST":               OperationTypesReplaceHost,
	"IN_PLACE_UPGRADE":           OperationTypesInPlaceUpgrade,
}

var mappingOperationTypesEnumLowerCase = map[string]OperationTypesEnum{
	"create_sddc":                OperationTypesCreateSddc,
	"delete_sddc":                OperationTypesDeleteSddc,
	"create_cluster":             OperationTypesCreateCluster,
	"delete_cluster":             OperationTypesDeleteCluster,
	"create_esxi_host":           OperationTypesCreateEsxiHost,
	"delete_esxi_host":           OperationTypesDeleteEsxiHost,
	"upgrade_hcx":                OperationTypesUpgradeHcx,
	"downgrade_hcx":              OperationTypesDowngradeHcx,
	"cancel_downgrade_hcx":       OperationTypesCancelDowngradeHcx,
	"refresh_hcx_license_status": OperationTypesRefreshHcxLicenseStatus,
	"swap_billing":               OperationTypesSwapBilling,
	"replace_host":               OperationTypesReplaceHost,
	"in_place_upgrade":           OperationTypesInPlaceUpgrade,
}

// GetOperationTypesEnumValues Enumerates the set of values for OperationTypesEnum
func GetOperationTypesEnumValues() []OperationTypesEnum {
	values := make([]OperationTypesEnum, 0)
	for _, v := range mappingOperationTypesEnum {
		values = append(values, v)
	}
	return values
}

// GetOperationTypesEnumStringValues Enumerates the set of values in String for OperationTypesEnum
func GetOperationTypesEnumStringValues() []string {
	return []string{
		"CREATE_SDDC",
		"DELETE_SDDC",
		"CREATE_CLUSTER",
		"DELETE_CLUSTER",
		"CREATE_ESXI_HOST",
		"DELETE_ESXI_HOST",
		"UPGRADE_HCX",
		"DOWNGRADE_HCX",
		"CANCEL_DOWNGRADE_HCX",
		"REFRESH_HCX_LICENSE_STATUS",
		"SWAP_BILLING",
		"REPLACE_HOST",
		"IN_PLACE_UPGRADE",
	}
}

// GetMappingOperationTypesEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOperationTypesEnum(val string) (OperationTypesEnum, bool) {
	enum, ok := mappingOperationTypesEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
