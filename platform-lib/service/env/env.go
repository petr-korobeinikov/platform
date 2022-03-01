package env

import (
	"bufio"
	"fmt"
	"os"
	"sync"
)

const File = ".platform/env/.env"

func (r *environmentRegistry) Register(name, value string) *environmentRegistry {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.elts[name] = value

	return r
}

func (r *environmentRegistry) RegisterMany(env map[string]string) *environmentRegistry {
	r.mu.Lock()
	defer r.mu.Unlock()

	for k, v := range env {
		r.elts[k] = v
	}

	return r
}

func (r *environmentRegistry) All() map[string]string {
	return r.elts
}

func (r *environmentRegistry) Clear() *environmentRegistry {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.elts = make(map[string]string)

	return r
}

func Registry() *environmentRegistry {
	once.Do(func() {
		instance = &environmentRegistry{}
		instance.elts = make(map[string]string)
	})

	return instance
}

func WriteEnvFile() error {
	f, err := os.Create(File)
	if err != nil {
		return err
	}

	w := bufio.NewWriter(f)
	for k, v := range Registry().All() {
		_, err = w.WriteString(fmt.Sprintf("%s=%s\n", k, v))
		if err != nil {
			return err
		}
	}

	err = w.Flush()
	if err != nil {
		return err
	}

	return nil
}

type (
	environmentRegistry struct {
		mu   sync.Mutex
		elts map[string]string
	}
)

var (
	instance *environmentRegistry
	once     sync.Once
)
