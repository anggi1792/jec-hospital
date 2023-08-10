package jecconfiguration

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func ConnectSqlx(config DbConfiguration) (db *sqlx.DB, err error) {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s TimeZone=Asia/Jakarta",
		config.Host,
		config.Port,
		config.Dbuser,
		config.Dbpassword,
		config.Dbname,
		config.Sslmode,
	)

	db, err = sqlx.Connect("postgres", dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return
}
