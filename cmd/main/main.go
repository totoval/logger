package main

import (
	"github.com/totoval/framework/helpers/toto"
	"github.com/totoval/logger/pkg"
	"github.com/totoval/logger/pkg/facade"
	"github.com/totoval/logger/pkg/structs"
)

//@todo register into global vars
var log facade.Logger

func main() {
	// use driver then config
	l := &pkg.Log{}
	if err := l.Use(pkg.DriverLogrus).Config(map[string]interface{}{
		"level": structs.LevelTrace,
	}); err != nil {
		panic(err)
	}
	// get the facade
	log = l.Component().(facade.Logger)

	log.Info("test", toto.V{"haha": 123}, toto.V{"1": 2})
}
