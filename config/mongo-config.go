package config

import (
	"os"

	"github.com/joho/godotenv"
)

type MongoConfig struct {
	DBLINK string
	DBNAME string
}

func InitConfig() *MongoConfig {
	var response = new(MongoConfig)
	response = readData()
	return response
}

func readData() *MongoConfig {
	var data = new(MongoConfig)

	data = readEnv()

	if data == nil {
		err := godotenv.Load(".env")
		data = readEnv()
		if err != nil || data == nil {
			return nil
		}
	}
	return data
}

func readEnv() *MongoConfig {
	var data = new(MongoConfig)
	var permit = true

	if val, found := os.LookupEnv("DBLINK"); found {
		data.DBLINK = val
	} else {
		permit = false
	}

	if val, found := os.LookupEnv("DBNAME"); found {
		data.DBNAME = val
	} else {
		permit = false
	}

	if !permit {
		return nil
	}
	return data
}