package config

import (
	"flag"
	"github.com/abbi-gaurav/kyma-connectors/components/github/event-gw/internal/logger"
)

type Opts struct {
	LogRequest       *bool
	Passwd           *string
	UserName         *string
	BasicAuthEnabled *bool
	AppName          *string
	EventPublishURL  *string
	SecretKey        *string
}

var GlobalConfig *Opts

func ParseFlags() {
	logRequest := flag.Bool("verbose", false, "log each incoming event request")
	baEnabled := flag.Bool("basic-auth-enabled", false, "Enable basic Auth")
	userName := flag.String("username", "", "Basic Auth username")
	password := flag.String("password", "", "Basic Auth Password")
	appName := flag.String("app-name", "github", "Application Name")
	eventPublishURL := flag.String("event-publish-url", "http://event-publish-service.kyma-system.svc.cluster.local:8080/v1/events", "URL to forward incoming events to Kyma Eventing")
	secretKey := flag.String("secret-key", "", "secret key used to verify event payload signature")
	flag.Parse()

	GlobalConfig = &Opts{
		LogRequest:       logRequest,
		BasicAuthEnabled: baEnabled,
		UserName:         userName,
		Passwd:           password,
		AppName:          appName,
		EventPublishURL:  eventPublishURL,
		SecretKey:        secretKey,
	}

	if *GlobalConfig.BasicAuthEnabled && (*GlobalConfig.UserName == "" || *GlobalConfig.Passwd == "") {
		logger.Logger.Panicw("invalid configuration - Missing credentials", "config", GlobalConfig)
	}

	if *GlobalConfig.AppName == "" {
		logger.Logger.Panicw("Invalid configuration - Missing APP Name", "config", GlobalConfig)
	}

	logger.Logger.Infow("App config", "config", GlobalConfig)
}
