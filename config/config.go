package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type DatabaseConfig struct {
	Username string
	Password string
	Host     string
	Port     int
	Database string
}

type ServerConfig struct {
	Host     string
	Hostname string
	Port     int
	LogFile  string
}

type FiwareConfig struct {
	OrionCB_Host    string
	OrionCB_Port    int
	Cygnus_Host     string
	Cygnus_Port     int
	Cygnus_Database string
}

type FlinkConfig struct {
	Host string
	Port int
}

type Config struct {
	Database DatabaseConfig
	Server   ServerConfig
	Fiware   FiwareConfig
	Flink    FlinkConfig
}

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func New() *Config {
	return &Config{
		Server: ServerConfig{
			Host:     getEnv("HOST", "127.0.0.1"),
			Hostname: getEnv("HOSTNAME", "beegons"),
			Port:     getEnvAsInt("PORT", 9000),
		},
		Database: DatabaseConfig{
			Username: getEnv("DB_USERNAME", ""),
			Password: getEnv("DB_PASSWORD", ""),
			Host:     getEnv("DB_HOST", "127.0.0.1"),
			Port:     getEnvAsInt("DB_PORT", 27027),
			Database: getEnv("DB_DATABASE", "defaultdb"),
		},
		Fiware: FiwareConfig{
			OrionCB_Host:    getEnv("ORION_CB_HOST", "localhost"),
			OrionCB_Port:    getEnvAsInt("ORION_CB_PORT", 1026),
			Cygnus_Host:     getEnv("CYGNUS_HOST", "localhost"),
			Cygnus_Port:     getEnvAsInt("CYGNUS_PORT", 5050),
			Cygnus_Database: getEnv("CYGNUS_DATABASE", "sth_default"),
		},
		Flink: FlinkConfig{
			Host: getEnv("FLINK_HOST", "localhost"),
			Port: getEnvAsInt("FLINK_PORT", 9001),
		},
	}
}

// Simple helper function to read an environment or return a default value
func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}

// Simple helper function to read an environment variable into integer or return a default value
func getEnvAsInt(name string, defaultVal int) int {
	valueStr := getEnv(name, "")
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}

	return defaultVal
}
