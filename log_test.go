package test

import (
	"fmt"
	"testing"

	"github.com/skycandyzhe/go-com/config"
	"github.com/skycandyzhe/go-com/logger"
)

func TestParseRule(t *testing.T) {
	fmt.Println(config.GetDefaultConf())
	logger.GetDefaultLogger().Info("info test")
	logger.GetDefaultLogger().Debug("debug test")
	logger.GetDefaultLogger().Error("debug test")
}
