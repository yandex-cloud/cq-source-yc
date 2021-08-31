package modelfromproto

type Table struct {
	// Should be in camel case
	Service      string
	Resource     string
	AbsolutePath []string
	RelativePath []string

	Multiplex string

	Columns   []*Column
	Relations []*Table

	Alias string
}

type Column struct {
	Name            string
	Description     string
	Type            string
	Resolver        string
	CreationOptions *CreationOptions
}

type CreationOptions struct {
	Nullable string
	Unique   string
}

type File struct {
	Table     *Table
	Relations []*Table
}

var defaultYCOptions = []Option{
	WithAlias("Id", ChangeColumn(
		&Column{
			Name:        "id",
			Type:        "schema.TypeString",
			Description: "ID of the resource.",
			Resolver:    "client.ResolveResourceId",
		},
	),
	),
	WithAlias("FolderId", ChangeColumn(
		&Column{
			Name:        "folder_id",
			Type:        "schema.TypeString",
			Description: "ID of the folder that the resource belongs to.",
			Resolver:    "client.ResolveFolderID",
		},
	),
	),
	WithAlias("CreatedAt", ChangeColumn(
		&Column{
			Name:     "created_at",
			Type:     "schema.TypeTimestamp",
			Resolver: "client.ResolveAsTime",
		},
	),
	),
	WithAlias("Labels", ChangeColumn(
		&Column{
			Name:        "labels",
			Type:        "schema.TypeJSON",
			Description: "Resource labels as `key:value` pairs. Maximum of 64 per resource.",
			Resolver:    "client.ResolveLabels",
		},
	),
	),
}

func ResourceFileFromProto(service, resource, pathToProto string, opts ...Option) (*File, error) {
	defaultOptions := defaultYCOptions
	defaultOptions = append(defaultOptions, opts...)
	co := NewCollapsedOptions(defaultOptions)

	b := TableBuilder{
		Service:       service,
		Multiplex:     "client.MultiplexBy(client.Folders)",
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

	return &File{Table: tableModel, Relations: expandRelations(tableModel)}, nil
}

func expandRelations(table *Table) (tables []*Table) {
	for _, relation := range table.Relations {
		tables = append(tables, expandRelations(relation)...)
		tables = append(tables, relation)
	}
	return
}
