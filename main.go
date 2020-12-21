package main

import (
	"awesomeProject1/routes"
	"fmt"
	"github.com/Kamva/mgm"
	"go.mongodb.org/mongo-driver/mongo/options"
	"net/http"
	"os"
)

func init() {
	err := mgm.SetDefaultConfig(nil, os.Getenv("DB_NAME"), options.Client().ApplyURI(os.Getenv("DB_URI")))
	if err != nil {
		panic(err)
	}
}

func main() {
	http.HandleFunc("/images", routes.HandleImage)

	fmt.Println("Listening...")

	err := http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("PORT")), nil)
	if err != nil {
		panic(err)
	}
}
