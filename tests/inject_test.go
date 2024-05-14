package tests

import (
	"testing"

	"github.com/demoManito/inject"
	"github.com/demoManito/inject/injector"
	"github.com/demoManito/inject/tests/testdata/service"
)

type handler struct {
	Service *service.Service `inject:"service"`
}

func TestInject_Injector(t *testing.T) {
	h := &handler{}
	inject.New(injector.New()).Inject(h)
	if h == nil {
		t.Fatal("h is nil")
	}
	if h.Service == nil {
		t.Fatal("h.Service is nil")
	}
	if h.Service.UserDao == nil {
		t.Fatal("h.Service.UserDao is nil")
	}
	if h.Service.UserDao.DB == nil {
		t.Errorf("h.Service.UserDao.DB is nil")
	}
}
