## inject (golang dependency injection)

### About

`inject` is a simple dependency injection library for Go. based on tag `inject` to inject dependencies into struct fields.

### Quick Start

```go
package main

import (
    "database/sql"
    "github.com/demoManito/inject"
)

type Handler struct {
	service *Service `inject:"service"`
}

type Service struct {
	dao *Dao `inject:"dao"`
}

type Dao struct {
    db *sql.DB `inject:"db"`
}

func main() {
	hanlder := &Handler{}
	inject.New(injector.New()).Inject(handler)
}
```
