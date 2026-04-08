package godash

import (
	"path/filepath"
	"strings"
)

// FixWindowPath return Window-Sytle Path, eg: C://a/b/c instead of /C/a/b/c
func FixWindowsPath(p string) string {
	if strings.HasPrefix(p, "/") && len(p) >= 3 && p[2] == '/' {
		drive := strings.ToUpper(string(p[1]))
		rest := p[2:]
		return drive + ":/" + rest
	}
	return filepath.Clean(p)
}
