package config

import (
	"bufio"
	"log"
	"os"
	"fmt"
	"strings"
	"sync"
)

// Config Public Interface ---------------------------------------------------------------
const (
	PortMain = iota
	PortMock

	PostgresUser
	PostgresPassword
	PostgresDB
	PostgresSSL
	PostgresTimezone
	DatabasePort
	DatabaseHost
	PgadminPW
	PgadminMail

	JwtTokenSecretKey
	JwtTokenExpireHours
	EncryptionSecretKey

	ReservedParameter
)

func Get(envVariableName int) string {
	if configMap == nil {
		initConfig()
	}
	variable := configMap[envVariableName]
	if variable == "" {
		errFmt := "Environment variable with index:%d is empty or it doesn't exists"
		errMsg := fmt.Sprintf(errFmt, envVariableName)
		log.Fatal(errMsg)
	}
	return configMap[envVariableName]
}

// Config Private Implementation ---------------------------------------------------------
var (
	configOnce sync.Once
	configMap  map[int]string
)

func init() {
	initEnv()
	defer initConfig()
}

func initConfig() {
	err := loadConfig()
	if err != nil {
		log.Fatal(err)
	}
}

func loadConfig() error {
	configOnce.Do(func() {
		configMap = make(map[int]string)
		configMap[PortMain] = getEnv("PORT_MAIN")
		configMap[PortMock] = getEnv("PORT_MOCK")

		configMap[PostgresUser] = getEnv("POSTGRES_USER")
		configMap[PostgresPassword] = getEnv("POSTGRES_PASSWORD")
		configMap[PostgresDB] = getEnv("POSTGRES_DB")
		configMap[PostgresSSL] = getEnv("POSTGRES_SSL")
		configMap[PostgresTimezone] = getEnv("POSTGRES_TIMEZONE")
		configMap[DatabasePort] = getEnv("DATABASE_PORT")
		configMap[DatabaseHost] = getEnv("DATABASE_HOST")
		configMap[PgadminPW] = getEnv("PGADMIN_PW")
		configMap[PgadminMail] = getEnv("PGADMIN_MAIL")

		configMap[JwtTokenSecretKey] = getEnv("JWT_TOKEN_SECRET_KEY")
		configMap[JwtTokenExpireHours] = getEnv("JWT_TOKEN_EXPIRE_HOURS")
		configMap[EncryptionSecretKey] = getEnv("ENCRYPTION_SECRET_KEY")

		configMap[ReservedParameter] = getEnv("RESERVED_PARAMETER")
	})
	return nil
}

// Env Private Implementation ------------------------------------------------------------
const (
	configFileName        = ".env"
	separatorOfConfigFile = "="
)

var (
	envOnce sync.Once
	envMap  map[string]string
)

func initEnv() {
	err := loadEnvFile()
	if err != nil {
		log.Fatal(err)
	}
}

func getEnv(envName string) string {
	if envMap == nil {
		initEnv()
	}
	return envMap[envName]
}

func loadEnvFile() error {
	envOnce.Do(func() {
		envMap = make(map[string]string)
		envFile, err := os.Open(configFileName)
		if err != nil {
			log.Fatal(err)
		}
		defer envFile.Close()

		scanner := bufio.NewScanner(envFile)
		for scanner.Scan() {
			line := scanner.Text()
			keyVal := strings.Split(line, separatorOfConfigFile)
			if len(keyVal) != 2 {
				continue
			}
			envMap[keyVal[0]] = keyVal[1]
		}
		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
	})
	return nil
}
