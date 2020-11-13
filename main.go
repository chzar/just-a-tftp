package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/kardianos/service"
)

var logger service.Logger
var srvConfig *config
var svcConfig *service.Config

type program struct{}

func (p *program) Start(s service.Service) error {
	// Start should not block. Do the actual work async.
	go p.run()
	return nil
}
func (p *program) run() {

	pwd, _ := os.Getwd()
	logger.Infof("Starting Server in %s...", pwd)

	err := BuildTftpServer(srvConfig.Directory, srvConfig.Readonly).ListenAndServe(srvConfig.ConnectionString) // blocks until s.Shutdown() is called
	if err != nil {
		logger.Errorf("Server: %v\n", err)
	}
}
func (p *program) Stop(s service.Service) error {
	// Stop should not block. Return with a few seconds.
	return nil
}

func main() {
	srvConfig = argparse()
	svcConfig = &service.Config{
		Name:             "justatftpd",
		DisplayName:      "justatftpd",
		Description:      "justatftpd",
		Arguments:        strings.Split(fmt.Sprintf("--ro=%s --dir=%s", strconv.FormatBool(srvConfig.Readonly), srvConfig.Directory), " "),
		WorkingDirectory: srvConfig.Directory,
		Executable:       os.Args[0],
	}

	prg := &program{}
	s, err := service.New(prg, svcConfig)
	if err != nil {
		log.Fatal(err)
	}

	for _, command := range []string{"install", "uninstall"} {
		for _, arg := range os.Args {
			if command == strings.ToLower(arg) {
				err = service.Control(s, command)
				if err != nil {
					log.Fatal(err)
				}
				return
			}
		}
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
