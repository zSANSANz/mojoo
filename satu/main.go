package main

import (
	"Satu/config"
	"Satu/routes"
	"fmt"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	config.InitDB()
	//config.DropTable() //reset tables
	config.InitialMigration()

	// seeders.Seed()
	// seeders.ItemSeed() 

	e := routes.New()

	e.Logger.Fatal(e.Start(":3000"))

}
