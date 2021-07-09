package main

import (
	"github.com/cloudquery/cq-provider-sdk/serve"
	"github.com/yandex-cloud/cq-provider-yandex/resources"
)

func main() {
	serve.Serve(&serve.Options{
		Name:     "yandex",
		Provider: resources.Provider(),
	})
}

//func main() {
//	//table, _ := resources.ComputeResource("resources/proto/disk.proto", resources.DefaultColumns{})
//	tableGen, _ := resources.NewTableGenerator(
//		"instance", "resources/proto/instance.proto",
//		resources.DefaultColumns{
//			"CreatedAt": {
//				Name:     "created_at",
//				Type:     schema.TypeTimestamp,
//				Resolver: client.ResolveAsTime,
//			},
//		},
//		[]string{"Status"},
//	)
//	table, _ := tableGen.Generate()
//	fmt.Println(table)
//}
