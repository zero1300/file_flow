package helper

import (
	"errors"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"math/rand"
	"strconv"
	"time"
)

func GetUserIdByToken(ctx *gin.Context) (int, error) {
	get, exists := ctx.Get("id")
	id := get.(int)
	if exists {
		return id, nil
	}
	return -1, errors.New("")
}

func GenerateVerificationCode() string {
	rand.Seed(time.Now().UnixNano())
	fourDigitNumber := rand.Intn(1000) + 1000
	code := strconv.Itoa(fourDigitNumber)
	return code
}

// GenPassword 生成密码
func GenPassword(password string) (string, error) {
	hashedPassword, errHash := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if errHash != nil {
		return "", errHash
	}
	return string(hashedPassword), nil
}

// ComparePassword 校验密码, false: 密码错误, true: 密码正确
func ComparePassword(password, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return false
	} else {
		return true
	}
}
