package config

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/spf13/viper"
)

type Configuration struct {
	Server struct {
		Port int
	}

	Mock struct {
		Port int
	}

	Database struct {
		Name     string
		User     string
		Password string
		Host     string
		Port     int
		Ssl      string
		Timezone string
	}
	Pgadmin struct {
		Mail string
		Pw   string
	}
	JwtToken struct {
		ExpireHours int
		SecretKey   string
	}
	Encryption struct {
		SecretKey string
	}
	Payment struct {
		RedirectUrl string
		Gateways    struct {
			Saderat struct {
				Terminal struct {
					Id string
				}
				Urls struct {
					Token   string
					Payment string
					Verify  string
				}
			}
		}
	}
	App struct {
		Reserved       string
		DebugMode      bool
		ImageLogo      string
		TicketFileName string
	}
	Auth struct {
		RequestLogoutHeader string
		TokenPrefix         string
	}
}

var Config *Configuration

func IsDebugMode() bool {
	Load()
	return Config.App.DebugMode
}

func Load() {
	if Config == nil {
		load()
	}
}

// ----------------------------------------------------------------------------------------
var (
	testModeEnabled bool
)

func init() {
	Load()
}

func load() {
	var err error
	var conf *Configuration
	appConfigFileName := os.Getenv("APP_CONFIG") // Defined in make file
	if appConfigFileName == "" {                 // If run out of makefile use default config
		appConfigFileName = "config-docker"
	}
	testModeEnabled = strings.HasSuffix(os.Args[0], ".test") // Run in test mode?
	if testModeEnabled {
		// The 'APP_ROOT' env-variable is defined in makefile and used by: 'make test'
		address := os.Getenv("APP_ROOT")
		if address == "" {
			panic(errors.New("the APP_ROOT environment variable is not defined, please use 'make test' instead of 'go test ./...'"))
		}
		conf, err = loadConfiguration(address, appConfigFileName, "yml")
	} else {
		conf, err = loadConfiguration(".", appConfigFileName, "yml")
	}
	if err != nil {
		log.Fatal("Loading viper config faild")
	}
	Config = conf
}

func loadConfiguration(configPath, configName, ConfigType string) (*Configuration, error) {
	var config *Configuration

	viper.AddConfigPath(configPath)
	viper.SetConfigName(configName)
	viper.SetConfigType(ConfigType)

	viper.AutomaticEnv()
	// first checks for SERVER.PORT
	viper.SetEnvKeyReplacer(strings.NewReplacer(`.`, `_`))
	// after above instruction, checks for SERVER_PORT

	err := viper.ReadInConfig()
	if err != nil {
		return nil, fmt.Errorf("could not read the config file: %v", err)
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		return nil, fmt.Errorf("could not unmarshal: %v", err)
	}

	return config, nil
}
