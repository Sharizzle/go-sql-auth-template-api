package user

import (
	"golang.org/x/crypto/bcrypt"
	"os"
	"github.com/dgrijalva/jwt-go"
	"time" 
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email string `gorm:"unique_index" json:"email"`
	Name string `json:"name"`
	Hash string `json:"-"` 
}

type JWTToken struct {
	Token string `json:"token"`
}

func (u User) hashPassword(password string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), 4)
	return string(bytes)
}

func (u User) checkPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Hash), []byte(password))
	return err == nil
}

func (u User) generateJWT() (JWTToken, error) {
	signingKey := []byte(os.Getenv("JWT_SECRET"))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp": time.Now().Add(time.Hour * 1 * 1).Unix(),
		"user_id": int(u.ID),
		"name": u.Name,
		"email": u.Email,
	})
	tokenString, err := token.SignedString(signingKey)
	return JWTToken{tokenString}, err
}