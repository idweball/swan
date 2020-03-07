package app

import (
	"flag"
	"swan/internal/config"
	"swan/internal/engine"
	"swan/internal/storage"
	"swan/pkg/log"
)

//Run 运行ConfigMap
func Run() {
	filename := flag.String("cfg", "config.toml", "the configuration file")
	flag.Parse()

	cfg, err := config.CfgParser(*filename)
	if err != nil {
		log.Fatalln(err)
	}

	s, err := storage.New(cfg.Storage.Type, cfg.Storage.Config)
	if err != nil {
		log.Fatalln(err)
	}

	e, err := engine.New(s, cfg.Templates)
	if err != nil {
		log.Fatalln(err)
	}

	err = e.Run()
	if err != nil {
		log.Fatalln(err)
	}
}
