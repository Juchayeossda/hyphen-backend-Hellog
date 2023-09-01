package configuration

import (
	"fmt"
	"hyphen-backend-hellog/exception"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewDatabase(config Config) *gorm.DB {
	username := config.Get("DATASOURCE_USERNAME")
	password := config.Get("DATASOURCE_PASSWORD")
	host := config.Get("DATASOURCE_HOST")
	port := config.Get("DATASOURCE_PORT")
	dbName := config.Get("DATASOURCE_DB_NAME")

	maxPoolOpen, err := strconv.Atoi(config.Get("DATASOURCE_POOL_MAXs"))
	maxPoolIdle, err := strconv.Atoi(config.Get("DATASOURCE_POOL_IDLE_CONN"))
	maxPollLifeTime, err := strconv.Atoi(config.Get("DATASOURCE_POOL_LIFE_TIME"))
	exception.InspectErr(err)

	loggerDb := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log 수준
			IgnoreRecordNotFoundError: true,        // 로거에 대한 ErrRecordNotFound 오류 무시
			Colorful:                  true,        // Enable color
		},
	)

	// dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"

	dsn := fmt.Sprintf("%s:%s+@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, password, host, port, dbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: loggerDb,
	})
	exception.InspectErr(err)

	sqlDB, err := db.DB()
	exception.InspectErr(err)

	sqlDB.SetMaxOpenConns(maxPoolOpen)
	sqlDB.SetMaxIdleConns(maxPoolIdle)
	sqlDB.SetConnMaxLifetime(time.Duration(rand.Int31n(int32(maxPollLifeTime))) * time.Millisecond)

	return db
}
