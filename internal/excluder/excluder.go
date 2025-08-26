package excluder

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type Excluder interface {
	IsExcluded(path string) bool
}

type excluder struct {
	excludePaths map[string]struct{}
}

// getDefaultExclusions returns the default exclusions
// This is now managed by SettingsService and should not be used directly
func getDefaultExclusions() []string {
	return []string{
		// Windows system directories
		`c:\windows`, `c:\program files`, `c:\program files (x86)`, `c:\programdata`,
		`c:\users\default`, `c:\users\public`, `$recycle.bin`, `c:\intel`,
		`c:\recovery`, `c:\perflogs`, `c:\windows.old`, `c:\system.sav`,

		// Windows-specific folders
		`system volume information`, `c:\pagefile.sys`, `c:\hiberfil.sys`,

		// Program cache directories
		`c:\users\*\appdata\local`, "cachestorage", "cache", "zxcvbndata",

		// Development directories
		"node_modules", "vendor", "sdk", "npm-cache", `go\pkg`, `venv`,
		`discover-iq-temp`,

		// Hidden files and directories
		".*",

		// Temporary and log directories
		`c:\temp`, `c:\tmp`, `c:\logs`, `c:\debug`, `c:\inetpub\logs`,

		// Others
		`*.log`, `*.so`, `*.mo`, `*.pl`, `*.temp`, `*.bin`, `*.lock`, `*.pdb`, `*.exp`, `*.lib`,
		`*.dll`, `*.d`, `*.rlib`, `*.rmeta`, `*.o`, `license*`, `readme*`, `*.ie5`, `credits.txt`,

		// macOS system directories
		`/Applications`, `/System`, `/Library`, `/private`, `/bin`, `/sbin`, `/usr`, `/dev`, `/etc`, `/tmp`, `/var`, `/opt`,
		`/Users/*/Library/*`,
		`/Users/Shared/*`,
		`/Users/*/Pictures/Photos Library.photoslibrary`,
		`/Users/*/Pictures/Photo Booth Library`,
		`/Users/*/Music/Music Library.musiclibrary`,
		`/Users/*/Movies/Final Cut Library.fcpbundle`,
		`/Users/*/Library/Containers/*`,
		`/Users/*/Library/Group Containers/*`,
		`.nofollow`, `.resolve`, `.vol`,

		// Specific files and patterns
		`*.so`, `*.dylib`, `*.o`, `*.bin`, `*.lock`, `*.tmp`, `*.conf`,
	}
}

func normalizePaths(paths []string) []string {
	normalized := make([]string, len(paths))
	for i, path := range paths {
		normalized[i] = filepath.Clean(strings.ToLower(path))
	}
	return normalized
}

func loadCustomPathsFromConfig() ([]string, error) {
	workingDir, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	file, err := os.Open(filepath.Join(workingDir, "exclude_config.json"))
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var cfg struct {
		ExcludePaths []string `json:"excludePaths"`
	}
	err = json.NewDecoder(file).Decode(&cfg)
	if err != nil {
		return nil, err
	}

	return cfg.ExcludePaths, nil
}

func NewExcluder() Excluder {
	customPaths, err := loadCustomPathsFromConfig()
	if err != nil {
		customPaths = []string{}
	}

	var ignoreFolders []string = getDefaultExclusions()

	mergedPaths := normalizePaths(append(ignoreFolders, customPaths...))
	excludeMap := make(map[string]struct{}, len(mergedPaths))
	for _, path := range mergedPaths {
		excludeMap[path] = struct{}{}
	}
	return &excluder{
		excludePaths: excludeMap,
	}
}

func (e *excluder) IsExcluded(path string) bool {
	normalizedPath := filepath.Clean(strings.ToLower(path))
	basePath := filepath.Base(normalizedPath)
	fmt.Printf("[EXLUDER]Base Path: %s\n", basePath)
	if basePath == "." {
		return false
	}
	if strings.HasPrefix(basePath, ".") {
		return true
	}
	ext := strings.TrimPrefix(filepath.Ext(basePath), ".")
	if ext != "" && e.isNumeric(ext) {
		return true
	}

	// Check exact path match
	if _, exists := e.excludePaths[normalizedPath]; exists {
		return true
	}

	for rule := range e.excludePaths {
		if match, _ := filepath.Match(rule, filepath.Base(normalizedPath)); match {
			return true
		}
		if e.matchedWildcard(rule, normalizedPath) {
			return true
		}
	}
	return false
}

func (e *excluder) isNumeric(s string) bool {
	for _, r := range s {
		if r < '0' || r > '9' {
			return false
		}
	}
	return len(s) > 0
}

func (e *excluder) matchedWildcard(rule, path string) bool {
	parts := strings.Split(rule, string(os.PathSeparator))
	pathParts := strings.Split(path, string(os.PathSeparator))
	if len(pathParts) < len(parts) {
		return false
	}

	for i, part := range parts {
		if part == "*" {
			continue
		}
		if !strings.EqualFold(part, pathParts[i]) {
			return false
		}
	}
	return true
}
