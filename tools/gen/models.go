package gen

type ResourceFileModel struct {
	Table     *TableModel
	Relations []*TableModel
}

type TableModel struct {
	// All fields below should be in camel case
	Service           string
	Resource          string
	AbsolutFieldPath  []string
	RelativeFieldPath []string

	Multiplex string

	Columns   []*ColumnModel
	Relations []*TableModel
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
