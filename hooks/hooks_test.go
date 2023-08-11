package hooks

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func seed(path string) []string {

	hooks := []string{"pre-commit",
		"prepare-commit-msg",
		"commit-msg",
		"post-commit",
		"pre-rebase",
		"post-checkout",
		"post-merge",
		"pre-push",
		"pre-receive",
		"update",
		"post-receive",
		"post-update",
		"pre-auto-gc",
		"post-rewrite",
	}
	files := []string{}
	for _, hook := range hooks {
		file := filepath.Join(path, strings.Join([]string{hook, "sample"}, "."))
		os.Create(file)
		files = append(files, file)
	}
	return files
}

func TestHooks(t *testing.T) {

	temp := t.TempDir()

	src := filepath.Join(temp, "hooks")

	os.Mkdir(src, 0755)

	_ = seed(src)

	defer os.RemoveAll(src)

	hooks, err := LoadHooksFrom(src)
	assert.NoError(t, err)
	assert.NotEmpty(t, hooks)

}
