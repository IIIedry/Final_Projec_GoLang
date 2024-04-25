package auth

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// Простая имитация базы данных пользователей
var users = map[string]string{
	"user1@example.com": "password1",
	"user2@example.com": "password2",
}

var jwtKeys = []byte("your_secret_key")

type CustomClaims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

// Функция для аутентификации пользователя и генерации JWT токена
func SignIn(email, password string) (string, error) {
	// Получение пароля для данного email из базы данных
	expectedPassword, ok := users[email]
	if !ok {
		return "", errors.New("user not found")
	}

	// Проверка совпадения паролей
	if password != expectedPassword {
		return "", errors.New("incorrect password")
	}

	// Создание токена JWT
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &CustomClaims{
		Email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// Функция для регистрации нового пользователя и генерации JWT токена
func SignUp(email, password, confirmPassword string) (string, error) {
	// Проверка наличия пользователя в базе данных
	if _, ok := users[email]; ok {
		return "", errors.New("user already exists")
	}

	// Проверка совпадения паролей
	if password != confirmPassword {
		return "", errors.New("passwords do not match")
	}

	// Добавление нового пользователя в базу данных
	users[email] = password

	// Создание токена JWT
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &CustomClaims{
		Email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
