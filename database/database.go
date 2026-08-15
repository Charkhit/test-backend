package database

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/yourname/backend/config"

	_ "github.com/jackc/pgx/v5/stdlib"
)

var DB *sql.DB

func Connect(cfg *config.Config) {
	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s",
		cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName, cfg.DBSSLMode,
	)

	db, err := sql.Open("pgx", dsn)
	if err != nil {
		log.Fatalf("เชื่อมต่อ database ไม่สำเร็จ: %v", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatalf("ping database ไม่สำเร็จ: %v", err)
	}

	log.Println("เชื่อมต่อ database สำเร็จ")
	DB = db
}
