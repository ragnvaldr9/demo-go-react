package vite

import "fmt"

const INVALID_MANIFEST_STRUCT = "Invalid manifest structure. Manifest should be Record<name, chunk>. See https://vitejs.dev/guide/backend-integration.html"
const MULTIPLE_ENTRY_ERR = "Found muiltiple entry points. Multipage apps currently are not supported"

type ReadOnlyErr struct {
	field string
}

func (e ReadOnlyErr) Error() string {
	return fmt.Sprintf("Field %v is readonly", e.field)
}

func createReadOnlyError(f string) ReadOnlyErr {
	return ReadOnlyErr{field: f}
}