package main

import (
	"MyNote/internal/infrastructure/config"
	"MyNote/internal/infrastructure/db"
)

func main() {

	cfg := config.Load("/media/Akunamatata/Pet-project/MyNote/config/config.yaml")

	db.ConnectDB(cfg)
}
