package config

import (
	"os"
	"fmt"
	"log"
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
	Jwt struct {
		Token struct {
			Expire struct {
				Hours int
			}
			Secret struct {
				Key string
			}
		}
	}
	Encryption struct {
		Secret struct {
			Key string
		}
	}
	Payment struct {
		Redirect struct {
			Url string
		}
		Gateways struct {
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
		Reserved string
		DebugMode bool
		ImageLogo string
		TicketFileName string
	}
}

var Config *Configuration

func IsTestMode() {
	if testModeEnabled {
		fmt.Println("run under go test")
	} else {
		fmt.Println("normal run")
	}
}

func IsDebugMode() bool {
	return Config.App.DebugMode
}

func Load() {
	if Config == nil {
		load()
	}
}
//----------------------------------------------------------------------------------------
var (
	testModeEnabled bool
)

func init() {
	Load()
}

func load() {
	var err error
	var conf *Configuration
	testModeEnabled = strings.HasSuffix(os.Args[0], ".test")
	if testModeEnabled {
		address := os.Getenv("PWD")
		fmt.Println("  XX:", address)
		conf, err = loadConfiguration(address, "config", "yml")
	} else {
		conf, err = loadConfiguration(".", "config", "yml")
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
