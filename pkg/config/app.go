package config

import(
	"os"
	"github.com/joho/godotenv"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db * gorm.DB
)

func getEnvVariable(key string) string {
	err := godotenv.Load(".env")
  
	if err != nil {
	  panic("Error loading .env file")
	}
  
	return os.Getenv(key)
  }

func Connect() {
	connStr := getEnvVariable("DB_USER")+":"+getEnvVariable("DB_PASSWORD")+"@tcp("+getEnvVariable("DB_HOST")+":"+getEnvVariable("DB_PORT")+")/"+getEnvVariable("DB_NAME")+"?charset=utf8mb4&parseTime=True&loc=Local"
	d, err := gorm.Open("mysql", connStr)

	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}