package ycmodelbuilder

import "github.com/yandex-cloud/cq-provider-yandex/tools/gen/ycmodel"

type Alias interface {
	ApplyToColumn(column *ycmodel.Column)
	ApplyToTable(table *ycmodel.Table)
}

type UnimplementedAlias struct{}

func (u UnimplementedAlias) ApplyToColumn(*ycmodel.Column) {}

func (u UnimplementedAlias) ApplyToTable(*ycmodel.Table) {}

type changeName struct {
	UnimplementedAlias
	Name string
}

func ChangeName(name string) Alias {
	return changeName{Name: name}
}

func (c changeName) ApplyToColumn(column *ycmodel.Column) {
	column.Name = c.Name
}

func (c changeName) ApplyToTable(table *ycmodel.Table) {
	table.Alias = c.Name
}

type changeColumn struct {
	UnimplementedAlias
	column *ycmodel.Column
}

func ChangeColumn(column *ycmodel.Column) Alias {
	return changeColumn{column: column}
}

func (r changeColumn) ApplyToColumn(column *ycmodel.Column) {
	*column = *r.column
}
