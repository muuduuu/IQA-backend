package database

import(

	"log"
	"time"


	"github.com/muuduuu/IQA-BACKEND/configs"
    "github.com/muuduuu/IQA-BACKEND/models"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
    "gorm.io/gorm/logger"

)


var DB *gorm.DB

func ConnectDB(){
	configs.LoadEnv()
	dsn:= configs.GetEnv("DATABASE_URL")

	newLogger := logger.New(
		log.New(log.Writer(),"\r\n",log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel: logger.Info,
			IgnoreRecordNotFoundError: true,
			Colorful: true,
		},
	)

	db, err := gorm.Open(postgres.Open(dsn),&&gorm.Config{
		Logger: newLogger
	})

	if err != nil{
		log.Fatal("Failed to connect: %v",err)
	}

	DB=db

	log.Println("Success!!")

}







