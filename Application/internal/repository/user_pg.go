package repository

import (
	"Application"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
	"golang.org/x/crypto/bcrypt"
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

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return "0", err
	}

	var name string
	row := tx.QueryRow(ctx, "INSERT INTO users (name , email, username, password) VALUES ($1, $2, $3, $4) RETURNING name", user.Name, user.Email, user.Username, hashedPassword)
	if err = row.Scan(&name); err != nil {
		tx.Rollback(ctx)
		return "0", err
	}
	if err != nil {
		return "0", err
	}
	return name, tx.Commit(ctx)
}

func (r *UsersPgx) AuthenticateUser(username, password string, ctx *gin.Context) (*Application.User, error) {
	var user Application.User
	tx, err := r.conn.Begin(ctx)
	if err != nil {
		return nil, err
	}

	err = tx.QueryRow(ctx, "SELECT * FROM users WHERE username = $1 AND password = $2", username, password).
		Scan(&user.ID, &user.Name, &user.Email, &user.Username, &user.Password, &user.Role)
	if err != nil {
		return nil, err
	}

	tx.Commit(ctx)
	return &user, nil
}

func (r *UsersPgx) GetUserById(id int, ctx *gin.Context) (Application.User, error) {
	var user Application.User
	tx, err := r.conn.Begin(ctx)
	if err != nil {
		return user, err
	}
	row := tx.QueryRow(ctx, "SELECT id, name, email, username, role FROM users WHERE id = $1", id)
	err = row.Scan(&user.ID, &user.Name, &user.Email, &user.Username, &user.Role)
	if err != nil {
		return user, err
	}
	tx.Commit(ctx)
	return user, nil
}
