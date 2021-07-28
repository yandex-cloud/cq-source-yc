package main

import (
	"log"

	"github.com/cloudquery/cq-provider-sdk/provider/docs"
	"github.com/yandex-cloud/cq-provider-yandex/resources"
)

func main() {
	err := docs.GenerateDocs(resources.Provider(), "./docs")
	if err != nil {
		log.Fatalf("Failed to geneerate docs: %s", err)
	}
}
