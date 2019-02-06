package service

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/prometheus/common/log"
	"github.com/ubccr/goipa"
	. "hiveon-api/utils"
)

type User struct {
	Name     string `json:"name"`
	Password string `json:"email"`
	jwt.StandardClaims
}

func getAdminLdapConnection() *ipa.Client {
	conn := ipa.NewClient(GetConfig().GetString("ldap.host"), "")
	err := conn.RemoteLogin(GetConfig().GetString("ldap.user"), GetConfig().GetString("ldap.password"))
	if err != nil {
		log.Error(err)
	}
	return conn
}
func getUserLdapConnection(username string, password string) (*ipa.Client, error) {
	conn := ipa.NewClient(GetConfig().GetString("ldap.host"), "")
	err := conn.RemoteLogin(username,password)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return conn, nil
}
func Register(username string, password string) (string, error) {
	conn := getAdminLdapConnection()

	if _, err := conn.UserShow(username); err == nil {
		return "", errors.New("user already exists")
	}

	_, err := conn.UserAdd(username, password, "test", "1", "2", "3", false)
	if err != nil {
		log.Error(err)
		return "", err
	}

	token, err := generateToken(username, password)
	return token, err

}

func generateToken(username string, password string) (string, error) {
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), &User{
		Name:     username,
		Password: password,
	})
	tokenString, err := token.SignedString([]byte(GetConfig().GetString("ldap.key")))
	return tokenString, err
}

func Login(username string, password string) (string, error) {
	if _, err:= getUserLdapConnection(username, password); err != nil{
		return "", err
	}
	return generateToken(username, password)
}
