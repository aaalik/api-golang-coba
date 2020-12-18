package main

import (
	"github.com/aaalik/coba-golang/bootstrap"
	routers "github.com/aaalik/coba-golang/routers"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	db := bootstrap.Database{}
	db.Connect()
	defer db.CloseConnection()

	routers.SetupRouter()
}
