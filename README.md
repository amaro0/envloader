# envloader

Simple go utility for env config loading and validation

Get with:
```bash
go get github.com/amaro0/envloader
```

### How to use
```go 
package main

import (
	"fmt"
	"github.com/amaro0/envloader"
)

type someConf struct {
	Port    string `env:"PORT" envDefault:"3002" validate:"numeric"`
	GinMode string `env:"GIN_MODE" envDefault:"debug" validate:"oneof=debug release"`
}

func main() {
	err, conf := envloader.Load(someConf{})

	if err != nil {
		panic(err)
	}

	someConfInstance := conf.(*someConf)

	fmt.Println(someConfInstance)
}
```

You can tag your struct with options from both [go-playground/validator](https://github.com/go-playground/validator) and [caarlos0/env](https://github.com/caarlos0/env)
