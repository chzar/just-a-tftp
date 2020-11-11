package main

import (
	"log"
	"os"

	"github.com/kardianos/service"
)

var logger service.Logger

type program struct{}

func (p *program) Start(s service.Service) error {
	// Start should not block. Do the actual work async.
	go p.run()
	return nil
}
func (p *program) run() {
	config, err := LoadConfig()
	if err != nil {
		logger.Errorf("config: %v\n", err)
		os.Exit(1)
	}

	logger.Info("Starting Server...")
	err = BuildTftpServer(config.Directory, config.Readonly).ListenAndServe(config.ConnectionString) // blocks until s.Shutdown() is called
	if err != nil {
		logger.Errorf("server: %v\n", err)
	}
}
func (p *program) Stop(s service.Service) error {
	// Stop should not block. Return with a few seconds.
	return nil
}

func main() {
	svcConfig := &service.Config{
		Name:        "justatftpd",
		DisplayName: "justatftpd",
		Description: "justatftpd",
	}

	prg := &program{}
	s, err := service.New(prg, svcConfig)
	if err != nil {
		log.Fatal(err)
	}
	logger, err = s.Logger(nil)
	if err != nil {
		log.Fatal(err)
	}
	err = s.Run()
	if err != nil {
		logger.Error(err)
	}
}
