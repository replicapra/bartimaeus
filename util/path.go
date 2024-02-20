package util

import "path/filepath"

func GetAbsPath(relPath string) (absPath string) {
	absPath, err := filepath.Abs(relPath)
	CheckErr(err)
	return
}
