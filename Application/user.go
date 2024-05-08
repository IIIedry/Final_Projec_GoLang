package Application

type User struct {
	ID       int    `json:"ID" db:"id"`
	Name     string `json:"Name" db:"name"`
	Email    string `json:"Email" db:"email"`
	Username string `json:"Username" db:"username"`
	Password string `json:"Password" db:"password"`
	Role     string `json:"Role" db:"role"`
}
