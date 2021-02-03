package main

import (
	"fmt"
	"github.com/amaro0/envloader"
)

type someConf struct {
	Port    string `env:"PORT" envDefault:"3002" validate:"numeric"`
	GinMode string `env:"GIN_MODE" envDefault:"debug" validate:"oneof=debug release`
}


func main(){
	err, conf := envloader.Load(someConf{})

	if err != nil{
		panic(err)
	}

	someConfInstance := conf.(*someConf)

	fmt.Println(someConfInstance)
}
