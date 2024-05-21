## Inject (golang dependency injection)

### About Inject

`inject` is a simple dependency injection library for Go. based on tag `inject` to inject dependencies into struct fields.

### Quick Start

- dao/dao.go
```go
package dao

import (
	"database/sql"

	"github.com/demoManito/inject"
)

func init() {
	inject.Register(func(injector inject.Injector) error {
		injector.Register("db", &sql.DB{})
		injector.Register("dao", &Dao{})
		return nil
	})
}

type Dao struct {
	DB *sql.DB `inject:"db"`
}
```

- service/service.go
```go
package service

import (
	"github.com/demoManito/inject"

	"github.com/xxx/xxx/dao"
)

func init() {
	inject.Register(func(injector inject.Injector) error {
		injector.Register("service", &Service{})
		return nil
	})
}

type Service struct {
	Dao *dao.Dao `inject:"dao"`
}
```

- main.go

```go
package main

import (
	"database/sql"

	"github.com/demoManito/inject"
	"github.com/demoManito/inject/injector"

	"github.com/xxx/xxx/service"
)

type Handler struct {
	Service *service.Service `inject:"service"`
}

func main() {
	hanlder := &Handler{}
	inject.New(injector.New()).Inject(handler)

	// eg: use handler, private fields can also be injected
	handler.Service.Dao.DB.Ping()
}
```

### Easy Example
[Easy Example](https://github.com/demoManito/inject/tree/main/tests)
