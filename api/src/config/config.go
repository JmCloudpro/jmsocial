package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv" //to import this package, go to api main folder and run in cmd :  go get github.com/joho/godotenv
)

var (
	Dbcon     = ""   //dbcon string connect to db
	Dbport    = 0    //db port
	SecretKey []byte //its the key used to assign the authentication token

)

// Load  loads environment variables
func Load() {
	var err error
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	Dbport, err = strconv.Atoi(os.Getenv("API_PORT"))
	if err != nil {
		Dbport = 9000
	}
	Dbcon = fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWD"),
		os.Getenv("DB_NAME"),
	)
	SecretKey = []byte(os.Getenv("SECRET_KEY"))
}
