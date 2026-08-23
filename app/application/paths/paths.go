package paths

import (
	"path/filepath"
	"runtime"
)

var (
	_, file, _, _ = runtime.Caller(0)
	Root = filepath.Join(filepath.Dir(file), "../../..")
)
