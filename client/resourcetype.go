package client

import (
	"github.com/yandex-cloud/go-genproto/yandex/cloud/billing/v1"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/organizationmanager/v1"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/resourcemanager/v1"
)

type ResourceType string

const (
	ResourceTypeOrganization   ResourceType = "organization-manager.organization"
	ResourceTypeCloud          ResourceType = "resource-manager.cloud"
	ResourceTypeFolder         ResourceType = "resource-manager.folder"
	ResourceTypeBillingAccount ResourceType = "billing.account"
)

// TODO: codegen
func ResourceTypeFromProto(p any) (ResourceType, bool) {
	switch p.(type) {
	case organizationmanager.Organization:
		return ResourceTypeOrganization, true
	case resourcemanager.Cloud:
		return ResourceTypeCloud, true
	case billing.BillingAccount:
		return ResourceTypeBillingAccount, true
	default:
		return ResourceType(""), false
	}
}
