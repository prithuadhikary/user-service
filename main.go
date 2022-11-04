package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/prithuadhikary/user-service/controller"
	"github.com/prithuadhikary/user-service/domain"
	"github.com/prithuadhikary/user-service/repository"
	"github.com/prithuadhikary/user-service/service"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"os"
	"time"
)

func main() {
	db, err := InitialiseDB(&DbConfig{
		User:     "postgres",
		Password: "password",
		DbName:   "groot",
		Host:     "localhost",
		Port:     "5432",
		Schema:   "users",
	})
	if err != nil {
		panic(err)
	}
	err = db.AutoMigrate(&domain.User{})
	if err != nil {
		panic(err)
	}

	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)

	engine := gin.Default()

	controller.NewUserController(engine, userService)

	log.Fatal(engine.Run(":8080"))
}

func InitialiseDB(dbConfig *DbConfig) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=require TimeZone=Asia/Kolkata", dbConfig.Host, dbConfig.User, dbConfig.Password, dbConfig.DbName, dbConfig.Port)
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			Colorful:                  true,        // Disable color
		},
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: newLogger,
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   dbConfig.Schema + ".",
			SingularTable: false,
		},
	})
	if err != nil {
		return nil, err
	}
	return db, err
}

type DbConfig struct {
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	DbName   string `mapstructure:"dbName"`
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	Schema   string `mapstructure:"schema"`
}
