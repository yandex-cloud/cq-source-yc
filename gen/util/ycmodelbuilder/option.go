package ycmodelbuilder

type CollapsedOptions struct {
	Paths         []string
	IgnoredFields map[string]struct{}
	Aliases       map[string]Alias
}

type Option interface {
	Apply(co *CollapsedOptions)
}

func NewCollapsedOptions(opts []Option) CollapsedOptions {
	co := CollapsedOptions{
		Paths:         []string{"."},
		IgnoredFields: map[string]struct{}{},
		Aliases:       map[string]Alias{},
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
	co.Paths = w.paths
}

func WithProtoPaths(paths ...string) Option {
	return withProtoPaths{paths: paths}
}

type withIgnoredColumns struct {
	ignoredFields []string
}

func (w withIgnoredColumns) Apply(co *CollapsedOptions) {
	for _, ignoredColumn := range w.ignoredFields {
		co.IgnoredFields[ignoredColumn] = struct{}{}
	}
}

func WithIgnored(ignoredFields ...string) Option {
	return withIgnoredColumns{ignoredFields: ignoredFields}
}

type withAlias struct {
	path  string
	alias Alias
}

func (w withAlias) Apply(co *CollapsedOptions) {
	co.Aliases[w.path] = w.alias
}

func WithAlias(path string, alias Alias) Option {
	return withAlias{path: path, alias: alias}
}
