package cmd

import (
	"os"
	"path/filepath"
)

func GetHooksDir() string {
	pwd, _ := os.Getwd()

	return filepath.Join(pwd, ".git", "hooks")
}
