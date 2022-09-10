package utills

import "golang.org/x/crypto/bcrypt"

func PasswordUtilsHashing(password string) (string, error) {
	hashedby,err := bcrypt.GenerateFromPassword([]byte(password),15)
	if err != nil{
		return "",err
	}
	return string(hashedby),nil
}

func PasswordUtilsDecodeHashing(password string,hashedpassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedpassword), []byte(password))
	return err == nil 
}