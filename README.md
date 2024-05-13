## Inject (golang dependency injection)

### About Inject

`inject` is a simple dependency injection library for Go. based on tag `inject` to inject dependencies into struct fields.

### Quick Start

```go
package main

import (
    "database/sql"

    "github.com/demoManito/inject"
    "github.com/demoManito/injector"
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

func init() {
	inject.
}

func main() {
	hanlder := &Handler{}
	inject.New(injector.New()).Inject(handler)
}
```
