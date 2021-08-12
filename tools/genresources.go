package main

import (
	"fmt"

	"github.com/yandex-cloud/cq-provider-yandex/tools/gen/base"
	"github.com/yandex-cloud/cq-provider-yandex/tools/gen/serverless"
	ycmb "github.com/yandex-cloud/cq-provider-yandex/tools/gen/ycmodelbuilder"
)

func main() {
	var err error

	err = base.Generate(
		"KMS",
		"SymmetricKey",
		"yandex/cloud/kms/v1/symmetric_key.proto",
		"resources",
		ycmb.WithProtoPaths("cloudapi"),

		// Field doesn't exist in used version of sdk
		ycmb.WithIgnoredColumns("PrimaryVersion.HostedByHsm"),
	)

	if err != nil {
		fmt.Println(err)
	}

	err = base.Generate(
		"Compute",
		"Image",
		"yandex/cloud/compute/v1/image.proto",
		"resources",
		ycmb.WithProtoPaths("cloudapi"),
	)

	if err != nil {
		fmt.Println(err)
	}

	err = base.Generate(
		"Compute",
		"Instance",
		"yandex/cloud/compute/v1/instance.proto",
		"resources",
		ycmb.WithProtoPaths("cloudapi"),
		ycmb.WithAlias(
			"NetworkInterfaces.PrimaryV4Address.DnsRecords",
			ycmb.ChangeName("yandex_compute_instance_net_interface_ipv4_dns_records"),
		),
		ycmb.WithAlias(
			"NetworkInterfaces.PrimaryV4Address.OneToOneNat.DnsRecords",
			ycmb.ChangeName("yandex_compute_instance_net_interface_ipv4_1_1_nat_dns_records"),
		),
		ycmb.WithAlias(
			"NetworkInterfaces.PrimaryV6Address.DnsRecords",
			ycmb.ChangeName("yandex_compute_instance_net_interface_ipv6_dns_records"),
		),
		ycmb.WithAlias(
			"NetworkInterfaces.PrimaryV6Address.OneToOneNat.DnsRecords",
			ycmb.ChangeName("yandex_compute_instance_net_interface_ipv6_1_1_nat_dns_records"),
		),
	)

	if err != nil {
		fmt.Println(err)
	}

	err = base.Generate(
		"Compute",
		"Disk",
		"yandex/cloud/compute/v1/disk.proto",
		"resources",
		ycmb.WithProtoPaths("cloudapi"),
	)

	if err != nil {
		fmt.Println(err)
	}

	err = base.Generate(
		"VPC",
		"Network",
		"yandex/cloud/vpc/v1/network.proto",
		"resources",
		ycmb.WithProtoPaths("cloudapi"),
	)

	if err != nil {
		fmt.Println(err)
	}

	err = base.Generate(
		"VPC",
		"Subnet",
		"yandex/cloud/vpc/v1/subnet.proto",
		"resources",
		ycmb.WithProtoPaths("cloudapi"),
	)

	if err != nil {
		fmt.Println(err)
	}

	err = base.Generate(
		"VPC",
		"Address",
		"yandex/cloud/vpc/v1/address.proto",
		"resources",
		ycmb.WithProtoPaths("cloudapi"),
		ycmb.WithAlias(
			"Address.ExternalIpv4Address.Requirements.DdosProtectionProvider",
			ycmb.ChangeName("addr_ext_ipv_4_addr_requirements_ddos_protect_prov"),
		),
		ycmb.WithAlias(
			"Address.ExternalIpv4Address.Requirements.OutgoingSmtpCapability",
			ycmb.ChangeName("addr_ext_ipv_4_addr_requirements_out_smtp_cap"),
		),
	)

	if err != nil {
		fmt.Println(err)
	}

	err = base.Generate(
		"IAM",
		"ServiceAccount",
		"yandex/cloud/iam/v1/service_account.proto",
		"resources",
		ycmb.WithProtoPaths("cloudapi"),
	)

	if err != nil {
		fmt.Println(err)
	}

	err = base.Generate(
		"K8S",
		"Cluster",
		"yandex/cloud/k8s/v1/cluster.proto",
		"resources",
		ycmb.WithProtoPaths("cloudapi", "cloudapi/third_party/googleapis"),
		ycmb.WithIgnoredColumns("Master.MaintenancePolicy"),
	)

	if err != nil {
		fmt.Println(err)
	}

	err = base.Generate(
		"VPC",
		"SecurityGroup",
		"yandex/cloud/vpc/v1/security_group.proto",
		"resources",
		ycmb.WithProtoPaths("cloudapi"),
	)

	if err != nil {
		fmt.Println(err)
	}

	err = base.Generate(
		"CertificateManager",
		"Certificate",
		"yandex/cloud/certificatemanager/v1/certificate.proto",
		"resources",
		ycmb.WithProtoPaths("cloudapi"),
	)

	if err != nil {
		fmt.Println(err)
	}

	err = serverless.Generate(
		"ApiGateway",
		"yandex/cloud/serverless/apigateway/v1/apigateway.proto",
		"resources",
		ycmb.WithProtoPaths("cloudapi"),
	)

	if err != nil {
		fmt.Println(err)
	}

	//err = base.Generate(
	//	"OrganizationManager",
	//	"Organization",
	//	"yandex/cloud/organizationmanager/v1/organization.proto",
	//	"resources",
	//	ycmb.WithProtoPaths("cloudapi"),
	//)
	//
	//if err != nil {
	//	fmt.Println(err)
	//}

	//err = base.Generate(
	//	"OrganizationManager",
	//	"Federation",
	//	"yandex/cloud/organizationmanager/v1/saml/federation.proto",
	//	"resources",
	//	ycmb.WithProtoPaths("cloudapi"),
	//)
	//
	//if err != nil {
	//	fmt.Println(err)
	//}
}
