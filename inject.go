package inject

import (
	"sync"
)

var (
	lock  = &sync.Mutex{}
	funcs = make([]Func, 0, 20)
)

type (
	Func func(Injector) error
)

// Injector is the interface that New must implement.
type Injector interface {
	Register(key string, value any)
	Load(key string) (value any, ok bool)
	Delete(key string)
	Range(f func(key, value any) bool)
	Inject(val any)
}

// Register register the Func instance.
func Register(injectFunc Func) {
	lock.Lock()
	defer lock.Unlock()
	funcs = append(funcs, injectFunc)
}

// New returns the Injector instance. If the Injector instance is nil, it will be initialized.
func New(injector Injector) Injector {
	lock.Lock()
	for _, f := range funcs {
		err := f(injector)
		if err != nil {
			panic(err)
		}
	}
	lock.Unlock()

	injector.Range(func(_, value any) bool {
		injector.Inject(value) // inject the value field
		return true
	})

	return injector
}
