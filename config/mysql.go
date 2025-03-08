package config

import (
	"errors"
	"log"
	"os"
	"time"

	_ "github.com/apache/calcite-avatica-go/v5"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var ErrRecordNotFound = errors.New("record not found")

func ConnectMySQL() (*gorm.DB, error) {
	connectionString := CONFIG["MYSQL_USER"] + ":" + CONFIG["MYSQL_PASS"] + "@tcp(" + CONFIG["MYSQL_HOST"] + ":" + CONFIG["MYSQL_PORT"] + ")/" + CONFIG["MYSQL_SCHEMA"] + "?parseTime=true&loc=Local"

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			Colorful:                  true,        // Disable color
		},
	)
	mysqlConn, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{Logger: newLogger})

	if err != nil {
		log.Println("Error connect to MySQL: ", err.Error())
		return nil, err
	}

	log.Println("MySQL connection success")
	return mysqlConn, nil

}
