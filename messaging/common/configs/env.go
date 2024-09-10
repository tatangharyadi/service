package configs

import (
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/spf13/viper"
)

type Env struct {
	AppEnv                    string `mapstructure:"APP_ENV"`
	AppPort                   string `mapstructure:"APP_PORT"`
	FirebaseProjectId         string `mapstructure:"FIREBASE_PROJECT_ID"`
	FirebaseServiceAccountKey string `mapstructure:"FIREBASE_SERVICE_ACCOUNT_KEY"`
}

var logger zerolog.Logger

func InitEnv() (*Env, zerolog.Logger) {
	zerolog.LevelFieldName = "severity"
	logger = zerolog.New(os.Stderr).Level(zerolog.InfoLevel)

	env := Env{}
	viper.BindEnv("APP_ENV")
	viper.BindEnv("APP_PORT")
	viper.BindEnv("FIREBASE_PROJECT_ID")
	viper.BindEnv("FIREBASE_SERVICE_ACCOUNT_KEY")

	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			logger.Info().Msg("The .env file not found")
		} else {
			logger.Error().Msg("Error reading .env file")
		}
	}

	viper.AutomaticEnv()
	err = viper.Unmarshal(&env)
	if err != nil {
		logger.Fatal().Msg("Error unmarshalling env")
	}

	if env.AppEnv == "DEV" {
		output := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}
		logger = zerolog.New(output).
			Level(zerolog.Level(zerolog.DebugLevel)).
			With().Timestamp().
			Logger()
	}

	logger.Info().Msgf("Starting %s mode:%s", env.AppEnv, env.AppPort)

	return &env, logger
}
