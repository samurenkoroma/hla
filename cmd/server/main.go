package main

import (
	"github.com/samurenkoroma/lha/app"
	"github.com/samurenkoroma/lha/internal/config"
)

func main() {
	app.NewApplication(config.MustLoad()).
		Configure().
		Run()
}
