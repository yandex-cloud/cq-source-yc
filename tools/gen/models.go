package gen

type ResourceFileModel struct {
	Table     *TableModel
	Relations []*TableModel
}

type TableModel struct {
	// Should be in camel case
	Service      string
	Resource     string
	AbsolutePath []string
	RelativePath []string

	Multiplex string

	Columns   []*ColumnModel
	Relations []*TableModel

	Alias string
}

type ColumnModel struct {
	Name        string
	Description string
	Type        string
	Resolver    string
}

type ResourceTestFileModel struct {
	Service  string
	Resource string
}
