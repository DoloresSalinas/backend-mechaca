package config

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/jackc/pgx/v5"
)

var Conn *pgx.Conn

func InitDB() {
	err := godotenv.Load()
	if err != nil {
		log.Println("⚠️ No se pudo cargar .env, usando entorno actual")
	}

	user := os.Getenv("user")
	pass := os.Getenv("password")
	host := os.Getenv("host")
	port := os.Getenv("port")
	dbname := os.Getenv("dbname")

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", user, pass, host, port, dbname)

	conn, err := pgx.Connect(context.Background(), dsn)
	if err != nil {
		log.Fatalf("❌ No se pudo conectar a Supabase: %v", err)
	}

	Conn = conn
	log.Println("✅ Conectado a la base de datos Supabase")
}
