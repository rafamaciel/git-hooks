package hooks

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var HOOKS = []string{
	"pre-commit",
	"pre-push.sample",
}

func setup(temp string) {
	src := filepath.Join(temp, "hooks")
	os.Mkdir(src, 0755)
	for _, hook := range HOOKS {
		os.Create(filepath.Join(temp, hook))
	}
}

func teardown(temp string) {
	os.RemoveAll(temp)
}

func TestHook(t *testing.T) {
	temp := t.TempDir()

	setup(temp)

	for _, file := range HOOKS {
		hook, err := NewHook(filepath.Join(temp, file))
		assert.NoError(t, err)
		ext := filepath.Ext(file)
		if ext == "" {
			assert.Equal(t, file, hook.Kind)
			assert.Equal(t, ENABLED, hook.Status)

			hook.Disable()

			assert.Equal(t, DISABLED, hook.Status)

		} else {
			kind := strings.TrimSuffix(file, ext)
			assert.Equal(t, kind, hook.Kind)
			assert.Equal(t, DISABLED, hook.Status)
		}
	}

	defer teardown(temp)
}
