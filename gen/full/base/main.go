package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/iancoleman/strcase"
	"github.com/jinzhu/inflection"
	"github.com/yandex-cloud/cq-provider-yandex/gen/util"
	"github.com/yandex-cloud/cq-provider-yandex/gen/util/modelfromproto"
)

func generate(service, resource, pathToProto string, opts ...modelfromproto.Option) {
	opts = append(opts, modelfromproto.WithProtoPaths("cloudapi", "cloudapi/third_party/googleapis"))

	resourceFileModel, err := modelfromproto.ResourceFileFromProto(service, resource, pathToProto, opts...)
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		return
	}

	out := filepath.Join(util.ResourcesDir, util.ToFlat(service)+"_"+strcase.ToSnake(inflection.Plural(resource))+".go")

	util.SilentExecute(util.TemplatesDir{
		MainFile: "base_resource.go.tmpl",
		Path:     "templates",
	}, resourceFileModel, out)
}

func main() {
	generate("Compute", "Disk", "yandex/cloud/compute/v1/disk.proto")

	generate("Compute", "Image", "yandex/cloud/compute/v1/image.proto")

	generate("Compute", "Instance", "yandex/cloud/compute/v1/instance.proto",
		modelfromproto.WithAlias("NetworkInterfaces.PrimaryV4Address.DnsRecords",
			modelfromproto.ChangeName("yandex_compute_instance_net_interface_ipv4_dns_records")),
		modelfromproto.WithAlias("NetworkInterfaces.PrimaryV4Address.OneToOneNat.DnsRecords",
			modelfromproto.ChangeName("yandex_compute_instance_net_interface_ipv4_1_1_nat_dns_records")),
		modelfromproto.WithAlias("NetworkInterfaces.PrimaryV6Address.DnsRecords",
			modelfromproto.ChangeName("yandex_compute_instance_net_interface_ipv6_dns_records")),
		modelfromproto.WithAlias("NetworkInterfaces.PrimaryV6Address.OneToOneNat.DnsRecords",
			modelfromproto.ChangeName("yandex_compute_instance_net_interface_ipv6_1_1_nat_dns_records")))

	generate("VPC", "Network", "yandex/cloud/vpc/v1/network.proto")

	generate("VPC", "Subnet", "yandex/cloud/vpc/v1/subnet.proto")

	generate("VPC", "Address", "yandex/cloud/vpc/v1/address.proto",
		modelfromproto.WithAlias("Address.ExternalIpv4Address.Requirements.DdosProtectionProvider",
			modelfromproto.ChangeName("addr_ext_ipv_4_addr_requirements_ddos_protect_prov")),
		modelfromproto.WithAlias("Address.ExternalIpv4Address.Requirements.OutgoingSmtpCapability",
			modelfromproto.ChangeName("addr_ext_ipv_4_addr_requirements_out_smtp_cap")))

	generate("VPC", "SecurityGroup", "yandex/cloud/vpc/v1/security_group.proto")

	generate("IAM", "ServiceAccount", "yandex/cloud/iam/v1/service_account.proto")

	generate("K8S", "Cluster", "yandex/cloud/k8s/v1/cluster.proto",
		modelfromproto.WithIgnored("Master.MaintenancePolicy"))

	generate("CertificateManager", "Certificate",
		"yandex/cloud/certificatemanager/v1/certificate.proto", modelfromproto.WithProtoPaths("cloudapi"))

	generate("KMS", "SymmetricKey", "yandex/cloud/kms/v1/symmetric_key.proto")
}
