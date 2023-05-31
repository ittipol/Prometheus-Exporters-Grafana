package main

import (
	"fmt"
	"prometheus/database"
	"strings"
	"time"

	"github.com/spf13/viper"
	"gorm.io/gorm"
)

func init() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
}

type product struct {
	ID       int
	Name     string `gorm:"size:100"`
	Quantity int
}

func main() {

	initTimeZone()
	db := initDbConnection()

	db.AutoMigrate(product{})

	for i := 0; i < 2000; i++ {
		db.Create(&product{
			Name:     fmt.Sprintf("Name %v", i),
			Quantity: i,
		})
	}

}

func initTimeZone() {
	// LoadLocation looks for the IANA Time Zone database
	// List of tz database time zones
	// https: //en.wikipedia.org/wiki/List_of_tz_database_time_zones
	location, err := time.LoadLocation("Asia/Bangkok")
	if err != nil {
		panic(err)
	}

	// init system time zone
	time.Local = location

	// timeInUTC := time.Date(2018, 8, 30, 12, 0, 0, 0, time.UTC)
	// fmt.Println(timeInUTC.In(location))
}

func initDbConnection() *gorm.DB {
	return database.GetDbConnection(
		viper.GetString("database.username"),
		viper.GetString("database.password"),
		viper.GetString("database.host"),
		viper.GetInt("database.port"),
		viper.GetString("database.db"),
		false,
	)
}
