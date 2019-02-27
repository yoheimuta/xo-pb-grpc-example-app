package exppath_test

import (
	"path/filepath"
	"runtime"
)

// ProjectRootPath creates a path to project root.
func ProjectRootPath(elem ...string) string {
	_, this, _, ok := runtime.Caller(0)
	if !ok {
		return ""
	}
	root := filepath.Dir(filepath.Dir(filepath.Dir(this)))
	if len(elem) == 0 {
		return root
	}
	return filepath.Join(append([]string{root}, elem...)...)
}

// ScriptRootPath creates a path to _script.
func ScriptRootPath(elem ...string) string {
	return ProjectRootPath(append([]string{"_script"}, elem...)...)
}
