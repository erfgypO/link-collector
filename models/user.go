package models

import (
	"errors"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"strings"
)

type User struct {
	Id          uint          `json:"id" gorm:"primary_key"`
	Username    string        `json:"username" gorm:"unique"`
	Password    []byte        `json:"-"`
	AccessToken []AccessToken `json:"accessTokens"`
	Links       []Link        `json:"links"`
}

type AccessToken struct {
	Id     uint   `json:"id" gorm:"primary_key"`
	Token  string `json:"token" gorm:"unique"`
	UserID uint   `json:"userId"`
}

type UserLoginDto struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required,min=8"`
}

func containsUser(username string) bool {
	var count int64

	DB.Model(&User{}).Where("username = ?", strings.ToLower(username)).Count(&count)

	return count != 0
}

func CreateUser(login UserLoginDto) (User, error) {
	var user User

	if containsUser(login.Username) {
		return user, errors.New("username is already in use")
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(login.Password), bcrypt.DefaultCost)
	if err != nil {
		return User{}, err
	}

	user = User{
		Username: strings.ToLower(login.Username),
		Password: passwordHash,
	}

	DB.Create(&user)

	return user, nil
}

func CreateAccessToken(login UserLoginDto) (AccessToken, error) {
	var user User
	var accessToken AccessToken
	result := DB.First(&user, "username = ?", strings.ToLower(login.Username))

	if result.Error != nil || bcrypt.CompareHashAndPassword(user.Password, []byte(login.Password)) != nil {
		return accessToken, errors.New("invalid username or password")
	}

	accessToken = AccessToken{
		UserID: user.Id,
		Token:  strings.ReplaceAll(uuid.NewString()+uuid.NewString(), "-", ""),
	}

	DB.Create(&accessToken)

	return accessToken, nil
}

func DeleteAccessToken(accessToken string) {
	DB.Delete(&AccessToken{}, "token = ?", accessToken)
}

func GetUserIdByToken(token string) (uint, error) {
	var accessToken AccessToken

	result := DB.First(&accessToken, "token = ?", token)

	if result.Error != nil {
		return 0, errors.New("invalid access token")
	}

	return accessToken.UserID, nil
}
