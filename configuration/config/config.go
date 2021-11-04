package config

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/joeshaw/envdecode"
	"github.com/joho/godotenv"
	"github.com/pkg/errors"
)

// Config holds configuration for the project.
type Config struct {
	Port           string `env:"PORT,default=6666"`
	Env            string `env:"ENV,default=development"`
	Database       DatabaseConfig
	JWTConfig      JWTConfig
	InternalConfig InternalConfig
}

// DatabaseConfig holds configuration for database.
type DatabaseConfig struct {
	Host     string `env:"DATABASE_HOST,default=localhost"`
	Port     string `env:"DATABASE_PORT,default=5432"`
	Username string `env:"DATABASE_USERNAME,required"`
	Password string `env:"DATABASE_PASSWORD,required"`
	Name     string `env:"DATABASE_NAME,required"`
}

// JWTConfig holds configuration for JWT secret key
type JWTConfig struct {
	SecretKey string `env:"JWT_SECRET_KEY,required"`
}

// InternalConfig holds configuration for internal communication between microservices.
type InternalConfig struct {
	Username string `env:"SVC_USERNAME,required"`
	Password string `env:"SVC_PASSWORD,required"`
}

// NewConfig creates an instance of Config.
// It needs the path of the env file to be used.
func NewConfig(env string) (*Config, error) {
	var config Config
	if err := godotenv.Load(env); err != nil {
		log.Println(errors.Wrap(err, "[NewConfig] error reading .env file, defaulting to OS environment variables"))
	}

	if err := envdecode.Decode(&config); err != nil {
		return nil, errors.Wrap(err, "[NewConfig] error decoding env")
	}

	return &config, nil
}

func OpenDatabase(config *Config) *gorm.DB {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=require",
		config.Database.Host,
		config.Database.Port,
		config.Database.Username,
		config.Database.Password,
		config.Database.Name)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	CheckError(err)
	return db
}

// func CreateConnection() *sql.DB {
// 	// load .env file
// 	err := godotenv.Load("user.env")

// 	if err != nil {
// 		log.Fatalf("Error loading .env file")
// 	}

// 	// Kita buka koneksi ke db
// 	db, err := sql.Open("postgres", os.Getenv("POSTGRES_URL"))

// 	if err != nil {
// 		panic(err)
// 	}

// 	// check the connection
// 	err = db.Ping()

// 	if err != nil {
// 		panic(err)
// 	}
// 	return db
// }

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
