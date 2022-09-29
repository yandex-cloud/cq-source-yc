package recipes

import (
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/certificatemanager/v1"
)

func CertificateManager() []*Resource {
	return []*Resource{
		{
			Service:      "certificatemanager",
			SubService:   "certificates",
			Struct:       new(certificatemanager.Certificate),
			SkipFields:   []string{id},
			ExtraColumns: codegen.ColumnDefinitions{idCol},
			Multiplex:    multiplexFolder,
		},
	}
}
