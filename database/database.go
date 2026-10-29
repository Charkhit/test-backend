package database

import (
	"database/sql"
	"log"

	"github.com/chakhit/backend/config"
	_ "github.com/microsoft/go-mssqldb"
)

func Connect(cfg *config.Config) *sql.DB {
	dsn := cfg.DBUrl
	db, err := sql.Open("sqlserver", dsn)
	if err != nil {
		log.Fatalf("เชื่อมต่อ database ไม่สำเร็จ: %v", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatalf("ping database ไม่สำเร็จ: %v", err)
	}

	log.Println("เชื่อมต่อ database สำเร็จ")
	return db
}
