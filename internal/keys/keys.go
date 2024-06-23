package keys

import (
	"fmt"
	"strings"

	gonanoid "github.com/matoous/go-nanoid/v2"
	"golang.org/x/crypto/bcrypt"
)

// CreateUserId UserIdを作成する
func CreateUserId(shardKey string) (string, error) {
	uuid, err := gonanoid.New(20)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s:%s", shardKey, uuid), nil
}

// CreateRoomId RoomIdを作成する
func CreateRoomId() (string, error) {
	uuid, err := gonanoid.New(20)
	if err != nil {
		return "", err
	}

	return uuid, nil
}

// CreateStateOauthCookie RoomIdを作成する
func CreateStateOauthCookie() (string, error) {
	uuid, err := gonanoid.New(20)
	if err != nil {
		return "", err
	}

	return uuid, nil
}

// GetShardKeyByUserId ユーザIDからシャードキーを取得する
func GetShardKeyByUserId(userID string) string {
	return strings.Split(userID, ":")[0]
}

// CreatePassword パスワードを作成する
func CreatePassword() (string, error) {
	password, err := gonanoid.New(20)
	if err != nil {
		return "", err
	}

	return password, nil
}

// CreateHashPassword ハッシュパスワードを作成する
func CreateHashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}

// CheckPassword パスワードを検証する
func CheckPassword(password, hashPassword string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(password)); err != nil {
		return false
	}

	return true
}
