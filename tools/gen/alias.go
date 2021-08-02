package gen

type Alias interface {
	ApplyToColumn(column *ColumnModel)
	ApplyToTable(table *TableModel)
}

type UnimplementedAlias struct{}

func (u UnimplementedAlias) ApplyToColumn(*ColumnModel) {}

func (u UnimplementedAlias) ApplyToTable(*TableModel) {}

type changeName struct {
	UnimplementedAlias
	Name string
}

func ChangeName(name string) Alias {
	return changeName{Name: name}
}

func (c changeName) ApplyToColumn(column *ColumnModel) {
	column.Name = c.Name
}

func (c changeName) ApplyToTable(table *TableModel) {
	table.Alias = c.Name
}

type changeColumn struct {
	UnimplementedAlias
	column *ColumnModel
}

func ChangeColumn(column *ColumnModel) Alias {
	return changeColumn{column: column}
}

func (r changeColumn) ApplyToColumn(column *ColumnModel) {
	*column = *r.column
}
