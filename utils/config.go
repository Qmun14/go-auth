package utils

import "github.com/spf13/viper"

// Config stores untuk semua konfigurasi aplikasi,
// nilai nya dibaca oleh viper dari sebuah config file atau environment variables.
type Config struct {
	ServerAddress       string `mapstructure:"SERVER_ADDRESS"`
	ClientAddress       string `mapstructure:"CLIENT_ADDRESS"`
	EmailSenderName     string `mapstructure:"EMAIL_SENDER_NAME"`
	EmailSenderAddress  string `mapstructure:"EMAIL_SENDER_ADDRESS"`
	EmailSenderPassword string `mapstructure:"EMAIL_SENDER_PASSWORD"`
	DBHost              string `mapstructure:"DB_HOST"`
	DBUser              string `mapstructure:"DB_USER"`
	DBPass              string `mapstructure:"DB_PASS"`
	DBName              string `mapstructure:"DB_NAME"`
	DBPort              string `mapstructure:"DB_PORT"`
}

// LoadConfig membaca configurasi dari file atau environment variables.
func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()
	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
