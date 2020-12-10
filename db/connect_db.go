package db

import (
	"database/sql"
	"fmt"
	"log"

	"os"

	_ "github.com/go-sql-driver/mysql"

	"github.com/joho/godotenv"
)

func ConnectSQL() (*sql.DB, error) {
	err := godotenv.Load("./.env")
	if err != nil {
		log.Fatal(err)
	}

	//.envファイルからdataSourceNameを作成してDBに接続する
	dsn := fmt.Sprintf("%s:%s@tcp(db:3306)/%s", os.Getenv("DB_ROOT_USER"), os.Getenv("DB_ROOT_PASS"), os.Getenv("DB_DATABASE"))
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}

	return db, err
}
