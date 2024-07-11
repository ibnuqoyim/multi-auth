package configs

import (
	"context"
	"log"
	"time"

	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Config struct {
	MongoURI             string `mapstructure:"MONGO_URI"`
	RedisAddr            string `mapstructure:"REDIS_ADDR"`
	JWTSecret            string `mapstructure:"JWT_SECRET"`
	GoogleClientID       string `mapstructure:"GOOGLE_CLIENT_ID"`
	GoogleClientSecret   string `mapstructure:"GOOGLE_CLIENT_SECRET"`
	FacebookClientID     string `mapstructure:"FACEBOOK_CLIENT_ID"`
	FacebookClientSecret string `mapstructure:"FACEBOOK_CLIENT_SECRET"`
	GitHubClientID       string `mapstructure:"GITHUB_CLIENT_ID"`
	GitHubClientSecret   string `mapstructure:"GITHUB_CLIENT_SECRET"`
}

func LoadConfig() (config Config, err error) {
	viper.AddConfigPath(".")
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error while reading config file %s", err)
	}
	err = viper.Unmarshal(&config)
	return
}

func ConnectDB(uri string) (*mongo.Database, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}

	db := client.Database("login-service")
	return db, nil
}
