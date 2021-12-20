package main

//"fmt"
//"heitor/programas/ordena"
import (
	"heitor/programas/library"
	"log"
	"os"

	"github.com/joho/godotenv"
)

//checar diferencas com https://blog.logrocket.com/making-http-requests-in-go/
func main() {
	//calculadora.Calculadora()
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	host := os.Getenv("HOST")
	dbport := os.Getenv("DBPORT")
	user := os.Getenv("USER")
	dbName := os.Getenv("NAME")
	password := os.Getenv("PASSWORD")

	db := library.SetEnviroment(host, user, dbName, password, dbport)
	library.Library(db)
}
