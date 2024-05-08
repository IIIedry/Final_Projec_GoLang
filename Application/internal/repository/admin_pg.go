package repository

import (
	"Application"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
	"log"
)

type AdminPgx struct {
	conn *pgx.Conn
}

func NewAdminPg(db *pgx.Conn) *AdminPgx {
	return &AdminPgx{conn: db}
}

func (r *AdminPgx) GetAllUser(ctx *gin.Context) ([]Application.User, error) {
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

func (r *AdminPgx) UpdateUserRole(userID int, newRole string, ctx *gin.Context) error {
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

func (r *AdminPgx) IsAdmin(username, password string, ctx *gin.Context) (bool, error) {
	var isAdmin bool
	tx, err := r.conn.Begin(ctx)
	if err != nil {
		return false, err
	}

	err = tx.QueryRow(ctx, "SELECT EXISTS (SELECT 1 FROM users WHERE username = $1 AND password = $2 AND role = 'admin')", username, password).Scan(&isAdmin)
	if err != nil {
		return false, err
	}
	log.Println(isAdmin)
	tx.Commit(ctx)
	return isAdmin, nil
}
