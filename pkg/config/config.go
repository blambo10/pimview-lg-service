package config

import (
	"github.com/gobuffalo/envy"
	l "github.com/sirupsen/logrus"
	"strconv"
)

type Config struct {
	logging Logging
	mqtt    MQTT
	webos   WebOS
}

type MQTT struct {
	Address  string
	Port     int
	Username string
	Password string
	ClientID string
}

type WebOS struct {
	Host string
}

type Logging struct {
	Level int
}

func GetLogEnvs() Logging {
	level, err := strconv.Atoi(envy.Get("LOGGING_LEVEL", "4"))

	if err != nil {
		level = 4
	}

	if level > len(l.AllLevels)-1 || level < 0 {
		level = 4
	}

	return Logging{
		Level: level,
	}
}

func NewConfig() *Config {
	envy.Load()

	port, err := strconv.Atoi(envy.Get("MQTT_BROKER_PORT", ""))

	if err != nil {
		port = 1883
	}

	return &Config{
		logging: GetLogEnvs(),
		mqtt: MQTT{
			Address:  envy.Get("MQTT_BROKER_ADDRESS", "localhost"),
			Port:     port,
			Username: envy.Get("MQTT_BROKER_USERNAME", ""),
			Password: envy.Get("MQTT_BROKER_PASSWORD", ""),
			ClientID: envy.Get("MQTT_BROKER_CLIENTID", "pimview-lg"),
		},
		webos: WebOS{
			Host: envy.Get("WEB_OS_ADDRESS", ""),
		},
	}
}

func GetLogger() Logging {
	return NewConfig().logging
}

func GetMQTT() MQTT {
	return NewConfig().mqtt
}

func GetWebOS() WebOS {
	return NewConfig().webos
}
