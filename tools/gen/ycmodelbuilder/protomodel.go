package ycmodelbuilder

import "github.com/yandex-cloud/cq-provider-yandex/tools/gen/ycmodel"

var defaultYCOptions = []Option{
	WithAlias("Id", ChangeColumn(
		&ycmodel.Column{
			Name:        "id",
			Type:        "schema.TypeString",
			Description: "ID of the resource.",
			Resolver:    "client.ResolveResourceId",
		},
	),
	),
	WithAlias("FolderId", ChangeColumn(
		&ycmodel.Column{
			Name:        "folder_id",
			Type:        "schema.TypeString",
			Description: "ID of the folder that the resource belongs to.",
			Resolver:    "client.ResolveFolderID",
		},
	),
	),
	WithAlias("CreatedAt", ChangeColumn(
		&ycmodel.Column{
			Name:     "created_at",
			Type:     "schema.TypeTimestamp",
			Resolver: "client.ResolveAsTime",
		},
	),
	),
	WithAlias("Labels", ChangeColumn(
		&ycmodel.Column{
			Name:        "labels",
			Type:        "schema.TypeJSON",
			Description: "Resource labels as `key:value` pairs. Maximum of 64 per resource.",
			Resolver:    "client.ResolveLabels",
		},
	),
	),
}

func ResourceFileFromProto(service, resource, pathToProto string, opts ...Option) (*ycmodel.File, error) {
	defaultOptions := defaultYCOptions
	defaultOptions = append(defaultOptions, opts...)
	co := NewCollapsedOptions(defaultOptions)

	b := TableBuilder{
		Service:       service,
		Multiplex:     "client.FolderMultiplex",
		IgnoredFields: co.IgnoredFields,
		Aliases:       co.Aliases,
	}

	err := b.WithMessageFromProto(resource, pathToProto, co.Paths...)
	if err != nil {
		return nil, err
	}

	tableModel, err := b.Build()
	if err != nil {
		return nil, err
	}

	return &ycmodel.File{Table: tableModel, Relations: expandRelations(tableModel)}, nil
}

func expandRelations(table *ycmodel.Table) (tables []*ycmodel.Table) {
	for _, relation := range table.Relations {
		tables = append(tables, expandRelations(relation)...)
		tables = append(tables, relation)
	}
	return
}
