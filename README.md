## Inject (golang dependency injection)

### About Inject

`inject` is a simple dependency injection library for Go. based on tag `inject` to inject dependencies into struct fields.

### Quick Start

- service/service.go
```go
package service

import (
    "github.com/demoManito/inject"
    
    "github.com/xxx/dao"
)

type Service struct {
    dao *Dao `inject:"dao"`
}

func init() {
    inject.Register(func(injector inject.Injector) error {
        injector.Register("service", &Service{})
        return nil
    })
}
```

- dao/dao.go
```go
package dao

import (
    "database/sql"

    "github.com/demoManito/inject"
)

type Dao struct {
    db *sql.DB `inject:"db"`
}

func init() {
    inject.Register(func(injector inject.Injector) error {
        injector.Register("dao", &Dao{db: &sql.DB{}})
        return nil
    })
}
```

- main.go
```go
package main

import (
    "database/sql"

    "github.com/demoManito/inject"
    "github.com/demoManito/inject/injector"

    _ "github.com/xxx/dao"
    _ "github.com/xxx/service"
)

type Handler struct {
    service *Service `inject:"service"`
}

func main() {
    hanlder := &Handler{}
    inject.New(injector.New()).Inject(handler)
    
    // eg: use handler
    handler.service.dao.db.Ping()
}
```
