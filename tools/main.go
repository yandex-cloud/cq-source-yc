package main

import "fmt"

func main() {
	var err error

	err = Generate("Kms", "SymmetricKey", "yandex/cloud/kms/v1/symmetric_key.proto", WithProtoPaths("../cloudapi"))

	if err != nil {
		fmt.Println(err)
	}

	err = Generate("Compute", "Image", "yandex/cloud/compute/v1/image.proto", WithProtoPaths("../cloudapi"))

	if err != nil {
		fmt.Println(err)
	}

	err = Generate("Compute", "Instance", "yandex/cloud/compute/v1/instance.proto",
		WithProtoPaths("../cloudapi"),
		WithIgnoredColumns("NetworkInterfaces.PrimaryV4Address", "NetworkInterfaces.PrimaryV6Address"))

	if err != nil {
		fmt.Println(err)
	}

	err = Generate("Compute", "Disk", "yandex/cloud/compute/v1/disk.proto", WithProtoPaths("../cloudapi"))

	if err != nil {
		fmt.Println(err)
	}
	err = Generate("Vpc", "Network", "yandex/cloud/vpc/v1/network.proto", WithProtoPaths("../cloudapi"))

	if err != nil {
		fmt.Println(err)
	}

	err = Generate("Vpc", "Subnet", "yandex/cloud/vpc/v1/subnet.proto", WithProtoPaths("../cloudapi"))

	if err != nil {
		fmt.Println(err)
	}

	err = Generate("Vpc", "Address", "yandex/cloud/vpc/v1/address.proto", WithProtoPaths("../cloudapi"))

	if err != nil {
		fmt.Println(err)
	}

	err = Generate("Iam", "ServiceAccount", "yandex/cloud/iam/v1/service_account.proto", WithProtoPaths("../cloudapi"))

	if err != nil {
		fmt.Println(err)
	}
}
