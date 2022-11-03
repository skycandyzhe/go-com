package server

import (
	"time"

	"github.com/kardianos/service"
	"github.com/skycandyzhe/go-com/logger"
)

type Program struct {
	IsStop bool
}

func (p *Program) Start(s service.Service) error {
	// Start should not block. Do the actual work async.
	p.IsStop = false
	go p.run()
	return nil
}
func (p *Program) run() {
	logger.SetupLogger(logger.GetDefaultLogger())

	for p.IsStop {
		time.Sleep(time.Minute)
	}
}
func (p *Program) Stop(s service.Service) error {
	// Stop should not block. Return with a few seconds.
	p.IsStop = true
	return nil
}

/*
	var svcConfig = &service.Config{
		Name:        "das-agent",
		DisplayName: "das-agent",
		Description: "das-agent",
		Dependencies: []string{
			"Requires=network.target",
			"After=network-online.target syslog.target"},
		Arguments: []string{"crontab", "-crontab=true"},
	}
*/
func ControlServerInstall(prg *Program, config *service.Config, log logger.MyLoggerInterface) {

	s, err := service.New(prg, config)
	if err != nil {
		log.Errorf("error create  server %v", config)
		logger.Logger.Error(err)
		return
	}
	err = service.Control(s, "install")
	if err != nil {
		log.Errorf("%s install server failure  %v", &config.Name)

	} else {
		log.Infof("%s install server success")

	}
}
func ControlUninstall(prg *Program, config *service.Config, log logger.MyLoggerInterface) {

	s, err := service.New(prg, config)
	if err != nil {
		logger.Logger.Error(err)
		return
	}
	err = service.Control(s, "uninstall")
	if err != nil {
		log.Errorf("%s uninstall server failure  %v", &config.Name)
	} else {
		log.Infof("%s uninstall server success")
	}

}

func ControlServerStart(prg *Program, config *service.Config, log logger.MyLoggerInterface) {

	s, err := service.New(prg, config)
	if err != nil {
		log.Errorf("error open  server %v", config)
		logger.Logger.Error(err)
		return
	}
	err = service.Control(s, "start")
	if err != nil {
		log.Errorf("%s start server failure  %v", &config.Name)

	} else {
		log.Infof("%s start server success")

	}
}
func ControlServerStop(prg *Program, config *service.Config, log logger.MyLoggerInterface) {

	s, err := service.New(prg, config)
	if err != nil {
		log.Errorf("error open  server %v", config)
		logger.Logger.Error(err)
		return
	}
	err = service.Control(s, "stop")
	if err != nil {
		log.Errorf("%s stop server failure  %v", &config.Name)

	} else {
		log.Infof("%s stop server success")

	}
}
