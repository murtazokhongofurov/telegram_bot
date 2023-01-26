package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"gitlab.com/telegram_bot/bot"
	"gitlab.com/telegram_bot/config"
	"gitlab.com/telegram_bot/storage"
)

func main() {
	cfg := config.Load(".")

	psqlUrl := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.Postgres.Host,
		cfg.Postgres.Port,
		cfg.Postgres.User,
		cfg.Postgres.Password,
		cfg.Postgres.Database,
	)

	psqlConn, err := sql.Open("postgres", psqlUrl)
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	strg := storage.NewStoragePg(psqlConn)

	botHandler := bot.New(cfg, strg)

	botHandler.Start()
}
