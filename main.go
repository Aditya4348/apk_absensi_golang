package main

import (
	"apk_absesnsi_harian/database"
	"apk_absesnsi_harian/route"
	"fmt"
	"log"
	"net/http"
)

func main() {
	err := database.ConnectDB()
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	route.InitRoutes(database.DB)

	fmt.Println("Server is running on port 8010")
	log.Fatal(http.ListenAndServe(":8010", nil))
}
