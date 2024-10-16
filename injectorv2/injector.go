package injectorv2

import (
	"reflect"
	"sync"
	"unsafe"
)

const (
	tag = "inject"
)

var (
	once = &sync.Once{}

	injector *Injector
)

// Injector is a struct that stores objects.
type Injector struct {
	objs map[string]any
}

// New singleton pattern returns a new Injector instance.
func New() *Injector {
	once.Do(func() {
		injector = &Injector{
			objs: make(map[string]any, 20),
		}
	})
	return injector
}

// Register the object to the Injector instance.
func (i *Injector) Register(key string, value any) {
	i.objs[key] = value
}

// Load returns the object from the Injector instance.
func (i *Injector) Load(key string) (value any, ok bool) {
	val, ok := i.objs[key]
	return val, ok
}

// Delete deletes the object from the Injector instance.
func (i *Injector) Delete(key string) {
	delete(i.objs, key)
}

// Range calls f sequentially for each key and value present in the Injector.
func (i *Injector) Range(f func(key, value any) bool) {
	for key, value := range i.objs {
		_ = f(key, value)
	}
}

// Inject injects the object to the value field.
func (i *Injector) Inject(val any) {
	value := reflect.ValueOf(val)
	for {
		if value.Kind() == reflect.Ptr || value.Kind() == reflect.Interface {
			value = value.Elem()
		} else {
			break
		}
	}
	if value.Kind() != reflect.Struct {
		return
	}
	for j := 0; j < value.NumField(); j++ {
		injectKey := value.Type().Field(j).Tag.Get(tag)
		injectVal, ok := i.Load(injectKey)
		if !ok {
			continue
		}
		field := value.Field(j)
		if field.CanSet() {
			field.Set(reflect.ValueOf(injectVal))
		} else {
			field = reflect.NewAt(field.Type(), unsafe.Pointer(field.UnsafeAddr())).Elem()
			field.Set(reflect.ValueOf(injectVal))
		}
	}
}
