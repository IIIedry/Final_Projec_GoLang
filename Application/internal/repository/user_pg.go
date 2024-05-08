package repository

import (
	"Application"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
	"log"
)

type UsersPgx struct {
	conn *pgx.Conn
}

func NewUsersPg(conn *pgx.Conn) *UsersPgx {
	return &UsersPgx{conn: conn}
}

func (r *UsersPgx) CreateUser(user Application.User, ctx *gin.Context) (string, error) {
	tx, err := r.conn.Begin(ctx)
	if err != nil {
		return "0", err
	}
	defer tx.Rollback(ctx)

	var name string
	row := tx.QueryRow(ctx, "INSERT INTO users (id, name , email, username, password, role) VALUES ($1, $2, $3, $4, $5, $6) RETURNING name", user.ID, user.Name, user.Email, user.Username, user.Password, user.Role)
	if err = row.Scan(&name); err != nil {
		tx.Rollback(ctx)
		return "0", err
	}
	if err != nil {
		return "0", err
	}
	return name, tx.Commit(ctx)
}
func (r *UsersPgx) GetAllUser(ctx *gin.Context) ([]Application.User, error) {
	var users []Application.User
	tx, err := r.conn.Begin(ctx)
	rows, err := tx.Query(ctx, "SELECT * FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var user Application.User
		err = rows.Scan(&user.ID, &user.Name, &user.Email, &user.Username, &user.Password, &user.Role)
		log.Println(user)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
		log.Println(users)
	}
	tx.Commit(ctx)
	return users, nil
}
