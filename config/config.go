package config

import (
	"log"
	"os"

	"github.com/VictorAvelar/mollie-api-go/mollie"
	"github.com/joho/godotenv"
)

func GetClient() *mollie.Client {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading env file")
	}

	config := mollie.NewConfig(true, os.Getenv("MOLLIE_API_TOKEN"))
	client, err := mollie.NewClient(nil, config)
	if err != nil {
		log.Fatal(err)
	}

	return client
}
