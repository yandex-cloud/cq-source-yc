package gen

type CollapsedOptions struct {
	paths           []string
	defaultColumns  map[string]*ColumnModel
	ignoredFields   map[string]struct{}
	relationAliases map[string]string
}

type Option interface {
	Apply(co *CollapsedOptions)
}

func NewCollapsedOptions(opts []Option) CollapsedOptions {
	co := CollapsedOptions{
		paths:           []string{"."},
		defaultColumns:  map[string]*ColumnModel{},
		ignoredFields:   map[string]struct{}{},
		relationAliases: map[string]string{},
	}
	for _, opt := range opts {
		opt.Apply(&co)
	}
	return co
}

type withProtoPaths struct {
	paths []string
}

func (w withProtoPaths) Apply(co *CollapsedOptions) {
	co.paths = w.paths
}

func WithProtoPaths(paths ...string) Option {
	return withProtoPaths{paths: paths}
}

type withDefaultColumns struct {
	defaultColumns map[string]*ColumnModel
}

func (w withDefaultColumns) Apply(co *CollapsedOptions) {
	co.defaultColumns = w.defaultColumns
}

func WithDefaultColumns(defaultColumns map[string]*ColumnModel) Option {
	return withDefaultColumns{defaultColumns: defaultColumns}
}

type withIgnoredColumns struct {
	ignoredFields []string
}

func (w withIgnoredColumns) Apply(co *CollapsedOptions) {
	for _, ignoredColumn := range w.ignoredFields {
		co.ignoredFields[ignoredColumn] = struct{}{}
	}
}

func WithIgnoredColumns(ignoredFields ...string) Option {
	return withIgnoredColumns{ignoredFields: ignoredFields}
}

type withAlias struct {
	path  string
	alias string
}

func (w withAlias) Apply(co *CollapsedOptions) {
	co.relationAliases[w.path] = w.alias
}

func WithAlias(path, alias string) Option {
	return withAlias{path: path, alias: alias}
}
