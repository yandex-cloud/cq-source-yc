package util

import (
	"strings"

	"github.com/iancoleman/strcase"
)

const (
	ResourcesDir = "resources"
)

// ToFlat converts a string to flat case
func ToFlat(s string) string {
	return strings.ToLower(strcase.ToCamel(s))
}
