package config

import (
	"fmt"
	"log"
	"strings"

	"github.com/spf13/viper"
)

var Config *Configuration

func init() {
	conf, err := loadConfiguration(".", "config", "yml")
	if err != nil {
		log.Fatal("Loading viper config faild")
	}
	Config = conf
}

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
