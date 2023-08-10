package jectools

import (
	"database/sql"
	"os"
)

// Get Environtment Value based on Key
func GetEnv(key string) (value string) {
	value = os.Getenv(key)

	return
}

func NewSQLNullString(s string) sql.NullString {
	if len(s) == 0 {
		return sql.NullString{}
	}

	return sql.NullString{
		String: s,
		Valid:  true,
	}
}
