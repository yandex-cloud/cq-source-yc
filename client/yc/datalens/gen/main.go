// Command gen generates Go models from the vendored DataLens OpenAPI spec
// (https://api.datalens.tech/json/).
//
// The published spec is OpenAPI 3.1, which oapi-codegen does not support yet
// (https://github.com/oapi-codegen/oapi-codegen/issues/373), so the spec is
// downgraded to 3.0 first:
//   - `"type": [T, "null"]` becomes `"type": T` + `"nullable": true`
//   - `const` becomes a single-value `enum`
//   - `prefixItems` (tuples) is dropped, leaving untyped array items
//
// The downgraded spec is then fed to oapi-codegen (a go.mod tool) configured
// by oapi-codegen.yaml. To refresh models run `make gen-datalens` (downloads the
// spec too) or `go generate ./...`.
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

const specFile = "openapi.json"

func main() {
	raw, err := os.ReadFile(specFile)
	if err != nil {
		log.Fatalf("read spec: %v", err)
	}

	downgraded, err := downgradeTo30(raw)
	if err != nil {
		log.Fatalf("downgrade spec to OpenAPI 3.0: %v", err)
	}

	tmp := filepath.Join(os.TempDir(), "datalens-openapi-3.0.json")
	if err := os.WriteFile(tmp, downgraded, 0o644); err != nil {
		log.Fatalf("write downgraded spec: %v", err)
	}
	defer func() { _ = os.Remove(tmp) }()

	cmd := exec.Command("go", "tool", "oapi-codegen", "--config", "oapi-codegen.yaml", tmp)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		log.Fatalf("oapi-codegen: %v", err)
	}
}

func downgradeTo30(raw []byte) ([]byte, error) {
	var spec map[string]any
	if err := json.Unmarshal(raw, &spec); err != nil {
		return nil, fmt.Errorf("parse spec: %w", err)
	}

	walk(spec)
	spec["openapi"] = "3.0.3"

	return json.Marshal(spec)
}

func walk(node any) {
	switch n := node.(type) {
	case map[string]any:
		if types, ok := n["type"].([]any); ok {
			nullable := false
			var rest []any
			for _, t := range types {
				if t == "null" {
					nullable = true
				} else {
					rest = append(rest, t)
				}
			}
			if len(rest) > 0 {
				n["type"] = rest[0]
			} else {
				delete(n, "type")
			}
			if nullable {
				n["nullable"] = true
			}
		}
		if c, ok := n["const"]; ok {
			delete(n, "const")
			n["enum"] = []any{c}
		}
		if _, ok := n["prefixItems"]; ok {
			delete(n, "prefixItems")
			if _, ok := n["items"]; !ok {
				n["items"] = map[string]any{}
			}
		}
		for _, v := range n {
			walk(v)
		}
	case []any:
		for _, v := range n {
			walk(v)
		}
	}
}
