package service

import (
	"github.com/demoManito/inject"
	"github.com/demoManito/inject/tests/testdata/dao"
)

func init() {
	inject.Register(func(injector inject.Injector) error {
		injector.Register("service", &Service{})
		return nil
	})
}

type Service struct {
	UserDao *dao.UserDao `inject:"user_dao"`
}
