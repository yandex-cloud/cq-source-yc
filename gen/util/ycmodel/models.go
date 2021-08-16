package ycmodel

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
	Name        string
	Description string
	Type        string
	Resolver    string
}

type File struct {
	Table     *Table
	Relations []*Table
}
