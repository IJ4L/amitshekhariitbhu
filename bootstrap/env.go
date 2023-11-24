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
	containerAppDirectory := "/app" // Gantilah dengan working directory sesuai dengan struktur direktori di dalam container
	dirInContainer := filepath.Join(containerAppDirectory, "bootstrap")

	// Gabungkan direktori dalam container dengan path relatif ke file konfigurasi
	configFilePath := filepath.Join(dirInContainer, ".env")

	// Mengatur lokasi file konfigurasi
	viper.SetConfigFile(configFilePath)

	// Baca konfigurasi dari file
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("Tidak dapat membaca file konfigurasi:", err)
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
