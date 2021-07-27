package main

import (
	"fmt"

	"github.com/yandex-cloud/cq-provider-yandex/tools/gen"
)

func main() {
	var err error

	err = gen.Generate(
		"Kms",
		"SymmetricKey",
		"yandex/cloud/kms/v1/symmetric_key.proto",
		"resources",
		gen.WithProtoPaths("cloudapi"),
	)

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

	err = gen.Generate(
		"Compute",
		"Instance",
		"yandex/cloud/compute/v1/instance.proto",
		"resources",
		gen.WithProtoPaths("cloudapi"),
		gen.WithIgnoredColumns("NetworkInterfaces.PrimaryV4Address", "NetworkInterfaces.PrimaryV6Address"),
	)

	if err != nil {
		fmt.Println(err)
	}

	err = gen.Generate(
		"Compute",
		"Disk",
		"yandex/cloud/compute/v1/disk.proto",
		"resources",
		gen.WithProtoPaths("cloudapi"),
		gen.WithIgnoredColumns("SourceImageId", "SourceSnapshotId"), // to avoid tests fail due to empty columns corresponding oneof fields
	)

	if err != nil {
		fmt.Println(err)
	}

	err = gen.Generate(
		"Vpc",
		"Network",
		"yandex/cloud/vpc/v1/network.proto",
		"resources",
		gen.WithProtoPaths("cloudapi"),
	)

	if err != nil {
		fmt.Println(err)
	}

	err = gen.Generate(
		"Vpc",
		"Subnet",
		"yandex/cloud/vpc/v1/subnet.proto",
		"resources",
		gen.WithProtoPaths("cloudapi"),
	)

	if err != nil {
		fmt.Println(err)
	}

	err = gen.Generate(
		"Vpc",
		"Address",
		"yandex/cloud/vpc/v1/address.proto",
		"resources",
		gen.WithProtoPaths("cloudapi"),
	)

	if err != nil {
		fmt.Println(err)
	}

	err = gen.Generate(
		"Iam",
		"ServiceAccount",
		"yandex/cloud/iam/v1/service_account.proto",
		"resources",
		gen.WithProtoPaths("cloudapi"),
	)

	if err != nil {
		fmt.Println(err)
	}
}
