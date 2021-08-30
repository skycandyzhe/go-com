package test

import (
	"fmt"
	"testing"

	"github.com/skycandyzhe/go-com/config"
	"github.com/skycandyzhe/go-com/logger"
)

func TestParseRule(t *testing.T) {
	fmt.Println(config.Conf)
	logger.Logger.Info("info test")
	logger.Logger.Debug("debug test")
	logger.Logger.Error("debug test")
}
