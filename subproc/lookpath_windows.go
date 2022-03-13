package subproc

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/goark/errs"
)

// LookPath searches for an executable named file in the
// directories named by the PATH environment variable.
// If file contains a slash, it is tried directly and the PATH is not consulted.
// LookPath also uses PATHEXT environment variable to match
// a suitable candidate.
// The result may be an absolute path or a path relative to the current directory.
func LookPath(file string) (string, error) {
	var exts []string
	x := os.Getenv(`PATHEXT`)
	if x != "" {
		for _, e := range strings.Split(strings.ToLower(x), `;`) {
			if e == "" {
				continue
			}
			if e[0] != '.' {
				e = "." + e
			}
			exts = append(exts, e)
		}
	} else {
		exts = []string{".com", ".exe", ".bat", ".cmd"}
	}

	if strings.ContainsAny(file, `:\/`) {
		if f, err := findExecutable(file, exts); err == nil {
			return f, nil
		} else {
			return "", errs.Wrap(err, errs.WithContext("file", file))
		}
	}
	// if f, err := findExecutable(filepath.Join(".", file), exts); err == nil {
	// 	return f, nil
	// }
	path := os.Getenv("path")
	for _, dir := range filepath.SplitList(path) {
		if f, err := findExecutable(filepath.Join(dir, file), exts); err == nil {
			return f, nil
		}
	}
	return "", errs.Wrap(exec.ErrNotFound, errs.WithContext("file", file))
}

func chkStat(file string) error {
	d, err := os.Stat(file)
	if err != nil {
		return err
	}
	if d.IsDir() {
		return os.ErrPermission
	}
	return nil
}

func hasExt(file string) bool {
	i := strings.LastIndex(file, ".")
	if i < 0 {
		return false
	}
	return strings.LastIndexAny(file, `:\/`) < i
}

func findExecutable(file string, exts []string) (string, error) {
	if len(exts) == 0 {
		return file, chkStat(file)
	}
	if hasExt(file) {
		if chkStat(file) == nil {
			return file, nil
		}
	}
	for _, e := range exts {
		if f := file + e; chkStat(f) == nil {
			return f, nil
		}
	}
	return "", os.ErrNotExist
}

/**
 * These functions are fork from "os/exec" standard package
 * Go standers packages licensed under MIT license
 */
