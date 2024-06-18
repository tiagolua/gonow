package main

import (
	"errors"

	"github.com/tiagolua/gonow.git/config"
	"github.com/tiagolua/gonow.git/router"
)

func main() {
	// initialize config
    err error := config.Init()	
	if err != nil{
		panic(error)
	}

	//initialize router
	router.Initializer()
}
