package models

import "github.com/dgrijalva/jwt-go"

type UserConn struct {
	Name     string `json:"name"`
	Password string `json:"password"`

	Shared
}

type Token struct {
	Name string `json:"name"`
	Id   string `json:"id"`
	jwt.StandardClaims
}

type UserInfo struct {
	Name string `json:"name"`
	Id   string `json:"id"`
}

func (uc UserConn) ToUserInfo() UserInfo {
	return UserInfo{
		Id:   uc.Id,
		Name: uc.Name,
	}
}
