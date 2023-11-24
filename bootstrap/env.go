package bootstrap

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

type Env struct {
	AppEnv                 string `mapstructure:"APP_ENV"`
	ServerAddress          string `mapstructure:"SERVER_ADDRESS"`
	ContextTimeout         int    `mapstructure:"CONTEXT_TIMEOUT"`
	Host                   string `mapstructure:"DB_HOST"`
	Port                   string `mapstructure:"DB_PORT"`
	User                   string `mapstructure:"DB_USER"`
	Password               string `mapstructure:"DB_PASS"`
	Database               string `mapstructure:"DB_NAME"`
	AccessTokenExpiryHour  int    `mapstructure:"ACCESS_TOKEN_EXPIRY_HOUR"`
	RefreshTokenExpiryHour int    `mapstructure:"REFRESH_TOKEN_EXPIRY_HOUR"`
	AccessTokenSecret      string `mapstructure:"ACCESS_TOKEN_SECRET"`
	RefreshTokenSecret     string `mapstructure:"REFRESH_TOKEN_SECRET"`
}

func NewEnv() *Env {
	env := Env{}
	
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("Tidak dapat mendapatkan working directory:", err)
	}

	fmt.Println(dir)

	// Gabungkan direktori saat ini dengan path relatif ke file konfigurasi
	configFilePath := filepath.Join(dir, "bootstrap", ".env")

	// Mengatur lokasi file konfigurasi
	viper.SetConfigFile(configFilePath)

	err = viper.ReadInConfig()
	if err != nil {
		log.Fatal("Can't find the file .env : ", err)
	}

	err = viper.Unmarshal(&env)
	if err != nil {
		log.Fatal("Environment can't be loaded: ", err)
	}

	fmt.Println(env)

	if env.AppEnv == "development" {
		log.Println("The App is running in development env")
	}

	return &env
}
