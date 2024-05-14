package dao

import (
	"database/sql"

	"github.com/demoManito/inject"
)

func init() {
	inject.Register(func(injector inject.Injector) error {
		injector.Register("db", &sql.DB{})
		injector.Register("user_dao", &UserDao{})
		return nil
	})
}

type UserDao struct {
	DB *sql.DB `inject:"db"`
}
