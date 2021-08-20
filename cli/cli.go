package cli

import (
	"flag"
	"github.com/fsnotify/fsnotify"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func Init() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("/etc/elastic-apm-operator/")
	viper.AddConfigPath("$HOME/.elastic-apm-operator")
	viper.AddConfigPath(".")
	viper.WatchConfig()
	viper.OnConfigChange(func(event fsnotify.Event) {
		log.Infof("Config file changed: %s", event.Name)
	})
	err := viper.ReadInConfig()
	if err != nil {
		log.Debugf("Config file not found or not readable: %s \n", err)
	}
}

// Parse function processes the CLI input and makes them available to the application
func Parse(port *int, tlsCert *string, tlsKey *string, loglevel *log.Level) {
	flag.Int("port", 8443, "port where to listen for requests")
	flag.String("tlsCert", "~/.elastic-apm-operator/tls.crt", "TLS certificate to use for HTTPS")
	flag.String("tlsKey", "~/.elastic-apm-operator/tls.key", "TLS key to use for HTTPS")
	flag.String("log-level", "INFO", "log level of the application")

	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.Parse()

	err := viper.BindPFlags(pflag.CommandLine)
	if err != nil {
		log.Error("Error while processing command line options")
		return
	}

	*port = viper.GetInt("port")
	*tlsCert = viper.GetString("tlsCert")
	*tlsKey = viper.GetString("tlsKey")
	*loglevel = parseLogLevel(viper.GetString("log-level"))
}

// parseLogLevel parses a supplied log level in string format, and defaults to INFO in case it fails
func parseLogLevel(logLevel string) log.Level {
	if level, err := log.ParseLevel(logLevel); err == nil {
		return level
	} else {
		return log.InfoLevel
	}
}
