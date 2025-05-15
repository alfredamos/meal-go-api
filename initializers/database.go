package initializers

import (
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	var err error
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	//mysql:"//root:Aladjopape1961!@localhost:3306/meal-order-db"
	//dsn := "user:Aladjopape1961!@tcp(127.0.0.1:3306)/meal-order-db?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := os.Getenv("DATABASE_URL")
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to database!")
	}

}