package utills

import (
	"fmt"

	"github.com/dgrijalva/jwt-go"
)

var SECRET_KEY = "XABACD!DTE"

func GenerateJwtUtils(claims *jwt.MapClaims) (string,error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,claims)
	webToken,err := token.SignedString([]byte(SECRET_KEY))
	if err != nil{
		return "",err
	}
	return webToken,nil
}

func VerifyJwt(tokenJwt string )(*jwt.Token, error){
	token ,err := jwt.Parse(tokenJwt,func(token *jwt.Token) (interface{}, error) {
		if _, isValid := token.Method.(*jwt.SigningMethodHMAC) ; !isValid{
			return nil , fmt.Errorf("unexpected signing method %v",token.Header["alg"])
		}
		return []byte(SECRET_KEY),nil
	})

	if err != nil{
		return nil,err
	}
	return token,nil
}

func DecodeToken(tokenString string) (jwt.MapClaims,error){
	token ,err := VerifyJwt(tokenString)
	if err != nil{
		return nil,err
	}
	claims ,isValid := token.Claims.(jwt.MapClaims)
	
	if isValid && token.Valid{
		return claims,nil
	}
	return nil,fmt.Errorf("Invalid Token")
}



