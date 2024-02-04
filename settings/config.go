package settings

import "os"

type ConfigService struct {
	Host         string
	Port         string
	PhoneForAuth string
	MtsApiKey    string
}

type ConfigDB struct {
	DBUser     string
	DBPassword string
	DBName     string
}

func InitConfigService() ConfigService {
	config := ConfigService{
		Host:         os.Getenv("HOST"),
		Port:         os.Getenv("PORT"),
		PhoneForAuth: os.Getenv("PHONE"),
		MtsApiKey:    os.Getenv("OTP_MTS_KEY"),
	}

	return config
}

func InitConfigDB() ConfigDB {
	config := ConfigDB{
		DBUser:     os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBName:     os.Getenv("DB_NAME"),
	}

	return config
}
