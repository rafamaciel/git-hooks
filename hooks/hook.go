package hooks

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

const (
	ENABLED  = 1
	DISABLED = 0
)

type Hook struct {
	Kind   string
	Status int
	path   string
}

func (this *Hook) Enable() {

	_, err := os.Stat(this.path)

	enabled := strings.Replace(this.path, ".sample", "", -1)
	fmt.Println(this.path)
	if err == nil {
		os.Rename(this.path, enabled)
		this.Status = ENABLED
	}

}

func (this *Hook) Disable() {
	disabled := strings.Join([]string{this.path, "sample"}, ".")

	_, err := os.Stat(this.path)

	if err == nil {
		os.Rename(this.path, disabled)
		this.Status = DISABLED
	}

}

func NewHook(path string) (Hook, error) {

	file, err := os.Stat(path)

	if err != nil {
		return Hook{}, err
	}

	kind := file.Name()

	ext := filepath.Ext(kind)

	if ext == "" {
		return Hook{
			Status: ENABLED,
			Kind:   kind,
			path:   path,
		}, nil
	}

	name := strings.TrimSuffix(kind, ext)

	return Hook{
		Status: DISABLED,
		Kind:   name,
		path:   path,
	}, nil

}
