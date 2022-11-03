package main

import (
	"api/src/config"
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

/*
Funcion init runs before function main, it will be run once one to generate the secret for que token ecryption.
After generate, you put in .env file and comment this function
func init() {
	key := make([]byte, 64)
	if _, err := rand.Read(key); err != nil {
		log.Fatal(err)
	}
	stringBase64 := base64.StdEncoding.EncodeToString(key)
	fmt.Println(stringBase64)

}
*/

func main() {

	config.Load()

	r := router.Generate()
	fmt.Printf("Running on port: %d", config.Dbport)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Dbport), r))

}
