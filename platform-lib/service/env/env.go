package env

import "sync"

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

func Registry() *environmentRegistry {
	once.Do(func() {
		instance = &environmentRegistry{}
		instance.elts = make(map[string]string)
	})

	return instance
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
