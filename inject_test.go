package inject

import (
	"testing"

	"github.com/demoManito/inject/injector"
)

type Handler struct {
	service *Service `inject:"service"`
}

type Service struct {
	user *User `inject:"user"`
}

type User struct {
	Name string `inject:"name"`
}

func TestNew(t *testing.T) {
	Register(func(injector Injector) error {
		injector.Register("service", &Service{})
		return nil
	})
	Register(func(injector Injector) error {
		injector.Register("user", &User{})
		return nil
	})
	Register(func(injector Injector) error {
		injector.Register("name", "demoManito")
		return nil
	})
	if len(funcs) != 3 {
		t.Fatalf("funcs length is %d, not 3", len(funcs))
	}

	handler := &Handler{}
	New(injector.NewInjector()).Inject(handler)
	if handler == nil {
		t.Fatal("handler is nil")
	}
	if handler.service == nil {
		t.Fatal("handler.service is nil")
	}
	if handler.service.user == nil {
		t.Fatal("handler.service.user is nil")
	}
	if handler.service.user.Name != "demoManito" {
		t.Errorf("handler.service.user.Name is %s, not demoManito", handler.service.user.Name)
	}
}
