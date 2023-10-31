package models

import (
	"github.com/golang-jwt/jwt"
)

type userData struct{
	Username string
	Password string
}

type ArticlesData struct{
	Articles []ArticleData `json:"articles"`
}


type ArticleData struct{
	Author string `json:"author"`
	Title string  `json:"title"`
	Description string `json:"description"`
	Url string `json:"url"`
	UrlToImage string `json:"urlToImage"`
}

type CustomClaims struct {
	Username string `json:"User"`
	jwt.StandardClaims
}


var SecretKey = []byte("SecretYouShouldHide")