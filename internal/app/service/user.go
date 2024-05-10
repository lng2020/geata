package service

import (
	"geata/internal/app/model"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type Credential struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type User struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	Role     string `json:"role"`
}

func generateToken(user *model.User) (string, error) {
	claims := jwt.MapClaims{
		"user_id": user.ID,
		"role":    user.RoleID,
		"exp":     time.Now().Add(time.Hour * 72).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte("your_secret_key"))
}

func generatePasswordHash(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func roleTypeToRole(rt model.RoleType) string {
	if rt == model.RoleAdmin {
		return "admin"
	}
	return "user"
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func Login(c *gin.Context) {
	var credential Credential
	if err := c.ShouldBindJSON(&credential); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := model.GetUserByUsername(Engine, credential.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if user == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	if checkPasswordHash(credential.Password, user.PasswordHash) {
		token, err := generateToken(user)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"token": token,
			"user": &User{
				ID:       user.ID,
				Username: user.Username,
				Role:     roleTypeToRole(model.RoleType(user.RoleID)),
			},
		})
		return
	}
	c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
}

func Register(c *gin.Context) {
	var credential Credential
	if err := c.ShouldBindJSON(&credential); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	exist, err := model.IfUsernameExist(Engine, credential.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if exist {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Username already exists"})
		return
	}

	passwordHash, err := generatePasswordHash(credential.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	hasUser, err := model.HasUser(Engine)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	roleID := model.RoleUser
	if !hasUser {
		roleID = model.RoleAdmin
	}

	user := model.User{
		Username:     credential.Username,
		PasswordHash: passwordHash,
		RoleID:       int64(roleID),
	}

	err = model.CreateUser(Engine, &user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	token, err := generateToken(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": token,
		"user": &User{
			ID:       user.ID,
			Username: user.Username,
			Role:     roleTypeToRole(model.RoleType(user.RoleID)),
		},
	})
}
