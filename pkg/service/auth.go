package service

import (
	"crypto/sha1"
	"fmt"
	"time"

	"github.com/Kusoff/todo-app"
	"github.com/Kusoff/todo-app/pkg/repository"
	"github.com/dgrijalva/jwt-go"
)

const (
	solt       = "nvehibehsm4637653mnejbnl"     // Соль для хэширования паролей
	signingKey = "grkji#&*#35FSFJija#4353KSFjh" // Ключ для подписи JWT
	tokenTTL   = 12 * time.Hour                 // Время жизни токена (12 часов)
)

type tokenClaims struct {
	jwt.StandardClaims
	UserId int `json: "user_id"`
}

type AuthService struct {
	repo repository.Authorithation
}

func NewAuthServise(repo repository.Authorithation) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user todo.User) (int, error) {
	user.Password = s.generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

func (s *AuthService) GenerateToken(username, password string) (string, error) {
	//get user from DB
	user, err := s.repo.GetUser(username, s.generatePasswordHash(password))
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt: time.Now().Unix(),
		},
		user.Id})

	return token.SignedString([]byte(signingKey))

}

func (s *AuthService) generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(solt)))
}
