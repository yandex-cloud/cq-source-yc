package main

import (
	"log"

	"github.com/cloudquery/cq-provider-sdk/provider/docs"
	"github.com/yandex-cloud/cq-provider-yandex/resources/provider"
)

func main() {
	err := docs.GenerateDocs(provider.Provider(), "./docs", false)
	if err != nil {
		log.Fatalf("Failed to geneerate docs: %s", err)
	}
}
