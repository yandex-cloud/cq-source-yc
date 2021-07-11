package tools

import (
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	"github.com/jhump/protoreflect/desc/protoparse"
)

type Option interface {
	Apply(tg *TableGenerator)
}

// proto files

type withProtoFile struct {
	protoFile string
}

func (w withProtoFile) Apply(tg *TableGenerator) {
	parser := protoparse.Parser{IncludeSourceCodeInfo: true}
	protoFiles, err := parser.ParseFiles(w.protoFile)
	if err != nil {
		return
	}
	tg.protoFile = protoFiles[0]
}

func WithProtoFile(protoFile string) Option {
	return withProtoFile{protoFile: protoFile}
}

// table name

type withTableName struct {
	tableName string
}

func (w withTableName) Apply(tg *TableGenerator) {
	tg.tableName = w.tableName
}

func WithTableName(tableName string) Option {
	return withTableName{tableName: tableName}
}

// default columns

type withDefaultColumns struct {
	defaultCols DefaultColumns
}

func (w withDefaultColumns) Apply(tg *TableGenerator) {
	tg.defaultCols = w.defaultCols
}

func WithDefaultColumns(defaultCols DefaultColumns) Option {
	return withDefaultColumns{defaultCols: defaultCols}
}

// ignored columns

type withIgnoredColumns struct {
	ignoreCols IgnoredColumns
}

func (w withIgnoredColumns) Apply(tg *TableGenerator) {
	ignoreColsMap := make(map[string]bool)
	for _, col := range w.ignoreCols {
		ignoreColsMap[col] = true
	}
	tg.ignoreCols = ignoreColsMap
}

func WithIgnoredColumns(ignoreCols ...string) Option {
	return withIgnoredColumns{ignoreCols: ignoreCols}
}

// with fetcher

type withFetcher struct {
	fetcher schema.TableResolver
}

func (w withFetcher) Apply(tg *TableGenerator) {
	tg.fetcher = w.fetcher
}

func WithFetcher(fetcher schema.TableResolver) Option {
	return withFetcher{fetcher: fetcher}
}
