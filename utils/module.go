package utils

import (
	"os"
	"path/filepath"
	"strings"
)

func GetModules() []string {
	result := []string{}
	dir := "/etc/modulefiles/"
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		path = strings.ReplaceAll(path, dir, "")
		if info.IsDir() && len(path) > 0 {
			result = append(result, path)
		}
		if info.Mode().IsRegular() {
			matched, err := filepath.Match("**/.*", path)
			if err != nil {
				return err
			}
			if !matched {
				result = append(result, path)
			}
		}
		return nil
	})
	if err != nil {
		return []string{}
	}
	return result
}
