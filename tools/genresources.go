package main

import (
	"fmt"

	"github.com/yandex-cloud/cq-provider-yandex/tools/gen"
)

func main() {
	var err error

	err = gen.Generate(
		"KMS",
		"SymmetricKey",
		"yandex/cloud/kms/v1/symmetric_key.proto",
		"resources",
		gen.WithProtoPaths("cloudapi"),

		// Field doesn't exist in corresponding struct
		gen.WithIgnoredColumns("PrimaryVersion.HostedByHsm"),
	)

	if err != nil {
		fmt.Println(err)
	}

	err = gen.GenerateTests("KMS", "SymmetricKey", "resources")

	if err != nil {
		fmt.Println(err)
	}

	err = gen.Generate(
		"Compute",
		"Image",
		"yandex/cloud/compute/v1/image.proto",
		"resources",
		gen.WithProtoPaths("cloudapi"),
	)

	if err != nil {
		fmt.Println(err)
	}

	err = gen.GenerateTests("Compute", "Image", "resources")

	if err != nil {
		fmt.Println(err)
	}

	err = gen.Generate(
		"Compute",
		"Instance",
		"yandex/cloud/compute/v1/instance.proto",
		"resources",
		gen.WithProtoPaths("cloudapi"),

		// TODO: names of corresponding columns greater then PostgreSQL limit (64 bytes)
		gen.WithIgnoredColumns("NetworkInterfaces.PrimaryV4Address", "NetworkInterfaces.PrimaryV6Address"),
	)

	if err != nil {
		fmt.Println(err)
	}

	err = gen.GenerateTests("Compute", "Instance", "resources")

	if err != nil {
		fmt.Println(err)
	}

	err = gen.Generate(
		"Compute",
		"Disk",
		"yandex/cloud/compute/v1/disk.proto",
		"resources",
		gen.WithProtoPaths("cloudapi"),

		// TODO: test framework from sdk fails on empty columns, so oneof fields ignored for a while
		gen.WithIgnoredColumns("Source.SourceImageId", "Source.SourceSnapshotId"),
	)

	if err != nil {
		fmt.Println(err)
	}

	err = gen.GenerateTests("Compute", "Disk", "resources")

	if err != nil {
		fmt.Println(err)
	}

	err = gen.Generate(
		"VPC",
		"Network",
		"yandex/cloud/vpc/v1/network.proto",
		"resources",
		gen.WithProtoPaths("cloudapi"),
	)

	if err != nil {
		fmt.Println(err)
	}

	err = gen.GenerateTests("VPC", "Network", "resources")

	if err != nil {
		fmt.Println(err)
	}

	err = gen.Generate(
		"VPC",
		"Subnet",
		"yandex/cloud/vpc/v1/subnet.proto",
		"resources",
		gen.WithProtoPaths("cloudapi"),
	)

	if err != nil {
		fmt.Println(err)
	}

	err = gen.GenerateTests("VPC", "Subnet", "resources")

	if err != nil {
		fmt.Println(err)
	}

	err = gen.Generate(
		"VPC",
		"Address",
		"yandex/cloud/vpc/v1/address.proto",
		"resources",
		gen.WithProtoPaths("cloudapi"),

		// TODO: test framework from sdk fails on empty columns, so oneof fields ignored for a while
		gen.WithIgnoredColumns("Address.ExternalIpv4Address"),
	)

	if err != nil {
		fmt.Println(err)
	}

	err = gen.GenerateTests("VPC", "Address", "resources")

	if err != nil {
		fmt.Println(err)
	}

	err = gen.Generate(
		"IAM",
		"ServiceAccount",
		"yandex/cloud/iam/v1/service_account.proto",
		"resources",
		gen.WithProtoPaths("cloudapi"),
	)

	if err != nil {
		fmt.Println(err)
	}

	err = gen.GenerateTests("IAM", "ServiceAccount", "resources")

	if err != nil {
		fmt.Println(err)
	}
}
