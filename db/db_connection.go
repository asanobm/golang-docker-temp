package db

import (
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // Use PostgreSQL in gorm
	"github.com/joho/godotenv"
)

// Conn function can get a database connection.
func Conn() *gorm.DB {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Environment variables load error.")
	}
	db, err := gorm.Open(
		"postgres",
		"host="+os.Getenv("DATABASE_HOST") +" "+
		"port="+os.Getenv("DATABASE_PORT") +" "+
		"dbname="+os.Getenv("DATABASE_NAME") +" "+
		"user="+os.Getenv("DATABASE_USER") +" "+
		"password="+os.Getenv("DATABASE_PASSWORD") +" "+
		"sslmode=disable")
	if err != nil {
		log.Panic(err)
	}
	log.Println("Connected Database")
	return db
}

