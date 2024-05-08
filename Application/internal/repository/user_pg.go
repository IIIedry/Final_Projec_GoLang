package repository

import (
	"Application"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
	"golang.org/x/crypto/bcrypt"
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

func (r *UsersPgx) GetAllUser(ctx *gin.Context) ([]Application.User, error) {
	var users []Application.User
	tx, err := r.conn.Begin(ctx)
	rows, err := tx.Query(ctx, "SELECT id, name, email, username, role FROM users;")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var user Application.User
		err = rows.Scan(&user.ID, &user.Name, &user.Email, &user.Username, &user.Role)
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

func (r *UsersPgx) IsAdmin(username, password string, ctx *gin.Context) (bool, error) {
	var isAdmin bool
	tx, err := r.conn.Begin(ctx)
	if err != nil {
		return false, err
	}

	err = tx.QueryRow(ctx, "SELECT EXISTS (SELECT 1 FROM users WHERE username = $1 AND password = $2 AND role = 'admin')", username, password).Scan(&isAdmin)
	if err != nil {
		return false, err
	}

	tx.Commit(ctx)
	return isAdmin, nil
}

func (r *UsersPgx) UpdateUserRole(userID int, newRole string, ctx *gin.Context) error {
	tx, err := r.conn.Begin(ctx)
	if err != nil {
		return err
	}

	_, err = tx.Exec(ctx, "UPDATE users SET role = $1 WHERE id = $2", newRole, userID)
	if err != nil {
		tx.Rollback(ctx)
		return err
	}

	err = tx.Commit(ctx)
	if err != nil {
		return err
	}

	return nil
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
