package hooks

import (
	"os"
	"path/filepath"
)

func LoadHooksFrom(src string) ([]Hook, error) {
	hooks := []Hook{}
	entries, err := os.ReadDir(src)

	for _, entry := range entries {
		if !entry.IsDir() {
			path := filepath.Join(src, entry.Name())
			hook, err := NewHook(path)
			if err == nil {
				hooks = append(hooks, hook)
			}
		}
	}

	if err != nil {
		return nil, err
	}

	return hooks, nil
}
