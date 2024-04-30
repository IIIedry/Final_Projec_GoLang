package repository

import (
	"context"
	"github.com/jackc/pgx/v4"
)

type Config struct {
	User string
	Pass string
	Host string
	Port string
	DB   string
}

func NewConnection(conf Config) (*pgx.Conn, error) {
	connString := "postgres://" + conf.User + ":" + conf.Pass + "@" + conf.Host + ":" + conf.Port + "/" + conf.DB
	conn, err := pgx.Connect(context.Background(), connString)
	if err != nil {
		return nil, err
	}
	err = conn.Ping(context.Background())
	if err != nil {
		return nil, err
	}
	return conn, nil
}
