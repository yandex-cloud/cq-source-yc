package tools

import (
	"fmt"

	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	"github.com/jhump/protoreflect/desc"
	"github.com/jhump/protoreflect/desc/protoparse"
)

type Option interface {
	Apply(tg *collapsedOptions) error
}

type withProtoFile struct {
	messageName string
	protoFile   string
	paths       []string
}

func (w withProtoFile) Apply(co *collapsedOptions) error {
	parser := protoparse.Parser{IncludeSourceCodeInfo: true, ImportPaths: w.paths}
	protoFiles, err := parser.ParseFiles(w.protoFile)
	if err != nil {
		return err
	}
	protoFile := protoFiles[0]
	co.message = protoFile.FindMessage(protoFile.GetPackage() + "." + w.messageName)
	if co.message == nil {
		return fmt.Errorf("message %v not found", w.messageName)
	}
	return nil
}

func WithProtoFile(messageName string, protoFile string, paths ...string) Option {
	return withProtoFile{messageName: messageName, protoFile: protoFile, paths: paths}
}

type withMessage struct {
	message *desc.MessageDescriptor
}

func (w withMessage) Apply(co *collapsedOptions) error {
	co.message = w.message
	return nil
}

func WithMessage(message *desc.MessageDescriptor) Option {
	return withMessage{message: message}
}

type withTableName struct {
	tableName string
}

func (w withTableName) Apply(co *collapsedOptions) error {
	co.tableName = w.tableName
	return nil
}

func WithTableName(tableName string) Option {
	return withTableName{tableName: tableName}
}

type withDefaultColumns struct {
	defaultColumns map[string]schema.Column
}

func (w withDefaultColumns) Apply(co *collapsedOptions) error {
	co.defaultColumns = w.defaultColumns
	return nil
}

func WithDefaultColumns(defaultCols map[string]schema.Column) Option {
	return withDefaultColumns{defaultColumns: defaultCols}
}

type withIgnoredColumns struct {
	ignoredColumns []string
}

func (w withIgnoredColumns) Apply(co *collapsedOptions) error {
	ignoreColsMap := make(map[string]struct{})
	for _, col := range w.ignoredColumns {
		ignoreColsMap[col] = struct{}{}
	}
	co.ignoredFields = ignoreColsMap
	return nil
}

func WithIgnoredColumns(ignoreCols ...string) Option {
	return withIgnoredColumns{ignoredColumns: ignoreCols}
}

type withResolver struct {
	resolver schema.TableResolver
}

func (w withResolver) Apply(co *collapsedOptions) error {
	co.resolver = w.resolver
	return nil
}

func WithResolver(resolver schema.TableResolver) Option {
	return withResolver{resolver: resolver}
}

type withMultiplex struct {
	multipex func(meta schema.ClientMeta) []schema.ClientMeta
}

func (w withMultiplex) Apply(co *collapsedOptions) error {
	co.multiplex = w.multipex
	return nil
}

func WithMultiplex(multiplex func(meta schema.ClientMeta) []schema.ClientMeta) Option {
	return withMultiplex{multipex: multiplex}
}
