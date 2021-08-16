package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/iancoleman/strcase"
	"github.com/jinzhu/inflection"
	"github.com/yandex-cloud/cq-provider-yandex/gen/util"
	ycmb "github.com/yandex-cloud/cq-provider-yandex/gen/util/ycmodelbuilder"
)

func generate(service, resource, pathToProto string, opts ...ycmb.Option) {
	opts = append(opts, ycmb.WithProtoPaths("cloudapi", "cloudapi/third_party/googleapis"))

	resourceFileModel, err := ycmb.ResourceFileFromProto(service, resource, pathToProto, opts...)
	if err != nil {
		fmt.Fprint(os.Stderr)
		return
	}

	out := filepath.Join(util.ResourcesDir, util.ToTogether(service)+"_"+strcase.ToSnake(inflection.Plural(resource))+".go")

	util.SilentExecute(util.TemplatesDir{
		MainFile: "resource.go.tmpl",
		Path:     "templates",
	}, resourceFileModel, out)
}

func main() {
	generate("Compute", "Disk", "yandex/cloud/compute/v1/disk.proto")

	generate("Compute", "Image", "yandex/cloud/compute/v1/image.proto")

	generate("Compute", "Instance", "yandex/cloud/compute/v1/instance.proto",
		ycmb.WithAlias("NetworkInterfaces.PrimaryV4Address.DnsRecords",
			ycmb.ChangeName("yandex_compute_instance_net_interface_ipv4_dns_records")),
		ycmb.WithAlias("NetworkInterfaces.PrimaryV4Address.OneToOneNat.DnsRecords",
			ycmb.ChangeName("yandex_compute_instance_net_interface_ipv4_1_1_nat_dns_records")),
		ycmb.WithAlias("NetworkInterfaces.PrimaryV6Address.DnsRecords",
			ycmb.ChangeName("yandex_compute_instance_net_interface_ipv6_dns_records")),
		ycmb.WithAlias("NetworkInterfaces.PrimaryV6Address.OneToOneNat.DnsRecords",
			ycmb.ChangeName("yandex_compute_instance_net_interface_ipv6_1_1_nat_dns_records")))

	generate("VPC", "Network", "yandex/cloud/vpc/v1/network.proto")

	generate("VPC", "Subnet", "yandex/cloud/vpc/v1/subnet.proto")

	generate("VPC", "Address", "yandex/cloud/vpc/v1/address.proto",
		ycmb.WithAlias("Address.ExternalIpv4Address.Requirements.DdosProtectionProvider",
			ycmb.ChangeName("addr_ext_ipv_4_addr_requirements_ddos_protect_prov")),
		ycmb.WithAlias("Address.ExternalIpv4Address.Requirements.OutgoingSmtpCapability",
			ycmb.ChangeName("addr_ext_ipv_4_addr_requirements_out_smtp_cap")))

	generate("VPC", "SecurityGroup", "yandex/cloud/vpc/v1/security_group.proto")

	generate("IAM", "ServiceAccount", "yandex/cloud/iam/v1/service_account.proto")

	generate("K8S", "Cluster", "yandex/cloud/k8s/v1/cluster.proto",
		ycmb.WithIgnored("Master.MaintenancePolicy"))

	generate("CertificateManager", "Certificate",
		"yandex/cloud/certificatemanager/v1/certificate.proto", ycmb.WithProtoPaths("cloudapi"))

	generate("KMS", "SymmetricKey", "yandex/cloud/kms/v1/symmetric_key.proto")
}
