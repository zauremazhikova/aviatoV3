package config

const (
	FlightStopMaxNumber = 5
)

type StorageConfig struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Database string `json:"database"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func GetConfig() StorageConfig {

	return StorageConfig{
		Host:     "localhost",
		Port:     8080,
		Database: "aviato",
		Username: "postgres",
		Password: "778977",
	}

}
